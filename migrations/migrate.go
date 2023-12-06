package main

import (
	"CRUD-API-GO/initializers"
	"CRUD-API-GO/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Employee{})
}
