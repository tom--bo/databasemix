-- Create a limited privilege user for testing minimum required privileges
-- This user should have only the minimum permissions needed for dbmix to work

-- Create dbmix_limited user with minimum privileges
CREATE USER IF NOT EXISTS 'dbmix_limited'@'%' IDENTIFIED BY 'limited_pass';

-- Grant SELECT on the target database
GRANT SELECT ON testdb.* TO 'dbmix_limited'@'%';
GRANT SELECT ON testdb2.* TO 'dbmix_limited'@'%';

-- Grant SELECT on system tables needed by dbmix
GRANT SELECT ON mysql.user TO 'dbmix_limited'@'%';

-- Grant SELECT on performance_schema tables (using wildcard to cover all needed tables)
GRANT SELECT ON performance_schema.* TO 'dbmix_limited'@'%';

-- Optional: Uncomment for replication testing
-- GRANT REPLICATION CLIENT ON *.* TO 'dbmix_limited'@'%';

FLUSH PRIVILEGES;
