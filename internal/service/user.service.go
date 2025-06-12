package service

import "github.com/buiminhhoat/go-ecommerce-backend-api/internal/repositories"

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: repositories.NewUserRepository(),
	}
}

func (us *UserService) GetInfoUser() string {
	return us.userRepository.GetInfoUser()
}
