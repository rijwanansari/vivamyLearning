package cmd

import (
	"github.com/rijwanansari/vivaLearning/config"
	"github.com/spf13/cobra"
	"github.com/vivasoft-ltd/golang-course-utils/logger"
)

var rootCmd = &cobra.Command{
	Use: "root",
}

func Execute() {
	// Load the configuration
	config.LoadConfig()

	// Initialize the logger
	initLogger()

	// Initialize the database connection
	//conn.InitDB()

	rootCmd.AddCommand(serveCmd)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func initLogger() {
	logger.SetFileLogger(config.Logger().FilePath)
}
