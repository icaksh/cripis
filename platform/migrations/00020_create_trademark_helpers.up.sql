CREATE TABLE trademark_status
(
    id     SERIAL PRIMARY KEY,
    status VARCHAR(255) NOT NULL
);

CREATE TABLE trademark_classes
(
    id   SERIAL PRIMARY KEY,
    class VARCHAR(255)   NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE trademark_registration_status
(
    id     SERIAL PRIMARY KEY,
    status VARCHAR(255) NOT NULL
);