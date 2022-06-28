package models

type User struct {
	ID       string     `json:"id"`
	Username string     `json:"user"`
	Email    string     `json:"email"`
	Contacts []*Contact `json:"contacts"`
}
