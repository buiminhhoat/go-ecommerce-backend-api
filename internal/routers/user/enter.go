package user

type UserRouterGroup struct {
	UserRouter
	ProductRouter
	TicketRouter
}

var RouterGroupApp = new(UserRouterGroup)
