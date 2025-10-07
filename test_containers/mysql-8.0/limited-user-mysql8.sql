-- Additional grants for MySQL 8.0+ specific tables

-- Grant SELECT on MySQL 8.0+ specific tables
GRANT SELECT ON mysql.role_edges TO 'dbmix_limited'@'%';
GRANT SELECT ON mysql.component TO 'dbmix_limited'@'%';

FLUSH PRIVILEGES;
