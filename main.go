package main

import (
    "github.com/gin-gonic/gin"
    "circleconnect-user/database"
    "circleconnect-user/routes"
)

func main() {
    database.InitDB()
    r := gin.Default()
    routes.SetupRoutes(r)
    r.Run(":8081")
}
 