package main

import (
	"github.com/nikzayn/go-kafka/users/src/database"
	"github.com/nikzayn/go-kafka/users/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	database.Connect()

	db, err := gorm.Open(mysql.Open("root:root@tcp(host.docker.internal:33066)/ambassador"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var users []models.User
	db.Find(&users)

	for _, user := range users {
		database.DB.Create(user)
	}
}
