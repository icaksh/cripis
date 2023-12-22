package queries

import (
	"github.com/google/uuid"
	"github.com/icaksh/cripis/app/models"
	"github.com/jmoiron/sqlx"
)

type TrademarkQueries struct {
	*sqlx.DB
}

type GetTrademarkByUser struct {
	models.Trademark
	models.TrademarkRegistration
}

func (q *TrademarkQueries) CreateTrademark(v *models.Trademark) error {
	query := `INSERT INTO trademarks(id, created_at, updated_at, expired_at, number, name, class, holder_id, registration_id, approved_at, approved_by, file, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	_, err := q.Exec(query, v.ID, v.CreatedAt, v.UpdatedAt, v.ExpiredAt, v.RegisterNumber, v.Name, v.Class, v.Holder, v.RegistrationId, v.ApprovedAt, v.ApprovedBy, v.File, v.Status)
	if err != nil {
		// Return only error.
		return err
	}

	return nil
}

func (q *TrademarkQueries) GetAllTrademarks() ([]models.Trademark, error) {
	result := []models.Trademark{}
	query := `SELECT * FROM trademarks ORDER BY trademarks.created_at`

	err := q.Select(&result, query)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (q *TrademarkQueries) GetAllTrademarksByName(name string) ([]models.Trademark, error) {
	result := []models.Trademark{}
	query := `SELECT * FROM trademarks WHERE trademarks.name=$1 ORDER BY trademarks.created_at`
	err := q.Select(&result, query, name)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (q *TrademarkQueries) GetTrademarkByUser(userId uuid.UUID) ([]GetTrademarkByUser, error) {
	result := []GetTrademarkByUser{}
	query := `SELECT * FROM trademarks INNER JOIN trademark_registrations ON  trademarks.registration_id = trademark_registrations.id WHERE trademark_registrations.register_id=$1 ORDER BY trademarks.created_at`
	err := q.Get(&result, query, userId)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (q *TrademarkQueries) GetTrademarkById(id uuid.UUID) ([]GetTrademarkByUser, error) {
	result := []GetTrademarkByUser{}
	query := `SELECT * FROM trademarks WHERE trademarks.id=$1 ORDER BY trademarks.created_at`
	err := q.Get(&result, query, id)
	if err != nil {
		return result, err
	}
	return result, nil
}
