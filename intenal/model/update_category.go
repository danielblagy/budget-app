package model

type UpdateCategory struct {
	ID   int64  `db:"id" json:"id" validate:"required"`
	Name string `db:"name" json:"name" validate:"required,min=1,max=128"`
}
