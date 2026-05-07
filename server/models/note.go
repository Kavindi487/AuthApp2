package models

import "time"

type Note struct {
    ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
    UserID    uint      `json:"user_id" gorm:"not null"`
    Title     string    `json:"title" gorm:"not null"`
    Content   string    `json:"content" gorm:"not null"`
    CreatedAt time.Time `json:"created_at"`
}