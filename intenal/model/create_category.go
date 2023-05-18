package model

type CreateCategory struct {
	Name string `json:"name" validate:"required,min=1,max=128"`
}
