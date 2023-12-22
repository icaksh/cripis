CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS user_profiles
(
    user_id     UUID,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP    NULL,
    card_number BIGINT NOT NULL,
    dob       TIME NOT NULL,
    sex         INT NOT NULL,
    address     VARCHAR(255) NOT NULL,
    village     INT NOT NULL,
    district    INT NOT NULL,
    regency     INT NOT NULL,
    province    INT NOT NULL,
    postal_code INT NOT NULL,
    PRIMARY KEY (user_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (village, district, regency, province) REFERENCES villages(id, district_id, regency_id, province_id)
);