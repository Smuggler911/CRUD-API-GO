package initializers

import (
	"github.com/joho/godotenv"
	"log"
)

// загрузка переменных среды
func LoadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
