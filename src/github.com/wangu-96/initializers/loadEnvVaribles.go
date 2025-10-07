package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load("C:/Users/HP ZBOOK/Documents/Go-Projects/Go-Practice/src/github.com/wangu-96/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
