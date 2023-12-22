package models

type TrademarkClass struct {
	ID          int    `db:"id" json:"id"`
	Class       string `db:"class" json:"class"`
	Description string `db:"description" json:"description"`
}
