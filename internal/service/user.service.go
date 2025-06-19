package service

import (
	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/repositories"
	"github.com/buiminhhoat/go-ecommerce-backend-api/pkg/response"
)

// type UserService struct {
// 	userRepository *repositories.UserRepository
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepository: repositories.NewUserRepository(),
// 	}
// }

// func (us *UserService) GetInfoUser() string {
// 	return us.userRepository.GetInfoUser()
// }

// INTERFACE_VERSION

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepository repositories.IUserRepository
}

func NewUserService(
	userRepository repositories.IUserRepository,
) IUserService {
	return &userService{
		userRepository: userRepository,
	}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	if us.userRepository.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}
	return response.ErrCodeSuccess
}
