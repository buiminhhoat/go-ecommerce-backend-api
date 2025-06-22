package controller

import (
	"fmt"

	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/service"
	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/vo"
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
	var params vo.UserRegistratorRequest
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.ErrorResponse(c, response.ErrCodeParamInvalid, err.Error())
	}
	fmt.Printf("Email params: %v", params.Email)
	result := uc.userService.Register(params.Email, params.Purpose)
	response.SuccessResponse(c, result, nil)
}

// func (uc *UserController) GetUserById(c *gin.Context) {
// 	fmt.Println("---> My Handler")
// 	response.SuccessResponse(c, 20001, []string{"buiminhhoat", "Bùi Minh Hoạt"})
// }
