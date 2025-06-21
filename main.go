package main

import (
	"backend/config"
	"backend/core/handlers"
	"backend/core/middlewares"
	"backend/core/repositories"
	"backend/core/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	config.NewAppInitEnvironment()
}
func main() {
	db := config.NewAppDatabase()
	app := fiber.New()
	app.Use(cors.New(config.CorsConfig()))

	// repositories
	quoteRepo := repositories.NewQuoteRepository(db, "quotes")
	userRepo := repositories.NewUserRepository(db, "users")
	// services
	quoteService := services.NewQuoteService(quoteRepo)
	userService := services.NewUserService(userRepo)
	// handlers
	quoteHandler := handlers.NewQuoteHandler(quoteService)
	userHandler := handlers.NewUserHandler(userService)
	// routes
	app.Post("/register", userHandler.CreateUser)
	app.Post("/signin", userHandler.SignIn)
	app.Put("/user/:id/:qouteID", middlewares.AccessToken, userHandler.UpdateVote)

	app.Get("/quote", middlewares.AccessToken, quoteHandler.GetQuotes)
	app.Post("/quote", middlewares.AccessToken, quoteHandler.CreateQuote)
	app.Put("/quote/:id", middlewares.AccessToken, quoteHandler.UpdateQuote)
	app.Delete("/quote/:id", middlewares.AccessToken, quoteHandler.DeleteQuote)
	app.Listen("localhost:3000")
}
