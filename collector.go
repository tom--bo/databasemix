package main

// Collector is the common interface for database information collection.
// MySQLCollector and PostgreSQLCollector implement this interface.
type Collector interface {
	CollectAll() (*DatabaseInfo, error)
}
