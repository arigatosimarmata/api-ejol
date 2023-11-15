package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func InitLoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error load file env : %s", err)
	}

}
