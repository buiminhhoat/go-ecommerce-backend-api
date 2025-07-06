package auth

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func ExtractBearerToken(c *gin.Context) (string, bool) {
	authHeader := c.GetHeader("Authorization")
	fmt.Println(authHeader)
	if strings.HasPrefix(authHeader, "Bearer") {
		return strings.TrimPrefix(authHeader, "Bearer "), true
	}
	return "", false
}
