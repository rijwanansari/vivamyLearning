package cmd

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/rijwanansari/vivaLearning/config"
	"github.com/rijwanansari/vivaLearning/conn"
	"github.com/rijwanansari/vivaLearning/controllers"
	repository "github.com/rijwanansari/vivaLearning/repositories"
	"github.com/rijwanansari/vivaLearning/routes"
	"github.com/rijwanansari/vivaLearning/server"
	"github.com/rijwanansari/vivaLearning/services"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: Serve,
}

func Serve(cmd *cobra.Command, args []string) {
	// Print configuration values loaded from .env file
	appConfig := config.App()
	dbConfig := config.Db()

	fmt.Printf("=== Configuration Loaded from .env ===\n")
	fmt.Printf("App Name: %s\n", appConfig.Name)
	fmt.Printf("App Port: %d\n", appConfig.Port)
	fmt.Printf("DB Host: %s\n", dbConfig.Host)
	fmt.Printf("DB Port: %s\n", dbConfig.Port)
	fmt.Printf("DB User: %s\n", dbConfig.User)
	fmt.Printf("DB Schema: %s\n", dbConfig.Schema)
	fmt.Printf("DB Debug: %t\n", dbConfig.Debug)
	fmt.Printf("=====================================\n\n")

	// Initialize database connection
	conn.InitDB()

	//database client
	dbClient := conn.Db()

	// repository
	userRepo := repository.NewUserRepository(dbClient)

	// services
	userService := services.NewUserService(userRepo)
	// auth service
	authService := services.NewAuthService(userRepo)

	//get controller
	authController := controllers.NewAuthController(userService, authService)

	// Initialize the server
	echoServer := echo.New()
	server := server.New(echoServer)

	//register routes
	routes := routes.New(echoServer, authController)
	routes.Init()

	// Start the server
	server.Start(config.App().Port)
}
