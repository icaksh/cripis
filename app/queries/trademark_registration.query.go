package queries

import (
	"github.com/google/uuid"
	"github.com/icaksh/cripis/app/models"
)

func (q *TrademarkQueries) CreateTrademarkRegistration(tr *models.TrademarkRegistration, t *models.Trademark) error {
	queryRegistration := `
	INSERT INTO trademark_registrations(status, trademark_id, created_by, sme_certificate, register_signature) 
	VALUES (1, $1, $2, $3, $4);`

	queryTrademark := `
	INSERT INTO trademarks(status, id, created_by, registration_number, trademark_name ,trademark_class ,owner_name ,address,village,district,regency,province,image)
	VALUES (1, $1 ,$2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`

	_, err := q.Exec(queryTrademark, tr.ID, tr.CreatedBy, t.RegisterNumber, t.TrademarkName, t.Class, t.OwnerName, t.Address, t.Village, t.District, t.Regency, t.Province, t.Image)
	if err != nil {
		return err
	}

	_, err = q.Exec(queryRegistration, tr.ID, tr.CreatedBy, tr.SMECertificate, tr.RegisterSignature)
	if err != nil {
		return err
	}

	q.LogQueries.CreateLog(tr.CreatedBy, "Create Trademark Registration")
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
