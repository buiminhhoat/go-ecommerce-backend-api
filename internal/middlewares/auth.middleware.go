package middlewares

import (
	"context"
	"fmt"
	"log"

	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/utils/auth"
	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the request url path
		uri := c.Request.URL.Path
		log.Println("URI request: ", uri)
		jwtToken, ok := auth.ExtractBearerToken(c)
		if !ok {
			fmt.Println("jwtToken: ", jwtToken)
			fmt.Println("ok: ", ok)
			c.AbortWithStatusJSON(401, gin.H{"code": 40001, "err": "Unauthorized", "description": ""})
			return
		}

		// Validate jwt token by subject
		claims, err := auth.VerifyTokenSubject(jwtToken)

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"code": 40001, "err": "Invalid token", "description": ""})
			return
		}
		// Update claims to context
		log.Println("claims::: UUID::", claims.Subject)
		ctx := context.WithValue(c.Request.Context(), "subjectUUID", claims.Subject)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
