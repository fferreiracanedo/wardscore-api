package services

import (
	"errors"
	"wardscore-api/internal/database"
	"wardscore-api/internal/models"
)

type ReplayService struct{}

func NewReplayService() *ReplayService {
    return &ReplayService{}
}

// GetByID busca replay por ID
func (rs *ReplayService) GetByID(id uint) (*models.Replay, error) {
    var replay models.Replay
    result := database.DB.Preload("User").Preload("Analysis").First(&replay, id)
    if result.Error != nil {
        return nil, errors.New("replay não encontrado")
    }
    return &replay, nil
}

// GetByUserID busca replays por usuário
func (rs *ReplayService) GetByUserID(userID uint) ([]models.Replay, error) {
    var replays []models.Replay
    result := database.DB.Where("user_id = ?", userID).
        Preload("Analysis").
        Order("created_at DESC").
        Find(&replays)
    
    if result.Error != nil {
        return nil, result.Error
    }
    
    return replays, nil
}

// Create cria novo replay
func (rs *ReplayService) Create(replay *models.Replay) (*models.Replay, error) {
    // Verificar se já existe replay com mesmo Match ID
    var count int64
    database.DB.Model(&models.Replay{}).Where("match_id = ?", replay.MatchID).Count(&count)
    if count > 0 {
        return nil, errors.New("replay com este Match ID já existe")
    }

    result := database.DB.Create(replay)
    if result.Error != nil {
        return nil, result.Error
    }

    return replay, nil
}

// Update atualiza replay
func (rs *ReplayService) Update(replay *models.Replay) (*models.Replay, error) {
    result := database.DB.Save(replay)
    if result.Error != nil {
        return nil, result.Error
    }
    return replay, nil
}

// Delete remove replay
func (rs *ReplayService) Delete(id uint) error {
    result := database.DB.Delete(&models.Replay{}, id)
    if result.Error != nil {
        return result.Error
    }
    
    if result.RowsAffected == 0 {
        return errors.New("replay não encontrado")
    }
    
    return nil
}

// GetPendingReplays busca replays para processar
func (rs *ReplayService) GetPendingReplays() ([]models.Replay, error) {
    var replays []models.Replay
    result := database.DB.Where("status = ?", models.StatusUploaded).
        Preload("User").
        Find(&replays)
    
    if result.Error != nil {
        return nil, result.Error
    }
    
    return replays, nil
}

// UpdateStatus atualiza status do replay
func (rs *ReplayService) UpdateStatus(id uint, status models.ReplayStatus) error {
    result := database.DB.Model(&models.Replay{}).Where("id = ?", id).Update("status", status)
    if result.Error != nil {
        return result.Error
    }
    
    if result.RowsAffected == 0 {
        return errors.New("replay não encontrado")
    }
    
    return nil
}