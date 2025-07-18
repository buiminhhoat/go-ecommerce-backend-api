package user

import "github.com/gin-gonic/gin"

type TicketRouter struct{}

func (pr *TicketRouter) InitTicketRouter(Router *gin.RouterGroup) {
	// public router
	ticketRouterPublic := Router.Group("/ticket")
	{
		// ticketRouterPublic.GET("/search")
		ticketRouterPublic.GET("/item/:id")
	}
	// private router
}
