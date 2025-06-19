package controller

import (
	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/service"
	"github.com/buiminhhoat/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(
	userService service.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("", "")
	response.SuccessResponse(c, result, nil)
}

// func (uc *UserController) GetUserById(c *gin.Context) {
// 	fmt.Println("---> My Handler")
// 	response.SuccessResponse(c, 20001, []string{"buiminhhoat", "Bùi Minh Hoạt"})
// }
