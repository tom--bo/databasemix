-- Plugin and component testing
-- Note: This script will work for MySQL 8.0+ where components are available

-- Install some common plugins (these are usually already available)
-- Authentication plugins are typically already loaded

-- For MySQL 8.0+, install some components if available
-- These may not work in all environments, but will demonstrate component collection

-- Set some plugin-related variables for testing
-- Note: validate_password plugin may not be installed by default
-- SET GLOBAL validate_password.policy = 'MEDIUM';
-- SET GLOBAL validate_password.length = 8;

-- Note: The actual installation of plugins/components depends on MySQL version
-- and available plugin files. The dbmix tool will collect whatever is available.