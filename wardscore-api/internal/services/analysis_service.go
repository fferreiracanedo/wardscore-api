package services

import (
	"errors"
	"math/rand"
	"wardscore-api/internal/database"
	"wardscore-api/internal/models"
)

type AnalysisService struct{}

func NewAnalysisService() *AnalysisService {
    return &AnalysisService{}
}

// GetByID busca análise por ID
func (as *AnalysisService) GetByID(id uint) (*models.Analysis, error) {
    var analysis models.Analysis
    result := database.DB.Preload("User").Preload("Replay").First(&analysis, id)
    if result.Error != nil {
        return nil, errors.New("análise não encontrada")
    }
    return &analysis, nil
}

// GetByReplayID busca análise por replay ID
func (as *AnalysisService) GetByReplayID(replayID uint) (*models.Analysis, error) {
    var analysis models.Analysis
    result := database.DB.Where("replay_id = ?", replayID).
        Preload("User").Preload("Replay").First(&analysis)
    if result.Error != nil {
        return nil, errors.New("análise não encontrada")
    }
    return &analysis, nil
}

// GetByUserID busca análises do usuário
func (as *AnalysisService) GetByUserID(userID uint) ([]models.Analysis, error) {
    var analyses []models.Analysis
    result := database.DB.Where("user_id = ?", userID).
        Preload("Replay").
        Order("created_at DESC").
        Find(&analyses)
    
    if result.Error != nil {
        return nil, result.Error
    }
    
    return analyses, nil
}

// ProcessReplay processa um replay e cria análise
func (as *AnalysisService) ProcessReplay(replayID uint) (*models.Analysis, error) {
    // Buscar replay
    var replay models.Replay
    result := database.DB.First(&replay, replayID)
    if result.Error != nil {
        return nil, errors.New("replay não encontrado")
    }

    // Verificar se já foi processado
    var existingAnalysis models.Analysis
    if database.DB.Where("replay_id = ?", replayID).First(&existingAnalysis).Error == nil {
        return nil, errors.New("replay já foi processado")
    }

    // Marcar replay como processando
    replay.MarkAsProcessing()
    database.DB.Save(&replay)

    // SIMULAÇÃO: Criar análise com dados aleatórios
    // Em produção, aqui seria o processamento real do arquivo .rofl
    analysis := &models.Analysis{
        UserID:              replay.UserID,
        ReplayID:            replayID,
        WardScore:           50 + rand.Float64()*50, // 50-100
        WardsPlaced:         10 + rand.Intn(20),     // 10-30
        WardsDestroyed:      5 + rand.Intn(15),      // 5-20
        VisionScore:         20 + rand.Intn(80),     // 20-100
        ControlWardsPlaced:  3 + rand.Intn(7),       // 3-10
    }

    // Calcular métricas derivadas
    if replay.Duration > 0 {
        analysis.WardsPerMinute = float64(analysis.WardsPlaced) / (float64(replay.Duration) / 60.0)
    }
    
    if analysis.WardsPlaced > 0 {
        analysis.VisionControlRatio = float64(analysis.WardsDestroyed) / float64(analysis.WardsPlaced)
    }

    // Criar análise
    result = database.DB.Create(analysis)
    if result.Error != nil {
        // Marcar replay como falhou
        replay.MarkAsFailed()
        database.DB.Save(&replay)
        return nil, result.Error
    }

    // Marcar replay como completado
    replay.MarkAsCompleted()
    database.DB.Save(&replay)

    return analysis, nil
}

// Create cria nova análise
func (as *AnalysisService) Create(analysis *models.Analysis) (*models.Analysis, error) {
    result := database.DB.Create(analysis)
    if result.Error != nil {
        return nil, result.Error
    }
    return analysis, nil
}