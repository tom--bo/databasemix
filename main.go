package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Configuration for database connection
type Config struct {
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
}


func main() {
	// Parse command line arguments
	config := parseFlags()

	// Connect to MySQL
	db, err := connectToMySQL(config)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	defer db.Close()

	// Create MySQL collector
	collector, err := NewMySQLCollector(db, config)
	if err != nil {
		log.Fatalf("Failed to create MySQL collector: %v", err)
	}

	// Collect all MySQL information
	info, err := collector.CollectAll()
	if err != nil {
		log.Fatalf("Failed to collect MySQL information: %v", err)
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

func parseFlags() *Config {
	config := &Config{}

	flag.StringVar(&config.Host, "host", "localhost", "Database host")
	flag.StringVar(&config.Port, "port", "3306", "MySQL port")
	flag.StringVar(&config.User, "user", "root", "MySQL user")
	flag.StringVar(&config.Password, "password", "", "Database password")
	flag.StringVar(&config.Database, "database", "", "MySQL database name (if not specified, all accessible databases will be analyzed)")
	flag.BoolVar(&config.Replication, "replication", false, "Include replication information")
	flag.BoolVar(&config.ExceptTables, "except-tables", false, "Exclude tables and views")
	flag.BoolVar(&config.ExceptStoredProcedures, "except-stored-procedures", false, "Exclude stored procedures and functions")
	flag.BoolVar(&config.ExceptVariables, "except-variables", false, "Exclude variables/configuration parameters")
	flag.BoolVar(&config.OnlyModifiedVariables, "only-modified-variables", false, "Show only modified variables (default: show all)")
	flag.BoolVar(&config.ExceptUsers, "except-users", false, "Exclude user accounts")
	flag.BoolVar(&config.ExceptRoles, "except-roles", false, "Exclude user roles")
	flag.BoolVar(&config.ExceptPlugins, "except-plugins", false, "Exclude installed plugins")
	flag.StringVar(&config.Format, "format", "markdown", "Output format: markdown, xml, plaintext")
	flag.StringVar(&config.OutputFile, "outfile", "dbmix-output", "Output filename (if not specified, output goes to stdout)")

	flag.Parse()


	// Try to get values from environment variables if not provided via flags
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
		// Default to markdown for invalid formats
		fmt.Printf("Warning: Unsupported format '%s'. Using default 'markdown' format.\n", config.Format)
		config.Format = "markdown"
	}

	return config
}

// connectToMySQL connects to MySQL database
func connectToMySQL(config *Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.User, config.Password, config.Host, config.Port, config.Database)

	// Connect to MySQL
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Test the connection
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

// Main data structure that holds all MySQL information
type DatabaseInfo struct {
	ConnectionInfo  *ConnectionInfo
	Tables          []TableInfo
	Users           []UserAccount
	Routines        []RoutineInfo
	Variables       []Variable
	Roles           []UserRole
	Plugins         []Plugin
	Components      []Component
	ReplicationInfo *ReplicationInfo
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

// User account information
type UserAccount struct {
	User            string
	Host            string
	SSLType         string
	Plugin          string
	AccountLocked   string
	PasswordExpired string
	Grants          []string
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
