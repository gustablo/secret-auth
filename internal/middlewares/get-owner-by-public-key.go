package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustablo/secret-auth/internal/models"
)

func GetOwnerByPublicKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		publick := c.Request.Header.Get("public-key")
		if publick == "" {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		id, err := models.Owner{Key: publick}.SelectByKey()

		if err != nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Set("ownerId", id)
		c.Next()
	}
}
