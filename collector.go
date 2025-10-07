package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

// MySQLCollector handles MySQL information collection with version-specific logic
type MySQLCollector struct {
	db      *sql.DB
	version *MySQLVersion
	config  *Config
}

// NewMySQLCollector creates a new MySQL collector
func NewMySQLCollector(db *sql.DB, config *Config) (*MySQLCollector, error) {
	collector := &MySQLCollector{
		db:     db,
		config: config,
	}

	// Detect MySQL version
	version, err := DetectMySQLVersion(db)
	if err != nil {
		return nil, fmt.Errorf("failed to detect MySQL version: %v", err)
	}
	collector.version = version

	return collector, nil
}


// CollectAll collects all MySQL information
func (c *MySQLCollector) CollectAll() (*DatabaseInfo, error) {
	info := &DatabaseInfo{}

	// Collect connection information
	if err := c.collectConnectionInfo(info); err != nil {
		return nil, fmt.Errorf("failed to collect connection info: %v", err)
	}

	// Collect tables and views unless excluded
	if !c.config.ExceptTables {
		if err := c.collectTables(info); err != nil {
			return nil, fmt.Errorf("failed to collect tables: %v", err)
		}
	}

	// Collect users unless excluded
	if !c.config.ExceptUsers {
		if err := c.collectUsers(info); err != nil {
			return nil, fmt.Errorf("failed to collect users: %v", err)
		}
	}

	// Collect roles unless excluded (MySQL 8.0+ only)
	if !c.config.ExceptRoles && c.version.IsMySQL8OrLater() {
		if err := c.collectMySQL8Features(info); err != nil {
			return nil, fmt.Errorf("failed to collect MySQL 8.0+ features: %v", err)
		}
	}

	// Collect stored functions and procedures unless excluded
	if !c.config.ExceptStoredProcedures {
		if err := c.collectRoutines(info); err != nil {
			return nil, fmt.Errorf("failed to collect routines: %v", err)
		}
	}

	// Collect variables unless excluded
	if !c.config.ExceptVariables {
		if err := c.collectVariables(info); err != nil {
			return nil, fmt.Errorf("failed to collect variables: %v", err)
		}
	}

	// Collect plugins and components unless excluded
	if !c.config.ExceptPlugins {
		if err := c.collectPlugins(info); err != nil {
			return nil, fmt.Errorf("failed to collect plugins: %v", err)
		}
		// Also collect components for MySQL 8.0+ if plugins are not excluded
		if c.version.IsMySQL8OrLater() {
			if err := c.collectComponents(info); err != nil {
				return nil, fmt.Errorf("failed to collect components: %v", err)
			}
		}
	}

	// Collect replication info if requested
	if c.config.Replication {
		if err := c.collectReplicationInfo(info); err != nil {
			return nil, fmt.Errorf("failed to collect replication info: %v", err)
		}
	}

	return info, nil
}

// collectConnectionInfo collects connection and version information
func (c *MySQLCollector) collectConnectionInfo(info *DatabaseInfo) error {
	// Get version
	var version string
	err := c.db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		return err
	}

	// Create connection info
	info.ConnectionInfo = &ConnectionInfo{
		Host:     c.config.Host,
		Port:     c.config.Port,
		User:     c.config.User,
		Database: c.config.Database,
		Version:  version,
	}

	return nil
}

// collectTables collects table and view information
func (c *MySQLCollector) collectTables(info *DatabaseInfo) error {
	// Get all databases accessible to the user
	databases, err := c.getDatabases()
	if err != nil {
		return err
	}

	for _, dbName := range databases {
		tables, err := c.getTablesForDatabase(dbName)
		if err != nil {
			continue // Skip databases we can't access
		}
		info.Tables = append(info.Tables, tables...)
	}
	return nil
}

// getDatabases returns list of accessible databases
func (c *MySQLCollector) getDatabases() ([]string, error) {
	var databases []string

	// If specific database is requested, use only that
	if c.config.Database != "" {
		databases = append(databases, c.config.Database)
		return databases, nil
	}

	// Otherwise get all accessible databases
	rows, err := c.db.Query("SHOW DATABASES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dbName string
		if err := rows.Scan(&dbName); err != nil {
			continue
		}
		// Skip system databases unless specifically requested
		if dbName != "information_schema" && dbName != "performance_schema" &&
			dbName != "mysql" && dbName != "sys" {
			databases = append(databases, dbName)
		}
	}

	return databases, nil
}

// getTablesForDatabase collects tables and views for a specific database
func (c *MySQLCollector) getTablesForDatabase(dbName string) ([]TableInfo, error) {
	var tables []TableInfo

	query := `
		SELECT TABLE_NAME, TABLE_TYPE, ENGINE, AUTO_INCREMENT, CREATE_TIME, UPDATE_TIME, 
		       TABLE_COLLATION, COALESCE(SUBSTRING_INDEX(TABLE_COLLATION, '_', 1), '') as TABLE_CHARSET,
		       ROW_FORMAT, TABLE_COMMENT, CREATE_OPTIONS
		FROM information_schema.TABLES 
		WHERE TABLE_SCHEMA = ? 
		ORDER BY TABLE_NAME`

	rows, err := c.db.Query(query, dbName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var table TableInfo
		var autoIncrement sql.NullInt64
		var createTime, updateTime sql.NullString
		var engine, collation, charset, rowFormat, comment, createOptions sql.NullString

		err := rows.Scan(&table.Name, &table.Type, &engine, &autoIncrement,
			&createTime, &updateTime, &collation, &charset, &rowFormat, &comment, &createOptions)
		if err != nil {
			continue
		}

		table.Database = dbName
		table.Schema = dbName
		if engine.Valid {
			table.Engine = engine.String
		}
		if autoIncrement.Valid {
			table.AutoIncrement = autoIncrement.Int64
		}
		if createTime.Valid && createTime.String != "" {
			if t, err := time.Parse("2006-01-02 15:04:05", createTime.String); err == nil {
				table.CreatedAt = t
			}
		}
		if updateTime.Valid && updateTime.String != "" {
			if t, err := time.Parse("2006-01-02 15:04:05", updateTime.String); err == nil {
				table.UpdatedAt = t
			}
		}
		if collation.Valid {
			table.Collation = collation.String
		}
		if charset.Valid {
			table.Charset = charset.String
		}
		if rowFormat.Valid {
			table.RowFormat = rowFormat.String
		}
		if comment.Valid {
			table.Comment = comment.String
		}
		if createOptions.Valid {
			table.CreateOptions = createOptions.String
		}

		// Get DDL
		ddl, err := c.getTableDDL(dbName, table.Name, table.Type)
		if err == nil {
			table.DDL = ddl
		}

		tables = append(tables, table)
	}
	return tables, nil
}

// getTableDDL gets the CREATE statement for a table or view
func (c *MySQLCollector) getTableDDL(schema, name, tableType string) (string, error) {
	var query string
	if tableType == "VIEW" {
		query = fmt.Sprintf("SHOW CREATE VIEW `%s`.`%s`", schema, name)
		// SHOW CREATE VIEW returns 4 columns: View, Create View, character_set_client, collation_connection
		var viewName, ddl, charsetClient, collationConnection string
		err := c.db.QueryRow(query).Scan(&viewName, &ddl, &charsetClient, &collationConnection)
		if err != nil {
			return "", err
		}
		return ddl, nil
	} else {
		query = fmt.Sprintf("SHOW CREATE TABLE `%s`.`%s`", schema, name)
		// SHOW CREATE TABLE returns 2 columns: Table, Create Table
		var tableName, ddl string
		err := c.db.QueryRow(query).Scan(&tableName, &ddl)
		if err != nil {
			return "", err
		}
		return ddl, nil
	}
}

// collectUsers collects user account information
func (c *MySQLCollector) collectUsers(info *DatabaseInfo) error {
	var query string
	if c.version.IsMySQL8OrLater() {
		query = `
			SELECT User, Host, plugin, account_locked, password_expired
			FROM mysql.user 
			ORDER BY User, Host`
	} else {
		query = `
			SELECT User, Host, plugin, account_locked, 'N' as password_expired
			FROM mysql.user 
			ORDER BY User, Host`
	}

	rows, err := c.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var user UserAccount
		var plugin, locked, expired sql.NullString

		err := rows.Scan(&user.User, &user.Host, &plugin, &locked, &expired)
		if err != nil {
			continue
		}

		if plugin.Valid {
			user.Plugin = plugin.String
		}
		if locked.Valid {
			user.AccountLocked = locked.String
		}
		if expired.Valid {
			user.PasswordExpired = expired.String
		}

		// Get grants for this user
		grants, err := c.getUserGrants(user.User, user.Host)
		if err == nil {
			user.Grants = grants
		}

		info.Users = append(info.Users, user)
	}

	return nil
}

// getUserGrants gets grants for a specific user
func (c *MySQLCollector) getUserGrants(user, host string) ([]string, error) {
	var grants []string

	query := fmt.Sprintf("SHOW GRANTS FOR `%s`@`%s`", user, host)
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var grant string
		if err := rows.Scan(&grant); err != nil {
			continue
		}
		grants = append(grants, grant)
	}

	return grants, nil
}

// collectRoutines collects all stored procedures and functions
func (c *MySQLCollector) collectRoutines(info *DatabaseInfo) error {
	// Collect both functions and procedures
	query := `
		SELECT ROUTINE_NAME, ROUTINE_TYPE, ROUTINE_SCHEMA, DEFINER, CREATED, LAST_ALTERED,
		       SQL_DATA_ACCESS, SECURITY_TYPE, ROUTINE_DEFINITION, DTD_IDENTIFIER
		FROM information_schema.ROUTINES
		WHERE ROUTINE_SCHEMA NOT IN ('information_schema', 'performance_schema', 'mysql', 'sys')
		ORDER BY ROUTINE_SCHEMA, ROUTINE_TYPE, ROUTINE_NAME`

	rows, err := c.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var routine RoutineInfo
		var definer, dataAccess, securityType, definition, returns sql.NullString
		var created, lastAltered sql.NullString

		err := rows.Scan(&routine.Name, &routine.Type, &routine.Schema, &definer,
			&created, &lastAltered, &dataAccess, &securityType,
			&definition, &returns)
		if err != nil {
			continue
		}

		if definer.Valid {
			routine.Definer = definer.String
		}
		if created.Valid && created.String != "" {
			if t, err := time.Parse("2006-01-02 15:04:05", created.String); err == nil {
				routine.Created = t
			}
		}
		if lastAltered.Valid && lastAltered.String != "" {
			if t, err := time.Parse("2006-01-02 15:04:05", lastAltered.String); err == nil {
				routine.LastAltered = t
			}
		}
		if dataAccess.Valid {
			routine.DataAccess = dataAccess.String
		}
		if securityType.Valid {
			routine.SecurityType = securityType.String
		}
		if definition.Valid {
			routine.Definition = definition.String
		}
		if returns.Valid {
			routine.Returns = returns.String
		}

		// Get parameters
		params, err := c.getRoutineParameters(routine.Schema, routine.Name)
		if err == nil {
			routine.Parameters = params
		}

		info.Routines = append(info.Routines, routine)
	}

	return nil
}

// getRoutineParameters gets parameters for a routine
func (c *MySQLCollector) getRoutineParameters(schema, name string) (string, error) {
	query := `
		SELECT GROUP_CONCAT(
			CONCAT(PARAMETER_MODE, ' ', PARAMETER_NAME, ' ', DTD_IDENTIFIER)
			ORDER BY ORDINAL_POSITION
			SEPARATOR ','
		) as parameters
		FROM information_schema.PARAMETERS
		WHERE SPECIFIC_SCHEMA = ? AND SPECIFIC_NAME = ?`

	var params sql.NullString
	err := c.db.QueryRow(query, schema, name).Scan(&params)
	if err != nil {
		return "", err
	}

	if params.Valid {
		return params.String, nil
	}
	return "", nil
}

// collectVariables collects all global variables
func (c *MySQLCollector) collectVariables(info *DatabaseInfo) error {
	// Try performance_schema first for MySQL 5.7+
	if c.version.SupportsPerformanceSchemaVariablesInfo() {
		return c.collectVariablesFromPerformanceSchema(info)
	}
	// Fallback for older versions
	return c.collectVariablesFromInformationSchema(info)
}

// collectVariablesFromPerformanceSchema collects variables from performance_schema (MySQL 5.7+)
func (c *MySQLCollector) collectVariablesFromPerformanceSchema(info *DatabaseInfo) error {
	var query string
	
	if c.config.OnlyModifiedVariables {
		// Only get modified variables
		query = `
			SELECT vi.VARIABLE_NAME, gv.VARIABLE_VALUE, vi.VARIABLE_SOURCE, vi.DEFAULT_VALUE
			FROM performance_schema.variables_info vi
			JOIN performance_schema.global_variables gv ON vi.VARIABLE_NAME = gv.VARIABLE_NAME
			WHERE vi.VARIABLE_SOURCE != 'COMPILED'
			ORDER BY vi.VARIABLE_NAME`
	} else {
		// Get all variables
		query = `
			SELECT vi.VARIABLE_NAME, gv.VARIABLE_VALUE, 
			       COALESCE(vi.VARIABLE_SOURCE, 'COMPILED') as VARIABLE_SOURCE, 
			       COALESCE(vi.DEFAULT_VALUE, '') as DEFAULT_VALUE
			FROM performance_schema.global_variables gv
			LEFT JOIN performance_schema.variables_info vi ON vi.VARIABLE_NAME = gv.VARIABLE_NAME
			ORDER BY gv.VARIABLE_NAME`
	}

	rows, err := c.db.Query(query)
	if err != nil {
		// Fallback to information_schema
		return c.collectVariablesFromInformationSchema(info)
	}
	defer rows.Close()

	for rows.Next() {
		var variable Variable
		var source, defaultValue sql.NullString

		err := rows.Scan(&variable.Name, &variable.CurrentValue, &source, &defaultValue)
		if err != nil {
			continue
		}

		if source.Valid {
			variable.Source = source.String
			variable.IsModified = (source.String != "COMPILED")
		}
		if defaultValue.Valid {
			variable.DefaultValue = defaultValue.String
		}

		info.Variables = append(info.Variables, variable)
	}

	return nil
}

// collectVariablesFromInformationSchema collects variables from information_schema (older MySQL versions)
func (c *MySQLCollector) collectVariablesFromInformationSchema(info *DatabaseInfo) error {
	var query string
	
	if c.config.OnlyModifiedVariables {
		// For older versions, collect commonly modified variables
		return c.collectCommonModifiedVariables(info)
	}
	
	// Get all variables from performance_schema
	query = `SELECT VARIABLE_NAME, VARIABLE_VALUE 
			 FROM performance_schema.global_variables 
			 ORDER BY VARIABLE_NAME`
	
	rows, err := c.db.Query(query)
	if err != nil {
		// If performance_schema doesn't exist, try SHOW VARIABLES
		return c.collectVariablesFromShowVariables(info)
	}
	defer rows.Close()

	for rows.Next() {
		var variable Variable
		err := rows.Scan(&variable.Name, &variable.CurrentValue)
		if err != nil {
			continue
		}
		
		// For older versions, we can't easily determine if it's modified
		variable.Source = "UNKNOWN"
		variable.IsModified = false
		
		info.Variables = append(info.Variables, variable)
	}

	return nil
}

// collectCommonModifiedVariables collects commonly modified variables for older MySQL versions
func (c *MySQLCollector) collectCommonModifiedVariables(info *DatabaseInfo) error {
	// Collect some commonly modified variables with known defaults
	commonVars := map[string]string{
		"character_set_server":          "latin1",
		"collation_server":              "latin1_swedish_ci",
		"max_connections":               "151",
		"innodb_buffer_pool_size":       "134217728",
		"datadir":                       "/var/lib/mysql/",
		"socket":                        "/var/run/mysqld/mysqld.sock",
		"pid_file":                      "/var/run/mysqld/mysqld.pid",
		"secure_file_priv":              "",
		"skip_name_resolve":             "OFF",
		"default_authentication_plugin": "mysql_native_password",
	}

	for varName, defaultValue := range commonVars {
		var value string
		// Try performance_schema first, then fallback to information_schema
		query := "SELECT VARIABLE_VALUE FROM performance_schema.global_variables WHERE VARIABLE_NAME = ?"
		err := c.db.QueryRow(query, varName).Scan(&value)
		if err != nil {
			query = "SELECT VARIABLE_VALUE FROM performance_schema.global_variables WHERE VARIABLE_NAME = ?"
			err = c.db.QueryRow(query, varName).Scan(&value)
			if err != nil {
				continue
			}
		}

		// Only add if value differs from default
		if value != defaultValue {
			variable := Variable{
				Name:         varName,
				CurrentValue: value,
				DefaultValue: defaultValue,
				Source:       "DYNAMIC",
				IsModified:   true,
			}
			info.Variables = append(info.Variables, variable)
		}
	}

	return nil
}

// collectMySQL8Features collects MySQL 8.0+ specific features
func (c *MySQLCollector) collectMySQL8Features(info *DatabaseInfo) error {
	// Collect roles only
	if err := c.collectRoles(info); err != nil {
		// Don't fail completely if roles collection fails
		fmt.Printf("Warning: Failed to collect roles: %v\n", err)
	}

	return nil
}

// collectRoles collects MySQL 8.0 roles
func (c *MySQLCollector) collectRoles(info *DatabaseInfo) error {
	// Get all roles
	query := `
		SELECT User as role_name, Host as role_host
		FROM mysql.user 
		WHERE account_locked = 'Y' AND password_expired != 'Y'
		ORDER BY User, Host`

	rows, err := c.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var role UserRole
		err := rows.Scan(&role.RoleName, &role.RoleHost)
		if err != nil {
			continue
		}

		// Get role grants
		grants, err := c.getUserGrants(role.RoleName, role.RoleHost)
		if err == nil {
			role.Grants = grants
		}

		// Get role members (users who have this role)
		members, err := c.getRoleMembers(role.RoleName, role.RoleHost)
		if err == nil {
			role.Members = members
		}

		info.Roles = append(info.Roles, role)
	}

	return nil
}

// getRoleMembers gets users who have a specific role
func (c *MySQLCollector) getRoleMembers(roleName, roleHost string) ([]string, error) {
	var members []string

	query := `
		SELECT CONCAT(TO_USER, '@', TO_HOST) as member
		FROM mysql.role_edges 
		WHERE FROM_USER = ? AND FROM_HOST = ?`

	rows, err := c.db.Query(query, roleName, roleHost)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var member string
		if err := rows.Scan(&member); err != nil {
			continue
		}
		members = append(members, member)
	}

	return members, nil
}

// collectComponents collects MySQL 8.0 components
func (c *MySQLCollector) collectComponents(info *DatabaseInfo) error {
	query := "SELECT * FROM mysql.component"
	rows, err := c.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var component Component
		err := rows.Scan(&component.ComponentID, &component.ComponentGroupID, &component.ComponentURN)
		if err != nil {
			continue
		}
		info.Components = append(info.Components, component)
	}

	return nil
}

// collectPlugins collects plugin information
func (c *MySQLCollector) collectPlugins(info *DatabaseInfo) error {
	query := `
		SELECT PLUGIN_NAME, PLUGIN_VERSION, PLUGIN_STATUS, PLUGIN_TYPE, 
		       PLUGIN_LIBRARY, PLUGIN_DESCRIPTION
		FROM information_schema.PLUGINS
		ORDER BY PLUGIN_TYPE, PLUGIN_NAME`

	rows, err := c.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Default plugins to exclude (built-in MySQL plugins)
	excludePlugins := map[string]bool{
		"mysqlx_cache_cleaner":        true,
		"sha2_cache_cleaner":          true,
		"caching_sha2_password":       true,
		"mysql_native_password":       true,
		"sha256_password":             true,
		"mysqlx":                      true,
		"ngram":                       true,
		"ARCHIVE":                     true,
		"binlog":                      true,
		"BLACKHOLE":                   true,
		"CSV":                         true,
		"FEDERATED":                   true,
		"InnoDB":                      true,
		"MEMORY":                      true,
		"MRG_MYISAM":                  true,
		"MyISAM":                      true,
		"ndbcluster":                  true,
		"ndbinfo":                     true,
		"PERFORMANCE_SCHEMA":          true,
		"TempTable":                   true,
		"daemon_keyring_proxy_plugin": true,
		// InnoDB Information Schema plugins
		"INNODB_BUFFER_PAGE":               true,
		"INNODB_BUFFER_PAGE_LRU":           true,
		"INNODB_BUFFER_POOL_STATS":         true,
		"INNODB_CACHED_INDEXES":            true,
		"INNODB_CMP":                       true,
		"INNODB_CMPMEM":                    true,
		"INNODB_CMPMEM_RESET":              true,
		"INNODB_CMP_PER_INDEX":             true,
		"INNODB_CMP_PER_INDEX_RESET":       true,
		"INNODB_CMP_RESET":                 true,
		"INNODB_COLUMNS":                   true,
		"INNODB_FT_BEING_DELETED":          true,
		"INNODB_FT_CONFIG":                 true,
		"INNODB_FT_DEFAULT_STOPWORD":       true,
		"INNODB_FT_DELETED":                true,
		"INNODB_FT_INDEX_CACHE":            true,
		"INNODB_FT_INDEX_TABLE":            true,
		"INNODB_INDEXES":                   true,
		"INNODB_METRICS":                   true,
		"INNODB_SESSION_TEMP_TABLESPACES":  true,
		"INNODB_TABLES":                    true,
		"INNODB_TABLESPACES":               true,
		"INNODB_TABLESTATS":                true,
		"INNODB_TEMP_TABLE_INFO":           true,
		"INNODB_TRX":                       true,
		"INNODB_VIRTUAL":                   true,
		"ndb_transid_mysql_connection_map": true,
	}

	for rows.Next() {
		var plugin Plugin
		var library, description sql.NullString

		err := rows.Scan(&plugin.Name, &plugin.Version, &plugin.Status,
			&plugin.Type, &library, &description)
		if err != nil {
			continue
		}

		// Skip default plugins
		if excludePlugins[plugin.Name] {
			continue
		}

		if library.Valid {
			plugin.Library = library.String
		}
		if description.Valid {
			plugin.Description = description.String
		}

		info.Plugins = append(info.Plugins, plugin)
	}

	return nil
}

// collectReplicationInfo collects replication information
func (c *MySQLCollector) collectReplicationInfo(info *DatabaseInfo) error {
	replicationInfo := &ReplicationInfo{}

	// Get basic replication status
	if err := c.getReplicationStatus(replicationInfo); err == nil {
		info.ReplicationInfo = replicationInfo
	}

	// Get replica status if available
	if err := c.getReplicaStatus(replicationInfo); err == nil {
		info.ReplicationInfo = replicationInfo
	}

	// Get semi-sync status
	if err := c.getSemiSyncStatus(replicationInfo); err == nil {
		info.ReplicationInfo = replicationInfo
	}

	// Get group replication info
	if err := c.getGroupReplicationInfo(replicationInfo); err == nil {
		info.ReplicationInfo = replicationInfo
	}

	return nil
}

// getReplicationStatus gets basic replication status
func (c *MySQLCollector) getReplicationStatus(info *ReplicationInfo) error {
	status := &ReplicationStatus{}

	// Get server ID
	var serverID int
	err := c.db.QueryRow("SELECT @@server_id").Scan(&serverID)
	if err == nil {
		status.ServerID = serverID
	}

	// Get server UUID
	var serverUUID string
	err = c.db.QueryRow("SELECT @@server_uuid").Scan(&serverUUID)
	if err == nil {
		status.ServerUUID = serverUUID
	}

	// Check if binary logging is enabled
	var logBin string
	err = c.db.QueryRow("SELECT @@log_bin").Scan(&logBin)
	if err == nil {
		status.LogBinEnabled = (logBin == "1" || strings.ToUpper(logBin) == "ON")
	}

	// Get binary log format
	var binlogFormat string
	err = c.db.QueryRow("SELECT @@binlog_format").Scan(&binlogFormat)
	if err == nil {
		status.BinlogFormat = binlogFormat
	}

	// Get GTID mode
	var gtidMode string
	err = c.db.QueryRow("SELECT @@gtid_mode").Scan(&gtidMode)
	if err == nil {
		status.GTIDMode = gtidMode
	}

	// Get current binary log file and position
	err = c.db.QueryRow("SHOW MASTER STATUS").Scan(&status.CurrentLogFile, &status.CurrentLogPos, nil, nil)
	// Don't fail if this doesn't work (might not be a master)

	info.ReplicationStatus = status
	return nil
}

// getReplicaStatus gets replica status information
func (c *MySQLCollector) getReplicaStatus(info *ReplicationInfo) error {
	// Try to get replica status
	rows, err := c.db.Query("SHOW REPLICA STATUS")
	if err != nil {
		// Try older syntax
		rows, err = c.db.Query("SHOW SLAVE STATUS")
		if err != nil {
			return err
		}
	}
	defer rows.Close()

	// Process replica status if available
	// This is complex as SHOW REPLICA STATUS returns many columns
	// For now, just check if we have any replica status
	if rows.Next() {
		// If we get here, this server is a replica
		// In a full implementation, you'd parse all the columns
		info.ReplicaStatus = []ReplicaStatus{{
			// Simplified replica status
			SlaveIORunning:  "Unknown", // Would need to parse actual columns
			SlaveSQLRunning: "Unknown",
		}}
	}

	return nil
}

// getSemiSyncStatus gets semi-synchronous replication status
func (c *MySQLCollector) getSemiSyncStatus(info *ReplicationInfo) error {
	status := &SemiSyncStatus{}

	// Check if semi-sync master is enabled
	var masterEnabled string
	err := c.db.QueryRow("SELECT @@rpl_semi_sync_master_enabled").Scan(&masterEnabled)
	if err == nil {
		status.MasterEnabled = (masterEnabled == "1" || strings.ToUpper(masterEnabled) == "ON")
	}

	// Check if semi-sync slave is enabled
	var slaveEnabled string
	err = c.db.QueryRow("SELECT @@rpl_semi_sync_slave_enabled").Scan(&slaveEnabled)
	if err == nil {
		status.SlaveEnabled = (slaveEnabled == "1" || strings.ToUpper(slaveEnabled) == "ON")
	}

	info.SemiSyncStatus = status
	return nil
}

// getGroupReplicationInfo gets group replication information
func (c *MySQLCollector) getGroupReplicationInfo(info *ReplicationInfo) error {
	groupInfo := &GroupReplicationInfo{}

	// Check if group replication is started
	var memberState string
	err := c.db.QueryRow("SELECT MEMBER_STATE FROM performance_schema.replication_group_members WHERE MEMBER_ID = @@server_uuid").Scan(&memberState)
	if err == nil {
		groupInfo.MemberState = memberState
	}

	// Get group name
	var groupName string
	err = c.db.QueryRow("SELECT @@group_replication_group_name").Scan(&groupName)
	if err == nil {
		groupInfo.GroupName = groupName
	}

	// Get single primary mode
	var singlePrimary string
	err = c.db.QueryRow("SELECT @@group_replication_single_primary_mode").Scan(&singlePrimary)
	if err == nil {
		groupInfo.SinglePrimaryMode = (singlePrimary == "1" || strings.ToUpper(singlePrimary) == "ON")
	}

	info.GroupReplicationInfo = groupInfo
	return nil
}

// collectVariablesFromShowVariables collects variables using SHOW VARIABLES (fallback for very old MySQL)
func (c *MySQLCollector) collectVariablesFromShowVariables(info *DatabaseInfo) error {
	if c.config.OnlyModifiedVariables {
		// For very old versions, collect commonly modified variables
		return c.collectCommonModifiedVariables(info)
	}
	
	// Use SHOW VARIABLES as last resort
	rows, err := c.db.Query("SHOW VARIABLES")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var variable Variable
		err := rows.Scan(&variable.Name, &variable.CurrentValue)
		if err != nil {
			continue
		}
		
		// For very old versions, we can't determine if it's modified
		variable.Source = "UNKNOWN"
		variable.IsModified = false
		
		info.Variables = append(info.Variables, variable)
	}

	return nil
}
