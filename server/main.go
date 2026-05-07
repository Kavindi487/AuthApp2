package main

import (
    "server/handlers"
    "server/middleware"
    "server/models"
    "server/repository"
    "server/services"
    "context"
    "fmt"
    "log"
    "time"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func main() {
    // MySQL
    dsn := "root:123456789@tcp(127.0.0.1:3307)/auth_db?parseTime=true"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to MySQL:", err)
    }
    fmt.Println("MySQL connected!")
    db.AutoMigrate(&models.User{}, &models.Note{})

    // MongoDB
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    mongoClient, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatal("MongoDB connection failed:", err)
    }
    err = mongoClient.Ping(ctx, nil)
    if err != nil {
        log.Fatal("MongoDB ping failed:", err)
    }
    fmt.Println("MongoDB connected!")
    logCollection := mongoClient.Database("auth_db").Collection("login_logs")

    // Wire layers
    userRepo    := repository.NewUserRepository(db)
    logRepo     := repository.NewLogRepository(logCollection)
    noteRepo    := repository.NewNoteRepository(db)
    authService := services.NewAuthService(userRepo, logRepo)
    noteService := services.NewNoteService(noteRepo)
    authHandler := handlers.NewAuthHandler(authService)
    noteHandler := handlers.NewNoteHandler(noteService)

    r := gin.Default()

    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:4200"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    api := r.Group("/api")
    {
        // Public routes
        api.POST("/register", authHandler.Register)
        api.POST("/login",    authHandler.Login)

        // Protected routes — JWT middleware runs first
        protected := api.Group("/")
        protected.Use(middleware.JWTMiddleware())
        {
            protected.POST("/notes",       noteHandler.CreateNote)
            protected.GET("/notes",        noteHandler.GetMyNotes)
            protected.DELETE("/notes/:id", noteHandler.DeleteNote)
        }
    }

    fmt.Println("Server running on http://localhost:8080")
    r.Run(":8080")
}