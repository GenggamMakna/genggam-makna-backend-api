package main

import (
	"genggam-makna-api/config"
	"genggam-makna-api/models"
)


func main() {
	db := config.InitDB()

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	err := db.AutoMigrate(&models.Client{}, &models.Users{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}
}	