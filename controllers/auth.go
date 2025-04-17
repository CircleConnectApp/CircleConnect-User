package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"circleconnect-user/database"
	"circleconnect-user/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/golang-jwt/jwt/v4"
)

var RedisClient *redis.Client
var ctx = context.Background()

func Login(c *gin.Context) {
	var loginData struct {
		Email string `json:"email"`
	}

	if err := c.BindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	cacheKey := fmt.Sprintf("user:login:%s", loginData.Email)
	cachedUser, err := RedisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		// Return cached user if found
		var user models.User
		if err := json.Unmarshal([]byte(cachedUser), &user); err == nil {
			token := generateToken(user)
			c.JSON(http.StatusOK, gin.H{"user": user, "token": token, "cached": true})
			return
		}
	}

	var user models.User
	if err := database.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Save to Redis cache
	userJson, _ := json.Marshal(user)
	RedisClient.Set(ctx, cacheKey, userJson, 10*time.Minute)

	token := generateToken(user)
	c.JSON(http.StatusOK, gin.H{"user": user, "token": token, "cached": false})
}

func generateToken(user models.User) string {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET_KEY")
	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString
}
