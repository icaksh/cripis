CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE trademarks
(
    id         UUID      PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP NULL,
    expired_at     VARCHAR(255) NULL,
    number     VARCHAR(255) NOT NULL,
    name       VARCHAR(255) NOT NULL,
    class       INT NOT NULL,
    holder_id    UUID NOT NULL,
    registration_id UUID NOT NULL,
    approved_at DATE NULL,
    approved_by DATE NULL,
    file       VARCHAR(255) NOT NULL,
    status     INT          NOT NULL,
    FOREIGN KEY (class) REFERENCES trademark_classes(id),
    FOREIGN KEY (holder_id) REFERENCES trademark_holders(id),
    FOREIGN KEY (status) REFERENCES trademark_status(id),
    FOREIGN KEY (registration_id) REFERENCES trademark_registrations(id)
);

