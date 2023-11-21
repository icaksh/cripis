package models

type SearchPDKI struct {
	Keyword string `query:"keyword" validate:"required"`
	Page int `query:"page" validate:"required"`
	ShowFilter bool `query:"show_filter" validate:"required"`
	Type string `query:"type" validate:"required"`
	OrderState string `query:"order_state" validate:"required"`
}