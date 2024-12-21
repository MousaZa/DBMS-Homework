-- ===========================================
-- TABLES
-- ===========================================

-- Create the users table
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    role VARCHAR(50) NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password TEXT NOT NULL
);

-- Create the customers table inheriting from users
CREATE TABLE customers (
    can_borrow BOOLEAN DEFAULT TRUE
) INHERITS (users);

-- Create the admins table inheriting from users
CREATE TABLE admins (
    library_name VARCHAR(255) NOT NULL
) INHERITS (users);

-- Create the categories table
CREATE TABLE categories (
    category_id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL
);

-- Create the books table
CREATE TABLE books (
    book_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    category_id INT REFERENCES categories(category_id) ON DELETE SET NULL,
    cover_url TEXT,
    language VARCHAR(50),
    likes INTEGER DEFAULT 0,
    summary TEXT,
    available BOOLEAN DEFAULT TRUE
);

-- Create the borrows table
CREATE TABLE borrows (
    borrow_id SERIAL PRIMARY KEY,
    book_id INT NOT NULL REFERENCES books(book_id) ON DELETE CASCADE,
    user_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    start_date DATE NOT NULL,
    end_date DATE, -- Automatically set
    status VARCHAR(50) NOT NULL -- e.g., 'borrowed', 'returned', 'late'
);

-- Create the likes table
CREATE TABLE likes (
    like_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    book_id INT NOT NULL REFERENCES books(book_id) ON DELETE CASCADE,
    UNIQUE (user_id, book_id) -- Ensures no duplicate likes
);

-- Create the notifications table
CREATE TABLE notifications (
    notification_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    status VARCHAR(50) NOT NULL -- e.g., 'sent', 'read'
);

-- Create the commercial_notifications table inheriting from notifications
CREATE TABLE commercial_notifications (
    message TEXT NOT NULL
) INHERITS (notifications);

-- Create the late_notifications table inheriting from notifications
CREATE TABLE late_notifications (
    book_id INT NOT NULL REFERENCES books(book_id) ON DELETE CASCADE
) INHERITS (notifications);

-- ===========================================
-- FUNCTIONS AND TRIGGERS
-- ===========================================

-- 1. Function to update the likes count in books
CREATE OR REPLACE FUNCTION update_likes_count()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE books
    SET likes = (SELECT COUNT(*) FROM likes WHERE book_id = NEW.book_id)
    WHERE book_id = NEW.book_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger AFTER INSERT on the likes table
CREATE TRIGGER trigger_after_insert_likes
AFTER INSERT ON likes
FOR EACH ROW
EXECUTE FUNCTION update_likes_count();

-- Trigger AFTER DELETE on the likes table
CREATE TRIGGER trigger_after_delete_likes
AFTER DELETE ON likes
FOR EACH ROW
EXECUTE FUNCTION update_likes_count();

-- 2. Function to prevent borrowing of unavailable books
CREATE OR REPLACE FUNCTION prevent_unavailable_books()
RETURNS TRIGGER AS $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM books WHERE book_id = NEW.book_id AND available = TRUE
    ) THEN
        RAISE EXCEPTION 'This book is currently unavailable for borrowing.';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger BEFORE INSERT on the borrows table
CREATE TRIGGER trigger_prevent_unavailable_books
BEFORE INSERT ON borrows
FOR EACH ROW
EXECUTE FUNCTION prevent_unavailable_books();

-- 3. Function to automatically set the borrow end date (e.g., 14 days later)
CREATE OR REPLACE FUNCTION set_borrow_end_date()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.end_date IS NULL THEN
        NEW.end_date := NEW.start_date + INTERVAL '14 days';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger BEFORE INSERT on the borrows table
CREATE TRIGGER trigger_set_borrow_end_date
BEFORE INSERT ON borrows
FOR EACH ROW
EXECUTE FUNCTION set_borrow_end_date();
