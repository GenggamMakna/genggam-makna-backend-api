package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Email     string
	Username  string
	Password  string
	FirstName string
	LastName  string
	Contact   string
}

type Client struct {
	gorm.Model
	IP        string
	Browser   string
	Version   string
	OS        string
	Device    string
	Origin    string
	API       string
}
