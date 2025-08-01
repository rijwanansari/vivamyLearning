package conn

import (
	"fmt"
	"log"
	"time"

	"github.com/rijwanansari/vivaLearning/config"
	"github.com/rijwanansari/vivaLearning/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB() {
	// Get DB config
	dbConfig := config.Db()

	// Build DSN
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Pass,
		dbConfig.Schema,
	)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(func() logger.LogLevel {
			if dbConfig.Debug {
				return logger.Info
			}
			return logger.Silent
		}()),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Setup connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get generic DB from GORM: %v", err)
	}
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConn)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Duration(dbConfig.MaxConnLifetime) * time.Second)

	// Ping test
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	fmt.Println("Connected to the database successfully")

	// Auto Migrate models
	err = db.AutoMigrate(
		&domain.User{},
		&domain.Course{},
		&domain.Lesson{},
		&domain.UserCourse{},
	)
	if err != nil {
		log.Fatalf("Auto migration failed: %v", err)
	}
}

func Db() *gorm.DB {
	return db
}

func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Unable to get DB instance for closing:", err)
		return
	}
	sqlDB.Close()
}
