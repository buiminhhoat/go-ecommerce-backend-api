package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping/", Pong)
	}

	return r
}

func Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "buiminhhoat")
	id := c.Query("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"name":    name,
		"id":      id,
		"user":    []string{"buiminhhoat", "hoatbm4"},
	})
}
