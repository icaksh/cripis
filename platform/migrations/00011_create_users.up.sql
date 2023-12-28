CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users
(
    id          UUID,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW(),
    email       VARCHAR(255) NOT NULL,
    first_name  VARCHAR(255) NOT NULL,
    last_name   VARCHAR(255) NOT NULL,
    password    VARCHAR(255) NOT NULL,
    roles       INT          NOT NULL,
    status      INT          NOT NULL,
    verified    INT          NOT NULL,
    card_number BIGINT       NOT NULL,
    dob         TIMESTAMP    NOT NULL,
    sex         INT          NOT NULL,
    address     VARCHAR(255) NOT NULL,
    village     INT          NOT NULL,
    district    INT          NOT NULL,
    regency     INT          NOT NULL,
    province    INT          NOT NULL,
    postal_code INT          NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (roles) REFERENCES user_roles (id),
    FOREIGN KEY (verified) REFERENCES user_verified (id),
    FOREIGN KEY (village, district, regency, province) REFERENCES villages(id, district_id, regency_id, province_id)
);