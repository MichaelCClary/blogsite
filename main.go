package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/michaelcclary/blogsite/controller"
	"github.com/michaelcclary/blogsite/model"
	"github.com/michaelcclary/blogsite/model/database"
	"log"
	"os"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadDatabase() {
	database.Connect()
	err := database.Database.AutoMigrate(&model.User{})
	if err != nil {
		log.Panicf("Auto Migrate User Failed:%v", err)
	}
	err = database.Database.AutoMigrate(&model.Entry{})
	if err != nil {
		log.Panicf("Auto Migrate Entry Failed:%v", err)
	}
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	port := os.Getenv("PORT")
	err := router.Run(":" + port)
	if err != nil {
		log.Panicf("Auto Migrate Entry Failed:%v", err)
	}
	fmt.Printf("Server running on port %v", port)
}