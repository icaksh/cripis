CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    email VARCHAR (255) NOT NULL,
    username VARCHAR (255) NOT NULL,
    password VARCHAR (255) NOT NULL,
    phone BIGINT NOT NULL,
    level INT NOT NULL,
    status INT NOT NULL,
    verified INT NOT NULL
);

CREATE TABLE user_profile (
    user_id UUID,
    first_name VARCHAR (255) NOT NULL,
    last_name VARCHAR (255) NOT NULL,
    address VARCHAR (255),
    city VARCHAR (255),
    province VARCHAR (255),
    postal_code INT
);