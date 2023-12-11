package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload" // Autoload .env file
)

var (
	MongoURL string
	DbName   string
	DbType   string
	DbHost   string
	DbPass   string
	DbUser   string
	DbPort   string

	RedisHost string
	RedisPort string
	RedisPass string

	ResultBackend string

	BrokerUrl string
)

func init() {
	MongoURL = os.Getenv("DB_URL")
	DbName = os.Getenv("DB_NAME")
	DbType = os.Getenv("DB_TYPE")
	DbHost = os.Getenv("DB_HOST")
	DbPass = os.Getenv("DB_PASS")
	DbUser = os.Getenv("DB_USER")
	DbPort = os.Getenv("DB_PORT")

	RedisHost = os.Getenv("REDIS_HOST")
	RedisPort = os.Getenv("REDIS_PORT")
	RedisPass = os.Getenv("REDIS_PASS")

	ResultBackend = os.Getenv("RESULT_BACKEND")
	BrokerUrl = os.Getenv("BROKER_URL")
}
