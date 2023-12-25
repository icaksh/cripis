package queries

import (
	"fmt"
	"github.com/icaksh/cripis/app/models"
	"github.com/jmoiron/sqlx"
)

type CivilQueries struct {
	*sqlx.DB
}

func (q *CivilQueries) GetProvinces() ([]models.Province, error) {
	province := []models.Province{}
	query := `SELECT * FROM provinces`

	err := q.Select(&province, query)
	if err != nil {
		return province, err
	}
	fmt.Println(province)
	return province, nil
}

func (q *CivilQueries) GetProvince(id int) (models.Province, error) {
	province := models.Province{}
	query := `SELECT * FROM provinces WHERE id=$1`

	err := q.Get(&province, query, id)
	if err != nil {
		return province, err
	}
	fmt.Println(province)
	return province, nil
}

func (q *CivilQueries) GetRegencies(provinceId int) ([]models.Regency, error) {
	regency := []models.Regency{}
	query := `SELECT * FROM regencies WHERE province_id=$1`

	err := q.Select(&regency, query, provinceId)
	if err != nil {
		return regency, err
	}

	return regency, nil
}

func (q *CivilQueries) GetRegency(id int, provinceId int) (models.Regency, error) {
	regency := models.Regency{}
	query := `SELECT * FROM regencies WHERE id=$1 AND province_id=$2`

	err := q.Get(&regency, query, id, provinceId)
	if err != nil {
		return regency, err
	}

	return regency, nil
}

func (q *CivilQueries) GetDistricts(provinceId int, regencyId int) ([]models.District, error) {
	district := []models.District{}
	query := `SELECT * FROM districts WHERE province_id=$1 AND regency_id=$2`

	err := q.Select(&district, query, provinceId, regencyId)
	if err != nil {
		return district, err
	}

	return district, nil
}

func (q *CivilQueries) GetDistrict(id int, provinceId int, regencyId int) (models.District, error) {
	district := models.District{}
	query := `SELECT * FROM districts WHERE id=$1 AND province_id=$2 AND regency_id=$3`

	err := q.Get(&district, query, id, provinceId, regencyId)
	if err != nil {
		return district, err
	}

	return district, nil
}

func (q *CivilQueries) GetVillages(provinceId int, regencyId int, districtId int) ([]models.Village, error) {
	village := []models.Village{}
	query := `SELECT * FROM villages WHERE province_id=$1 AND regency_id=$2 AND district_id=$3`

	err := q.Select(&village, query, provinceId, regencyId, districtId)
	if err != nil {
		return village, err
	}

	return village, nil
}

func (q *CivilQueries) GetAddressFromDak(provinceId int, regencyId int, districtId int, villageId int) (models.AddressFromDak, error) {
	village := models.AddressFromDak{}
	query := `SELECT CONCAT (villages.name, ', ', districts.name, ', ', regencies.name,', ', provinces.name) AS address FROM villages INNER JOIN districts ON villages.district_id=districts.id INNER JOIN regencies ON villages.regency_id=regencies.id INNER JOIN provinces ON villages.province_id=provinces.id WHERE villages.province_id=$1 AND villages.regency_id=$2 AND villages.district_id=$3 AND villages.id=$4 AND districts.province_id=$1 and districts.regency_id=$2 and districts.id=$3 and regencies.province_id=$1 and regencies.id=$2 and provinces.id=$1`

	err := q.Get(&village, query, provinceId, regencyId, districtId, villageId)
	if err != nil {
		return village, err
	}

	return village, nil
}
