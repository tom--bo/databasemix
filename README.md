# DatabaseMix

`DatabaseMix` summarizes database information such as ACL, Table definition, variables, etc...
(90% of this code is written by AI)

## Features

- **Database support**: MySQL 5.7+, 8.0+, 8.4+ / PostgreSQL 16, 17, 18
- **Retrieves comprehensive database information**:
  - Database version and connection details
  - Tables with metadata (engine, collation, charset, row format, auto_increment, etc.)
  - Table DDL statements (`CREATE TABLE`)
  - Views and their DDL statements (`CREATE VIEW`)
  - Stored functions and procedures with metadata and definitions
- **Security information**:
  - User accounts and their attributes
  - User privileges (`GRANTS`)
  - Roles and role grants (MySQL 8.0+ / PostgreSQL)
- **System configuration**:
  - Global variables (all or only modified with `-only-modified-variables`)
  - Installed plugins (MySQL) / Extensions (PostgreSQL)
  - Components (MySQL 8.0+)
  - Replication information (optional, MySQL only, with `-replication`)
- **Multiple output formats**: Markdown (default), XML, Plaintext
- **Environment variable support**:
  - MySQL: `MYSQL_HOST`, `MYSQL_PORT`, `MYSQL_USER`, `MYSQL_PASSWORD`, `MYSQL_DATABASE`
  - PostgreSQL: `PGHOST`, `PGPORT`, `PGUSER`, `PGPASSWORD`, `PGDATABASE`

## Requirements

- Go 1.22 or later
- Target database: MySQL 5.7 / 8.0 / 8.4, PostgreSQL 16 / 17 / 18

## Installation

```bash
go install github.com/tom--bo/databasemix@latest
```

## Building from Source

```bash
git clone https://github.com/tom--bo/databasemix
cd databasemix
make build
```

## Usage

```bash
# MySQL
./databasemix -type mysql -host localhost -port 3306 -user root -password yourpassword

# PostgreSQL
./databasemix -type postgres -host localhost -port 5432 -user postgres -password yourpassword -database mydb
```

When `-type` is omitted, the database type is auto-detected from the port number (3306 → MySQL, 5432 → PostgreSQL).

### Command Line Arguments

| Flag | Default | Description |
|------|---------|-------------|
| `-type` | (auto) | Database type: `mysql`, `postgres` (auto-detected from port if omitted) |
| `-host` | `localhost` | Database host |
| `-port` | `3306`/`5432` | Database port (default depends on type) |
| `-user` | `root`/`postgres` | Database user (default depends on type) |
| `-password` | | Database password |
| `-database` | | Database name (optional; all accessible databases if omitted) |
| `-replication` | `false` | Include replication information (MySQL only) |
| `-except-tables` | `false` | Exclude tables and views |
| `-except-stored-procedures` | `false` | Exclude stored procedures and functions |
| `-except-variables` | `false` | Exclude variables/configuration parameters |
| `-only-modified-variables` | `false` | Show only modified variables |
| `-except-users` | `false` | Exclude user accounts |
| `-except-roles` | `false` | Exclude user roles |
| `-except-plugins` | `false` | Exclude installed plugins (MySQL only) |
| `-except-extensions` | `false` | Exclude installed extensions (PostgreSQL only) |
| `-format` | `markdown` | Output format (`markdown`/`xml`/`plaintext`) |
| `-outfile` | `dbmix-output` | Output filename (extension added based on format) |

## Output

The output file contains:

1. **File Summary** - Database type, version, and file structure overview
2. **Variables** - Configuration parameters (optionally only modified ones)
3. **Tables** - Metadata and full DDL
4. **Views** - View definitions with DDL
5. **Stored Functions & Procedures** - Definitions with metadata
6. **User Accounts** - Usernames, authentication details, and privileges
7. **Roles** - Role definitions, privileges, and member assignments
8. **Plugins** (MySQL) / **Extensions** (PostgreSQL) - Installed plugins/extensions
9. **Replication Info** (MySQL, optional) - Replica status, semi-sync, group replication

## Testing

Docker containers are provided for testing against multiple database versions.

### Prerequisites

- Docker and Docker Compose
- GNU Make

### Make Targets

```bash
make build              # Build the binary
make test               # Run all tests (MySQL + PostgreSQL, all versions)
make test-mysql         # Run all MySQL tests (5.7, 8.0, 8.4)
make test-postgres      # Run all PostgreSQL tests (16, 17, 18)
make test-mysql-8.0     # Run test for a specific MySQL version
make test-postgres-17   # Run test for a specific PostgreSQL version
make containers-up      # Start all test containers
make containers-down    # Stop all test containers
make clean              # Remove binary and test output
make help               # Show all available targets
```

### Test Data

- **MySQL** (`test_containers/mysql-common/`): databases, tables, views, triggers, stored procedures/functions, sample data, multiple test users with different privilege levels
- **PostgreSQL** (`test_containers/postgres-common/`): databases, tables, views, functions, sample data, roles and users

Connection credentials: root(postgres)/rootpass, testuser/testpass, readonly/readpass, admin/adminpass

### Individual Container Management

Each version directory has a `run.sh` script for manual management:

```bash
cd test_containers/mysql-8.0
./run.sh start    # Start container
./run.sh stop     # Stop container
./run.sh shell    # Open database shell
./run.sh logs     # Show container logs
./run.sh status   # Show container status
./run.sh clean    # Remove container and volumes
```

## Feature Mapping Table with MySQL

| Feature | MySQL | PostgreSQL | Notes |
|---------|-------|------------|-------|
| List of tables | `information_schema.TABLES` | `information_schema.tables` | Almost the same |
| Table DDL | `SHOW CREATE TABLE` | SQL assembly | In PostgreSQL, it needs to be assembled |
| View DDL | `SHOW CREATE VIEW` | `pg_get_viewdef()` | |
| Users | `mysql.user` | `pg_roles (rolcanlogin=true)` | |
| Roles | `mysql.user` + `role_edges` | `pg_roles (rolcanlogin=false)` + `pg_auth_members` | |
| Privileges | `SHOW GRANTS FOR` | `information_schema.role_table_grants`, etc. | PostgreSQL uses multiple sources |
| Variables | `performance_schema.global_variables` | `pg_settings` | |
| Procedures | `information_schema.ROUTINES` | `pg_proc` + `pg_get_functiondef()` | |
| Plugins | `information_schema.PLUGINS` | N/A | MySQL only |
| Extensions | N/A | `pg_extension` | PostgreSQL only |
| Replication | `SHOW REPLICA STATUS`, etc. | `pg_stat_replication`, etc. | Planned for future support |
