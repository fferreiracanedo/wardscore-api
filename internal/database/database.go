package database

import (
	"log"
	"time"
	"wardscore-api/internal/config"
	"wardscore-api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


var DB *gorm.DB


func Connect(){
	var err error

	gormConfig := &gorm.Config{}

	if config.AppConfig.Debug {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	DB, err = gorm.Open(postgres.Open(config.AppConfig.DatabaseURL), gormConfig)

	if err != nil {
		log.Fatal("❌ Falha ao conectar com PostgreSQL:", err)
	}

	sqlDB, err := DB.DB()
    if err != nil {
        log.Fatal("❌ Falha ao configurar connection pool:", err)
    }

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)


	log.Println("✅ Conectado ao PostgreSQL")
}

func Migrate() {
	log.Println("⏳ Executando migrations...")


	err := DB.AutoMigrate(
		&models.User{},
        &models.Replay{},
        &models.Analysis{},
	)

	if err != nil {
        log.Fatal("❌ Falha nas migrations:", err)
    }

	log.Println("✅ Migrations executadas com sucesso")
}


func GetDB()  *gorm.DB{
	return DB
}