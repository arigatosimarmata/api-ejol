package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"bitbucket.bri.co.id/scm/ejol/api-ejol/config"
	"bitbucket.bri.co.id/scm/ejol/api-ejol/internal/handlers"
	"bitbucket.bri.co.id/scm/ejol/api-ejol/internal/repository"
	"bitbucket.bri.co.id/scm/ejol/api-ejol/internal/usecases"
	"bitbucket.bri.co.id/scm/ejol/api-ejol/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	utils.InitLoadEnv()
	app := fiber.New()

	fileLogName := os.Getenv("EJOL_DIRECTORY_LOG") + os.Getenv("SERVICE_LOG") + time.Now().Format("20060102") + ".log"

	logFile, err := os.OpenFile(fileLogName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	app.Use(logger.New(logger.Config{
		Format:       "${pid} ${time} ${method} ${path} ${body} ${status} ${latency}\n",
		TimeFormat:   "2006-01-02 15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       logFile,
	}))

	// Initialize Database
	db, err := config.DBConn("")
	if err != nil {
		fmt.Println("failed to init connection", err)
		return
	}

	// Initialize Repository
	ejolRepository := repository.NewEjolRepository(db)

	// Initialize Use Case
	ejolUsecase := usecases.NewEjolUseCase(ejolRepository)

	// Initialize Handler
	ejolHandler := handlers.NewEjolHandler(ejolUsecase)

	// Define Routes
	app.Post("/api/ej-nfs", ejolHandler.GetEjlogNFS)
	app.Post("/api/ej-db", ejolHandler.GetEjlogDB)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}
