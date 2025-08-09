package cmd

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	// repositories
	userRepo := repository.NewUserRepository(dbClient)
	courseRepo := repository.NewCourseRepository(dbClient)
	lessonRepo := repository.NewLessonRepository(dbClient)
	userCourseRepo := repository.NewUserCourseRepository(dbClient)

	// services
	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userRepo)
	courseService := services.NewCourseService(courseRepo, userCourseRepo, lessonRepo)
	lessonService := services.NewLessonService(lessonRepo, courseRepo, userCourseRepo)

	// controllers
	authController := controllers.NewAuthController(userService, authService)
	courseController := controllers.NewCourseController(courseService, lessonService)
	lessonController := controllers.NewLessonController(lessonService)

	// Initialize the server
	echoServer := echo.New()

	// Enable CORS for frontend (localhost:3000)
	echoServer.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowCredentials: true,
		MaxAge:           int((12 * time.Hour).Seconds()),
	}))

	server := server.New(echoServer)

	//register routes
	routes := routes.New(echoServer, authController, courseController, lessonController)
	routes.Init()

	// Start the server
	server.Start(config.App().Port)
}
