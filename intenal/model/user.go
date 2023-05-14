package model

type User struct {
	Username string `db:"username"`
	Email    string `db:"email"`
	FullName string `db:"full_name"`
}
