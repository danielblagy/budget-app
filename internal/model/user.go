package model

const UserUsernameMinLength = 10
const UserUsernameMaxLength = 128

type User struct {
	Username string `db:"username" json:"username" validate:"required,min=10,max=128"`
	Email    string `db:"email" json:"email" validate:"required,email"`
	FullName string `db:"full_name" json:"full_name" validate:"required,min=1,max=500"`
	Password string `db:"password_hash" json:"password,omitempty" validate:"required,min=8,max=64"`
}
