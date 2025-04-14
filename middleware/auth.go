package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"strings"
	"time"
)

var jwtSecret []byte

func init() {  
	secret := os.Getenv("JWT_SECRET_KEY") 
	jwtSecret = []byte(secret)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		// fmt.Println("Authorization Header:", authHeader)

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		fmt.Println("Token String:", tokenString)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				fmt.Printf("Unexpected signing method: %v\n", token.Header["alg"])
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil {
			fmt.Println("Token parse error:", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			// fmt.Println("Invalid token claims")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		} 
		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				// fmt.Println("Token expired at", exp)
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
				return
			}
		} else {
			fmt.Println("Missing exp field in token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			return
		}



		fmt.Println("Decoded JWT claims:", claims)

		if userID, ok := claims["user_id"]; ok {
			c.Set("user_id", userID)
		}
		if role, ok := claims["role"]; ok {
			c.Set("role", role)
		}

		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")
		// fmt.Println("User role in AdminMiddleware:", role) 
		if role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Admin only"})
			return
		}
		c.Next()
	}
}
