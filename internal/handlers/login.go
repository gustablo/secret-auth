package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/gustablo/secret-auth/internal/commons"
	"github.com/gustablo/secret-auth/internal/models"
)

func Login(c *gin.Context) {
	var user models.User
	var err error

	err = c.BindJSON(&user)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	foundUser, err := user.SelectByEmailOrUserName(c.GetString("ownerId"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Credentials not found"})
		return
	}

	err = commons.Compare(foundUser.Password, user.Password)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	token := models.Token{}

	accessToken, _ := token.Generate(jwt.MapClaims{
		"id":       foundUser.ID,
		"username": foundUser.Username,
		"email":    foundUser.Email,
	})

	refreshToken, _ := token.Generate(jwt.MapClaims{
		"id": foundUser.ID,
	})

	token.AccessToken = accessToken
	token.RefreshToken = refreshToken

	c.IndentedJSON(http.StatusOK, token)
}
