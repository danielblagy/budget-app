package model

import "time"

type Entry struct {
	ID          int64     `db:"id" json:"id"`
	UserID      string    `db:"user_id" json:"user_id"`
	CategoryID  *int64    `db:"category_id" json:"category_id"`
	Amount      float64   `db:"amount" json:"amount"`
	Date        time.Time `db:"date" json:"date"`
	Description string    `db:"description" json:"description"`
	Type        EntryType `db:"type" json:"type"`
}

type EntryType string

const (
	EntryTypeIncome  EntryType = "income"
	EntryTypeExpense EntryType = "expense"
)
