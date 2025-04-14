package controllers

import (
    "circleconnect-user/database"
    "circleconnect-user/models"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

func GetMyProfile(c *gin.Context) {
    id := c.GetFloat64("user_id")
    var user models.User
    if err := database.DB.First(&user, uint(id)).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

func UpdateMyProfile(c *gin.Context) {
    id := c.GetFloat64("user_id")
    var user models.User
    if err := database.DB.First(&user, uint(id)).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.BindJSON(&user)
    database.DB.Save(&user)
    c.JSON(http.StatusOK, user)
}

func GetUserByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var user models.User
    if err := database.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

func GetMyCommunities(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Will implement once community membership is available."})
}

func ChangeUserRole(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var user models.User
    if err := database.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    var input struct {
        Role string `json:"role"`
    }
    c.BindJSON(&input)
    user.Role = input.Role
    database.DB.Save(&user)
    c.JSON(http.StatusOK, user)
}
