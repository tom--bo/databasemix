-- Additional grants for MySQL 8.4+ specific tables
-- MySQL 8.4 removed mysql_native_password, uses caching_sha2_password by default

-- Recreate the user with MySQL 8.4 compatible authentication
-- The user may have been created by 07-limited-user.sql but with incompatible plugin
DROP USER IF EXISTS 'dbmix_limited'@'%';
CREATE USER 'dbmix_limited'@'%' IDENTIFIED WITH caching_sha2_password BY 'limited_pass';

-- Grant SELECT on the target databases
GRANT SELECT ON testdb.* TO 'dbmix_limited'@'%';
GRANT SELECT ON testdb2.* TO 'dbmix_limited'@'%';

-- Grant SELECT on system tables needed by dbmix
GRANT SELECT ON mysql.user TO 'dbmix_limited'@'%';

-- Grant SELECT on MySQL 8.0+ specific tables
GRANT SELECT ON mysql.role_edges TO 'dbmix_limited'@'%';
GRANT SELECT ON mysql.component TO 'dbmix_limited'@'%';

-- Grant SELECT on performance_schema tables
GRANT SELECT ON performance_schema.* TO 'dbmix_limited'@'%';

FLUSH PRIVILEGES;
