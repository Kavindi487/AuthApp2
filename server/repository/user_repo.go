package repository

import (
    "server/models"
    "gorm.io/gorm"
)

type UserRepository struct {
    DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
    return r.DB.Create(user).Error
}

func (r *UserRepository) FindUserByEmail(email string) (*models.User, error) {
    var user models.User
    result := r.DB.Where("email = ?", email).First(&user)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}