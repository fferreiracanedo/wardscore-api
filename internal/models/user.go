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

    RiotID   string `json:"riot_id" gorm:"uniqueIndex;not null"`
    GameName string `json:"game_name" gorm:"not null"`
    TagLine  string `json:"tag_line" gorm:"not null"`
    PUUID    string `json:"puuid" gorm:"uniqueIndex"`

    Email     string `json:"email" gorm:"uniqueIndex;not null"`
    AvatarURL string `json:"avatar_url"`
    IsPro     bool   `json:"is_pro" gorm:"default:false"`
    Region    string `json:"region" gorm:"default:'BR1'"`

    // Relacionamentos com ponteiros para evitar referência circular
    Replays  []*Replay  `json:"replays,omitempty" gorm:"foreignKey:UserID"`
    Analyses []*Analysis `json:"analyses,omitempty" gorm:"foreignKey:UserID"`
    Rankings []*Ranking `json:"rankings,omitempty" gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
    return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
    if u.Region == "" {
        u.Region = "BR1"
    }
    return nil
}