CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS user_roles
(
    id   INT,
    role VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS user_verified
(
    id   INT,
    status VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);