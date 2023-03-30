package main

import (
	"go-rest-api/config"
	"go-rest-api/models"
	"go-rest-api/routes"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	db := config.SetupDB()
	db.AutoMigrate(&models.User{})

	r := routes.SetupRoutes(db)
	r.Run(":8080")
}
