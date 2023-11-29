package utils

import (
	"log"
	"os"

	"bitbucket.bri.co.id/scm/ejol/api-ejol/config"
	"github.com/joho/godotenv"
)

func InitLoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error load file env : %s", err)
	}

	config.NFS_DAYS = os.Getenv("NFS_DAYS")
}
