package model

type Login struct {
	Username string `json:"username" validate:"required,min=10,max=128"`
	Password string `json:"password" validate:"required,min=8,max=64"`
}
