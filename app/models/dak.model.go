package models

type Province struct {
	Id   int    `db:"id" json:"id" validate:"required"`
	Name string `db:"name" json:"name" validate:"required"`
}

type Regency struct {
	Id         int    `db:"id" json:"id" validate:"required"`
	ProvinceId int    `db:"province_id" json:"province_id" validate:"required"`
	Name       string `db:"name" json:"name" validate:"required"`
}

type District struct {
	Id         int    `db:"id" json:"id" validate:"required"`
	ProvinceId int    `db:"province_id" json:"province_id" validate:"required"`
	RegencyId  int    `db:"regency_id" json:"regency_id" validate:"required"`
	Name       string `db:"name" json:"name" validate:"required"`
}

type Village struct {
	Id         int    `db:"id" json:"id" validate:"required"`
	ProvinceId int    `db:"province_id" json:"province_id" validate:"required"`
	RegencyId  int    `db:"regency_id" json:"regency_id" validate:"required"`
	DistrictId int    `db:"district_id" json:"district_id" validate:"required"`
	Name       string `db:"name" json:"name" validate:"required"`
}

type AddressFromDak struct {
	Id string `db:"address" json:"address" validate:"required"`
}

type ProvinceResponse struct {
	Id   int    `db:"id" json:"id" validate:"required"`
	Name string `db:"name" json:"name" validate:"required"`
}

type RegencyRequest struct {
	ProvinceId int `db:"province_id" param:"province_id"`
}

type RegencyResponse struct {
	Id   int    `db:"id" json:"id" validate:"required"`
	Name string `db:"name" json:"name" validate:"required"`
}

type DistrictRequest struct {
	ProvinceId int `db:"province_id" json:"province_id" validate:"required"`
	RegencyId  int `db:"regency_id" json:"regency_id" validate:"required"`
}

type DistrictResponse struct {
	Id   int    `db:"id" json:"id" validate:"required"`
	Name string `db:"name" json:"name" validate:"required"`
}

type VillageRequest struct {
	ProvinceId int `db:"province_id" json:"province_id" validate:"required"`
	RegencyId  int `db:"regency_id" json:"regency_id" validate:"required"`
	DistrictId int `db:"district_id" json:"district_id" validate:"required"`
}

type VillageResponse struct {
	Id   int    `db:"id" json:"id" validate:"required"`
	Name string `db:"name" json:"name" validate:"required"`
}
