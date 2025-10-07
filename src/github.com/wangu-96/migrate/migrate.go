package main

import (
	"github.com/wangu-96/initializers"
	"github.com/wangu-96/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
