package models

type TrademarkFromDJKI struct {
	Keyword string `query:"keyword" validate:"required"`
}

type TrademarkSimilarity struct {
	Keyword    string  `query:"keyword" validate:"required"`
	Class      string  `query:"class" validate:"required"`
	Similarity float64 `query:"similarity" validate:"required"`
}
