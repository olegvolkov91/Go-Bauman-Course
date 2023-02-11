package models

type User struct {
	Id           int    `json:"id"`
	Login        string `json:"login"`
	PasswordHash string `json:"password_hash"`
}
