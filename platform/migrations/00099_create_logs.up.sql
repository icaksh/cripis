CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS logs
(
    timestamp TIMESTAMP DEFAULT NOW(),
    user_id UUID    NULL,
    action VARCHAR(255) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);