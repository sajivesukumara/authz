package main

import (
	"authz/internal/initializers"
	"authz/internal/models"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()

}

func main() {
	
     initializers.DB.AutoMigrate(&models.User{})
}