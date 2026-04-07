-- Tables, views, and indexes for PostgreSQL test environment

-- Users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT true,
    metadata JSONB DEFAULT '{}'::jsonb
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_active ON users(is_active);

-- Products table
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    description TEXT,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    stock_quantity INTEGER NOT NULL DEFAULT 0,
    category VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_products_category ON products(category);
CREATE INDEX idx_products_price ON products(price);

-- Orders table
CREATE TYPE order_status AS ENUM ('pending', 'processing', 'shipped', 'delivered', 'cancelled');

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    total_amount NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    status order_status DEFAULT 'pending',
    notes TEXT
);

CREATE INDEX idx_orders_user_id ON orders(user_id);
CREATE INDEX idx_orders_status ON orders(status);

-- Order items table
CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INTEGER NOT NULL REFERENCES orders(id),
    product_id INTEGER NOT NULL REFERENCES products(id),
    quantity INTEGER NOT NULL DEFAULT 1,
    unit_price NUMERIC(10,2) NOT NULL
);

CREATE INDEX idx_order_items_order_id ON order_items(order_id);
CREATE INDEX idx_order_items_product_id ON order_items(product_id);

-- Data types test table
CREATE TABLE data_types_test (
    id SERIAL PRIMARY KEY,
    col_smallint SMALLINT,
    col_integer INTEGER,
    col_bigint BIGINT,
    col_numeric NUMERIC(15,5),
    col_real REAL,
    col_double DOUBLE PRECISION,
    col_varchar VARCHAR(255),
    col_char CHAR(10),
    col_text TEXT,
    col_boolean BOOLEAN,
    col_date DATE,
    col_time TIME,
    col_timestamp TIMESTAMP,
    col_timestamptz TIMESTAMPTZ,
    col_interval INTERVAL,
    col_uuid UUID,
    col_json JSON,
    col_jsonb JSONB,
    col_bytea BYTEA,
    col_inet INET,
    col_cidr CIDR,
    col_macaddr MACADDR,
    col_int_array INTEGER[],
    col_text_array TEXT[]
);

-- Logs table (partitioned by range)
CREATE TABLE logs (
    id SERIAL,
    log_date DATE NOT NULL,
    message TEXT,
    level VARCHAR(20)
) PARTITION BY RANGE (log_date);

CREATE TABLE logs_2024 PARTITION OF logs
    FOR VALUES FROM ('2024-01-01') TO ('2025-01-01');
CREATE TABLE logs_2025 PARTITION OF logs
    FOR VALUES FROM ('2025-01-01') TO ('2026-01-01');

-- Views
CREATE VIEW active_users AS
    SELECT id, username, email, created_at
    FROM users
    WHERE is_active = true;

CREATE VIEW order_summary AS
    SELECT u.username,
           COUNT(o.id) as total_orders,
           COALESCE(SUM(o.total_amount), 0) as total_spent,
           MAX(o.order_date) as last_order_date
    FROM users u
    LEFT JOIN orders o ON u.id = o.user_id
    GROUP BY u.username;

CREATE VIEW product_inventory AS
    SELECT name, stock_quantity, price,
           CASE
               WHEN stock_quantity = 0 THEN 'Out of Stock'
               WHEN stock_quantity < 10 THEN 'Low Stock'
               ELSE 'In Stock'
           END as stock_status
    FROM products;

-- Grant table access
GRANT SELECT ON ALL TABLES IN SCHEMA public TO readonly;
GRANT ALL ON ALL TABLES IN SCHEMA public TO testuser;
GRANT ALL ON ALL SEQUENCES IN SCHEMA public TO testuser;
GRANT ALL ON ALL TABLES IN SCHEMA public TO admin;
GRANT ALL ON ALL SEQUENCES IN SCHEMA public TO admin;
