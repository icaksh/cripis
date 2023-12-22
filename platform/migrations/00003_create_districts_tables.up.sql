CREATE TABLE districts
(
    province_id INT NOT NULL,
    regency_id INT NOT NULL,
    id   INT     NOT NULL,
    name VARCHAR NOT NULL,
    PRIMARY KEY (province_id, regency_id, id),
    FOREIGN KEY (province_id, regency_id) REFERENCES regencies(province_id, id)
)