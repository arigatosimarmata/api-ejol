package main

import (
	"fmt"
	"log"

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
	app.Use(logger.New())

	// Initialize Database
	db, err := config.ConnectDB()
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
