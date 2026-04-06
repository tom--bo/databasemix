package main

import (
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
)

// PostgreSQLVersion represents PostgreSQL version information
type PostgreSQLVersion struct {
	FullVersion string
	Major       int
	Minor       int
}

// DetectPostgreSQLVersion detects the PostgreSQL version from the database connection
func DetectPostgreSQLVersion(db *sql.DB) (*PostgreSQLVersion, error) {
	var versionString string
	err := db.QueryRow("SELECT version()").Scan(&versionString)
	if err != nil {
		return nil, fmt.Errorf("failed to get version: %v", err)
	}

	return ParsePostgreSQLVersion(versionString)
}

// ParsePostgreSQLVersion parses a PostgreSQL version string
// e.g. "PostgreSQL 16.3 (Debian 16.3-1.pgdg120+1) on x86_64-pc-linux-gnu..."
func ParsePostgreSQLVersion(versionString string) (*PostgreSQLVersion, error) {
	version := &PostgreSQLVersion{
		FullVersion: versionString,
	}

	re := regexp.MustCompile(`PostgreSQL\s+(\d+)(?:\.(\d+))?`)
	matches := re.FindStringSubmatch(versionString)
	if len(matches) < 2 {
		return nil, fmt.Errorf("unable to parse PostgreSQL version string: %s", versionString)
	}

	var err error
	version.Major, err = strconv.Atoi(matches[1])
	if err != nil {
		return nil, fmt.Errorf("invalid major version: %s", matches[1])
	}

	if len(matches) >= 3 && matches[2] != "" {
		version.Minor, err = strconv.Atoi(matches[2])
		if err != nil {
			return nil, fmt.Errorf("invalid minor version: %s", matches[2])
		}
	}

	return version, nil
}

// String returns a string representation of the version
func (v *PostgreSQLVersion) String() string {
	if v.Minor > 0 {
		return fmt.Sprintf("PostgreSQL %d.%d", v.Major, v.Minor)
	}
	return fmt.Sprintf("PostgreSQL %d", v.Major)
}

// IsAtLeast checks if this version is at least the specified major version
func (v *PostgreSQLVersion) IsAtLeast(major int) bool {
	return v.Major >= major
}
