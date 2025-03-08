package cmd

import (
    "log"
    "pedidos/src/core"
    "pedidos/src/server"
    messageDeps "pedidos/src/internal/infrastructure"
)

func Init() {
    db, err := core.ConnectDB()
    if err != nil {
        log.Fatal("Error al conectar a la base de datos:", err)
    }

    rabbitMQ, err := core.NewRabbitMQConnection()
    if err != nil {
        log.Fatal("Error al conectar a RabbitMQ:", err)
    }
    defer rabbitMQ.Close()


    messageDependencies := messageDeps.NewDependencies(db, rabbitMQ)
    messageConsumer := messageDependencies.GetConsumer()
    messageRoutes := messageDependencies.

    go messageConsumer.Start()

    server.Run(messageRoutes)
}