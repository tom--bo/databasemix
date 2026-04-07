-- Create test databases and users for PostgreSQL

-- Create test users
CREATE USER testuser WITH PASSWORD 'testpass';
CREATE USER readonly WITH PASSWORD 'readpass';
CREATE USER admin WITH PASSWORD 'adminpass' CREATEDB CREATEROLE;

-- Create second test database
CREATE DATABASE testdb2 OWNER postgres;

-- Grant privileges
GRANT ALL PRIVILEGES ON DATABASE testdb TO testuser;
GRANT CONNECT ON DATABASE testdb TO readonly;
GRANT ALL PRIVILEGES ON DATABASE testdb TO admin;

-- Grant schema privileges (applied to testdb via init scripts)
GRANT ALL ON SCHEMA public TO testuser;
GRANT USAGE ON SCHEMA public TO readonly;
GRANT ALL ON SCHEMA public TO admin;
