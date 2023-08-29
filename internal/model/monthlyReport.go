package model

type MonthlyReportItem struct {
	CategoryId   int64        `db:"category_id" json:"category_id"`
	CategoryName string       `db:"category_name" json:"category_name"`
	Amount       float64      `db:"amount" json:"amount"`
	Type         CategoryType `db:"type" `
}

type MonthlyReport struct {
	User     string                    `json:"user"`
	Year     int                       `json:"year"`
	Month    int                       `json:"month"`
	Incomes  []*MonthlyReportInputItem `json:"incomes"`
	Expenses []*MonthlyReportInputItem `json:"expenses"`
}

type MonthlyReportInputItem struct {
	CategoryId   int64   `db:"category_id" json:"category_id"`
	CategoryName string  `db:"category_name" json:"category_name"`
	Amount       float64 `db:"amount" json:"amount"`
}
