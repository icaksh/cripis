package main

import (
	"fmt"
	"github.com/icaksh/cripis/app/utils"
	"github.com/icaksh/cripis/pkg/configs"
	"github.com/icaksh/cripis/pkg/middleware"
	"github.com/icaksh/cripis/pkg/routes"
	recaptcha "github.com/jansvabik/fiber-recaptcha"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)
	fmt.Println("Registered routes:")
	for _, route := range app.Stack() {
		fmt.Printf("%s %s\n", route)
	}
	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	recaptcha.SecretKey = os.Getenv("RECAPTCHA_SECRET")
	// Routes.
	// routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	routes.MainRoutes(app) // Register a public routes for app.
	// routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with graceful shutdown).
	utils.StartServerWithGracefulShutdown(app)
}
