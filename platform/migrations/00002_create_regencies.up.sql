CREATE TABLE regencies
(
    province_id INT NOT NULL,
    id   INT     NOT NULL,
    name VARCHAR NOT NULL,
    PRIMARY KEY (province_id,id),
    FOREIGN KEY (province_id) REFERENCES provinces(id)
)