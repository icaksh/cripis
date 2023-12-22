package queries

import (
	"github.com/google/uuid"
	"github.com/icaksh/cripis/app/models"
)

func (q *TrademarkQueries) CreateTrademarkHolders(v *models.TrademarkHolder) error {
	query := `INSERT INTO trademark_holders(id, created_at, updated_at, register_id, full_name, address, village, district, regency, province) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := q.Exec(query, v.Id, v.CreatedAt, v.UpdatedAt, v.RegisterId, v.FullName, v.Address, v.Village, v.District, v.Regency, v.Province)
	if err != nil {
		return err
	}
	return nil
}

func (q *TrademarkQueries) GetAllTrademarkHolders() ([]models.TrademarkHolder, error) {
	result := []models.TrademarkHolder{}
	query := `SELECT * FROM trademark_holders ORDER BY trademark_holders.created_at`

	err := q.Get(&result, query)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (q *TrademarkQueries) GetTrademarkHolders(userId uuid.UUID) (models.TrademarkHolder, error) {
	result := models.TrademarkHolder{}
	query := `SELECT * FROM trademark_holders WHERE trademark_holders.id=$1`

	err := q.Get(&result, query, userId)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (q *TrademarkQueries) DeleteTrademarkHolders(id uuid.UUID) error {
	query := `DELETE FROM trademark_holders WHERE trademark_holders.id=$1`

	_, err := q.Exec(query, id)
	if err != nil {
		// Return only error.
		return err
	}

	return nil
}
