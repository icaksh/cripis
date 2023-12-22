CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS trademark_holders
(
    id     UUID,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP    NULL,
    register_id UUID NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    address     VARCHAR(255) NOT NULL,
    village     INT NOT NULL,
    district    INT NOT NULL,
    regency     INT NOT NULL,
    province    INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (village, district, regency, province) REFERENCES villages(id, district_id, regency_id, province_id),
    FOREIGN KEY (register_id) REFERENCES users(id)
);