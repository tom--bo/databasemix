# MySQL Mix Test Containers

This directory contains Docker-based test environments for testing MySQL Mix against different MySQL versions.

## Directory Structure

```
test-containers/
├── common/              # Shared SQL scripts for all versions
│   ├── 01-database-setup.sql
│   ├── 02-tables-views.sql
│   └── 03-procedures-functions.sql
├── mysql-5.7/          # MySQL 5.7 specific
│   └── run.sh
├── mysql-8.0/          # MySQL 8.0 specific
│   ├── run.sh
│   └── roles-setup.sql
├── mysql-8.4/          # MySQL 8.4 specific
│   └── run.sh
└── test.sh         # Run tests
```

## Prerequisites

- Docker installed and running
- Go installed (for building mysqlmix)
- Port availability: 3357 (5.7), 3380 (8.0), 3384 (8.4)

## Usage

### Execute Tests
```bash
bash test.sh 8.4
```

This will:
1. Start each MySQL container
2. Run mysqlmix against target MySQL 
3. Generate version-specific reports (mysql_checkup_5.7.md, etc.)
4. Stop containers after testing
5. Display a summary of test results

### Test Individual Versions

Each version has its own script with these commands:

```bash
# MySQL 5.7
cd mysql-5.7
./run.sh start    # Start container
./run.sh test     # Run mysqlmix test
./run.sh connect  # Connect to MySQL CLI
./run.sh status   # Show container status
./run.sh logs     # Show container logs
./run.sh stop     # Stop and remove container

# MySQL 8.0
cd mysql-8.0
./run.sh start    # Start container with roles
./run.sh test     # Run mysqlmix test (includes role checking)
./run.sh connect  # Connect to MySQL CLI
./run.sh stop     # Stop and remove container

# MySQL 8.4
cd mysql-8.4
./run.sh start    # Start container with latest features
./run.sh test     # Run mysqlmix test
./run.sh connect  # Connect to MySQL CLI
./run.sh stop     # Stop and remove container
```

## Connection Details

All containers use:
- Root password: `rootpass`
- Test database: `testdb`
- Test users: `testuser`, `readonly`, `admin`

Ports:
- MySQL 5.7: `3357`
- MySQL 8.0: `3380`
- MySQL 8.4: `3384`



