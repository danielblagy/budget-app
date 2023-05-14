package model

type User struct {
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
	FullName string `db:"full_name" json:"full_name"`
}
