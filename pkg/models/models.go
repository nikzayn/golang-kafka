package models

type User struct {
	ID   int    `json:"id"`
	Name string `json:"string"`
}

type Notification struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Message string `json:"message"`
}
