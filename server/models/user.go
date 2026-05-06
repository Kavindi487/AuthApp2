package models

import "time"

type User struct {
    ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
    Email        string    `json:"email" gorm:"unique;not null"`
    PasswordHash string    `json:"-" gorm:"not null"`
    CreatedAt    time.Time `json:"created_at"`
}


type LoginLog struct {
    UserID    uint      `bson:"user_id"`
    Email     string    `bson:"email"`
    Status    string    `bson:"status"`    // "success" or "failed"
    IP        string    `bson:"ip"`
    CreatedAt time.Time `bson:"created_at"`
}