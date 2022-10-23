package endersapi

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func GetEnvVars(envKey string) string {
	modEnv := os.Getenv("GOMOD")
	pathMod := strings.Split(modEnv, "/")
	pkgPath := strings.Join(pathMod[:len(pathMod)-1], "/")

	envVars, err := godotenv.Read(pkgPath + "/.env")
	if err != nil {
		log.Printf("Error loading .env file! \n%v", err)
	}

	return envVars[envKey]
}
