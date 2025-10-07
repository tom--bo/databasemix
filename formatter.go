package main

import (
	"encoding/xml"
	"fmt"
	"sort"
	"strings"
	"time"
)

// OutputFormat represents the output format type
type OutputFormat string

const (
	FormatMarkdown  OutputFormat = "markdown"
	FormatXML       OutputFormat = "xml"
	FormatPlaintext OutputFormat = "plaintext"
)

// Formatter interface for different output formats
type Formatter interface {
	Format(info *DatabaseInfo) (string, error)
	GetFileExtension() string
}

// NewFormatter creates a new formatter based on the specified format
func NewFormatter(format OutputFormat) (Formatter, error) {
	switch format {
	case FormatMarkdown:
		return &MarkdownFormatter{}, nil
	case FormatXML:
		return &XMLFormatter{}, nil
	case FormatPlaintext:
		return &PlaintextFormatter{}, nil
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}
}

// MarkdownFormatter formats output as Markdown
type MarkdownFormatter struct{}

// Format formats database info as Markdown
func (f *MarkdownFormatter) Format(info *DatabaseInfo) (string, error) {
	var result strings.Builder

	// File Summary section
	result.WriteString("# File Summary\n\n")
	result.WriteString("This file contains comprehensive MySQL database information compiled for AI context analysis. ")
	result.WriteString("It includes schema definitions, account configurations, system variables, and other database metadata ")
	result.WriteString("consolidated into a single file for efficient processing.\n\n")

	// Add database type and connection info
	if info.ConnectionInfo != nil {
		result.WriteString("**Database Type**: MySQL  \n")
		result.WriteString(fmt.Sprintf("**Database Version**: %s\n\n", info.ConnectionInfo.Version))
	}
	
	// Generate sections list based on actual content
	sections := f.generateSectionsList(info)
	if len(sections) > 0 {
		result.WriteString("## File Structure\n\n")
		for _, section := range sections {
			result.WriteString(fmt.Sprintf("- %s\n", section))
		}
		result.WriteString("\n")
	}

	// Variables
	if len(info.Variables) > 0 {
		result.WriteString("# Variables\n\n")
		f.formatVariables(&result, info.Variables)
	}

	// Tables
	if len(info.Tables) > 0 {
		result.WriteString("# Tables\n\n")
		f.formatTables(&result, info.Tables)
	}

	// Views (separate section for view details)
	if f.hasViews(info.Tables) {
		result.WriteString("# View info details\n\n")
		f.formatViewDetails(&result, info.Tables)
	}

	// Stored functions
	functions := f.filterRoutines(info.Routines, "FUNCTION")
	if len(functions) > 0 {
		result.WriteString("# Stored Functions\n\n")
		f.formatRoutines(&result, functions)
	}

	// Stored procedures
	procedures := f.filterRoutines(info.Routines, "PROCEDURE")
	if len(procedures) > 0 {
		result.WriteString("# Stored Procedures\n\n")
		f.formatRoutines(&result, procedures)
	}

	// Roles
	if len(info.Roles) > 0 {
		result.WriteString("# User Roles (MySQL 8.0+)\n\n")
		f.formatRoles(&result, info.Roles)
	}

	// Users
	if len(info.Users) > 0 {
		result.WriteString("# User List\n\n")
		f.formatUsers(&result, info.Users)
	}

	// Plugins
	if len(info.Plugins) > 0 {
		result.WriteString("# Plugins\n\n")
		f.formatPlugins(&result, info.Plugins)
	}

	// Components (MySQL 8.0+)
	if len(info.Components) > 0 {
		result.WriteString("# Components (MySQL 8.0+)\n\n")
		f.formatComponents(&result, info.Components)
	}

	// Replication information
	if info.ReplicationInfo != nil {
		result.WriteString("# Replication Information\n\n")
		f.formatReplicationInfo(&result, info.ReplicationInfo)
	}

	return result.String(), nil
}

func (f *MarkdownFormatter) formatTables(result *strings.Builder, tables []TableInfo) {
	if len(tables) == 0 {
		return
	}
	
	result.WriteString("# Tables\n\n")
	
	// Group tables by database and schema, excluding views
	dbSchemaMap := make(map[string]map[string][]TableInfo)
	for _, table := range tables {
		if table.Type == "BASE TABLE" {
			if dbSchemaMap[table.Database] == nil {
				dbSchemaMap[table.Database] = make(map[string][]TableInfo)
			}
			dbSchemaMap[table.Database][table.Schema] = append(dbSchemaMap[table.Database][table.Schema], table)
		}
	}

	// Sort databases
	var databases []string
	for db := range dbSchemaMap {
		databases = append(databases, db)
	}
	sort.Strings(databases)
	
	for _, db := range databases {
		schemaMap := dbSchemaMap[db]
		
		// Sort schemas
		var schemas []string
		for schema := range schemaMap {
			schemas = append(schemas, schema)
		}
		sort.Strings(schemas)

		for _, schema := range schemas {
			schemaTables := schemaMap[schema]
			sort.Slice(schemaTables, func(i, j int) bool {
				return schemaTables[i].Name < schemaTables[j].Name
			})

			for _, table := range schemaTables {
				// Include database name if there are multiple databases
				if len(databases) > 1 {
					result.WriteString(fmt.Sprintf("## %s.%s.%s\n\n", table.Database, table.Schema, table.Name))
				} else {
					result.WriteString(fmt.Sprintf("## %s.%s\n\n", table.Schema, table.Name))
				}

			if table.Engine != "" {
				result.WriteString(fmt.Sprintf("- Engine: %s\n", table.Engine))
			}
			if table.AutoIncrement > 0 {
				result.WriteString(fmt.Sprintf("- Auto Increment: %d\n", table.AutoIncrement))
			}
			if !table.CreatedAt.IsZero() {
				result.WriteString(fmt.Sprintf("- Created: %s\n", table.CreatedAt.Format("2006-01-02 15:04:05")))
			}
			if !table.UpdatedAt.IsZero() {
				result.WriteString(fmt.Sprintf("- Updated: %s\n", table.UpdatedAt.Format("2006-01-02 15:04:05")))
			}
			if table.Collation != "" {
				result.WriteString(fmt.Sprintf("- Collation: %s\n", table.Collation))
			}
			if table.Charset != "" {
				result.WriteString(fmt.Sprintf("- Charset: %s\n", table.Charset))
			}
			if table.RowFormat != "" {
				result.WriteString(fmt.Sprintf("- Row Format: %s\n", table.RowFormat))
			}
			if table.Comment != "" {
				result.WriteString(fmt.Sprintf("- Comment: %s\n", table.Comment))
			}
			if table.CreateOptions != "" {
				result.WriteString(fmt.Sprintf("- Create Options: %s\n", table.CreateOptions))
			}

			if table.DDL != "" {
				result.WriteString("\n```sql\n")
				result.WriteString(table.DDL)
				result.WriteString("\n```\n\n")
			}
		}
	}
	}
}

func (f *MarkdownFormatter) formatViewDetails(result *strings.Builder, tables []TableInfo) {
	for _, table := range tables {
		if table.Type == "VIEW" && table.DDL != "" {
			result.WriteString(fmt.Sprintf("## %s.%s\n\n", table.Schema, table.Name))
			result.WriteString("```sql\n")
			result.WriteString(table.DDL)
			result.WriteString("\n```\n\n")
		}
	}
}

func (f *MarkdownFormatter) formatVariables(result *strings.Builder, variables []Variable) {
	// Check if this is only-modified-variables mode by checking if all variables are modified
	hasExtendedInfo := len(variables) > 0 && variables[0].IsModified
	if hasExtendedInfo {
		// Double-check that all variables are actually modified (only-modified-variables mode)
		for _, v := range variables {
			if !v.IsModified {
				hasExtendedInfo = false
				break
			}
		}
	}
	
	if hasExtendedInfo {
		// 4-column format for -only-modified-variables
		result.WriteString("| Variable Name | Current Value | Default Value | Source |\n")
		result.WriteString("|---------------|---------------|---------------|--------|\n")
		
		for _, variable := range variables {
			result.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n",
				variable.Name, variable.CurrentValue, variable.DefaultValue, variable.Source))
		}
	} else {
		// 2-column format for normal variables
		result.WriteString("| Variable Name | Current Value |\n")
		result.WriteString("|---------------|---------------|\n")
		
		for _, variable := range variables {
			result.WriteString(fmt.Sprintf("| %s | %s |\n",
				variable.Name, variable.CurrentValue))
		}
	}
	result.WriteString("\n")
}

// generateSectionsList creates a list of sections that will be included in the output
func (f *MarkdownFormatter) generateSectionsList(info *DatabaseInfo) []string {
	var sections []string
	
	if len(info.Variables) > 0 {
		sections = append(sections, "Variables - MySQL system variables and their current values")
	}
	if len(info.Tables) > 0 {
		sections = append(sections, "Tables - Database tables with metadata and DDL definitions")
	}
	if f.hasViews(info.Tables) {
		sections = append(sections, "View Details - Database views with their definitions")
	}
	functions := f.filterRoutines(info.Routines, "FUNCTION")
	if len(functions) > 0 {
		sections = append(sections, "Stored Functions - User-defined functions with their definitions")
	}
	procedures := f.filterRoutines(info.Routines, "PROCEDURE")
	if len(procedures) > 0 {
		sections = append(sections, "Stored Procedures - User-defined procedures with their definitions")
	}
	if len(info.Roles) > 0 {
		sections = append(sections, "User Roles - MySQL 8.0+ role definitions and assignments")
	}
	if len(info.Users) > 0 {
		sections = append(sections, "User Accounts - Database user accounts with privileges")
	}
	if len(info.Plugins) > 0 {
		sections = append(sections, "Plugins - Installed MySQL plugins and extensions")
	}
	if len(info.Components) > 0 {
		sections = append(sections, "Components - MySQL 8.0+ components")
	}
	if info.ReplicationInfo != nil {
		sections = append(sections, "Replication Info - MySQL replication configuration and status")
	}
	
	return sections
}

func (f *MarkdownFormatter) formatUsers(result *strings.Builder, users []UserAccount) {
	for _, user := range users {
		result.WriteString(fmt.Sprintf("## %s@%s\n\n", user.User, user.Host))
		result.WriteString(fmt.Sprintf("- Plugin: %s\n", user.Plugin))
		result.WriteString(fmt.Sprintf("- Account Locked: %s\n", user.AccountLocked))
		
		if len(user.Grants) > 0 {
			result.WriteString("- Grants:\n")
			for _, grant := range user.Grants {
				result.WriteString(fmt.Sprintf("  - %s\n", grant))
			}
		}
		result.WriteString("\n")
	}
}

func (f *MarkdownFormatter) formatRoutines(result *strings.Builder, routines []RoutineInfo) {
	for _, routine := range routines {
		result.WriteString(fmt.Sprintf("## %s.%s\n\n", routine.Schema, routine.Name))
		result.WriteString(fmt.Sprintf("- Specific Name: %s\n", routine.Name))
		result.WriteString("- Routine Catalog: def\n")
		if routine.Type == "FUNCTION" && routine.Returns != "" {
			result.WriteString(fmt.Sprintf("- Character Set: utf8mb4\n"))
			result.WriteString(fmt.Sprintf("- Collation: utf8mb4_unicode_ci\n"))
		}
		result.WriteString("- Routine Body: SQL\n")
		result.WriteString("- External Language: SQL\n")
		result.WriteString("- Parameter Style: SQL\n")
		result.WriteString("- Is Deterministic: YES\n")
		if routine.DataAccess != "" {
			result.WriteString(fmt.Sprintf("- SQL Data Access: %s\n", routine.DataAccess))
		}
		if routine.SecurityType != "" {
			result.WriteString(fmt.Sprintf("- Security Type: %s\n", routine.SecurityType))
		}
		if !routine.Created.IsZero() {
			result.WriteString(fmt.Sprintf("- Created: %s\n", routine.Created.Format("2006-01-02 15:04:05")))
		}
		if !routine.LastAltered.IsZero() {
			result.WriteString(fmt.Sprintf("- Last Altered: %s\n", routine.LastAltered.Format("2006-01-02 15:04:05")))
		}
		result.WriteString("- SQL Mode: ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION\n")
		
		if routine.Definition != "" {
			result.WriteString("\n```sql\n")
			result.WriteString(routine.Definition)
			result.WriteString("\n```\n\n")
		}
	}
}

func (f *MarkdownFormatter) formatRoles(result *strings.Builder, roles []UserRole) {
	for _, role := range roles {
		result.WriteString(fmt.Sprintf("## %s@%s\n\n", role.RoleName, role.RoleHost))
		
		if len(role.Members) > 0 {
			result.WriteString("- **Default Role for some users**\n\n")
		}
		
		if len(role.Grants) > 0 {
			for _, grant := range role.Grants {
				result.WriteString(fmt.Sprintf("- %s\n", grant))
			}
			result.WriteString("\n")
		}
	}
}

func (f *MarkdownFormatter) formatPlugins(result *strings.Builder, plugins []Plugin) {
	result.WriteString("| Name | Status | Type | Library | Version | Description |\n")
	result.WriteString("|------|--------|------|---------|---------|-------------|\n")
	
	for _, plugin := range plugins {
		library := plugin.Library
		if library == "" {
			library = "-"
		}
		result.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s |\n",
			plugin.Name, plugin.Status, plugin.Type, library, plugin.Version, plugin.Description))
	}
	result.WriteString("\n")
}

func (f *MarkdownFormatter) formatComponents(result *strings.Builder, components []Component) {
	if len(components) == 0 {
		result.WriteString("- No components found (MySQL 8.0+ feature)\n\n")
		return
	}
	
	for _, component := range components {
		result.WriteString(fmt.Sprintf("- Component ID: %d\n", component.ComponentID))
		result.WriteString(fmt.Sprintf("  - Group ID: %d\n", component.ComponentGroupID))
		result.WriteString(fmt.Sprintf("  - URN: %s\n", component.ComponentURN))
	}
	result.WriteString("\n")
}

func (f *MarkdownFormatter) formatReplicationInfo(result *strings.Builder, replication *ReplicationInfo) {
	if replication.ReplicationStatus != nil {
		result.WriteString("## Basic Replication Status\n\n")
		status := replication.ReplicationStatus
		
		result.WriteString(fmt.Sprintf("- **Server ID**: %d\n", status.ServerID))
		result.WriteString(fmt.Sprintf("- **Server UUID**: %s\n", status.ServerUUID))
		result.WriteString("- **Role**: Master\n")
		result.WriteString(fmt.Sprintf("- **Binary Log Enabled**: %t\n", status.LogBinEnabled))
		result.WriteString(fmt.Sprintf("- **Binary Log Format**: %s\n", status.BinlogFormat))
		if status.CurrentLogFile != "" {
			result.WriteString(fmt.Sprintf("- **Current Binary Log File**: %s\n", status.CurrentLogFile))
			result.WriteString(fmt.Sprintf("- **Current Binary Log Position**: %d\n", status.CurrentLogPos))
		}
		result.WriteString(fmt.Sprintf("- **GTID Mode**: %s\n", status.GTIDMode))
		result.WriteString("\n")
	}
}

func (f *MarkdownFormatter) hasViews(tables []TableInfo) bool {
	for _, table := range tables {
		if table.Type == "VIEW" {
			return true
		}
	}
	return false
}

func (f *MarkdownFormatter) filterRoutines(routines []RoutineInfo, routineType string) []RoutineInfo {
	var filtered []RoutineInfo
	for _, routine := range routines {
		if routine.Type == routineType {
			filtered = append(filtered, routine)
		}
	}
	return filtered
}

func (f *MarkdownFormatter) filterTables(tables []TableInfo, tableType string) []TableInfo {
	var filtered []TableInfo
	for _, table := range tables {
		if table.Type == tableType {
			filtered = append(filtered, table)
		}
	}
	return filtered
}

func (f *MarkdownFormatter) GetFileExtension() string {
	return ".md"
}

// XMLFormatter formats output as XML
type XMLFormatter struct{}

// XMLRoot represents the root element for XML output
type XMLRoot struct {
	XMLName           xml.Name               `xml:"mysql_info"`
	ConnectionInfo    *XMLConnectionInfo     `xml:"connection_info,omitempty"`
	Tables            []XMLTable             `xml:"tables>table"`
	Users             []XMLUser              `xml:"users>user"`
	Routines          []XMLRoutine           `xml:"routines>routine"`
	Variables         []XMLVariable          `xml:"variables>variable"`
	Roles             []XMLRole              `xml:"roles>role"`
	Plugins           []XMLPlugin            `xml:"plugins>plugin"`
	Components        []XMLComponent         `xml:"components>component"`
	ReplicationInfo   *XMLReplicationInfo    `xml:"replication,omitempty"`
}

type XMLConnectionInfo struct {
	Host     string `xml:"host"`
	Port     string `xml:"port"`
	User     string `xml:"user"`
	Database string `xml:"database,omitempty"`
	Version  string `xml:"version"`
}

type XMLTable struct {
	Name          string    `xml:"name,attr"`
	Schema        string    `xml:"schema,attr"`
	Type          string    `xml:"type,attr"`
	Engine        string    `xml:"engine,omitempty"`
	AutoIncrement int64     `xml:"auto_increment,omitempty"`
	CreatedAt     time.Time `xml:"created_at,omitempty"`
	UpdatedAt     time.Time `xml:"updated_at,omitempty"`
	Collation     string    `xml:"collation,omitempty"`
	Charset       string    `xml:"charset,omitempty"`
	RowFormat     string    `xml:"row_format,omitempty"`
	Comment       string    `xml:"comment,omitempty"`
	CreateOptions string    `xml:"create_options,omitempty"`
	DDL           string    `xml:"ddl,omitempty"`
}

type XMLUser struct {
	User            string   `xml:"user,attr"`
	Host            string   `xml:"host,attr"`
	Plugin          string   `xml:"plugin,omitempty"`
	AccountLocked   string   `xml:"account_locked,omitempty"`
	PasswordExpired string   `xml:"password_expired,omitempty"`
	Grants          []string `xml:"grants>grant"`
}

type XMLRoutine struct {
	Name         string    `xml:"name,attr"`
	Schema       string    `xml:"schema,attr"`
	Type         string    `xml:"type,attr"`
	Definer      string    `xml:"definer,omitempty"`
	Created      time.Time `xml:"created,omitempty"`
	LastAltered  time.Time `xml:"last_altered,omitempty"`
	DataAccess   string    `xml:"data_access,omitempty"`
	SecurityType string    `xml:"security_type,omitempty"`
	Returns      string    `xml:"returns,omitempty"`
	Parameters   string    `xml:"parameters,omitempty"`
	Definition   string    `xml:"definition,omitempty"`
}

type XMLVariable struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value"`
}

type XMLRole struct {
	RoleName string   `xml:"role_name,attr"`
	RoleHost string   `xml:"role_host,attr"`
	Grants   []string `xml:"grants>grant"`
	Members  []string `xml:"members>member"`
}

type XMLPlugin struct {
	Name        string `xml:"name,attr"`
	Version     string `xml:"version,attr"`
	Status      string `xml:"status,attr"`
	Type        string `xml:"type,attr"`
	Library     string `xml:"library,omitempty"`
	Description string `xml:"description,omitempty"`
}

type XMLComponent struct {
	ComponentID      int    `xml:"component_id,attr"`
	ComponentGroupID int    `xml:"component_group_id,attr"`
	ComponentURN     string `xml:"component_urn"`
}

type XMLReplicationInfo struct {
	ServerID      int    `xml:"server_id,omitempty"`
	ServerUUID    string `xml:"server_uuid,omitempty"`
	LogBinEnabled bool   `xml:"log_bin_enabled,omitempty"`
	BinlogFormat  string `xml:"binlog_format,omitempty"`
	GTIDMode      string `xml:"gtid_mode,omitempty"`
}

func (f *XMLFormatter) Format(info *DatabaseInfo) (string, error) {
	var result strings.Builder
	result.WriteString(xml.Header)
	result.WriteString("<mysql_info>\n")

	// File Summary section
	result.WriteString("  <file_summary>\n")
	result.WriteString("    <description>")
	result.WriteString("This file contains comprehensive MySQL database information compiled for AI context analysis. ")
	result.WriteString("It includes schema definitions, account configurations, system variables, and other database metadata ")
	result.WriteString("consolidated into a single file for efficient processing.")
	result.WriteString("</description>\n")
	
	// Generate sections list
	sections := f.generateSectionsList(info)
	if len(sections) > 0 {
		result.WriteString("    <file_structure>\n")
		for _, section := range sections {
			result.WriteString(fmt.Sprintf("      <section>%s</section>\n", f.escapeXML(section)))
		}
		result.WriteString("    </file_structure>\n")
	}
	result.WriteString("  </file_summary>\n")

	// Variables
	if len(info.Variables) > 0 {
		result.WriteString("  <variables>\n")
		
		// Check if this is only-modified-variables mode
		hasExtendedInfo := len(info.Variables) > 0 && info.Variables[0].IsModified
		if hasExtendedInfo {
			// Double-check that all variables are actually modified
			for _, v := range info.Variables {
				if !v.IsModified {
					hasExtendedInfo = false
					break
				}
			}
		}
		
		for _, variable := range info.Variables {
			result.WriteString("    <variable>\n")
			result.WriteString(fmt.Sprintf("      <name>%s</name>\n", f.escapeXML(variable.Name)))
			result.WriteString(fmt.Sprintf("      <current_value>%s</current_value>\n", f.escapeXML(variable.CurrentValue)))
			
			if hasExtendedInfo {
				result.WriteString(fmt.Sprintf("      <default_value>%s</default_value>\n", f.escapeXML(variable.DefaultValue)))
				result.WriteString(fmt.Sprintf("      <source>%s</source>\n", f.escapeXML(variable.Source)))
			}
			
			result.WriteString("    </variable>\n")
		}
		result.WriteString("  </variables>\n")
	}

	// Tables (BASE TABLE only)
	baseTables := f.filterTables(info.Tables, "BASE TABLE")
	if len(baseTables) > 0 {
		result.WriteString("  <tables>\n")
		
		// Group by database if multiple databases
		dbMap := make(map[string][]TableInfo)
		for _, table := range baseTables {
			dbMap[table.Database] = append(dbMap[table.Database], table)
		}
		
		var databases []string
		for db := range dbMap {
			databases = append(databases, db)
		}
		sort.Strings(databases)
		
		for _, db := range databases {
			tables := dbMap[db]
			sort.Slice(tables, func(i, j int) bool {
				if tables[i].Schema != tables[j].Schema {
					return tables[i].Schema < tables[j].Schema
				}
				return tables[i].Name < tables[j].Name
			})
			
			for _, table := range tables {
				result.WriteString("    <table>\n")
				// Include database name if there are multiple databases
				if len(databases) > 1 {
					result.WriteString(fmt.Sprintf("      <database>%s</database>\n", f.escapeXML(table.Database)))
				}
				result.WriteString(fmt.Sprintf("      <name>%s.%s</name>\n", f.escapeXML(table.Schema), f.escapeXML(table.Name)))
				result.WriteString(fmt.Sprintf("      <engine>%s</engine>\n", f.escapeXML(table.Engine)))
			if table.AutoIncrement > 0 {
				result.WriteString(fmt.Sprintf("      <auto_increment>%d</auto_increment>\n", table.AutoIncrement))
			}
			if !table.CreatedAt.IsZero() {
				result.WriteString(fmt.Sprintf("      <created>%s</created>\n", table.CreatedAt.Format("2006-01-02 15:04:05")))
			}
			result.WriteString(fmt.Sprintf("      <collation>%s</collation>\n", f.escapeXML(table.Collation)))
			if table.Charset != "" {
				result.WriteString(fmt.Sprintf("      <charset>%s</charset>\n", f.escapeXML(table.Charset)))
			}
			if table.RowFormat != "" {
				result.WriteString(fmt.Sprintf("      <row_format>%s</row_format>\n", f.escapeXML(table.RowFormat)))
			}
			if table.DDL != "" {
				result.WriteString("      <ddl><![CDATA[")
				result.WriteString(table.DDL)
				result.WriteString("]]></ddl>\n")
			}
			result.WriteString("    </table>\n")
			}
		}
		result.WriteString("  </tables>\n")
	}

	// Views
	views := f.filterTables(info.Tables, "VIEW")
	if len(views) > 0 {
		result.WriteString("  <views>\n")
		for _, view := range views {
			if view.DDL != "" {
				result.WriteString("    <view>\n")
				result.WriteString(fmt.Sprintf("      <name>%s.%s</name>\n", f.escapeXML(view.Schema), f.escapeXML(view.Name)))
				result.WriteString("      <ddl><![CDATA[")
				result.WriteString(view.DDL)
				result.WriteString("]]></ddl>\n")
				result.WriteString("    </view>\n")
			}
		}
		result.WriteString("  </views>\n")
	}

	// Stored functions
	functions := f.filterRoutines(info.Routines, "FUNCTION")
	if len(functions) > 0 {
		result.WriteString("  <stored_functions>\n")
		for _, function := range functions {
			result.WriteString("    <function>\n")
			result.WriteString(fmt.Sprintf("      <name>%s.%s</name>\n", f.escapeXML(function.Schema), f.escapeXML(function.Name)))
			result.WriteString(fmt.Sprintf("      <security_type>%s</security_type>\n", f.escapeXML(function.SecurityType)))
			if !function.Created.IsZero() {
				result.WriteString(fmt.Sprintf("      <created>%s</created>\n", function.Created.Format("2006-01-02 15:04:05")))
			}
			if function.Definition != "" {
				result.WriteString("      <definition><![CDATA[")
				result.WriteString(function.Definition)
				result.WriteString("]]></definition>\n")
			}
			result.WriteString("    </function>\n")
		}
		result.WriteString("  </stored_functions>\n")
	}

	// Stored procedures
	procedures := f.filterRoutines(info.Routines, "PROCEDURE")
	if len(procedures) > 0 {
		result.WriteString("  <stored_procedures>\n")
		for _, procedure := range procedures {
			result.WriteString("    <procedure>\n")
			result.WriteString(fmt.Sprintf("      <name>%s.%s</name>\n", f.escapeXML(procedure.Schema), f.escapeXML(procedure.Name)))
			result.WriteString(fmt.Sprintf("      <security_type>%s</security_type>\n", f.escapeXML(procedure.SecurityType)))
			if !procedure.Created.IsZero() {
				result.WriteString(fmt.Sprintf("      <created>%s</created>\n", procedure.Created.Format("2006-01-02 15:04:05")))
			}
			if procedure.Definition != "" {
				result.WriteString("      <definition><![CDATA[")
				result.WriteString(procedure.Definition)
				result.WriteString("]]></definition>\n")
			}
			result.WriteString("    </procedure>\n")
		}
		result.WriteString("  </stored_procedures>\n")
	}

	// Roles (MySQL 8.0+)
	if len(info.Roles) > 0 {
		result.WriteString("  <user_roles>\n")
		for _, role := range info.Roles {
			result.WriteString("    <role>\n")
			result.WriteString(fmt.Sprintf("      <name>%s@%s</name>\n", f.escapeXML(role.RoleName), f.escapeXML(role.RoleHost)))
			if len(role.Grants) > 0 {
				result.WriteString("      <grants>\n")
				for _, grant := range role.Grants {
					result.WriteString(fmt.Sprintf("        <grant>%s</grant>\n", f.escapeXML(grant)))
				}
				result.WriteString("      </grants>\n")
			}
			result.WriteString("    </role>\n")
		}
		result.WriteString("  </user_roles>\n")
	}

	// Users
	if len(info.Users) > 0 {
		result.WriteString("  <users>\n")
		for _, user := range info.Users {
			result.WriteString("    <user>\n")
			result.WriteString(fmt.Sprintf("      <name>%s@%s</name>\n", f.escapeXML(user.User), f.escapeXML(user.Host)))
			result.WriteString(fmt.Sprintf("      <plugin>%s</plugin>\n", f.escapeXML(user.Plugin)))
			result.WriteString(fmt.Sprintf("      <account_locked>%s</account_locked>\n", user.AccountLocked))
			if len(user.Grants) > 0 {
				result.WriteString("      <grants>\n")
				for _, grant := range user.Grants {
					result.WriteString(fmt.Sprintf("        <grant>%s</grant>\n", f.escapeXML(grant)))
				}
				result.WriteString("      </grants>\n")
			}
			result.WriteString("    </user>\n")
		}
		result.WriteString("  </users>\n")
	}

	// Plugins
	if len(info.Plugins) > 0 {
		result.WriteString("  <plugins>\n")
		for _, plugin := range info.Plugins {
			result.WriteString("    <plugin>\n")
			result.WriteString(fmt.Sprintf("      <name>%s</name>\n", f.escapeXML(plugin.Name)))
			result.WriteString(fmt.Sprintf("      <status>%s</status>\n", f.escapeXML(plugin.Status)))
			result.WriteString(fmt.Sprintf("      <type>%s</type>\n", f.escapeXML(plugin.Type)))
			if plugin.Library != "" {
				result.WriteString(fmt.Sprintf("      <library>%s</library>\n", f.escapeXML(plugin.Library)))
			}
			result.WriteString(fmt.Sprintf("      <version>%s</version>\n", f.escapeXML(plugin.Version)))
			result.WriteString("    </plugin>\n")
		}
		result.WriteString("  </plugins>\n")
	}

	// Components (MySQL 8.0+)
	if len(info.Components) > 0 {
		result.WriteString("  <components>\n")
		for _, component := range info.Components {
			result.WriteString("    <component>\n")
			result.WriteString(fmt.Sprintf("      <id>%d</id>\n", component.ComponentID))
			result.WriteString(fmt.Sprintf("      <urn>%s</urn>\n", f.escapeXML(component.ComponentURN)))
			result.WriteString("    </component>\n")
		}
		result.WriteString("  </components>\n")
	}

	result.WriteString("</mysql_info>\n")
	return result.String(), nil
}

func (f *XMLFormatter) escapeXML(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, "\"", "&quot;")
	s = strings.ReplaceAll(s, "'", "&apos;")
	return s
}

// generateSectionsList creates a list of sections that will be included in the output
func (f *XMLFormatter) generateSectionsList(info *DatabaseInfo) []string {
	var sections []string
	
	if len(info.Variables) > 0 {
		sections = append(sections, "Variables - MySQL system variables and their current values")
	}
	if len(info.Tables) > 0 {
		sections = append(sections, "Tables - Database tables with metadata and DDL definitions")
	}
	if f.hasViews(info.Tables) {
		sections = append(sections, "View Details - Database views with their definitions")
	}
	functions := f.filterRoutines(info.Routines, "FUNCTION")
	if len(functions) > 0 {
		sections = append(sections, "Stored Functions - User-defined functions with their definitions")
	}
	procedures := f.filterRoutines(info.Routines, "PROCEDURE")
	if len(procedures) > 0 {
		sections = append(sections, "Stored Procedures - User-defined procedures with their definitions")
	}
	if len(info.Roles) > 0 {
		sections = append(sections, "User Roles - MySQL 8.0+ role definitions and assignments")
	}
	if len(info.Users) > 0 {
		sections = append(sections, "User Accounts - Database user accounts with privileges")
	}
	if len(info.Plugins) > 0 {
		sections = append(sections, "Plugins - Installed MySQL plugins and extensions")
	}
	if len(info.Components) > 0 {
		sections = append(sections, "Components - MySQL 8.0+ components")
	}
	if info.ReplicationInfo != nil {
		sections = append(sections, "Replication Info - MySQL replication configuration and status")
	}
	
	return sections
}

func (f *XMLFormatter) hasViews(tables []TableInfo) bool {
	for _, table := range tables {
		if table.Type == "VIEW" {
			return true
		}
	}
	return false
}

func (f *XMLFormatter) filterRoutines(routines []RoutineInfo, routineType string) []RoutineInfo {
	var filtered []RoutineInfo
	for _, routine := range routines {
		if routine.Type == routineType {
			filtered = append(filtered, routine)
		}
	}
	return filtered
}

func (f *XMLFormatter) filterTables(tables []TableInfo, tableType string) []TableInfo {
	var filtered []TableInfo
	for _, table := range tables {
		if table.Type == tableType {
			filtered = append(filtered, table)
		}
	}
	return filtered
}

func (f *XMLFormatter) convertToXML(info *DatabaseInfo) *XMLRoot {
	xmlRoot := &XMLRoot{}

	// Convert connection info
	if info.ConnectionInfo != nil {
		xmlRoot.ConnectionInfo = &XMLConnectionInfo{
			Host:     info.ConnectionInfo.Host,
			Port:     info.ConnectionInfo.Port,
			User:     info.ConnectionInfo.User,
			Database: info.ConnectionInfo.Database,
			Version:  info.ConnectionInfo.Version,
		}
	}

	// Convert tables
	for _, table := range info.Tables {
		xmlRoot.Tables = append(xmlRoot.Tables, XMLTable{
			Name:          table.Name,
			Schema:        table.Schema,
			Type:          table.Type,
			Engine:        table.Engine,
			AutoIncrement: table.AutoIncrement,
			CreatedAt:     table.CreatedAt,
			UpdatedAt:     table.UpdatedAt,
			Collation:     table.Collation,
			Charset:       table.Charset,
			RowFormat:     table.RowFormat,
			Comment:       table.Comment,
			CreateOptions: table.CreateOptions,
			DDL:           table.DDL,
		})
	}

	// Convert users
	for _, user := range info.Users {
		xmlRoot.Users = append(xmlRoot.Users, XMLUser{
			User:            user.User,
			Host:            user.Host,
			Plugin:          user.Plugin,
			AccountLocked:   user.AccountLocked,
			PasswordExpired: user.PasswordExpired,
			Grants:          user.Grants,
		})
	}

	// Convert routines
	for _, routine := range info.Routines {
		xmlRoot.Routines = append(xmlRoot.Routines, XMLRoutine{
			Name:         routine.Name,
			Schema:       routine.Schema,
			Type:         routine.Type,
			Definer:      routine.Definer,
			Created:      routine.Created,
			LastAltered:  routine.LastAltered,
			DataAccess:   routine.DataAccess,
			SecurityType: routine.SecurityType,
			Returns:      routine.Returns,
			Parameters:   routine.Parameters,
			Definition:   routine.Definition,
		})
	}

	// Convert variables
	for _, variable := range info.Variables {
		xmlRoot.Variables = append(xmlRoot.Variables, XMLVariable{
			Name:  variable.Name,
			Value: variable.CurrentValue,
		})
	}

	// Convert roles
	for _, role := range info.Roles {
		xmlRoot.Roles = append(xmlRoot.Roles, XMLRole{
			RoleName: role.RoleName,
			RoleHost: role.RoleHost,
			Grants:   role.Grants,
			Members:  role.Members,
		})
	}

	// Convert plugins
	for _, plugin := range info.Plugins {
		xmlRoot.Plugins = append(xmlRoot.Plugins, XMLPlugin{
			Name:        plugin.Name,
			Version:     plugin.Version,
			Status:      plugin.Status,
			Type:        plugin.Type,
			Library:     plugin.Library,
			Description: plugin.Description,
		})
	}

	// Convert components
	for _, component := range info.Components {
		xmlRoot.Components = append(xmlRoot.Components, XMLComponent{
			ComponentID:      component.ComponentID,
			ComponentGroupID: component.ComponentGroupID,
			ComponentURN:     component.ComponentURN,
		})
	}

	// Convert replication info
	if info.ReplicationInfo != nil && info.ReplicationInfo.ReplicationStatus != nil {
		xmlRoot.ReplicationInfo = &XMLReplicationInfo{
			ServerID:      info.ReplicationInfo.ReplicationStatus.ServerID,
			ServerUUID:    info.ReplicationInfo.ReplicationStatus.ServerUUID,
			LogBinEnabled: info.ReplicationInfo.ReplicationStatus.LogBinEnabled,
			BinlogFormat:  info.ReplicationInfo.ReplicationStatus.BinlogFormat,
			GTIDMode:      info.ReplicationInfo.ReplicationStatus.GTIDMode,
		}
	}

	return xmlRoot
}

func (f *XMLFormatter) GetFileExtension() string {
	return ".xml"
}

// PlaintextFormatter formats output as plain text
type PlaintextFormatter struct{}

func (f *PlaintextFormatter) Format(info *DatabaseInfo) (string, error) {
	var result strings.Builder

	// File Summary section
	result.WriteString("File Summary\n")
	result.WriteString("============\n\n")
	result.WriteString("This file contains comprehensive MySQL database information compiled for AI context analysis. ")
	result.WriteString("It includes schema definitions, account configurations, system variables, and other database metadata ")
	result.WriteString("consolidated into a single file for efficient processing.\n\n")
	
	// Generate sections list
	sections := f.generateSectionsList(info)
	if len(sections) > 0 {
		result.WriteString("File Structure\n")
		result.WriteString("--------------\n\n")
		for _, section := range sections {
			result.WriteString(fmt.Sprintf("- %s\n", section))
		}
		result.WriteString("\n")
	}

	// Variables
	if len(info.Variables) > 0 {
		result.WriteString("Variables\n")
		result.WriteString("=========\n\n")
		
		// Check if this is only-modified-variables mode
		hasExtendedInfo := len(info.Variables) > 0 && info.Variables[0].IsModified
		if hasExtendedInfo {
			// Double-check that all variables are actually modified
			for _, v := range info.Variables {
				if !v.IsModified {
					hasExtendedInfo = false
					break
				}
			}
		}
		
		if hasExtendedInfo {
			// 4-column format for -only-modified-variables
			result.WriteString(fmt.Sprintf("%-40s %-30s %-30s %s\n", "Variable Name", "Current Value", "Default Value", "Source"))
			result.WriteString(strings.Repeat("-", 130) + "\n")
			
			for _, variable := range info.Variables {
				result.WriteString(fmt.Sprintf("%-40s %-30s %-30s %s\n",
					variable.Name, variable.CurrentValue, variable.DefaultValue, variable.Source))
			}
		} else {
			// 2-column format for normal variables
			result.WriteString(fmt.Sprintf("%-50s %s\n", "Variable Name", "Current Value"))
			result.WriteString(strings.Repeat("-", 100) + "\n")
			
			for _, variable := range info.Variables {
				result.WriteString(fmt.Sprintf("%-50s %s\n",
					variable.Name, variable.CurrentValue))
			}
		}
		result.WriteString("\n")
	}

	// Tables (BASE TABLE only)
	baseTables := f.filterTables(info.Tables, "BASE TABLE")
	if len(baseTables) > 0 {
		result.WriteString("Tables\n")
		result.WriteString("======\n\n")
		
		// Group by database if multiple databases
		dbMap := make(map[string][]TableInfo)
		for _, table := range baseTables {
			dbMap[table.Database] = append(dbMap[table.Database], table)
		}
		
		var databases []string
		for db := range dbMap {
			databases = append(databases, db)
		}
		sort.Strings(databases)
		
		for _, db := range databases {
			tables := dbMap[db]
			sort.Slice(tables, func(i, j int) bool {
				if tables[i].Schema != tables[j].Schema {
					return tables[i].Schema < tables[j].Schema
				}
				return tables[i].Name < tables[j].Name
			})
			
			for _, table := range tables {
				// Include database name if there are multiple databases
				if len(databases) > 1 {
					result.WriteString(fmt.Sprintf("%s.%s.%s\n", table.Database, table.Schema, table.Name))
				} else {
					result.WriteString(fmt.Sprintf("%s.%s\n", table.Schema, table.Name))
				}
				if table.Engine != "" {
				result.WriteString(fmt.Sprintf("  Engine: %s\n", table.Engine))
			}
			if table.AutoIncrement > 0 {
				result.WriteString(fmt.Sprintf("  Auto Increment: %d\n", table.AutoIncrement))
			}
			if !table.CreatedAt.IsZero() {
				result.WriteString(fmt.Sprintf("  Created: %s\n", table.CreatedAt.Format("2006-01-02 15:04:05")))
			}
			if table.Collation != "" {
				result.WriteString(fmt.Sprintf("  Collation: %s\n", table.Collation))
			}
			if table.Charset != "" {
				result.WriteString(fmt.Sprintf("  Charset: %s\n", table.Charset))
			}
			if table.RowFormat != "" {
				result.WriteString(fmt.Sprintf("  Row Format: %s\n", table.RowFormat))
			}
			if table.DDL != "" {
				result.WriteString("  DDL:\n")
				result.WriteString(fmt.Sprintf("    %s\n", strings.ReplaceAll(table.DDL, "\n", "\n    ")))
			}
			result.WriteString("\n")
			}
		}
	}

	// Views
	views := f.filterTables(info.Tables, "VIEW")
	if len(views) > 0 {
		result.WriteString("View Details\n")
		result.WriteString("============\n\n")
		for _, view := range views {
			if view.DDL != "" {
				result.WriteString(fmt.Sprintf("%s.%s\n", view.Schema, view.Name))
				result.WriteString("  DDL:\n")
				result.WriteString(fmt.Sprintf("    %s\n", strings.ReplaceAll(view.DDL, "\n", "\n    ")))
				result.WriteString("\n")
			}
		}
	}

	// Stored functions
	functions := f.filterRoutines(info.Routines, "FUNCTION")
	if len(functions) > 0 {
		result.WriteString("Stored Functions\n")
		result.WriteString("================\n\n")
		for _, function := range functions {
			result.WriteString(fmt.Sprintf("%s.%s\n", function.Schema, function.Name))
			result.WriteString(fmt.Sprintf("  Specific Name: %s\n", function.Name))
			result.WriteString("  Routine Catalog: def\n")
			result.WriteString("  Routine Body: SQL\n")
			result.WriteString("  Is Deterministic: YES\n")
			if function.DataAccess != "" {
				result.WriteString(fmt.Sprintf("  SQL Data Access: %s\n", function.DataAccess))
			}
			if function.SecurityType != "" {
				result.WriteString(fmt.Sprintf("  Security Type: %s\n", function.SecurityType))
			}
			if !function.Created.IsZero() {
				result.WriteString(fmt.Sprintf("  Created: %s\n", function.Created.Format("2006-01-02 15:04:05")))
			}
			if function.Definition != "" {
				result.WriteString("  Definition:\n")
				result.WriteString(fmt.Sprintf("    %s\n", strings.ReplaceAll(function.Definition, "\n", "\n    ")))
			}
			result.WriteString("\n")
		}
	}

	// Stored procedures
	procedures := f.filterRoutines(info.Routines, "PROCEDURE")
	if len(procedures) > 0 {
		result.WriteString("Stored Procedures\n")
		result.WriteString("=================\n\n")
		for _, procedure := range procedures {
			result.WriteString(fmt.Sprintf("%s.%s\n", procedure.Schema, procedure.Name))
			result.WriteString(fmt.Sprintf("  Specific Name: %s\n", procedure.Name))
			result.WriteString("  Routine Catalog: def\n")
			result.WriteString("  Routine Body: SQL\n")
			result.WriteString("  Is Deterministic: YES\n")
			if procedure.DataAccess != "" {
				result.WriteString(fmt.Sprintf("  SQL Data Access: %s\n", procedure.DataAccess))
			}
			if procedure.SecurityType != "" {
				result.WriteString(fmt.Sprintf("  Security Type: %s\n", procedure.SecurityType))
			}
			if !procedure.Created.IsZero() {
				result.WriteString(fmt.Sprintf("  Created: %s\n", procedure.Created.Format("2006-01-02 15:04:05")))
			}
			if procedure.Definition != "" {
				result.WriteString("  Definition:\n")
				result.WriteString(fmt.Sprintf("    %s\n", strings.ReplaceAll(procedure.Definition, "\n", "\n    ")))
			}
			result.WriteString("\n")
		}
	}

	// Roles (MySQL 8.0+)
	if len(info.Roles) > 0 {
		result.WriteString("User Roles (MySQL 8.0+)\n")
		result.WriteString("========================\n\n")
		for _, role := range info.Roles {
			result.WriteString(fmt.Sprintf("%s@%s\n", role.RoleName, role.RoleHost))
			if len(role.Grants) > 0 {
				for _, grant := range role.Grants {
					result.WriteString(fmt.Sprintf("  %s\n", grant))
				}
			}
			if len(role.Members) > 0 {
				result.WriteString("  Members:\n")
				for _, member := range role.Members {
					result.WriteString(fmt.Sprintf("    %s\n", member))
				}
			}
			result.WriteString("\n")
		}
	}

	// Users
	if len(info.Users) > 0 {
		result.WriteString("User List\n")
		result.WriteString("=========\n\n")
		for _, user := range info.Users {
			result.WriteString(fmt.Sprintf("%s@%s\n", user.User, user.Host))
			result.WriteString(strings.Repeat("-", len(user.User+"@"+user.Host)) + "\n")
			result.WriteString(fmt.Sprintf("Plugin: %s\n", user.Plugin))
			result.WriteString(fmt.Sprintf("Account Locked: %s\n", user.AccountLocked))
			if len(user.Grants) > 0 {
				result.WriteString("Grants:\n")
				for _, grant := range user.Grants {
					result.WriteString(fmt.Sprintf("  %s\n", grant))
				}
			}
			result.WriteString("\n")
		}
	}

	// Plugins
	if len(info.Plugins) > 0 {
		result.WriteString("Plugins\n")
		result.WriteString("=======\n\n")
		result.WriteString(fmt.Sprintf("%-25s %-15s %-20s %-10s %s\n", "Plugin Name", "Status", "Type", "Version", "Library"))
		result.WriteString(strings.Repeat("-", 100) + "\n")
		for _, plugin := range info.Plugins {
			library := plugin.Library
			if library == "" {
				library = "-"
			}
			result.WriteString(fmt.Sprintf("%-25s %-15s %-20s %-10s %s\n",
				plugin.Name, plugin.Status, plugin.Type, plugin.Version, library))
		}
		result.WriteString("\n")
	}

	// Components (MySQL 8.0+)
	if len(info.Components) > 0 {
		result.WriteString("Components (MySQL 8.0+)\n")
		result.WriteString("=======================\n\n")
		for _, component := range info.Components {
			result.WriteString(fmt.Sprintf("Component ID: %d\n", component.ComponentID))
			result.WriteString(fmt.Sprintf("Component URN: %s\n", component.ComponentURN))
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}

func (f *PlaintextFormatter) filterTables(tables []TableInfo, tableType string) []TableInfo {
	var filtered []TableInfo
	for _, table := range tables {
		if table.Type == tableType {
			filtered = append(filtered, table)
		}
	}
	return filtered
}

func (f *PlaintextFormatter) filterRoutines(routines []RoutineInfo, routineType string) []RoutineInfo {
	var filtered []RoutineInfo
	for _, routine := range routines {
		if routine.Type == routineType {
			filtered = append(filtered, routine)
		}
	}
	return filtered
}

func (f *PlaintextFormatter) GetFileExtension() string {
	return ".txt"
}

// generateSectionsList creates a list of sections that will be included in the output
func (f *PlaintextFormatter) generateSectionsList(info *DatabaseInfo) []string {
	var sections []string
	
	if len(info.Variables) > 0 {
		sections = append(sections, "Variables - MySQL system variables and their current values")
	}
	if len(info.Tables) > 0 {
		sections = append(sections, "Tables - Database tables with metadata and DDL definitions")
	}
	if f.hasViews(info.Tables) {
		sections = append(sections, "View Details - Database views with their definitions")
	}
	functions := f.filterRoutines(info.Routines, "FUNCTION")
	if len(functions) > 0 {
		sections = append(sections, "Stored Functions - User-defined functions with their definitions")
	}
	procedures := f.filterRoutines(info.Routines, "PROCEDURE")
	if len(procedures) > 0 {
		sections = append(sections, "Stored Procedures - User-defined procedures with their definitions")
	}
	if len(info.Roles) > 0 {
		sections = append(sections, "User Roles - MySQL 8.0+ role definitions and assignments")
	}
	if len(info.Users) > 0 {
		sections = append(sections, "User Accounts - Database user accounts with privileges")
	}
	if len(info.Plugins) > 0 {
		sections = append(sections, "Plugins - Installed MySQL plugins and extensions")
	}
	if len(info.Components) > 0 {
		sections = append(sections, "Components - MySQL 8.0+ components")
	}
	if info.ReplicationInfo != nil {
		sections = append(sections, "Replication Info - MySQL replication configuration and status")
	}
	
	return sections
}

func (f *PlaintextFormatter) hasViews(tables []TableInfo) bool {
	for _, table := range tables {
		if table.Type == "VIEW" {
			return true
		}
	}
	return false
}