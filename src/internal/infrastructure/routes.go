package infrastructure

import (
    "github.com/gin-gonic/gin"
)

type MessageRoutes struct {
    GetMessagesController *MessageController
}

func NewMessageRoutes(getMessagesController *MessageController) *MessageRoutes {
    return &MessageRoutes{GetMessagesController: getMessagesController}
}

func (r *MessageRoutes) AttachRoutes(router *gin.Engine) {
    messagesGroup := router.Group("/messages")
    {
        messagesGroup.GET("/", func(c *gin.Context) {
            r.GetMessagesController.GetMessages(c.Writer, c.Request)
        })
    }
}