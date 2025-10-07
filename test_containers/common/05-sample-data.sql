-- Sample data for testing dbmix tool
USE testdb;

-- Insert sample data into users table
INSERT INTO users (username, email, created_at, is_active) VALUES
('john_doe', 'john@example.com', NOW() - INTERVAL 30 DAY, 1),
('jane_smith', 'jane@example.com', NOW() - INTERVAL 20 DAY, 1),
('bob_wilson', 'bob@example.com', NOW() - INTERVAL 10 DAY, 0),
('alice_brown', 'alice@example.com', NOW() - INTERVAL 5 DAY, 1);

-- Sample data for posts, categories, and post_categories tables (commented out as these tables don't exist yet)
-- INSERT INTO posts (title, content, author_id, created_at, status) VALUES
-- ('First Post', 'This is the content of the first post', 1, NOW() - INTERVAL 25 DAY, 'published'),
-- ('Second Post', 'Content for the second post', 2, NOW() - INTERVAL 15 DAY, 'published'),
-- ('Draft Post', 'This is a draft post', 1, NOW() - INTERVAL 5 DAY, 'draft'),
-- ('Another Post', 'More content here', 3, NOW() - INTERVAL 2 DAY, 'published');
--
-- INSERT INTO categories (name, description) VALUES
-- ('Technology', 'Posts about technology and programming'),
-- ('Lifestyle', 'Posts about lifestyle and daily life'),
-- ('Business', 'Posts about business and entrepreneurship');
--
-- INSERT INTO post_categories (post_id, category_id) VALUES
-- (1, 1), (1, 3),
-- (2, 2),
-- (3, 1),
-- (4, 2), (4, 3);

-- Insert sample data into user_profiles table (commented out as table doesn't exist yet)
-- INSERT INTO user_profiles (user_id, first_name, last_name, bio, avatar_url) VALUES
-- (1, 'John', 'Doe', 'Software developer and tech enthusiast', 'https://example.com/avatars/john.jpg'),
-- (2, 'Jane', 'Smith', 'Content writer and blogger', 'https://example.com/avatars/jane.jpg'),
-- (3, 'Bob', 'Wilson', 'Business consultant', 'https://example.com/avatars/bob.jpg'),
-- (4, 'Alice', 'Brown', 'UX designer', 'https://example.com/avatars/alice.jpg');

-- Update some test variables for demonstration
SET @old_sql_mode = @@sql_mode;
SET SESSION sql_mode = 'STRICT_TRANS_TABLES,NO_ZERO_DATE,NO_ZERO_IN_DATE,ERROR_FOR_DIVISION_BY_ZERO';

-- Session variable for testing (wait_timeout is a session-level variable)
SET SESSION wait_timeout = 3600;