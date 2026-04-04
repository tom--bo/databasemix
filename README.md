# DatabaseMix

A Go utility to connect to MySQL 5.7+ instances and retrieve comprehensive database information, outputting it to markdown, XML, or plaintext format.
(90% of this code is written by AI)

## Features

- **MySQL version support**: MySQL 5.7+, 8.0+, 8.4+ (including MariaDB and Percona detection)
- **Retrieves comprehensive database information**:
  - Database version and connection details
  - Tables with metadata (engine, collation, charset, row format, auto_increment, etc.)
  - Table DDL statements (`CREATE TABLE`)
  - Views and their DDL statements (`CREATE VIEW`)
  - Stored functions and procedures with metadata and definitions
  - Triggers
- **Security information**:
  - User accounts and their attributes
  - User privileges (`GRANTS`)
  - Roles and role grants (MySQL 8.0+)
- **System configuration**:
  - Global variables (all or only modified with `-only-modified-variables`)
  - Installed plugins (non-built-in only)
  - Components (MySQL 8.0+)
  - Replication information (optional, with `-replication`)
- **Multiple output formats**: Markdown (default), XML, Plaintext
- **Environment variable support**: `MYSQL_HOST`, `MYSQL_PORT`, `MYSQL_USER`, `MYSQL_PASSWORD`, `MYSQL_DATABASE`

## Requirements

- Go 1.22 or later
- Target database: MySQL 5.7 / 8.0 / 8.4

## Installation

```bash
go install github.com/tom--bo/databasemix@v0.1.0
```

## Building from Source

```bash
git clone https://github.com/tom--bo/databasemix
cd databasemix
go build -o databasemix .
```

## Usage

```bash
./databasemix -host=localhost -port=3306 -user=root -password=yourpassword
```

### Command Line Arguments

| Flag | Default | Description |
|------|---------|-------------|
| `-host` | `localhost` | Database host |
| `-port` | `3306` | MySQL port |
| `-user` | `root` | MySQL user |
| `-password` | | Database password |
| `-database` | | Database name (optional; all accessible databases if omitted) |
| `-replication` | `false` | Include replication information |
| `-except-tables` | `false` | Exclude tables and views |
| `-except-stored-procedures` | `false` | Exclude stored procedures and functions |
| `-except-variables` | `false` | Exclude variables/configuration parameters |
| `-only-modified-variables` | `false` | Show only modified variables (MySQL only) |
| `-except-users` | `false` | Exclude user accounts |
| `-except-roles` | `false` | Exclude user roles |
| `-except-plugins` | `false` | Exclude installed plugins |
| `-format` | `markdown` | Output format (`markdown`/`xml`/`plaintext`) |
| `-outfile` | `dbmix-output` | Output filename (extension added based on format) |

## Output

The output file (e.g., `dbmix-output.md`) contains:

1. **File Summary** - Database type, version, and file structure overview
2. **Variables** - Global configuration parameters (optionally only modified ones with default values and source)
3. **Tables** - Metadata (engine, charset, collation, row format, etc.) and full DDL
4. **Views** - View definitions with DDL
5. **Stored Functions & Procedures** - Definitions with metadata
6. **User Accounts** - Usernames, authentication details, and `GRANT` statements
7. **Roles** (MySQL 8.0+) - Role definitions, privileges, and member assignments
8. **Plugins** - Non-built-in installed plugins
9. **Replication Info** (optional) - Replica status, semi-sync, group replication

### Output Example

```markdown
# File Summary

**Database Type**: MySQL
**Database Version**: 8.0.45

# Variables

| Variable Name | Current Value |
|---------------|---------------|
| max_connections | 151 |

# Tables

## testdb.users

### Metadata
- Type: BASE TABLE
- Engine: InnoDB

### DDL
​```sql
CREATE TABLE `users` (...)
​```

# User Accounts

## root@localhost
- Grants:
  - GRANT ALL PRIVILEGES ON *.* TO `root`@`localhost`
```

## Testing

Docker containers are provided for testing against multiple MySQL versions.

### Prerequisites

- Docker and Docker Compose

### Available Test Containers

| Version | Port | Directory |
|---------|------|-----------|
| MySQL 5.7 | 3357 | `test_containers/mysql-5.7/` |
| MySQL 8.0 | 3380 | `test_containers/mysql-8.0/` |
| MySQL 8.4 | 3384 | `test_containers/mysql-8.4/` |

### Test Container Management

```bash
# Start a container
cd test_containers/mysql-8.0 && ./run.sh start

# Run databasemix against it
./databasemix -host=localhost -port=3380 -user=root -password=rootpass -database=testdb

# Other commands: stop, restart, logs, shell, status, clean
./run.sh stop
```

### Quick Test

```bash
cd test_containers
./quick-test.sh 8.0    # Accepts: 5.7, 8.0, 8.4
```

### Test Data

The `test_containers/common/` directory contains shared SQL init scripts that set up:
- 2 databases (`testdb`, `testdb2`) with tables, views, triggers
- Stored procedures and functions
- Sample data
- Multiple test users with different privilege levels

Version-specific setup (roles for 8.0+, auth plugin differences for 8.4) is in each version's directory.
