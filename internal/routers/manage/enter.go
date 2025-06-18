package manage

type ManageRouterGroup struct {
	UserRouter
	AdminRouter
}

var RouterGroupApp = new(ManageRouterGroup)
