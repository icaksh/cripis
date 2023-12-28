CREATE
    EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE trademarks
(
    id                  UUID PRIMARY KEY,
    created_at          TIMESTAMP DEFAULT NOW(),
    updated_at          TIMESTAMP DEFAULT NOW(),
    expired_at          TIMESTAMP NULL,
    created_by          UUID         NOT NULL,
    registration_number VARCHAR(255) NOT NULL,
    trademark_name      VARCHAR(255) NOT NULL,
    trademark_class     INT          NOT NULL,
    owner_name          VARCHAR(255) NOT NULL,
    address             VARCHAR(255) NOT NULL,
    village             INT          NOT NULL,
    district            INT          NOT NULL,
    regency             INT          NOT NULL,
    province            INT          NOT NULL,
    approved_at         TIMESTAMP    NULL,
    approved_by         UUID         NULL,
    file                VARCHAR(255) NULL,
    image               VARCHAR(255) NOT NULL,
    sme_certificate     VARCHAR(255) NULL,
    register_signature  VARCHAR(255) NULL,
    status              INT          NOT NULL,
    notes               TEXT        NULL,
    FOREIGN KEY (village, district, regency, province) REFERENCES villages (id, district_id, regency_id, province_id),
    FOREIGN KEY (created_by) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (trademark_class) REFERENCES trademark_classes (id),
    FOREIGN KEY (status) REFERENCES trademark_status (id)
);

