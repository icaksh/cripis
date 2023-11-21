CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE ip_status (
    id UUID PRIMARY KEY,
    status VARCHAR (255) NOT NULL,
)

CREATE TABLE ip_types (
    id int PRIMARY KEY,
    code VARCHAR (2) NOT NULL,
    name VARCHAR (255) NOT NULL,
)

CREATE TABLE ips (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    number VARCHAR (255) NOT NULL,
    name VARCHAR (255) NOT NULL,
    user_id VARCHAR (255) NOT NULL,
    type INT NOT NULL FOREIGN KEY REFERENCES intellectual_property_type(id),
    status INT NOT NULL FOREIGN KEY REFERENCES intellectual_property_status(id),
    file VARCHAR (255) NOT NULL
);