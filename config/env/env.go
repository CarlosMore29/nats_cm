package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvFile() {
	ENV := os.Getenv("ENV")
	if ENV != `prod` {
		var errEnv = godotenv.Load()
		if errEnv != nil {
			fmt.Println("Error loading .env file")
		}
	}
}
