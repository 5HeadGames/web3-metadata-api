package endersapi

import (
	"log"

	"github.com/joho/godotenv"
)

func GetEnvVars(envKey string) string {
	envVars, err := godotenv.Read()
	if err != nil {
		log.Printf("Error loading .env file! \n%v", err)
	}

	return envVars[envKey]
}
