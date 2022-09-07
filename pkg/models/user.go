package models

type User struct {
	Id    int    `db:"id" json:"id"`
	Email string `db:"email" json:"email"`
	Name  string `db:"name" json:"name"`
	Role  string `db:"role" json:"role"`
}
