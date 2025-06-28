package services

import (
    "encoding/json"
    "errors"
    "fmt"
    "time"
    "wardscore-api/internal/models"
    "wardscore-api/internal/database"
)

type UserService struct{}

func NewUserService() *UserService {
    return &UserService{}
}

// GetByID busca usuário por ID (com cache Redis)
func (us *UserService) GetByID(id uint) (*models.User, error) {
    cacheKey := fmt.Sprintf("user:%d", id)
    cached, err := database.GetCache(cacheKey)

    if err == nil && cached != "" {
        var user models.User
        if json.Unmarshal([]byte(cached), &user) == nil {
            return &user, nil
        }
    }

    var user models.User
    result := database.DB.First(&user, id)
    if result.Error != nil {
        return nil, errors.New("usuário não encontrado")
    }

    // Salvar no cache por 5 minutos
    if userJSON, err := json.Marshal(user); err == nil {
        database.SetCache(cacheKey, userJSON, 5*time.Minute)
    }

    return &user, nil
}

// GetByRiotID busca usuário por Riot ID
func (us *UserService) GetByRiotID(riotID string) (*models.User, error) {
    var user models.User
    result := database.DB.Where("riot_id = ?", riotID).First(&user)
    if result.Error != nil {
        return nil, errors.New("usuário não encontrado")
    }
    return &user, nil
}

// Create cria novo usuário
func (us *UserService) Create(user *models.User) (*models.User, error) {
    // Verificar duplicatas
    existingUser, _ := us.GetByRiotID(user.RiotID)
    if existingUser != nil {
        return nil, errors.New("usuário com este Riot ID já existe")
    }

    var count int64
    database.DB.Model(&models.User{}).Where("email = ?", user.Email).Count(&count)
    if count > 0 {
        return nil, errors.New("usuário com este email já existe")
    }

    result := database.DB.Create(user)
    if result.Error != nil {
        return nil, result.Error
    }

    return user, nil
}

// Update atualiza usuário
func (us *UserService) Update(user *models.User) (*models.User, error) {
    result := database.DB.Save(user)
    if result.Error != nil {
        return nil, result.Error
    }

    // Limpar cache
    cacheKey := fmt.Sprintf("user:%d", user.ID)
    database.DeleteCache(cacheKey)

    return user, nil
}

// Delete remove usuário
func (us *UserService) Delete(id uint) error {
    result := database.DB.Delete(&models.User{}, id)
    if result.Error != nil {
        return result.Error
    }

    if result.RowsAffected == 0 {
        return errors.New("usuário não encontrado")
    }

    // Limpar cache
    cacheKey := fmt.Sprintf("user:%d", id)
    database.DeleteCache(cacheKey)

    return nil
}

// GetAll lista usuários com paginação
func (us *UserService) GetAll(page, limit int) ([]models.User, int64, error) {
    var users []models.User
    var total int64

    // Contar total
    database.DB.Model(&models.User{}).Count(&total)

    // Buscar com paginação
    offset := (page - 1) * limit
    result := database.DB.Offset(offset).Limit(limit).Find(&users)

    if result.Error != nil {
        return nil, 0, result.Error
    }

    return users, total, nil
}

// GetUserWithReplays busca usuário com replays
func (us *UserService) GetUserWithReplays(id uint) (*models.User, error) {
    var user models.User
    result := database.DB.Preload("Replays").First(&user, id)
    if result.Error != nil {
        return nil, errors.New("usuário não encontrado")
    }
    return &user, nil
}