CREATE
    EXTENSION IF NOT EXISTS "uuid-ossp";

type TrademarkRegistration struct {
	ID                 uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt          time.Time `db:"created_at" json:"created_at"`
	UpdatedAt          time.Time `db:"updated_at" json:"updated_at"`
	RegistrationNumber string    `db:"registration_number" json:"registration_number"`
	RegisterId         uuid.UUID `db:"register_id" json:"register_id"`
	ApprovalId         uuid.UUID `db:"approval_id" json:"approval_id"`
	ApprovalDate       time.Time `db:"approval_date" json:"approval_date"`
	SMECertificate     string    `db:"sme_certificate" json:"sme_certificate"`
	RegisterSignature  string    `db:"register_signature" json:"register_signature"`
	Status             int16     `db:"status" json:"status" validate:"required"`
}

CREATE TABLE trademark_registrations
(
    id                  UUID PRIMARY KEY,
    created_at          TIMESTAMP DEFAULT NOW(),
    updated_at          TIMESTAMP    NULL,
    registration_number VARCHAR(255) NOT NULL,
    register_id         UUID         NOT NULL,
    approval_id         UUID         NULL,
    approval_date       DATE         NULL,
    sme_certificate     VARCHAR(255) NULL,
    register_signature  VARCHAR(255) NULL,
    status              INT          NOT NULL,
    FOREIGN KEY (register_id) REFERENCES users(id),
    FOREIGN KEY (approval_id) REFERENCES users(id),
    FOREIGN KEY (status) REFERENCES trademark_registration_status(id)
    );