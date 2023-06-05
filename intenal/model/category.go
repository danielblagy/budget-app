package model

type Category struct {
	ID     int64        `db:"id" json:"id"`
	UserID string       `db:"user_id" json:"user_id"`
	Name   string       `db:"name" json:"name" validate:"required,min=1,max=128"`
	Type   CategoryType `db:"type" json:"type" validate:"required"`
}

type CategoryType string

const (
	CategoryTypeIncome  CategoryType = "income"
	CategoryTypeExpense CategoryType = "expense"
)
