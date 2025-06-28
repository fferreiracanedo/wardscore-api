package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// AchievementCategory enum para categoria de conquista
type AchievementCategory string

const (
    CategoryWard        AchievementCategory = "ward"
    CategoryVision      AchievementCategory = "vision"
    CategoryImprovement AchievementCategory = "improvement"
    CategoryStreak      AchievementCategory = "streak"
    CategorySpecial     AchievementCategory = "special"
)

// AchievementRarity enum para raridade de conquista
type AchievementRarity string

const (
    RarityCommon    AchievementRarity = "common"
    RarityRare      AchievementRarity = "rare"
    RarityEpic      AchievementRarity = "epic"
    RarityLegendary AchievementRarity = "legendary"
)

// Achievement representa uma conquista
type Achievement struct {
    ID        uint           `json:"id" gorm:"primaryKey"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

    // Dados da conquista
    Title       string                `json:"title" gorm:"not null"`
    Description string                `json:"description" gorm:"not null"`
    Category    AchievementCategory   `json:"category" gorm:"not null"`
    Rarity      AchievementRarity     `json:"rarity" gorm:"not null"`
    Icon        string                `json:"icon"`
    
    // Critério e recompensa (JSON)
    Requirement json.RawMessage `json:"requirement" gorm:"type:jsonb"` // { "type": "score", "value": 80 }
    Reward      json.RawMessage `json:"reward" gorm:"type:jsonb"`      // { "xp": 100, "badge": "ward_master" }

    // Relacionamentos
    UserAchievements []UserAchievement `json:"user_achievements,omitempty" gorm:"foreignKey:AchievementID"`
}

func (Achievement) TableName() string {
    return "achievements"
}

// UserAchievement representa conquista de um usuário
type UserAchievement struct {
    ID        uint           `json:"id" gorm:"primaryKey"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

    // Dados do progresso
    Progress    int        `json:"progress" gorm:"default:0"` // 0-100
    UnlockedAt  *time.Time `json:"unlocked_at,omitempty"`
    
    // Relacionamentos
    UserID        uint        `json:"user_id" gorm:"not null;index"`
    User          User        `json:"user,omitempty" gorm:"foreignKey:UserID"`
    AchievementID uint        `json:"achievement_id" gorm:"not null;index"`
    Achievement   Achievement `json:"achievement,omitempty" gorm:"foreignKey:AchievementID"`
}

func (UserAchievement) TableName() string {
    return "user_achievements"
}

// IsUnlocked verifica se a conquista foi desbloqueada
func (ua *UserAchievement) IsUnlocked() bool {
    return ua.UnlockedAt != nil
}

// Unlock desbloqueia a conquista
func (ua *UserAchievement) Unlock() {
    now := time.Now()
    ua.UnlockedAt = &now
    ua.Progress = 100
}