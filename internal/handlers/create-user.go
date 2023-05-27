package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustablo/secret-auth/internal/models"
)

func CreateUser(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		return
	}

	err = user.Insert(c.GetString("ownerId"))

	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Error while processing user"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}
