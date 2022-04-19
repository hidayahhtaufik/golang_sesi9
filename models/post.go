package models

type User struct {
	ID     uint
	UserID uint   `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
