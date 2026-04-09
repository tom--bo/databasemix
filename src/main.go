package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// Configuration for database connection
type Config struct {
	DBType                 string // "mysql" or "postgres"
	Host                   string
	Port                   string
	User                   string
	Password               string
	Database               string
	Replication            bool
	Format                 string
	OutputFile             string
	ExceptTables           bool
	ExceptStoredProcedures bool
	ExceptVariables        bool
	OnlyModifiedVariables  bool
	ExceptUsers            bool
	ExceptRoles            bool
	ExceptPlugins          bool
	ExceptExtensions       bool // PostgreSQL only
	TLS                    string // TLS/SSL mode: MySQL(false,true,skip-verify,preferred), PostgreSQL(disable,require,verify-ca,verify-full)
}

func main() {
	// Parse command line arguments
	config, err := parseFlags()
	if err != nil {
		log.Fatalf("Error parsing flags: %v", err)
		os.Exit(1)
	}

	// Connect to database and create collector based on DB type
	var db *sql.DB
	var collector Collector

	switch config.DBType {
	case "mysql":
		db, err = connectToMySQL(config)
		if err != nil {
			log.Fatalf("Failed to connect to MySQL: %v", err)
		}
		defer db.Close()
		collector, err = NewMySQLCollector(db, config)
		if err != nil {
			log.Fatalf("Failed to create MySQL collector: %v", err)
		}
	case "postgres":
		db, err = connectToPostgreSQL(config)
		if err != nil {
			log.Fatalf("Failed to connect to PostgreSQL: %v", err)
		}
		defer db.Close()
		collector, err = NewPostgreSQLCollector(db, config)
		if err != nil {
			log.Fatalf("Failed to create PostgreSQL collector: %v", err)
		}
	}

	// Collect all database information
	info, err := collector.CollectAll()
	if err != nil {
		log.Fatalf("Failed to collect database information: %v", err)
	}

	// Create formatter based on requested format
	format := OutputFormat(config.Format)
	formatter, err := NewFormatter(format)
	if err != nil {
		log.Fatalf("Failed to create formatter: %v", err)
	}

	// Format the output
	output, err := formatter.Format(info)
	if err != nil {
		log.Fatalf("Failed to format output: %v", err)
	}

	// Determine output filename with appropriate extension
	outputFile := config.OutputFile
	if outputFile != "" {
		// Add file extension based on format if not already present
		var extension string
		switch config.Format {
		case "xml":
			extension = ".xml"
		case "plaintext":
			extension = ".txt"
		default: // markdown
			extension = ".md"
		}

		// Check if the file already has the correct extension
		if !strings.HasSuffix(strings.ToLower(outputFile), extension) {
			outputFile = outputFile + extension
		}
	}

	// Write output
	if outputFile == "" {
		// Write to stdout
		fmt.Print(output)
	} else {
		// Write to specified file
		err = os.WriteFile(outputFile, []byte(output), 0644)
		if err != nil {
			log.Fatalf("Failed to write to file: %v", err)
		}
		fmt.Printf("Database information has been written to %s\n", outputFile)
	}
}

func setDefaultConfigFromEnv(config *Config) {
	// Try to get values from environment variables
	// MySQL
	if config.Host == "localhost" && os.Getenv("MYSQL_HOST") != "" {
		config.Host = os.Getenv("MYSQL_HOST")
	}
	if os.Getenv("MYSQL_PORT") != "" {
		config.Port = os.Getenv("MYSQL_PORT")
	}
	if os.Getenv("MYSQL_USER") != "" {
		config.User = os.Getenv("MYSQL_USER")
	}
	if config.Password == "" && os.Getenv("MYSQL_PASSWORD") != "" {
		config.Password = os.Getenv("MYSQL_PASSWORD")
	}
	if config.Database == "" && os.Getenv("MYSQL_DATABASE") != "" {
		config.Database = os.Getenv("MYSQL_DATABASE")
	}
	// PostgreSQL
	if config.Host == "localhost" && os.Getenv("PGHOST") != "" {
		config.Host = os.Getenv("PGHOST")
	}
	if os.Getenv("PGPORT") != "" {
		config.Port = os.Getenv("PGPORT")
	}
	if os.Getenv("PGUSER") != "" {
		config.User = os.Getenv("PGUSER")
	}
	if config.Password == "" && os.Getenv("PGPASSWORD") != "" {
		config.Password = os.Getenv("PGPASSWORD")
	}
	if config.Database == "" && os.Getenv("PGDATABASE") != "" {
		config.Database = os.Getenv("PGDATABASE")
	}
}

func parseFlags() (*Config, error) {
	config := &Config{}
	setDefaultConfigFromEnv(config)

	flag.StringVar(&config.DBType, "type", "", "Database type: mysql, postgres (default: mysql)")
	flag.StringVar(&config.Host, "host", "localhost", "Database host, (default: localhost)")
	flag.StringVar(&config.Port, "port", "", "Database port (default: 3306 for mysql, 5432 for postgres)")
	flag.StringVar(&config.User, "user", "", "Database user (default: root for mysql, postgres for postgres)")
	flag.StringVar(&config.Password, "password", "", "Database password")
	flag.StringVar(&config.Database, "database", "", "Database name (if not specified, all accessible databases will be analyzed)")
	flag.BoolVar(&config.Replication, "replication", false, "Include replication information (MySQL only)")
	flag.BoolVar(&config.ExceptTables, "except-tables", false, "Exclude tables and views")
	flag.BoolVar(&config.ExceptStoredProcedures, "except-stored-procedures", false, "Exclude stored procedures and functions")
	flag.BoolVar(&config.ExceptVariables, "except-variables", false, "Exclude variables/configuration parameters")
	flag.BoolVar(&config.OnlyModifiedVariables, "only-modified-variables", false, "Show only modified variables (default: show all)")
	flag.BoolVar(&config.ExceptUsers, "except-users", false, "Exclude user accounts")
	flag.BoolVar(&config.ExceptRoles, "except-roles", false, "Exclude user roles")
	flag.BoolVar(&config.ExceptPlugins, "except-plugins", false, "Exclude installed plugins (MySQL only)")
	flag.BoolVar(&config.ExceptExtensions, "except-extensions", false, "Exclude installed extensions (PostgreSQL only)")
	flag.StringVar(&config.TLS, "tls", "", "TLS/SSL mode (MySQL: false,true,skip-verify,preferred / PostgreSQL: disable,require,verify-ca,verify-full)")
	flag.StringVar(&config.Format, "format", "markdown", "Output format: markdown, xml, plaintext")
	flag.StringVar(&config.OutputFile, "outfile", "dbmix-output", "Output filename (if not specified, output goes to stdout)")

	flag.Parse()

	// Resolve DB type
	dbType := strings.ToLower(strings.TrimSpace(config.DBType))
	switch dbType {
	case "mysql", "postgres", "postgresql":
		if dbType == "postgresql" {
			dbType = "postgres"
		}
		config.DBType = dbType
	case "":
		dbType = "mysql"
		config.DBType = dbType
	default:
		return nil, errors.New(fmt.Sprintf("unsupported database type '%s'", config.DBType))
	}
	// Set defaults based on DB type
	if config.Port == "" {
		if config.DBType == "postgres" {
			config.Port = "5432"
		} else {
			config.Port = "3306"
		}
	}
	if config.User == "" {
		if config.DBType == "postgres" {
			config.User = "postgres"
		} else {
			config.User = "root"
		}
	}

	// Validate TLS option based on DB type
	if config.TLS != "" {
		tlsVal := strings.ToLower(strings.TrimSpace(config.TLS))
		config.TLS = tlsVal
		switch config.DBType {
		case "mysql":
			switch tlsVal {
			case "true", "false", "skip-verify", "preferred":
				// valid
			default:
				return nil, fmt.Errorf("unsupported TLS mode '%s' for MySQL. Valid values: false, true, skip-verify, preferred", config.TLS)
			}
		case "postgres":
			switch tlsVal {
			case "disable", "require", "verify-ca", "verify-full":
				// valid
			default:
				return nil, fmt.Errorf("unsupported TLS mode '%s' for PostgreSQL. Valid values: disable, require, verify-ca, verify-full", config.TLS)
			}
		}
	}

	// Validate format and set default to markdown
	format := strings.ToLower(strings.TrimSpace(config.Format))
	switch format {
	case "", "markdown", "md":
		config.Format = "markdown"
	case "xml":
		config.Format = "xml"
	case "plaintext", "text", "txt":
		config.Format = "plaintext"
	default:
		fmt.Printf("Warning: Unsupported format '%s'. Using default 'markdown' format.\n", config.Format)
		config.Format = "markdown"
	}

	return config, nil
}

// connectToMySQL connects to MySQL database
func connectToMySQL(config *Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.User, config.Password, config.Host, config.Port, config.Database)

	// Add TLS parameter if specified
	if config.TLS != "" && config.TLS != "false" {
		dsn += "?tls=" + config.TLS
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// connectToPostgreSQL connects to PostgreSQL database
func connectToPostgreSQL(config *Config) (*sql.DB, error) {
	sslmode := "disable"
	if config.TLS != "" {
		sslmode = config.TLS
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, sslmode)
	if config.Database != "" {
		dsn += fmt.Sprintf(" dbname=%s", config.Database)
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Data structures remain the same, but moved here for completeness

// Connection information for display
type ConnectionInfo struct {
	Host     string
	Port     string
	User     string
	Database string
	Version  string
}

// Main data structure that holds all database information
type DatabaseInfo struct {
	DBType           string // "mysql" or "postgres"
	ConnectionInfo   *ConnectionInfo
	Tables           []TableInfo
	Users            []UserAccount
	Routines         []RoutineInfo
	Variables        []Variable
	Roles            []UserRole         // Kept for backward compat but no longer populated
	RoleRelatedVars  []RoleRelatedVar   // Role-related system variables
	Plugins          []Plugin
	Components       []Component
	ReplicationInfo  *ReplicationInfo
	Extensions       []Extension // PostgreSQL only
}

// Extension represents a PostgreSQL extension
type Extension struct {
	Name        string
	Version     string
	Description string
}

// Table information
type TableInfo struct {
	Database      string
	Schema        string
	Name          string
	Type          string // TABLE, VIEW
	Engine        string
	AutoIncrement int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Collation     string
	Charset       string
	RowFormat     string
	Comment       string
	CreateOptions string
	DDL           string
}

// RoleEdge represents a role assignment (role granted to a user/role)
type RoleEdge struct {
	FromRole    string // The role being granted (MySQL: 'role'@'host', PostgreSQL: 'rolename')
	ToUser      string // The user/role receiving the grant (MySQL: 'user'@'host', PostgreSQL: 'rolename')
	WithAdmin   bool   // WITH ADMIN OPTION
	IsDefault   bool   // Whether this is a default role for the target user
}

// RoleRelatedVar represents a role-related system variable
type RoleRelatedVar struct {
	Name  string
	Value string
}

// User account information
type UserAccount struct {
	User            string
	Host            string
	SSLType         string
	Plugin          string
	AccountLocked   string
	PasswordExpired string
	Grants          []string
	AssignedRoles   []RoleEdge // Roles assigned TO this user (this user is the recipient)
	GrantedTo       []RoleEdge // Users/roles this role is granted to (this user/role is the source)
	IsRole          bool       // Hint: true if this account looks like a role (MySQL: account_locked=Y, PG: !canlogin)
	// PostgreSQL specific
	IsSuperuser   bool
	CanCreateDB   bool
	CanCreateRole bool
	CanLogin      bool
	Inherit       bool
	ConnLimit     int
	ValidUntil    string
}

// Stored routine information (procedures and functions)
type RoutineInfo struct {
	Schema       string
	Name         string
	Type         string // FUNCTION, PROCEDURE
	Definer      string
	Created      time.Time
	LastAltered  time.Time
	DataAccess   string
	SecurityType string // DEFINER, INVOKER
	Returns      string
	Parameters   string
	Definition   string
}

// Variable information
type Variable struct {
	Name         string
	CurrentValue string
	DefaultValue string
	Source       string
	IsModified   bool
}

// User role information
type UserRole struct {
	RoleName string
	RoleHost string
	Grants   []string
	Members  []string
}

// Plugin information
type Plugin struct {
	Name        string
	Version     string
	Status      string
	Type        string
	Library     string
	Description string
}

// Component information (MySQL 8.0+)
type Component struct {
	ComponentID      int
	ComponentGroupID int
	ComponentURN     string
}

// Replication information
type ReplicationInfo struct {
	ReplicationStatus    *ReplicationStatus
	ReplicaStatus        []ReplicaStatus
	SemiSyncStatus       *SemiSyncStatus
	GroupReplicationInfo *GroupReplicationInfo
}

// Basic replication status
type ReplicationStatus struct {
	ServerID       int
	ServerUUID     string
	LogBinEnabled  bool
	BinlogFormat   string
	GTIDMode       string
	CurrentLogFile string
	CurrentLogPos  int64
}

// Replica status information
type ReplicaStatus struct {
	SlaveIORunning  string
	SlaveSQLRunning string
	MasterHost      string
	MasterPort      int
	MasterLogFile   string
	MasterLogPos    int64
	SecondsBehind   int64
}

// Semi-synchronous replication status
type SemiSyncStatus struct {
	MasterEnabled bool
	SlaveEnabled  bool
}

// Group replication information
type GroupReplicationInfo struct {
	GroupName         string
	MemberState       string
	SinglePrimaryMode bool
}
