package main

import (
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// MySQLVersion represents MySQL version information
type MySQLVersion struct {
	FullVersion string
	Major       int
	Minor       int
	Patch       int
	Variant     string // e.g., "mysql", "mariadb", "percona"
}

// DetectMySQLVersion detects the MySQL version from the database connection
func DetectMySQLVersion(db *sql.DB) (*MySQLVersion, error) {
	var versionString string
	err := db.QueryRow("SELECT VERSION()").Scan(&versionString)
	if err != nil {
		return nil, fmt.Errorf("failed to get version: %v", err)
	}

	return ParseMySQLVersion(versionString)
}

// ParseMySQLVersion parses a MySQL version string
func ParseMySQLVersion(versionString string) (*MySQLVersion, error) {
	version := &MySQLVersion{
		FullVersion: versionString,
		Variant:     "mysql", // default
	}

	// Detect variant
	lowerVersion := strings.ToLower(versionString)
	if strings.Contains(lowerVersion, "mariadb") {
		version.Variant = "mariadb"
	} else if strings.Contains(lowerVersion, "percona") {
		version.Variant = "percona"
	}

	// Extract version numbers using regex
	// Pattern matches versions like "8.0.42", "5.7.44", etc.
	re := regexp.MustCompile(`(\d+)\.(\d+)\.(\d+)`)
	matches := re.FindStringSubmatch(versionString)

	if len(matches) < 4 {
		return nil, fmt.Errorf("unable to parse version string: %s", versionString)
	}

	var err error
	version.Major, err = strconv.Atoi(matches[1])
	if err != nil {
		return nil, fmt.Errorf("invalid major version: %s", matches[1])
	}

	version.Minor, err = strconv.Atoi(matches[2])
	if err != nil {
		return nil, fmt.Errorf("invalid minor version: %s", matches[2])
	}

	version.Patch, err = strconv.Atoi(matches[3])
	if err != nil {
		return nil, fmt.Errorf("invalid patch version: %s", matches[3])
	}

	return version, nil
}

// IsMySQL8OrLater returns true if this is MySQL 8.0 or later
func (v *MySQLVersion) IsMySQL8OrLater() bool {
	if v.Variant != "mysql" {
		return false
	}
	return v.Major >= 8
}

// IsMySQL57OrLater returns true if this is MySQL 5.7 or later
func (v *MySQLVersion) IsMySQL57OrLater() bool {
	if v.Variant != "mysql" {
		return false
	}
	if v.Major > 5 {
		return true
	}
	if v.Major == 5 && v.Minor >= 7 {
		return true
	}
	return false
}

// IsMariaDB returns true if this is MariaDB
func (v *MySQLVersion) IsMariaDB() bool {
	return v.Variant == "mariadb"
}

// IsPercona returns true if this is Percona Server
func (v *MySQLVersion) IsPercona() bool {
	return v.Variant == "percona"
}

// SupportsRoles returns true if this version supports roles
func (v *MySQLVersion) SupportsRoles() bool {
	return v.IsMySQL8OrLater()
}

// SupportsComponents returns true if this version supports components
func (v *MySQLVersion) SupportsComponents() bool {
	return v.IsMySQL8OrLater()
}

// SupportsInformationSchemaViews returns true if this version has modern information_schema views
func (v *MySQLVersion) SupportsInformationSchemaViews() bool {
	return v.IsMySQL57OrLater()
}

// SupportsPerformanceSchemaVariablesInfo returns true if performance_schema.variables_info exists
func (v *MySQLVersion) SupportsPerformanceSchemaVariablesInfo() bool {
	return v.IsMySQL57OrLater()
}

// GetUserTableQuery returns the appropriate query to get user information based on version
func (v *MySQLVersion) GetUserTableQuery() string {
	if v.IsMySQL8OrLater() {
		return `
			SELECT User, Host, plugin, account_locked, password_expired
			FROM mysql.user 
			ORDER BY User, Host`
	} else {
		return `
			SELECT User, Host, plugin, account_locked, 'N' as password_expired
			FROM mysql.user 
			ORDER BY User, Host`
	}
}

// GetRoleQuery returns the appropriate query to get role information
func (v *MySQLVersion) GetRoleQuery() string {
	if !v.SupportsRoles() {
		return ""
	}

	// In MySQL 8.0, roles are users with account_locked='Y' and password_expired != 'Y'
	// But this is a simplified detection - actual role detection is more complex
	return `
		SELECT User as role_name, Host as role_host
		FROM mysql.user 
		WHERE account_locked = 'Y'
		ORDER BY User, Host`
}

// GetVariablesQuery returns the appropriate query to get modified variables
func (v *MySQLVersion) GetVariablesQuery() string {
	return `
			SELECT VARIABLE_NAME, VARIABLE_VALUE, VARIABLE_SOURCE
			FROM performance_schema.variables_info vi
			JOIN performance_schema.global_variables gv ON vi.VARIABLE_NAME = gv.VARIABLE_NAME
			WHERE VARIABLE_SOURCE != 'COMPILED'
			ORDER BY VARIABLE_NAME`
}

// GetReplicationStatusQuery returns the appropriate query for replication status
func (v *MySQLVersion) GetReplicationStatusQuery() string {
	if v.IsMySQL8OrLater() {
		return "SHOW REPLICA STATUS"
	} else {
		return "SHOW SLAVE STATUS"
	}
}

// String returns a string representation of the version
func (v *MySQLVersion) String() string {
	return fmt.Sprintf("%s %d.%d.%d (%s)",
		strings.Title(v.Variant), v.Major, v.Minor, v.Patch, v.FullVersion)
}

// CompareVersion compares this version with another version
// Returns: -1 if this < other, 0 if equal, 1 if this > other
func (v *MySQLVersion) CompareVersion(other *MySQLVersion) int {
	if v.Major != other.Major {
		if v.Major < other.Major {
			return -1
		}
		return 1
	}

	if v.Minor != other.Minor {
		if v.Minor < other.Minor {
			return -1
		}
		return 1
	}

	if v.Patch != other.Patch {
		if v.Patch < other.Patch {
			return -1
		}
		return 1
	}

	return 0
}

// IsAtLeast checks if this version is at least the specified version
func (v *MySQLVersion) IsAtLeast(major, minor, patch int) bool {
	if v.Major > major {
		return true
	}
	if v.Major < major {
		return false
	}

	if v.Minor > minor {
		return true
	}
	if v.Minor < minor {
		return false
	}

	return v.Patch >= patch
}
