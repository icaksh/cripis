CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users
(
    id         UUID,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP    NULL,
    email      VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name  VARCHAR(255) NOT NULL,
    password   VARCHAR(255) NOT NULL,
    phone      BIGINT       NOT NULL,
    roles      INT          NOT NULL,
    status     INT          NOT NULL,
    verified   INT          NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (roles) REFERENCES user_roles (id),
    FOREIGN KEY (verified) REFERENCES user_verified (id)
);