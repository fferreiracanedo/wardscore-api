package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)


type Config struct {
	Port        string
    Host        string
    DatabaseURL string
    RedisURL    string
    JWTSecret   string  // Corrigido para JWTSecret
    Debug       bool

    // Adicionar campos Riot que estão sendo usados
    RiotClientID     string
    RiotClientSecret string
    RiotAPIKey       string
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Arquivo .env não encontrado, usando variáveis do sistema")
	}

	AppConfig = Config{
        Port:        getEnv("PORT", "8080"),
        Host:        getEnv("HOST", "0.0.0.0"),
        DatabaseURL: getEnv("DATABASE_URL", ""),
        RedisURL:    getEnv("REDIS_URL", "redis://localhost:6379"),
        JWTSecret:   getEnv("JWT_SECRET", ""),
        Debug:       getEnvAsBool("DEBUG", true),
        RiotClientID:     getEnv("RIOT_CLIENT_ID", ""),
        RiotClientSecret: getEnv("RIOT_CLIENT_SECRET", ""),
        RiotAPIKey:       getEnv("RIOT_API_KEY", ""),
    }


	if AppConfig.DatabaseURL == "" {
		log.Fatal("❌ DATABASE_URL é obrigatória")
	}

	if AppConfig.JWTSecret == "" {
		log.Fatal("❌ JWT_SECRET é obrigatória")
	}

	log.Println("✅ Configurações carregadas com sucesso")

}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}


func getEnvAsBool(key string, defaultValue bool) bool {
	valStr := getEnv(key, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	return defaultValue

}