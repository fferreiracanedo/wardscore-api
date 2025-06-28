package models

import (
	"time"

	"gorm.io/gorm"
)

type Ranking struct {
    ID        uint           `json:"id" gorm:"primaryKey"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

    WardScore    float64 `json:"ward_score" gorm:"not null"`
    Position     int     `json:"position" gorm:"not null"`
    Region       string  `json:"region" gorm:"not null;index"`
    Tier         string  `json:"tier"`
    Division     string  `json:"division"`
    
    GamesPlayed    int     `json:"games_played" gorm:"default:0"`
    AverageScore   float64 `json:"average_score"`
    BestScore      float64 `json:"best_score"`
    WorstScore     float64 `json:"worst_score"`
    TotalWards     int     `json:"total_wards"`
    TotalVision    int     `json:"total_vision"`
    
    LastUpdated time.Time `json:"last_updated"`
    Season      string    `json:"season" gorm:"default:'2024'"`

    // Relacionamento com ponteiro
    UserID uint  `json:"user_id" gorm:"not null;index"`
    User   *User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

func (Ranking) TableName() string {
    return "rankings"
}

func (r *Ranking) BeforeCreate(tx *gorm.DB) error {
    if r.LastUpdated.IsZero() {
        r.LastUpdated = time.Now()
    }
    if r.Season == "" {
        r.Season = "2024"
    }
    return nil
}

func (r *Ranking) BeforeUpdate(tx *gorm.DB) error {
    r.LastUpdated = time.Now()
    return nil
}

func (r *Ranking) GetTierFromScore() (string, string) {
    switch {
    case r.WardScore >= 95:
        return "Challenger", "I"
    case r.WardScore >= 90:
        return "Grandmaster", "I"
    case r.WardScore >= 85:
        return "Master", "I"
    case r.WardScore >= 80:
        return "Diamond", "I"
    case r.WardScore >= 75:
        return "Diamond", "II"
    case r.WardScore >= 70:
        return "Platinum", "I"
    case r.WardScore >= 65:
        return "Platinum", "II"
    case r.WardScore >= 60:
        return "Gold", "I"
    case r.WardScore >= 55:
        return "Gold", "II"
    case r.WardScore >= 50:
        return "Silver", "I"
    case r.WardScore >= 45:
        return "Silver", "II"
    default:
        return "Bronze", "I"
    }
}

func (r *Ranking) UpdateTierFromScore() {
    r.Tier, r.Division = r.GetTierFromScore()
}