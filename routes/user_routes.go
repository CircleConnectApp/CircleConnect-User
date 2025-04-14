package routes

import (
    "github.com/gin-gonic/gin"
    "circleconnect-user/controllers"
    "circleconnect-user/middleware"
)

func SetupRoutes(r *gin.Engine) {
    user := r.Group("/users")
    user.Use(middleware.AuthMiddleware())
    {
        user.GET("/me", controllers.GetMyProfile)
        user.PUT("/me", controllers.UpdateMyProfile)
        user.GET("/me/communities", controllers.GetMyCommunities)
        user.GET("/:id", middleware.AdminMiddleware(), controllers.GetUserByID)
    }

    admin := r.Group("/admin")
    admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
    {
        admin.PUT("/users/:id/role", controllers.ChangeUserRole)
    }
}
