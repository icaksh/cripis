package queries

import (
	"github.com/google/uuid"
	"github.com/icaksh/cripis/app/models"
	"github.com/jmoiron/sqlx"
)

type TrademarkQueries struct {
	*sqlx.DB
	LogQueries
}

type GetTrademarkByUser struct {
	models.Trademark
}

func (q *TrademarkQueries) CreateTrademark(t *models.Trademark) error {
	queryTrademark := `
	INSERT INTO trademarks(status, id, created_by, registration_number, trademark_name ,trademark_class ,owner_name ,address,village,district,regency,province,image, sme_certificate, register_signature)
	VALUES (1, $1 ,$2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);`

	_, err := q.Exec(queryTrademark, t.ID, t.CreatedBy, t.RegisterNumber, t.TrademarkName, t.Class, t.OwnerName, t.Address, t.Village, t.District, t.Regency, t.Province, t.Image, t.SMECertificate, t.RegisterSignature)
	if err != nil {
		return err
	}

	return nil
}

func (q *TrademarkQueries) GetTrademarks() ([]models.Trademark, error) {
	result := []models.Trademark{}
	query := `SELECT * FROM trademarks ORDER BY trademarks.created_at`

	err := q.Select(&result, query)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (q *TrademarkQueries) GetTrademarksByName(name string) ([]models.Trademark, error) {
	result := []models.Trademark{}
	query := `SELECT * FROM trademarks WHERE trademarks.trademark_name ILIKE '%'||$1||'%' AND trademarks.status!=3 AND trademarks.status !=1 ORDER BY trademarks.created_at`
	err := q.Select(&result, query, name)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (q *TrademarkQueries) GetTrademarksByUser(userId uuid.UUID) ([]models.Trademark, error) {
	result := []models.Trademark{}
	query := `SELECT * FROM trademarks WHERE trademarks.created_by=$1 ORDER BY trademarks.created_at DESC`
	err := q.Select(&result, query, userId)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (q *TrademarkQueries) GetTrademark(id string) (models.Trademark, error) {
	result := models.Trademark{}
	query := `SELECT * FROM trademarks WHERE trademarks.id=$1 ORDER BY trademarks.created_at`

	err := q.Get(&result, query, id)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (q *TrademarkQueries) UpdateTrademark(v *models.Trademark) error {
	query := `UPDATE trademarks SET 
                    	updated_at=$2, trademark_name=$3, trademark_class=$4, owner_name=$5, address=$6, village=$7, district=$8, regency=$9, province=$10, image=$11, sme_certificate=$12, register_signature=$13, status=$14
                      WHERE id=$1 AND status!=2`

	_, err := q.Exec(query, v.ID, v.UpdatedAt, v.TrademarkName, v.Class, v.OwnerName, v.Address, v.Village, v.District, v.Regency, v.Province, v.Image, v.SMECertificate, v.RegisterSignature, v.Status)
	if err != nil {
		return err
	}
	return nil
}

func (q *TrademarkQueries) UpdateTrademarkStatus(v *models.Trademark) error {
	query := `UPDATE trademarks SET status=$2, notes=$3, approved_at=$4, approved_by=$5, expired_at=$6, file=$7 WHERE id=$1 AND status!=2`

	_, err := q.Exec(query, v.ID, v.Status, v.Notes, v.ApprovedAt, v.ApprovedBy, v.ExpiredAt, v.File)
	if err != nil {
		return err
	}
	return nil
}

func (q *TrademarkQueries) DeleteTrademark(id string) error {
	query := `DELETE FROM trademarks WHERE id=$1`
	_, err := q.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
