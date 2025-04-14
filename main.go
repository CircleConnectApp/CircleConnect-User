package main

import (
	"circleconnect-user/database"
	"circleconnect-user/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found - using system environment variables")
	}

 	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET_KEY not set in environment")
	}
	log.Println("JWT_SECRET_KEY loaded successfully")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"  
	}

	database.InitDB()
	r := gin.Default()
	routes.SetupRoutes(r)
	log.Printf("Server starting on port %s", port)
	r.Run(":" + port)
}