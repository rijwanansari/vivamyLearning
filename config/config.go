package config

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type AppConfig struct {
	Name string `json:"name"`
	Port int    `json:"port"`
}

type DbConfig struct {
	Host            string `json:"host"`
	Port            string `json:"port"`
	User            string `json:"user"`
	Pass            string `json:"pass"`
	Schema          string `json:"schema"`
	MaxIdleConn     int    `json:"maxIdleConn"`
	MaxOpenConn     int    `json:"maxOpenConn"`
	MaxConnLifetime int    `json:"maxConnLifetime"`
	Debug           bool   `json:"debug"`
}

type LoggerConfig struct {
	Level    string `json:"level"`
	FilePath string `json:"filePath"`
}
type RedisConfig struct {
	Host               string
	Port               string
	Pass               string
	Db                 int
	MandatoryPrefix    string
	AccessUuidPrefix   string
	RefreshUuidPrefix  string
	UserPrefix         string
	PermissionPrefix   string
	UserCacheTTL       time.Duration
	PermissionCacheTTL time.Duration
}

type Config struct {
	App    AppConfig    `json:"app"`
	Db     DbConfig     `json:"db"`
	Logger LoggerConfig `json:"logger"`
	Jwt    *JwtConfig   `json:"jwt"`
	Redis  *RedisConfig `json:"redis"`
}
type JwtConfig struct {
	AccessTokenSecret  string `json:"accessTokenSecret"`
	RefreshTokenSecret string `json:"refreshTokenSecret"`
	AccessTokenExpiry  int64  `json:"accessTokenExpiry"`  // in seconds
	RefreshTokenExpiry int64  `json:"refreshTokenExpiry"` // in seconds
}

var config Config

func LoadConfig() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env file :", err)
	} else {
		fmt.Println("Environment variables loaded from .env file")
	}

	// Bind environment variables for all configuration fields
	bindEnvironmentVariables()

	// Set default values
	setDefaults()

	// Try to load from environment variables first
	if err := viper.Unmarshal(&config); err == nil {
		fmt.Println("Config loaded from environment variables")
		return
	} else {
		fmt.Printf("Failed to unmarshal config from environment variables: %v\n", err)
	}

	// Fallback to Consul if environment variables are not sufficient
	loadFromConsul()
}

func bindEnvironmentVariables() {
	// App configuration
	_ = viper.BindEnv("app.name", "APP_NAME")
	_ = viper.BindEnv("app.port", "APP_PORT")

	// Database configuration
	_ = viper.BindEnv("db.host", "DB_HOST")
	_ = viper.BindEnv("db.port", "DB_PORT")
	_ = viper.BindEnv("db.user", "DB_USER")
	_ = viper.BindEnv("db.pass", "DB_PASS", "DB_PASSWORD")
	_ = viper.BindEnv("db.schema", "DB_SCHEMA", "DB_NAME")
	_ = viper.BindEnv("db.maxIdleConn", "DB_MAX_IDLE_CONN")
	_ = viper.BindEnv("db.maxOpenConn", "DB_MAX_OPEN_CONN")
	_ = viper.BindEnv("db.maxConnLifetime", "DB_MAX_CONN_LIFETIME")
	_ = viper.BindEnv("db.debug", "DB_DEBUG")

	// Logger configuration
	_ = viper.BindEnv("logger.level", "LOG_LEVEL")
	_ = viper.BindEnv("logger.filePath", "LOG_FILE_PATH")

	// JWT configuration
	_ = viper.BindEnv("jwt.accessTokenSecret", "JWT_ACCESS_TOKEN_SECRET")
	_ = viper.BindEnv("jwt.refreshTokenSecret", "JWT_REFRESH_TOKEN_SECRET")
	_ = viper.BindEnv("jwt.accessTokenExpiry", "JWT_ACCESS_TOKEN_EXPIRY")
	_ = viper.BindEnv("jwt.refreshTokenExpiry", "JWT_REFRESH_TOKEN_EXPIRY")

	// Consul configuration (for fallback)
	_ = viper.BindEnv("CONSUL_URL")
	_ = viper.BindEnv("CONSUL_PATH")
}

func setDefaults() {
	// App defaults
	viper.SetDefault("app.name", "did-api")
	viper.SetDefault("app.port", 8080)

	// Database defaults
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "5432")
	viper.SetDefault("db.maxIdleConn", 10)
	viper.SetDefault("db.maxOpenConn", 100)
	viper.SetDefault("db.maxConnLifetime", 3600) // 1 hour in seconds
	viper.SetDefault("db.debug", false)

	// Logger defaults
	viper.SetDefault("logger.level", "info")
	viper.SetDefault("logger.filePath", "logs/app.log")

	// JWT defaults
	viper.SetDefault("jwt.accessTokenSecret", "default-access-secret-change-in-production")
	viper.SetDefault("jwt.refreshTokenSecret", "default-refresh-secret-change-in-production")
	viper.SetDefault("jwt.accessTokenExpiry", 900)     // 15 minutes in seconds
	viper.SetDefault("jwt.refreshTokenExpiry", 604800) // 7 days in seconds

	//redis defaults
	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", "6379")
	viper.SetDefault("redis.pass", "")
	viper.SetDefault("redis.db", 2)
	viper.SetDefault("redis.mandatoryPrefix", "vivaLearning:")
	viper.SetDefault("redis.accessUuidPrefix", "access:")
	viper.SetDefault("redis.refreshUuidPrefix", "refresh:")
	viper.SetDefault("redis.userPrefix", "user:")
	viper.SetDefault("redis.permissionPrefix", "permission:")
	viper.SetDefault("redis.userCacheTTL", 5*time.Minute)
	viper.SetDefault("redis.permissionCacheTTL", 5*time.Minute)
}

func loadFromConsul() {
	consulURL := viper.GetString("CONSUL_URL")
	consulPath := viper.GetString("CONSUL_PATH")

	if len(consulURL) == 0 || len(consulPath) == 0 {
		panic("Neither environment variables nor Consul configuration is properly set")
	}

	viper.AddRemoteProvider("consul", consulURL, consulPath)
	viper.SetConfigType("json")
	if err := viper.ReadRemoteConfig(); err != nil {
		panic(fmt.Sprintf("Failed to read remote config from Consul: %v", err))
	}
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Sprintf("Failed to unmarshal config: %v", err))
	}

	fmt.Println("Config loaded from Consul")
}

func GetConfig() Config {
	return config
}

func App() *AppConfig {
	return &config.App
}

func Db() *DbConfig {
	return &config.Db
}

func Logger() *LoggerConfig {
	return &config.Logger
}

func Jwt() *JwtConfig {
	return config.Jwt
}

// Helper functions for JWT config to get time.Duration values
func (j *JwtConfig) GetAccessTokenExpiry() time.Duration {
	return time.Duration(j.AccessTokenExpiry) * time.Second
}

func (j *JwtConfig) GetRefreshTokenExpiry() time.Duration {
	return time.Duration(j.RefreshTokenExpiry) * time.Second
}

func Redis() *RedisConfig {
	return config.Redis
}
