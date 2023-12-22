CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE trademark_owners
(
    trademark_id         UUID,
    created_by  UUID NOT NULL,
    full_name   VARCHAR(255) NOT NULL,
    address     TEXT    NOT NULL,
    FOREIGN KEY (trademark_id) REFERENCES trademarks(id),
    FOREIGN KEY (created_by) REFERENCES  users(id)
);

