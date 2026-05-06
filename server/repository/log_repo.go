package repository

import (
    "server/models"
    "context"
    "time"

    "go.mongodb.org/mongo-driver/v2/mongo"
)

type LogRepository struct {
    Collection *mongo.Collection
}

func NewLogRepository(collection *mongo.Collection) *LogRepository {
    return &LogRepository{Collection: collection}
}

func (r *LogRepository) InsertLoginLog(userID uint, email, status, ip string) error {
    log := models.LoginLog{
        UserID:    userID,
        Email:     email,
        Status:    status,
        IP:        ip,
        CreatedAt: time.Now(),
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := r.Collection.InsertOne(ctx, log)
    return err
}