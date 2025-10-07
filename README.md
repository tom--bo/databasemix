# DatabaseMix 

A Go utility to connect to MySQL 5.7+ instances and retrieve comprehensive database information, outputting it to markdown, XML, or plaintext format.
(90% of this codes are written by AI)

## Features

- **MySQL version support**:
  - MySQL 5.7+
  - MySQL 8.0+
  - MySQL 8.4+
- **Retrieves comprehensive database information**:
  - Database version and connection details
  - Tables with metadata (engine/access method, collation, charset, row format, etc.)
  - Table DDL statements (CREATE TABLE)
  - Views and their DDL statements (CREATE VIEW)
  - Stored functions with metadata and definitions
  - Stored procedures with metadata and definitions
- **Security information**:
  - User accounts and their attributes
  - User privileges (GRANTS)
  - Roles and role grants (MySQL 8.0+)
- **System configuration**:
  - Global variables (all or only modified with `-only-modified-variables`)
- **Multiple output formats**:
  - Markdown (default)
  - XML
  - Plaintext

## Requirements

- Go 1.22 or later
- Target database:
  - MySQL 5.7
  - MySQL 8.0
  - MySQL 8.4

## Installation

```bash
go install github.com/tom--bo/databasemix@v0.1.0
```

## Usage

### Command Line Arguments

```bash
./databasemix -host=localhost -port=3306 -user=root -password=yourpassword
```

All command line arguments:

- `-host`: Database host (default: "localhost")
- `-port`: MySQL port (default: "3306")
- `-user`: MySQL user (default: "root")
- `-password`: Database password
- `-database`: Database name (optional, if not specified all accessible databases will be analyzed)
- `-replication`: Include replication information (default: false)
- `-except-tables`: Exclude tables and views (default: false)
- `-except-stored-procedures`: Exclude stored procedures and functions (default: false)
- `-except-variables`: Exclude variables/configuration parameters (default: false)
- `-only-modified-variables`: Show only modified variables (MySQL only, default: show all)
- `-except-users`: Exclude user accounts (default: false)
- `-except-roles`: Exclude user roles (default: false)
- `-format`: Output format (markdown/xml/plaintext, default: "markdown")
- `-outfile`: Output filename (default: "dbmix-output" with appropriate extension based on format)


## Output

The program generates an output file (default: `dbmix-output.md`, `dbmix-output.xml`, or `dbmix-output.txt` depending on the format) with comprehensive information about your database server. The output is organized into several sections:

### File Summary
- Database type and version
- Comprehensive overview of the file structure

### Tables
For each table:
- Full table name with schema
- Metadata:
  - Type, engine, charset, collation, row format, auto_increment, creation time
- Complete DDL (CREATE TABLE statement)

### Views
For each view:
- Full view name with schema
- Complete DDL (CREATE VIEW statement)

### Variables/Configuration Parameters
- Global variables (all or only modified with `-only-modified-variables`)
- Variable/parameter name and current value
- When using `-only-modified-variables` (MySQL): also shows default value and source

### User Accounts
For each user account:
- Username and host
- Authentication details (plugin, SSL type), account status
- Complete list of privileges (GRANT statements)

### Stored Functions and Procedures
For each routine:
- Name and schema
- Metadata (creation time, character set, deterministic status)
- Complete definition

### Roles
- MySQL 8.0+: Role name and host, mandatory/default status
- Privileges granted to the role
- Members who have been granted this role


```markdown
# File Summary

This file contains comprehensive MySQL database information compiled for AI context analysis...

**Database Type**: MySQL
**Database Version**: 8.0.39

## File Structure

- Variables - MySQL system variables and their current values
- Tables - Database tables with metadata and DDL definitions
- View Details - Database views with their definitions
- Stored Functions - User-defined functions with their definitions
- User Roles - MySQL 8.0+ role definitions and assignments
- User Accounts - Database user accounts with privileges
- Plugins - Installed MySQL plugins

# Configuration Parameters

| Variable Name | Current Value |
|---------------|---------------|
| max_connections | 100 |
| shared_buffers | 16384 |
| work_mem | 4096 |

# Tables

## public.employees

### Metadata
- Type: BASE TABLE
- Engine: InnoDB

### DDL
```sql
-- Table DDL not available in simplified collector
```

# User List

## admin_user@

- Grants:
  - GRANT developer TO admin_user

```

## Testing

The repository includes Docker containers for testing DB Mix against different database versions:

### Available Test Containers

**MySQL versions:**
- MySQL 5.7 (port 3357)
- MySQL 8.0 (port 3380)
- MySQL 8.4 (port 3384)

### Quick Start Examples

```bash
# MySQL 8.0
cd test_containers/mysql-8.0 && ./run.sh start && cd ../..
./databasemix -host=localhost -port=3380 -user=root -password=rootpass -database=testdb -outfile=mysql_report

# See all available test commands
./sample_command.sh

```
## Building from Source

```bash
# Clone the repository
git clone https://github.com/tom--bo/databasemix
cd databasemix

# Build
go build -o databasemix .

# Run
./databasemix --help
```


