package model

type CreateEntry struct {
	CategoryID  int64     `db:"category_id" json:"category_id" validate:"required"`
	Amount      float64   `db:"amount" json:"amount" validate:"required"`
	Date        string    `db:"date" json:"date" validate:"required"`
	Description string    `db:"description" json:"description"`
	Type        EntryType `db:"type" json:"type" validate:"required"`
}
