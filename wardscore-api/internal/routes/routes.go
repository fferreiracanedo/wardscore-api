package routes

import (
    "net/http"
    "wardscore-api/internal/controllers"
    "wardscore-api/internal/services"
    
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func SetupRoutes(r *gin.Engine) {
    // Configurar CORS
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-User-ID"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    r.Use(gin.Recovery())

    // Health check
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status":    "ok",
            "message":   "WardScore API funcionando! 🚀",
            "version":   "1.0.0",
            "endpoints": gin.H{
                "users":    "/api/v1/users",
                "replays":  "/api/v1/replays", 
                "analysis": "/api/v1/analysis",
            },
        })
    })

    // Inicializar services
    userService := services.NewUserService()
    replayService := services.NewReplayService()
    analysisService := services.NewAnalysisService()

    // Inicializar controllers
    userController := controllers.NewUserController(userService)
    replayController := controllers.NewReplayController(replayService)
    analysisController := controllers.NewAnalysisController(analysisService)

    // Grupo de rotas da API
    api := r.Group("/api/v1")
    {
        // ===== ROTAS DE USUÁRIO =====
        users := api.Group("/users")
        {
            users.POST("", userController.CreateUser)           // Criar usuário
            users.GET("", userController.GetAllUsers)           // Listar usuários
            users.GET("/profile", userController.GetProfile)    // Perfil do usuário
            users.PUT("/profile", userController.UpdateProfile) // Atualizar perfil
            users.DELETE("/:id", userController.DeleteUser)     // Deletar usuário
        }

        // ===== ROTAS DE REPLAY =====
        replays := api.Group("/replays")
        {
            replays.POST("/upload", replayController.UploadReplay)  // Upload replay
            replays.GET("", replayController.GetReplays)            // Listar replays
            replays.GET("/:id", replayController.GetReplay)         // Buscar replay específico
            replays.PUT("/:id", replayController.UpdateReplay)      // Atualizar replay
            replays.DELETE("/:id", replayController.DeleteReplay)   // Deletar replay
        }

        // ===== ROTAS DE ANÁLISE =====
        analysis := api.Group("/analysis")
        {
            analysis.GET("/:id", analysisController.GetAnalysis)           // Buscar análise
            analysis.POST("/process/:replay_id", analysisController.ProcessReplay) // Processar replay
            analysis.GET("/user/:user_id", analysisController.GetUserAnalyses)     // Análises do usuário
        }

        // ===== ROTAS DE ESTATÍSTICAS =====
        stats := api.Group("/stats")
        {
            stats.GET("/dashboard", func(c *gin.Context) {
                c.JSON(200, gin.H{
                    "message": "Dashboard stats - implementar depois",
                })
            })
        }

        // ===== ROTAS DE RANKING =====
        ranking := api.Group("/ranking")
        {
            ranking.GET("/global", func(c *gin.Context) {
                c.JSON(200, gin.H{
                    "message": "Global ranking - implementar depois",
                })
            })
            ranking.GET("/region/:region", func(c *gin.Context) {
                c.JSON(200, gin.H{
                    "message": "Regional ranking - implementar depois",
                    "region": c.Param("region"),
                })
            })
        }
    }
}