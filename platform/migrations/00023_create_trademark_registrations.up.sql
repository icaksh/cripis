CREATE
    EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE trademark_registrations
(
    trademark_id        UUID PRIMARY KEY,
    created_at          TIMESTAMP DEFAULT NOW(),
    updated_at          TIMESTAMP    NULL,
    created_by          UUID         NOT NULL,
    approval_id         UUID         NULL,
    approval_date       DATE         NULL,
    sme_certificate     VARCHAR(255) NULL,
    register_signature  VARCHAR(255) NULL,
    status              INT          NOT NULL,
    FOREIGN KEY (trademark_id) REFERENCES trademarks(id),
    FOREIGN KEY (created_by) REFERENCES users(id),
    FOREIGN KEY (approval_id) REFERENCES users(id),
    FOREIGN KEY (status) REFERENCES trademark_registration_status(id)
    );