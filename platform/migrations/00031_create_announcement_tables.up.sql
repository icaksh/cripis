CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS announcements
(
    id     INT PRIMARY KEY ,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP    NULL,
    title BIGINT NOT NULL,
    description TIME NOT NULL,
    image       VARCHAR(255) NOT NULL,
    created_by     UUID NOT NULL,
    FOREIGN KEY (created_by) REFERENCES users(id)
);