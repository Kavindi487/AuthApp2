package services

import (
    "server/models"
    "server/repository"
    "errors"

    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
    "time"
)

// Change this to any long random string in production
var jwtSecret = []byte("my-secret-key-change-in-production")

type AuthService struct {
    UserRepo *repository.UserRepository
    LogRepo  *repository.LogRepository
}

func NewAuthService(userRepo *repository.UserRepository, logRepo *repository.LogRepository) *AuthService {
    return &AuthService{UserRepo: userRepo, LogRepo: logRepo}
}

func (s *AuthService) RegisterUser(email, password string) (*models.User, error) {
    existing, _ := s.UserRepo.FindUserByEmail(email)
    if existing != nil {
        return nil, errors.New("email already registered")
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user := &models.User{
        Email:        email,
        PasswordHash: string(hashedPassword),
    }

    err = s.UserRepo.CreateUser(user)
    if err != nil {
        return nil, err
    }

    return user, nil
}

func (s *AuthService) LoginUser(email, password, ip string) (string, error) {
    // Find user by email
    user, err := s.UserRepo.FindUserByEmail(email)
    if err != nil {
        s.LogRepo.InsertLoginLog(0, email, "failed", ip)
        return "", errors.New("invalid email or password")
    }

    // Compare submitted password with stored hash
    err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
    if err != nil {
        s.LogRepo.InsertLoginLog(user.ID, email, "failed", ip)
        return "", errors.New("invalid email or password")
    }

    // Generate JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "email":   user.Email,
        "exp":     time.Now().Add(24 * time.Hour).Unix(), // expires in 24 hours
    })

    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }

    // Log success to MongoDB
    s.LogRepo.InsertLoginLog(user.ID, email, "success", ip)

    return tokenString, nil
}