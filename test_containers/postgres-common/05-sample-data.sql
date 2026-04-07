-- Sample data for PostgreSQL test environment

INSERT INTO users (username, email, is_active, metadata) VALUES
    ('john_doe', 'john@example.com', true, '{"role": "admin", "department": "engineering"}'),
    ('jane_smith', 'jane@example.com', true, '{"role": "user", "department": "marketing"}'),
    ('bob_wilson', 'bob@example.com', true, '{"role": "user", "department": "sales"}'),
    ('alice_jones', 'alice@example.com', false, '{"role": "user", "department": "hr"}'),
    ('charlie_brown', 'charlie@example.com', true, '{"role": "moderator"}');

INSERT INTO products (name, description, price, stock_quantity, category) VALUES
    ('Widget A', 'Standard widget', 19.99, 100, 'widgets'),
    ('Widget B', 'Premium widget', 29.99, 50, 'widgets'),
    ('Gadget X', 'Basic gadget', 49.99, 25, 'gadgets'),
    ('Gadget Y', 'Advanced gadget', 99.99, 10, 'gadgets'),
    ('Tool Z', 'Multi-purpose tool', 14.99, 200, 'tools');

INSERT INTO orders (user_id, total_amount, status) VALUES
    (1, 49.98, 'delivered'),
    (1, 99.99, 'shipped'),
    (2, 29.99, 'processing'),
    (3, 64.97, 'pending'),
    (5, 14.99, 'delivered');

INSERT INTO order_items (order_id, product_id, quantity, unit_price) VALUES
    (1, 1, 2, 19.99),
    (2, 4, 1, 99.99),
    (3, 2, 1, 29.99),
    (4, 1, 1, 19.99),
    (4, 3, 1, 49.99),
    (4, 5, 1, 14.99),
    (5, 5, 1, 14.99);

INSERT INTO logs (log_date, message, level) VALUES
    ('2024-06-15', 'System started', 'INFO'),
    ('2024-07-01', 'Database backup completed', 'INFO'),
    ('2024-12-25', 'High memory usage detected', 'WARNING'),
    ('2025-01-10', 'User login failed', 'ERROR'),
    ('2025-03-20', 'Routine maintenance', 'INFO');
