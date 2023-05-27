package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gustablo/secret-auth/config"
	"github.com/gustablo/secret-auth/internal/handlers"
	"github.com/gustablo/secret-auth/internal/middlewares"
)

func main() {
	config.OpenConn()
	config.ExecMigrations()

	router := gin.Default()

	authorized := router.Group("/")
	authorized.Use(middlewares.GetOwnerByPublicKey())
	{
		authorized.POST("/users", handlers.CreateUser)
		authorized.POST("/auth/login", handlers.Login)
	}

	router.Run(":8080")
}
