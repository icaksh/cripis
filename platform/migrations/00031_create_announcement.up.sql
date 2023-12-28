CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS announcements
(
    id          SERIAL PRIMARY KEY,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP    DEFAULT NOW(),
    title       VARCHAR(255) NOT NULL,
    description TEXT         NOT NULL,
    image       VARCHAR(255),
    created_by  UUID         NOT NULL,
    FOREIGN KEY (created_by) REFERENCES users (id) ON DELETE CASCADE
);