-- Stored functions and procedures for PostgreSQL test environment

-- Function: Calculate total spent by user
CREATE OR REPLACE FUNCTION calculate_user_total(p_user_id INTEGER)
RETURNS NUMERIC AS $$
DECLARE
    total NUMERIC(10,2);
BEGIN
    SELECT COALESCE(SUM(total_amount), 0) INTO total
    FROM orders
    WHERE user_id = p_user_id;
    RETURN total;
END;
$$ LANGUAGE plpgsql STABLE;

-- Function: Get product availability
CREATE OR REPLACE FUNCTION get_product_availability(p_product_id INTEGER)
RETURNS TEXT AS $$
DECLARE
    qty INTEGER;
BEGIN
    SELECT stock_quantity INTO qty
    FROM products
    WHERE id = p_product_id;

    IF qty IS NULL THEN
        RETURN 'Not Found';
    ELSIF qty = 0 THEN
        RETURN 'Out of Stock';
    ELSIF qty < 10 THEN
        RETURN 'Low Stock';
    ELSE
        RETURN 'In Stock';
    END IF;
END;
$$ LANGUAGE plpgsql STABLE;

-- Function: Format currency
CREATE OR REPLACE FUNCTION format_currency(amount NUMERIC)
RETURNS TEXT AS $$
BEGIN
    RETURN '$' || TO_CHAR(amount, 'FM999,999,990.00');
END;
$$ LANGUAGE plpgsql IMMUTABLE;

-- Procedure: Update product stock
CREATE OR REPLACE PROCEDURE update_product_stock(
    p_product_id INTEGER,
    p_quantity INTEGER
)
LANGUAGE plpgsql AS $$
BEGIN
    UPDATE products
    SET stock_quantity = stock_quantity + p_quantity
    WHERE id = p_product_id;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'Product % not found', p_product_id;
    END IF;
END;
$$;

-- Procedure: Generate sales report (returns result via RAISE NOTICE)
CREATE OR REPLACE PROCEDURE generate_sales_report(
    p_start_date DATE,
    p_end_date DATE
)
LANGUAGE plpgsql AS $$
DECLARE
    rec RECORD;
BEGIN
    FOR rec IN
        SELECT order_date::date as sale_date,
               COUNT(*) as order_count,
               SUM(total_amount) as daily_total
        FROM orders
        WHERE order_date >= p_start_date AND order_date < p_end_date
        GROUP BY order_date::date
        ORDER BY order_date::date
    LOOP
        RAISE NOTICE 'Date: %, Orders: %, Total: %', rec.sale_date, rec.order_count, rec.daily_total;
    END LOOP;
END;
$$;

-- Trigger function: Auto-update updated_at
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER before_user_update
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_timestamp();
