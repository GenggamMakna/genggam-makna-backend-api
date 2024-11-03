package main

import (
	"genggam-makna-api/config"
	"genggam-makna-api/models"
)


func main() {
	db := config.InitDB()

	err := db.AutoMigrate(&models.Client{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}
}	