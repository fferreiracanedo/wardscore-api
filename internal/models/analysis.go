package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Analysis struct {
    ID        uint           `json:"id" gorm:"primaryKey"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

    WardScore float64 `json:"ward_score" gorm:"not null"`
    Rank      string  `json:"rank" gorm:"not null"`

    WardsPlaced         int     `json:"wards_placed"`
    WardsDestroyed      int     `json:"wards_destroyed"`
    VisionScore         int     `json:"vision_score"`
    ControlWardsPlaced  int     `json:"control_wards_placed"`
    WardsPerMinute      float64 `json:"wards_per_minute"`
    VisionControlRatio  float64 `json:"vision_control_ratio"`

    GameStats   json.RawMessage `json:"game_stats,omitempty" gorm:"type:jsonb"`
    Insights    json.RawMessage `json:"insights,omitempty" gorm:"type:jsonb"`
    Suggestions json.RawMessage `json:"suggestions,omitempty" gorm:"type:jsonb"`
    HeatmapData json.RawMessage `json:"heatmap_data,omitempty" gorm:"type:jsonb"`

    // Relacionamentos com ponteiros
    UserID   uint    `json:"user_id" gorm:"not null;index"`
    User     *User   `json:"user,omitempty" gorm:"foreignKey:UserID"`
    ReplayID uint    `json:"replay_id" gorm:"uniqueIndex;not null"`
    Replay   *Replay `json:"replay,omitempty" gorm:"foreignKey:ReplayID"`
}

func (Analysis) TableName() string {
    return "analyses"
}

func (a *Analysis) GetRankFromScore() string {
    switch {
    case a.WardScore >= 95:
        return "S+"
    case a.WardScore >= 90:
        return "S"
    case a.WardScore >= 85:
        return "A+"
    case a.WardScore >= 80:
        return "A"
    case a.WardScore >= 70:
        return "B+"
    case a.WardScore >= 60:
        return "B"
    default:
        return "C"
    }
}

func (a *Analysis) BeforeCreate(tx *gorm.DB) error {
    if a.Rank == "" {
        a.Rank = a.GetRankFromScore()
    }
    
    if a.WardsPerMinute == 0 && a.WardsPlaced > 0 {
        var replay Replay
        if err := tx.First(&replay, a.ReplayID).Error; err == nil {
            if replay.Duration > 0 {
                a.WardsPerMinute = float64(a.WardsPlaced) / (float64(replay.Duration) / 60.0)
            }
        }
    }
    
    return nil
}