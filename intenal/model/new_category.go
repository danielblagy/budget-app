package model

type NewCategory struct {
	Name string `json:"name" validate:"required,min=1,max=128"`
}
