-- MySQL 8.0+ specific: Role setup

-- Create roles
CREATE ROLE IF NOT EXISTS 'app_read';
CREATE ROLE IF NOT EXISTS 'app_write';
CREATE ROLE IF NOT EXISTS 'app_admin';
CREATE ROLE IF NOT EXISTS 'developer';
CREATE ROLE IF NOT EXISTS 'analyst';

-- Grant privileges to roles
GRANT SELECT ON *.* TO 'app_read';
GRANT SELECT, INSERT, UPDATE, DELETE ON testdb.* TO 'app_write';
GRANT ALL PRIVILEGES ON testdb.* TO 'app_admin';
GRANT SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, INDEX, ALTER ON testdb.* TO 'developer';
GRANT SELECT ON *.* TO 'analyst';
GRANT EXECUTE ON testdb.* TO 'analyst';

-- Create users with roles
CREATE USER IF NOT EXISTS 'app_user1'@'%' IDENTIFIED BY 'apppass1';
CREATE USER IF NOT EXISTS 'app_user2'@'%' IDENTIFIED BY 'apppass2';
CREATE USER IF NOT EXISTS 'dev_user'@'%' IDENTIFIED BY 'devpass';
CREATE USER IF NOT EXISTS 'analyst_user'@'%' IDENTIFIED BY 'analystpass';

-- Grant roles to users
GRANT 'app_read' TO 'app_user1'@'%';
GRANT 'app_read', 'app_write' TO 'app_user2'@'%';
GRANT 'developer' TO 'dev_user'@'%';
GRANT 'analyst' TO 'analyst_user'@'%';

-- Set default roles
SET DEFAULT ROLE ALL TO 
    'app_user1'@'%',
    'app_user2'@'%',
    'dev_user'@'%',
    'analyst_user'@'%';

-- Create users with resource limits
CREATE USER IF NOT EXISTS 'limited_user'@'%' IDENTIFIED BY 'limitpass'
    WITH MAX_QUERIES_PER_HOUR 1000
    MAX_CONNECTIONS_PER_HOUR 100
    MAX_USER_CONNECTIONS 5;

GRANT 'app_read' TO 'limited_user'@'%';
SET DEFAULT ROLE ALL TO 'limited_user'@'%';

FLUSH PRIVILEGES;