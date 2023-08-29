package model

type MonthlyReportItem struct {
	CategoryId   int64        `db:"category_id" json:"category_id"`
	CategoryName string       `db:"category_name" json:"category_name"`
	Amount       float64      `db:"amount" json:"amount"`
	Type         CategoryType `db:"type" json:"type"`
}
