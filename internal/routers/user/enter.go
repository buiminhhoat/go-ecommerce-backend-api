package user

type UserRouterGroup struct {
	UserRouter
	ProductRouter
}

var RouterGroupApp = new(UserRouterGroup)
