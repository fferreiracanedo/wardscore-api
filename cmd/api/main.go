package main

import (
	"log"
	"wardscore-api/internal/config"
	"wardscore-api/internal/database"
	"wardscore-api/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("🚀 Iniciando WardScore API...")

	// 1. Carregar configurações do .env
	config.LoadConfig()

	// 2. Conectar ao PostgreSQL
	database.Connect()

	// 3. Conectar ao Redis
	database.ConnectRedis()

	// 4. Executar migrations
	database.Migrate()

	// 5. Configurar Gin
	if !config.AppConfig.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	
	r := gin.Default()

	// 6. Configurar todas as rotas
	routes.SetupRoutes(r)

	// 7. Iniciar servidor
	log.Printf("🌐 Servidor rodando em http://localhost:%s", config.AppConfig.Port)
	log.Printf("📊 Health check: http://localhost:%s/health", config.AppConfig.Port)
	log.Printf("📖 API Docs: http://localhost:%s/api/v1", config.AppConfig.Port)
	
	// Iniciar servidor HTTP
	if err := r.Run(":" + config.AppConfig.Port); err != nil {
		log.Fatal("❌ Falha ao iniciar servidor:", err)
	}
}