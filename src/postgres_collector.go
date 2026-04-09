package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

// PostgreSQLCollector handles PostgreSQL information collection
type PostgreSQLCollector struct {
	db      *sql.DB
	version *PostgreSQLVersion
	config  *Config
}

// NewPostgreSQLCollector creates a new PostgreSQL collector
func NewPostgreSQLCollector(db *sql.DB, config *Config) (*PostgreSQLCollector, error) {
	collector := &PostgreSQLCollector{
		db:     db,
		config: config,
	}

	version, err := DetectPostgreSQLVersion(db)
	if err != nil {
		return nil, fmt.Errorf("failed to detect PostgreSQL version: %v", err)
	}
	collector.version = version

	return collector, nil
}

// CollectAll collects all PostgreSQL information
func (c *PostgreSQLCollector) CollectAll() (*DatabaseInfo, error) {
	info := &DatabaseInfo{
		DBType: "postgres",
	}

	if err := c.collectConnectionInfo(info); err != nil {
		return nil, fmt.Errorf("failed to collect connection info: %v", err)
	}

	if !c.config.ExceptTables {
		if err := c.collectTables(info); err != nil {
			return nil, fmt.Errorf("failed to collect tables: %v", err)
		}
	}

	if !c.config.ExceptUsers {
		if err := c.collectAllRoles(info); err != nil {
			log.Printf("Warning: failed to collect users/roles: %v", err)
		}
	}

	if !c.config.ExceptStoredProcedures {
		if err := c.collectRoutines(info); err != nil {
			log.Printf("Warning: failed to collect routines: %v", err)
		}
	}

	if !c.config.ExceptVariables {
		if err := c.collectVariables(info); err != nil {
			log.Printf("Warning: failed to collect variables: %v", err)
		}
	}

	if !c.config.ExceptExtensions {
		if err := c.collectExtensions(info); err != nil {
			log.Printf("Warning: failed to collect extensions: %v", err)
		}
	}

	return info, nil
}

func (c *PostgreSQLCollector) collectConnectionInfo(info *DatabaseInfo) error {
	var version string
	err := c.db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		return err
	}

	info.ConnectionInfo = &ConnectionInfo{
		Host:     c.config.Host,
		Port:     c.config.Port,
		User:     c.config.User,
		Database: c.config.Database,
		Version:  version,
	}

	return nil
}

// getSchemas returns list of user schemas (excluding system schemas)
func (c *PostgreSQLCollector) getSchemas() ([]string, error) {
	query := `
		SELECT schema_name
		FROM information_schema.schemata
		WHERE schema_name NOT IN ('pg_catalog', 'information_schema', 'pg_toast')
		  AND schema_name NOT LIKE 'pg_temp_%'
		  AND schema_name NOT LIKE 'pg_toast_temp_%'
		ORDER BY schema_name`

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schemas []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			continue
		}
		schemas = append(schemas, name)
	}
	return schemas, nil
}

func (c *PostgreSQLCollector) collectTables(info *DatabaseInfo) error {
	schemas, err := c.getSchemas()
	if err != nil {
		return err
	}

	for _, schema := range schemas {
		tables, err := c.getTablesForSchema(schema)
		if err != nil {
			continue
		}
		info.Tables = append(info.Tables, tables...)
	}
	return nil
}

func (c *PostgreSQLCollector) getTablesForSchema(schema string) ([]TableInfo, error) {
	query := `
		SELECT t.table_name, t.table_type
		FROM information_schema.tables t
		WHERE t.table_schema = $1
		ORDER BY t.table_name`

	rows, err := c.db.Query(query, schema)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []TableInfo
	for rows.Next() {
		var table TableInfo
		var tableType string
		if err := rows.Scan(&table.Name, &tableType); err != nil {
			continue
		}

		table.Schema = schema
		table.Database = c.config.Database

		// Normalize type to match MySQL convention
		switch tableType {
		case "BASE TABLE":
			table.Type = "BASE TABLE"
		case "VIEW":
			table.Type = "VIEW"
		default:
			table.Type = tableType
		}

		// Get table metadata
		if table.Type == "BASE TABLE" {
			c.getTableMetadata(&table)
		}

		// Get DDL
		ddl, err := c.getTableDDL(schema, table.Name, table.Type)
		if err == nil {
			table.DDL = ddl
		}

		tables = append(tables, table)
	}
	return tables, nil
}

func (c *PostgreSQLCollector) getTableMetadata(table *TableInfo) {
	// Get row count estimate and table size info from pg_class
	query := `
		SELECT
			COALESCE(pg_size_pretty(pg_total_relation_size(c.oid)), '') as total_size,
			obj_description(c.oid, 'pg_class') as comment
		FROM pg_catalog.pg_class c
		JOIN pg_catalog.pg_namespace n ON n.oid = c.relnamespace
		WHERE n.nspname = $1 AND c.relname = $2`

	var totalSize sql.NullString
	var comment sql.NullString
	err := c.db.QueryRow(query, table.Schema, table.Name).Scan(&totalSize, &comment)
	if err == nil {
		if comment.Valid {
			table.Comment = comment.String
		}
		if totalSize.Valid {
			table.RowFormat = totalSize.String // Reuse RowFormat for size display
		}
	}
}

func (c *PostgreSQLCollector) getTableDDL(schema, name, tableType string) (string, error) {
	if tableType == "VIEW" {
		return c.getViewDDL(schema, name)
	}
	return c.buildTableDDL(schema, name)
}

func (c *PostgreSQLCollector) getViewDDL(schema, name string) (string, error) {
	var definition string
	err := c.db.QueryRow("SELECT pg_get_viewdef($1||'.'||$2, true)", schema, name).Scan(&definition)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("CREATE VIEW %s.%s AS\n%s", schema, name, definition), nil
}

func (c *PostgreSQLCollector) buildTableDDL(schema, name string) (string, error) {
	// Get columns
	colQuery := `
		SELECT column_name, data_type, character_maximum_length,
		       column_default, is_nullable, udt_name,
		       numeric_precision, numeric_scale
		FROM information_schema.columns
		WHERE table_schema = $1 AND table_name = $2
		ORDER BY ordinal_position`

	rows, err := c.db.Query(colQuery, schema, name)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var columns []string
	for rows.Next() {
		var colName, dataType, isNullable, udtName string
		var charMaxLen, numPrecision, numScale sql.NullInt64
		var colDefault sql.NullString

		if err := rows.Scan(&colName, &dataType, &charMaxLen, &colDefault, &isNullable, &udtName, &numPrecision, &numScale); err != nil {
			continue
		}

		col := fmt.Sprintf("    %s", colName)

		// Determine type display
		switch dataType {
		case "character varying":
			if charMaxLen.Valid {
				col += fmt.Sprintf(" varchar(%d)", charMaxLen.Int64)
			} else {
				col += " varchar"
			}
		case "character":
			if charMaxLen.Valid {
				col += fmt.Sprintf(" char(%d)", charMaxLen.Int64)
			} else {
				col += " char"
			}
		case "numeric":
			if numPrecision.Valid && numScale.Valid {
				col += fmt.Sprintf(" numeric(%d,%d)", numPrecision.Int64, numScale.Int64)
			} else if numPrecision.Valid {
				col += fmt.Sprintf(" numeric(%d)", numPrecision.Int64)
			} else {
				col += " numeric"
			}
		case "ARRAY":
			col += " " + udtName
		case "USER-DEFINED":
			col += " " + udtName
		default:
			col += " " + dataType
		}

		if isNullable == "NO" {
			col += " NOT NULL"
		}
		if colDefault.Valid {
			col += fmt.Sprintf(" DEFAULT %s", colDefault.String)
		}

		columns = append(columns, col)
	}

	// Get constraints
	constraintQuery := `
		SELECT
			conname,
			pg_get_constraintdef(c.oid, true) as condef
		FROM pg_catalog.pg_constraint c
		JOIN pg_catalog.pg_namespace n ON n.oid = c.connamespace
		JOIN pg_catalog.pg_class cl ON cl.oid = c.conrelid
		WHERE n.nspname = $1 AND cl.relname = $2
		ORDER BY contype, conname`

	cRows, err := c.db.Query(constraintQuery, schema, name)
	if err == nil {
		defer cRows.Close()
		for cRows.Next() {
			var conName, conDef string
			if err := cRows.Scan(&conName, &conDef); err != nil {
				continue
			}
			columns = append(columns, fmt.Sprintf("    CONSTRAINT %s %s", conName, conDef))
		}
	}

	ddl := fmt.Sprintf("CREATE TABLE %s.%s (\n%s\n);", schema, name, strings.Join(columns, ",\n"))

	// Get indexes (non-constraint)
	idxQuery := `
		SELECT indexdef
		FROM pg_indexes
		WHERE schemaname = $1 AND tablename = $2
		  AND indexname NOT IN (
			SELECT conname FROM pg_catalog.pg_constraint
			WHERE connamespace = (SELECT oid FROM pg_namespace WHERE nspname = $1)
			  AND conrelid = (SELECT oid FROM pg_class WHERE relname = $2 AND relnamespace = (SELECT oid FROM pg_namespace WHERE nspname = $1))
		  )`

	idxRows, err := c.db.Query(idxQuery, schema, name)
	if err == nil {
		defer idxRows.Close()
		for idxRows.Next() {
			var indexDef string
			if err := idxRows.Scan(&indexDef); err != nil {
				continue
			}
			ddl += "\n" + indexDef + ";"
		}
	}

	return ddl, nil
}

// collectAllRoles collects all PostgreSQL roles (both login and non-login) into a unified Users list
func (c *PostgreSQLCollector) collectAllRoles(info *DatabaseInfo) error {
	// Collect all roles (login users first, then non-login roles)
	query := `
		SELECT rolname, rolsuper, rolcreaterole, rolcreatedb,
		       rolcanlogin, rolreplication, rolinherit, rolconnlimit, rolvaliduntil
		FROM pg_catalog.pg_roles
		WHERE rolname NOT LIKE 'pg_%'
		ORDER BY rolcanlogin DESC, rolname`

	rows, err := c.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var user UserAccount
		var rolSuper, rolCreateRole, rolCreateDB, rolCanLogin, rolReplication, rolInherit bool
		var rolConnLimit int
		var rolValidUntil sql.NullString

		if err := rows.Scan(&user.User, &rolSuper, &rolCreateRole, &rolCreateDB,
			&rolCanLogin, &rolReplication, &rolInherit, &rolConnLimit, &rolValidUntil); err != nil {
			continue
		}

		user.Host = ""
		user.IsSuperuser = rolSuper
		user.CanCreateDB = rolCreateDB
		user.CanCreateRole = rolCreateRole
		user.CanLogin = rolCanLogin
		user.Inherit = rolInherit
		user.ConnLimit = rolConnLimit
		user.IsRole = !rolCanLogin
		if rolValidUntil.Valid {
			user.ValidUntil = rolValidUntil.String
		}

		// Build attribute summary
		var attrs []string
		if rolSuper {
			attrs = append(attrs, "SUPERUSER")
		}
		if rolCreateDB {
			attrs = append(attrs, "CREATEDB")
		}
		if rolCreateRole {
			attrs = append(attrs, "CREATEROLE")
		}
		if rolReplication {
			attrs = append(attrs, "REPLICATION")
		}
		if !rolInherit {
			attrs = append(attrs, "NOINHERIT")
		}
		if !rolCanLogin {
			attrs = append(attrs, "NOLOGIN")
		}
		if rolConnLimit >= 0 && rolCanLogin {
			attrs = append(attrs, fmt.Sprintf("CONNECTION LIMIT %d", rolConnLimit))
		}
		if len(attrs) > 0 {
			user.Plugin = strings.Join(attrs, ", ")
		}

		// Get grants (database-level privileges)
		grants, err := c.getUserGrants(user.User)
		if err == nil {
			user.Grants = grants
		}

		info.Users = append(info.Users, user)
	}

	// Collect role memberships and populate AssignedRoles/GrantedTo
	if err := c.collectRoleMemberships(info); err != nil {
		log.Printf("Warning: failed to collect role memberships: %v", err)
	}

	return nil
}

func (c *PostgreSQLCollector) getUserGrants(rolname string) ([]string, error) {
	var grants []string

	// Database-level privileges
	dbGrantQuery := `
		SELECT datname, datacl
		FROM pg_catalog.pg_database
		WHERE datacl IS NOT NULL AND datacl::text LIKE '%' || $1 || '%'`

	dbRows, err := c.db.Query(dbGrantQuery, rolname)
	if err == nil {
		defer dbRows.Close()
		for dbRows.Next() {
			var datname, datacl string
			if err := dbRows.Scan(&datname, &datacl); err != nil {
				continue
			}
			grants = append(grants, fmt.Sprintf("DATABASE %s: %s", datname, datacl))
		}
	}

	return grants, nil
}

// collectRoleMemberships collects pg_auth_members and populates AssignedRoles/GrantedTo
func (c *PostgreSQLCollector) collectRoleMemberships(info *DatabaseInfo) error {
	query := `
		SELECT r.rolname AS role_name,
		       m.rolname AS member_name,
		       am.admin_option
		FROM pg_catalog.pg_auth_members am
		JOIN pg_catalog.pg_roles r ON am.roleid = r.oid
		JOIN pg_catalog.pg_roles m ON am.member = m.oid
		ORDER BY r.rolname, m.rolname`

	rows, err := c.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Build lookup
	userIdx := make(map[string]int)
	for i := range info.Users {
		userIdx[info.Users[i].User] = i
	}

	for rows.Next() {
		var roleName, memberName string
		var adminOption bool
		if err := rows.Scan(&roleName, &memberName, &adminOption); err != nil {
			continue
		}

		edge := RoleEdge{
			FromRole:  roleName,
			ToUser:    memberName,
			WithAdmin: adminOption,
		}

		// Add to the member's AssignedRoles (member has been granted this role)
		if idx, ok := userIdx[memberName]; ok {
			info.Users[idx].AssignedRoles = append(info.Users[idx].AssignedRoles, edge)
		}

		// Add to the role's GrantedTo (this role has been granted to the member)
		if idx, ok := userIdx[roleName]; ok {
			info.Users[idx].GrantedTo = append(info.Users[idx].GrantedTo, edge)
		}
	}

	return nil
}

func (c *PostgreSQLCollector) collectRoutines(info *DatabaseInfo) error {
	query := `
		SELECT n.nspname as schema,
		       p.proname as name,
		       CASE p.prokind
		           WHEN 'f' THEN 'FUNCTION'
		           WHEN 'p' THEN 'PROCEDURE'
		           WHEN 'a' THEN 'AGGREGATE'
		           WHEN 'w' THEN 'WINDOW'
		           ELSE 'FUNCTION'
		       END as type,
		       pg_catalog.pg_get_userbyid(p.proowner) as definer,
		       pg_catalog.pg_get_functiondef(p.oid) as definition,
		       COALESCE(pg_catalog.pg_get_function_result(p.oid), '') as returns,
		       COALESCE(pg_catalog.pg_get_function_arguments(p.oid), '') as arguments,
		       CASE p.prosecdef WHEN true THEN 'DEFINER' ELSE 'INVOKER' END as security_type,
		       CASE p.provolatile
		           WHEN 'i' THEN 'IMMUTABLE'
		           WHEN 's' THEN 'STABLE'
		           WHEN 'v' THEN 'VOLATILE'
		       END as volatility
		FROM pg_catalog.pg_proc p
		JOIN pg_catalog.pg_namespace n ON n.oid = p.pronamespace
		WHERE n.nspname NOT IN ('pg_catalog', 'information_schema')
		  AND p.prokind IN ('f', 'p')
		ORDER BY n.nspname, p.prokind, p.proname`

	rows, err := c.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var routine RoutineInfo
		var definition sql.NullString
		var returns, arguments, securityType, volatility string

		if err := rows.Scan(&routine.Schema, &routine.Name, &routine.Type,
			&routine.Definer, &definition, &returns, &arguments, &securityType, &volatility); err != nil {
			continue
		}

		routine.Returns = returns
		routine.Parameters = arguments
		routine.SecurityType = securityType
		routine.DataAccess = volatility

		if definition.Valid {
			routine.Definition = definition.String
		}

		info.Routines = append(info.Routines, routine)
	}
	return nil
}

func (c *PostgreSQLCollector) collectVariables(info *DatabaseInfo) error {
	var query string
	if c.config.OnlyModifiedVariables {
		query = `
			SELECT name, setting, COALESCE(unit, '') as unit, source, boot_val,
			       sourcefile, sourceline
			FROM pg_catalog.pg_settings
			WHERE source != 'default'
			ORDER BY name`
	} else {
		query = `
			SELECT name, setting, COALESCE(unit, '') as unit, source, boot_val,
			       sourcefile, sourceline
			FROM pg_catalog.pg_settings
			ORDER BY name`
	}

	rows, err := c.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var variable Variable
		var unit, source string
		var bootVal, sourceFile sql.NullString
		var sourceLine sql.NullInt64

		if err := rows.Scan(&variable.Name, &variable.CurrentValue, &unit, &source, &bootVal,
			&sourceFile, &sourceLine); err != nil {
			continue
		}

		if unit != "" {
			variable.CurrentValue = variable.CurrentValue + " " + unit
		}

		variable.Source = source
		if sourceFile.Valid && sourceFile.String != "" {
			if sourceLine.Valid {
				variable.Source = fmt.Sprintf("%s (%s:%d)", source, sourceFile.String, sourceLine.Int64)
			} else {
				variable.Source = fmt.Sprintf("%s (%s)", source, sourceFile.String)
			}
		}
		variable.IsModified = (source != "default")
		if bootVal.Valid {
			variable.DefaultValue = bootVal.String
			if unit != "" {
				variable.DefaultValue = bootVal.String + " " + unit
			}
		}

		info.Variables = append(info.Variables, variable)
	}
	return nil
}

func (c *PostgreSQLCollector) collectExtensions(info *DatabaseInfo) error {
	query := `
		SELECT e.extname, e.extversion, COALESCE(c.description, '') as description
		FROM pg_catalog.pg_extension e
		LEFT JOIN pg_catalog.pg_description c ON c.objoid = e.oid
		ORDER BY e.extname`

	rows, err := c.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var ext Extension
		if err := rows.Scan(&ext.Name, &ext.Version, &ext.Description); err != nil {
			continue
		}
		info.Extensions = append(info.Extensions, ext)
	}
	return nil
}
