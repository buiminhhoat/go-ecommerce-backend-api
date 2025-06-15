package initialize

import (
	"fmt"

	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/controller"
	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Befor --> AA")
		c.Next()
		fmt.Println("Alter --> AA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Befor --> BB")
		c.Next()
		fmt.Println("Alter --> BB")
	}
}

func CC(c *gin.Context) {
	fmt.Println("Before --> CC")
	c.Next()
	fmt.Println("Alter --> CC")
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthenMiddleware(), BB(), CC)

	v1 := r.Group("/v1")
	{
		v1.GET("/ping/", controller.NewPongController().Pong)
		v1.GET("/user/1", controller.NewUserController().GetUserById)
	}

	return r
}
