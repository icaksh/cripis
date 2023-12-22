package queries

import (
	"github.com/google/uuid"
	"github.com/icaksh/cripis/app/models"
)

func (q *TrademarkQueries) CreateTrademarkRegistration(v *models.TrademarkRegistration) error {
	query := `INSERT INTO trademark_registrations(id, created_at, registration_number, register_id, sme_certificate, register_signature, status) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := q.Exec(query, v.ID, v.CreatedAt, v.RegistrationNumber, v.RegisterId, v.SMECertificate, v.RegisterSignature)
	if err != nil {
		// Return only error.
		return err
	}

	return nil
}

func (q *TrademarkQueries) GetTrademarkRegistrations(userId uuid.UUID) ([]models.TrademarkRegistration, error) {
	result := []models.TrademarkRegistration{}
	query := `SELECT * FROM trademark_registrations WHERE trademark_registrations.register_id=$1 ORDER BY trademark_registrations.created_at`

	err := q.Get(&result, query, userId)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (q *TrademarkQueries) AdGetTrademarkRegistrations() ([]models.TrademarkRegistration, error) {
	result := []models.TrademarkRegistration{}
	query := `SELECT * FROM trademark_registrations ORDER BY trademark_registrations.created_at`

	err := q.Get(&result, query)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (q *TrademarkQueries) GetTrademarkRegistration(id uuid.UUID, userId uuid.UUID) (models.TrademarkRegistration, error) {
	result := models.TrademarkRegistration{}
	query := `SELECT * FROM trademark_registrations WHERE trademark_registrations.id=$1 AND trademark_registrations.register_id=$2`

	err := q.Get(&result, query, id, userId)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (q *TrademarkQueries) AdGetTrademarkRegistration(id uuid.UUID) (models.TrademarkRegistration, error) {
	result := models.TrademarkRegistration{}
	query := `SELECT * FROM trademark_registrations WHERE trademark_registrations.id=$1`

	err := q.Get(&result, query, id)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (q *TrademarkQueries) DeleteTrademarkRegistration(id uuid.UUID, userId uuid.UUID) error {
	query := `DELETE FROM trademark_registrations WHERE trademark_registrations.id=$1 AND trademark_registrations.register_id=$2`

	_, err := q.Exec(query, id, userId)
	if err != nil {
		// Return only error.
		return err
	}

	return nil
}

func (q *TrademarkQueries) AdDeleteTrademarkRegistration(id uuid.UUID) error {
	query := `DELETE FROM trademark_registrations WHERE id=$1`

	_, err := q.Exec(query, id)
	if err != nil {
		// Return only error.
		return err
	}

	return nil
}
