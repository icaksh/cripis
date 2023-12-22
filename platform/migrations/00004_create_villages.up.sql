CREATE TABLE villages
(
    province_id INT NOT NULL,
    regency_id INT NOT NULL,
    district_id   INT     NOT NULL,
    id INT NOT NULL ,
    name VARCHAR NOT NULL,
    PRIMARY KEY (province_id, regency_id, district_id, id),
    FOREIGN KEY (province_id, regency_id, district_id) REFERENCES districts(province_id, regency_id, id)
)