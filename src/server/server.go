package server

import (
    "pedidos/src/internal/infrastructure"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func Run(messageRoutes *infrastructure.MessageRoutes) {
    router := gin.Default()

    messageRoutes.AttachRoutes(router)

    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"}, 
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Content-Type", "Authorization"},
        AllowCredentials: true,
    }))

    router.Run(":8081")
}