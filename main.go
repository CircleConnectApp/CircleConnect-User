package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
     "github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"circleconnect-user/controllers"
	"circleconnect-user/database"
	"circleconnect-user/routes"
)

var ctx = context.Background()
var rdb *redis.Client

func main() {
	
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", 
		Password: "",               
		DB:       0,                
	})

	
	ping, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err.Error())
		return
	}
	fmt.Println("Connected to Redis:", ping)

	controllers.RedisClient = client

	type Person struct {
		ID         string 
		Name       string `json:"name"`
		Age        int    `json:"age"`
		Occupation string `json:"occupation"`
	}

	elliotID:=uuid.NewString()

	person := Person{

		 ID:       elliotID ,
		Name:       "Elliot",
		Age:        30,
		Occupation: "Staff Software Engineer",
	}
	jsonString, err := json.Marshal(person)
	if err != nil {
		fmt.Printf("failed to marshal: %s\n", err.Error())
		return
	}

    elliotKey:=fmt.Sprintf("person:%s",&elliotID)
	err = client.Set(context.Background(), elliotKey, jsonString, 0).Err()
	if err != nil {
		fmt.Printf("Failed to set value in the redis instance: %s\n", err.Error())
		return
	}

	
	val, err := client.Get(context.Background(),elliotKey).Result()
	if err != nil {
		fmt.Printf("failed to get value from redis: %s\n", err.Error())
		return
	}

	fmt.Printf("value retrieved from redis: %s\n", val)

	if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
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
	r.POST("/login", controllers.Login) 

	
	routes.SetupRoutes(r)


	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
