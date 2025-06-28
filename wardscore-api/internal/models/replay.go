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

    FileName     string `json:"file_name" gorm:"not null"`
    OriginalName string `json:"original_name" gorm:"not null"`
    FilePath     string `json:"file_path" gorm:"not null"`
    FileSize     int64  `json:"file_size"`

    MatchID     string `json:"match_id" gorm:"uniqueIndex;not null"`
    GameMode    string `json:"game_mode"`
    GameVersion string `json:"game_version"`
    Duration    int    `json:"duration"`
    Champion    string `json:"champion"`
    Role        string `json:"role"`
    Queue       string `json:"queue"`

    Status ReplayStatus `json:"status" gorm:"default:'uploaded'"`

    UploadedAt  time.Time  `json:"uploaded_at"`
    ProcessedAt *time.Time `json:"processed_at,omitempty"`

    // Relacionamentos com ponteiros
    UserID   uint      `json:"user_id" gorm:"not null;index"`
    User     *User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
    Analysis *Analysis `json:"analysis,omitempty" gorm:"foreignKey:ReplayID"`
}

func (Replay) TableName() string {
    return "replays"
}

func (r *Replay) BeforeCreate(tx *gorm.DB) error {
    if r.UploadedAt.IsZero() {
        r.UploadedAt = time.Now()
    }
    return nil
}

func (r *Replay) IsProcessed() bool {
    return r.Status == StatusCompleted
}

func (r *Replay) CanBeProcessed() bool {
    return r.Status == StatusUploaded || r.Status == StatusFailed
}

func (r *Replay) MarkAsProcessing() {
    r.Status = StatusProcessing
}

func (r *Replay) MarkAsCompleted() {
    r.Status = StatusCompleted
    now := time.Now()
    r.ProcessedAt = &now
}

func (r *Replay) MarkAsFailed() {
    r.Status = StatusFailed
}