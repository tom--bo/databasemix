-- Common database setup for all MySQL versions

-- Create test database
CREATE DATABASE IF NOT EXISTS testdb CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS testdb2 CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

-- Create basic test users
CREATE USER IF NOT EXISTS 'testuser'@'%' IDENTIFIED BY 'testpass';
CREATE USER IF NOT EXISTS 'readonly'@'%' IDENTIFIED BY 'readpass';
CREATE USER IF NOT EXISTS 'admin'@'%' IDENTIFIED BY 'adminpass';

-- Grant privileges
GRANT ALL PRIVILEGES ON testdb.* TO 'testuser'@'%';
GRANT ALL PRIVILEGES ON testdb2.* TO 'testuser'@'%';
GRANT SELECT ON *.* TO 'readonly'@'%';
GRANT ALL PRIVILEGES ON *.* TO 'admin'@'%' WITH GRANT OPTION;

-- Use testdb for subsequent operations
USE testdb;