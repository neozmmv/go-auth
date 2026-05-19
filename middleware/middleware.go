package middleware

import (
	"fmt"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/neozmmv/go-auth/utils"
)

func Auth(c *gin.Context) {
	fmt.Println("Middleware executed")
	token := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
	_, err := utils.ValidateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
		return
	}
	c.Set("token", token)
	c.Next()
}
