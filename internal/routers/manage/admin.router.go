package manage

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (ur *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/register")
		adminRouterPublic.POST("/otp")
	}

	adminRouterPrivate := Router.Group("/admin/user")
	// adminRouterPrivate.Use(limiter())
	// adminRouterPrivate.Use(Authen())
	// adminRouterPrivate.Use(Permission())
	{
		adminRouterPrivate.GET("/active_user")
	}
}
