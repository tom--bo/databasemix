-- Roles setup for PostgreSQL test environment

-- Create roles (non-login)
CREATE ROLE app_read;
CREATE ROLE app_write;
CREATE ROLE app_admin;
CREATE ROLE developer;
CREATE ROLE analyst;

-- Grant privileges to roles
GRANT SELECT ON ALL TABLES IN SCHEMA public TO app_read;
GRANT INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO app_write;
GRANT ALL ON ALL TABLES IN SCHEMA public TO app_admin;
GRANT ALL ON ALL SEQUENCES IN SCHEMA public TO app_admin;

-- Assign roles to roles (hierarchy)
GRANT app_read TO app_write;
GRANT app_write TO app_admin;
GRANT app_read TO analyst;
GRANT app_admin TO developer;

-- Assign roles to users
GRANT app_read TO readonly;
GRANT app_write TO testuser;
GRANT developer TO admin;
