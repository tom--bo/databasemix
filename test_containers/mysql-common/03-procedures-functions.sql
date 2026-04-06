-- Common stored procedures and functions for all MySQL versions

USE testdb;

DELIMITER //

-- Stored Procedures
CREATE PROCEDURE get_user_orders(IN p_user_id INT)
BEGIN
    SELECT o.*, u.username, u.email
    FROM orders o
    JOIN users u ON o.user_id = u.id
    WHERE o.user_id = p_user_id
    ORDER BY o.order_date DESC;
END //

CREATE PROCEDURE update_product_stock(
    IN p_product_id INT,
    IN p_quantity INT,
    OUT p_success BOOLEAN
)
BEGIN
    DECLARE current_stock INT;
    DECLARE EXIT HANDLER FOR SQLEXCEPTION
    BEGIN
        SET p_success = FALSE;
        ROLLBACK;
    END;
    
    START TRANSACTION;
    
    SELECT stock_quantity INTO current_stock
    FROM products
    WHERE id = p_product_id
    FOR UPDATE;
    
    IF current_stock >= p_quantity THEN
        UPDATE products 
        SET stock_quantity = stock_quantity - p_quantity
        WHERE id = p_product_id;
        SET p_success = TRUE;
        COMMIT;
    ELSE
        SET p_success = FALSE;
        ROLLBACK;
    END IF;
END //

CREATE PROCEDURE generate_sales_report(
    IN p_start_date DATE,
    IN p_end_date DATE
)
BEGIN
    SELECT 
        DATE(o.order_date) as sale_date,
        COUNT(DISTINCT o.id) as order_count,
        COUNT(DISTINCT o.user_id) as unique_customers,
        SUM(o.total_amount) as daily_revenue
    FROM orders o
    WHERE o.order_date BETWEEN p_start_date AND p_end_date
        AND o.status = 'completed'
    GROUP BY DATE(o.order_date)
    ORDER BY sale_date;
END //

-- Stored Functions
CREATE FUNCTION calculate_user_total(p_user_id INT) 
RETURNS DECIMAL(10,2)
DETERMINISTIC
READS SQL DATA
BEGIN
    DECLARE total DECIMAL(10,2);
    SELECT COALESCE(SUM(total_amount), 0) INTO total
    FROM orders
    WHERE user_id = p_user_id AND status = 'completed';
    RETURN total;
END //

CREATE FUNCTION get_product_availability(p_product_id INT)
RETURNS VARCHAR(20)
DETERMINISTIC
READS SQL DATA
BEGIN
    DECLARE stock INT;
    SELECT stock_quantity INTO stock
    FROM products
    WHERE id = p_product_id;
    
    IF stock IS NULL THEN
        RETURN 'Unknown';
    ELSEIF stock = 0 THEN
        RETURN 'Out of Stock';
    ELSEIF stock < 10 THEN
        RETURN 'Low Stock';
    ELSE
        RETURN 'In Stock';
    END IF;
END //

CREATE FUNCTION format_currency(amount DECIMAL(10,2))
RETURNS VARCHAR(20)
DETERMINISTIC
NO SQL
BEGIN
    RETURN CONCAT('$', FORMAT(amount, 2));
END //

-- Triggers
CREATE TRIGGER before_user_update
BEFORE UPDATE ON users
FOR EACH ROW
BEGIN
    SET NEW.updated_at = CURRENT_TIMESTAMP;
END //

CREATE TRIGGER after_order_insert
AFTER INSERT ON orders
FOR EACH ROW
BEGIN
    INSERT INTO logs (log_date, level, message)
    VALUES (CURDATE(), 'INFO', CONCAT('New order created: ', NEW.id));
END //

CREATE TRIGGER before_product_delete
BEFORE DELETE ON products
FOR EACH ROW
BEGIN
    DECLARE order_count INT;
    SELECT COUNT(*) INTO order_count
    FROM order_items
    WHERE product_id = OLD.id;
    
    IF order_count > 0 THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Cannot delete product with existing orders';
    END IF;
END //

DELIMITER ;