package user

import (
	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/controller/account"
	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userController, _ := wire.InitUserRouterHandler()
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/login", account.Login.Login)
		userRouterPublic.POST("/otp")
	}

	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/get_info")
	}
}
