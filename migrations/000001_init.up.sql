CREATE TABLE users (
    id PRIMARY KEY,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    role INT NOT NULL

    -- created_at TIMESTAMP DEFAULT NOW()
);