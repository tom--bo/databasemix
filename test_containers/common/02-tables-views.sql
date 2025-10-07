-- Common table and view definitions for all MySQL versions

USE testdb;

-- Basic tables
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT TRUE,
    metadata JSON DEFAULT NULL,
    INDEX idx_email (email),
    INDEX idx_created (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE products (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    stock_quantity INT DEFAULT 0,
    category VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_category (category),
    INDEX idx_price (price)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE orders (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id INT NOT NULL,
    order_date DATETIME NOT NULL,
    total_amount DECIMAL(10,2) NOT NULL,
    status ENUM('pending', 'processing', 'completed', 'cancelled') DEFAULT 'pending',
    notes TEXT,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    INDEX idx_user_date (user_id, order_date),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE order_items (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    order_id BIGINT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    unit_price DECIMAL(10,2) NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id),
    INDEX idx_order (order_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Table with various data types
CREATE TABLE data_types_test (
    id INT PRIMARY KEY AUTO_INCREMENT,
    -- Numeric types
    tinyint_col TINYINT,
    smallint_col SMALLINT,
    mediumint_col MEDIUMINT,
    int_col INT,
    bigint_col BIGINT,
    decimal_col DECIMAL(10,2),
    float_col FLOAT,
    double_col DOUBLE,
    bit_col BIT(8),
    
    -- String types
    char_col CHAR(10),
    varchar_col VARCHAR(255),
    binary_col BINARY(16),
    varbinary_col VARBINARY(255),
    tinytext_col TINYTEXT,
    text_col TEXT,
    mediumtext_col MEDIUMTEXT,
    longtext_col LONGTEXT,
    
    -- Date and time types
    date_col DATE,
    time_col TIME,
    datetime_col DATETIME,
    timestamp_col TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    year_col YEAR,
    
    -- Other types
    enum_col ENUM('small', 'medium', 'large'),
    set_col SET('read', 'write', 'execute'),
    json_col JSON
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Partitioned table (supported in all versions)
CREATE TABLE logs (
    id BIGINT NOT NULL AUTO_INCREMENT,
    log_date DATE NOT NULL,
    level VARCHAR(10) NOT NULL,
    message TEXT,
    PRIMARY KEY (id, log_date)
) ENGINE=InnoDB
PARTITION BY RANGE (YEAR(log_date)) (
    PARTITION p2023 VALUES LESS THAN (2024),
    PARTITION p2024 VALUES LESS THAN (2025),
    PARTITION p2025 VALUES LESS THAN (2026)
);

-- Create views
CREATE VIEW active_users AS
SELECT id, username, email, created_at
FROM users
WHERE is_active = TRUE;

CREATE VIEW order_summary AS
SELECT 
    u.username,
    COUNT(o.id) as order_count,
    SUM(o.total_amount) as total_spent,
    MAX(o.order_date) as last_order_date
FROM users u
LEFT JOIN orders o ON u.id = o.user_id
GROUP BY u.id, u.username;

CREATE VIEW product_inventory AS
SELECT 
    p.id,
    p.name,
    p.category,
    p.price,
    p.stock_quantity,
    CASE 
        WHEN p.stock_quantity = 0 THEN 'Out of Stock'
        WHEN p.stock_quantity < 10 THEN 'Low Stock'
        ELSE 'In Stock'
    END as stock_status
FROM products p;

-- Insert sample data
INSERT INTO users (username, email, metadata) VALUES
('alice', 'alice@example.com', '{"role": "admin", "preferences": {"theme": "dark"}}'),
('bob', 'bob@example.com', '{"role": "user", "preferences": {"theme": "light"}}'),
('charlie', 'charlie@example.com', '{"role": "user", "preferences": {"notifications": true}}'),
('david', 'david@example.com', NULL),
('eve', 'eve@example.com', '{"role": "moderator"}');

INSERT INTO products (name, description, price, stock_quantity, category) VALUES
('Laptop Pro', 'High-performance laptop', 1299.99, 15, 'Electronics'),
('Wireless Mouse', 'Ergonomic wireless mouse', 29.99, 50, 'Electronics'),
('Coffee Maker', 'Automatic coffee maker', 89.99, 20, 'Appliances'),
('Desk Lamp', 'LED desk lamp', 39.99, 30, 'Furniture'),
('Notebook Set', 'Pack of 5 notebooks', 12.99, 100, 'Stationery');

INSERT INTO orders (user_id, order_date, total_amount, status) VALUES
(1, '2024-01-15 10:30:00', 1329.98, 'completed'),
(1, '2024-02-20 14:45:00', 29.99, 'completed'),
(2, '2024-01-20 09:15:00', 89.99, 'completed'),
(2, '2024-02-25 16:30:00', 52.98, 'processing'),
(3, '2024-02-28 11:00:00', 1299.99, 'pending');

INSERT INTO order_items (order_id, product_id, quantity, unit_price) VALUES
(1, 1, 1, 1299.99),
(1, 2, 1, 29.99),
(2, 2, 1, 29.99),
(3, 3, 1, 89.99),
(4, 4, 1, 39.99),
(4, 5, 1, 12.99),
(5, 1, 1, 1299.99);

INSERT INTO logs (log_date, level, message) VALUES
('2024-01-01', 'INFO', 'System started'),
('2024-01-02', 'ERROR', 'Connection timeout'),
('2024-02-01', 'INFO', 'User login'),
('2024-02-15', 'WARNING', 'High memory usage'),
('2025-01-01', 'INFO', 'New year system check');