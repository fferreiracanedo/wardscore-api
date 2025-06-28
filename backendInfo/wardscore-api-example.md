# 🔧 WARDSCORE API - EXEMPLOS DE CÓDIGO GO

## 📁 **ESTRUTURA INICIAL**

### `cmd/api/main.go`

```go
package main

import (
    "log"
    "wardscore-api/internal/database"
    "wardscore-api/internal/routes"
    "wardscore-api/internal/config"

    "github.com/gin-gonic/gin"
)

func main() {
    // Carregar configurações
    config.LoadConfig()

    // Conectar ao banco
    database.Connect()

    // Executar migrations
    database.Migrate()

    // Configurar Gin
    r := gin.Default()

    // Configurar rotas
    routes.SetupRoutes(r)

    // Iniciar servidor
    log.Println("🚀 Servidor iniciado na porta 8080")
    r.Run(":8080")
}
```

### `internal/config/config.go`

```go
package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    Port         string
    Host         string
    DatabaseURL  string
    JWTSecret    string
    RiotClientID string
    RiotSecret   string
}

var AppConfig Config

func LoadConfig() {
    err := godotenv.Load()
    if err != nil {
        log.Println("⚠️ Arquivo .env não encontrado, usando variáveis do sistema")
    }

    AppConfig = Config{
        Port:         getEnv("PORT", "8080"),
        Host:         getEnv("HOST", "0.0.0.0"),
        DatabaseURL:  getEnv("DATABASE_URL", ""),
        JWTSecret:    getEnv("JWT_SECRET", ""),
        RiotClientID: getEnv("RIOT_CLIENT_ID", ""),
        RiotSecret:   getEnv("RIOT_CLIENT_SECRET", ""),
    }

    log.Println("✅ Configurações carregadas")
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
```

### `internal/database/database.go`

```go
package database

import (
    "log"
    "wardscore-api/internal/config"
    "wardscore-api/internal/models"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    var err error
    DB, err = gorm.Open(postgres.Open(config.AppConfig.DatabaseURL), &gorm.Config{})

    if err != nil {
        log.Fatal("❌ Falha ao conectar com banco de dados:", err)
    }

    log.Println("✅ Conectado ao banco de dados")
}

func Migrate() {
    err := DB.AutoMigrate(
        &models.User{},
        &models.Replay{},
        &models.Analysis{},
    )

    if err != nil {
        log.Fatal("❌ Falha na migration:", err)
    }

    log.Println("✅ Migrations executadas")
}
```

### `internal/models/user.go`

```go
package models

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    ID        uint           `json:"id" gorm:"primaryKey"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

    // Riot Data
    RiotID     string `json:"riot_id" gorm:"uniqueIndex"`
    GameName   string `json:"game_name"`
    TagLine    string `json:"tag_line"`
    PUUID      string `json:"puuid" gorm:"uniqueIndex"`

    // Profile
    Email     string `json:"email" gorm:"uniqueIndex"`
    AvatarURL string `json:"avatar_url"`
    IsPro     bool   `json:"is_pro" gorm:"default:false"`

    // Relations
    Replays  []Replay  `json:"replays" gorm:"foreignKey:UserID"`
    Analyses []Analysis `json:"analyses" gorm:"foreignKey:UserID"`
}
```

### `internal/models/replay.go`

```go
package models

import (
    "time"
    "gorm.io/gorm"
)

type ReplayStatus string

const (
    StatusUploaded   ReplayStatus = "uploaded"
    StatusProcessing ReplayStatus = "processing"
    StatusCompleted  ReplayStatus = "completed"
    StatusFailed     ReplayStatus = "failed"
)

type Replay struct {
    ID        uint           `json:"id" gorm:"primaryKey"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

    // File Info
    FileName string `json:"file_name"`
    FilePath string `json:"file_path"`
    FileSize int64  `json:"file_size"`

    // Game Info
    MatchID     string `json:"match_id" gorm:"uniqueIndex"`
    GameMode    string `json:"game_mode"`
    GameVersion string `json:"game_version"`
    Duration    int    `json:"duration"` // em segundos

    Status ReplayStatus `json:"status" gorm:"default:'uploaded'"`

    // Relations
    UserID   uint     `json:"user_id"`
    User     User     `json:"user" gorm:"foreignKey:UserID"`
    Analysis Analysis `json:"analysis" gorm:"foreignKey:ReplayID"`
}
```

### `internal/models/analysis.go`

```go
package models

import (
    "time"
    "gorm.io/gorm"
    "encoding/json"
)

type Analysis struct {
    ID        uint           `json:"id" gorm:"primaryKey"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

    // Score Data
    WardScore float64 `json:"ward_score"`
    Rank      string  `json:"rank"`

    // Game Stats (JSON)
    GameStats json.RawMessage `json:"game_stats" gorm:"type:jsonb"`
    Insights  json.RawMessage `json:"insights" gorm:"type:jsonb"`

    // Relations
    UserID   uint   `json:"user_id"`
    User     User   `json:"user" gorm:"foreignKey:UserID"`
    ReplayID uint   `json:"replay_id" gorm:"uniqueIndex"`
    Replay   Replay `json:"replay" gorm:"foreignKey:ReplayID"`
}
```

### `internal/routes/routes.go`

```go
package routes

import (
    "wardscore-api/internal/controllers"
    "wardscore-api/internal/middleware"

    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func SetupRoutes(r *gin.Engine) {
    // CORS
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: true,
    }))

    // Health Check
    r.GET("/health", controllers.HealthCheck)

    // API v1
    v1 := r.Group("/api/v1")
    {
        // Public routes
        auth := v1.Group("/auth")
        {
            auth.POST("/register", controllers.Register)
            auth.POST("/login", controllers.Login)
            auth.GET("/riot", controllers.RiotAuth)
            auth.POST("/riot/callback", controllers.RiotCallback)
        }

        // Protected routes
        protected := v1.Group("/")
        protected.Use(middleware.AuthMiddleware())
        {
            // Users
            users := protected.Group("/users")
            {
                users.GET("/profile", controllers.GetProfile)
                users.PUT("/profile", controllers.UpdateProfile)
                users.DELETE("/profile", controllers.DeleteProfile)
            }

            // Replays
            replays := protected.Group("/replays")
            {
                replays.POST("/upload", controllers.UploadReplay)
                replays.GET("/", controllers.GetReplays)
                replays.GET("/:id", controllers.GetReplay)
                replays.DELETE("/:id", controllers.DeleteReplay)
            }

            // Analysis
            analysis := protected.Group("/analysis")
            {
                analysis.GET("/:id", controllers.GetAnalysis)
            }
        }
    }
}
```

### `internal/controllers/health.go`

```go
package controllers

import (
    "net/http"
    "wardscore-api/internal/database"

    "github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
    // Verificar conexão com banco
    sqlDB, err := database.DB.DB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status": "error",
            "message": "Database connection failed",
        })
        return
    }

    if err := sqlDB.Ping(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status": "error",
            "message": "Database ping failed",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status": "ok",
        "service": "wardscore-api",
        "version": "1.0.0",
        "database": "connected",
    })
}
```

### `internal/middleware/auth.go`

```go
package middleware

import (
    "net/http"
    "strings"
    "wardscore-api/internal/config"
    "wardscore-api/internal/utils"

    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Authorization header required",
            })
            c.Abort()
            return
        }

        tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

        claims, err := utils.ValidateJWT(tokenString, config.AppConfig.JWTSecret)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Invalid token",
            })
            c.Abort()
            return
        }

        c.Set("user_id", claims["user_id"])
        c.Next()
    }
}
```

### `internal/utils/jwt.go`

```go
package utils

import (
    "errors"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID uint, secret string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
        "iat":     time.Now().Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}

func ValidateJWT(tokenString, secret string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return []byte(secret), nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }

    return nil, errors.New("invalid token")
}
```

## 🚀 **COMANDOS PARA COMEÇAR**

```bash
# 1. Criar pasta da API
mkdir wardscore-api
cd wardscore-api

# 2. Inicializar módulo Go
go mod init wardscore-api

# 3. Instalar dependências principais
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
go get github.com/joho/godotenv
go get github.com/gin-contrib/cors

# 4. Criar estrutura de pastas
mkdir -p cmd/api
mkdir -p internal/{config,database,models,controllers,middleware,utils,services}
mkdir -p pkg
mkdir -p configs

# 5. Copiar arquivos de exemplo acima
# 6. Configurar .env
# 7. Testar
go run cmd/api/main.go
```

## 🔧 **DOCKER COMPOSE PARA DESENVOLVIMENTO**

### `docker-compose.yml`

```yaml
version: "3.8"

services:
  postgres:
    image: postgres:15
    environment:
      POSTGRES_DB: wardscore
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres123
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"

volumes:
  postgres_data:
```

Execute com: `docker-compose up -d`
