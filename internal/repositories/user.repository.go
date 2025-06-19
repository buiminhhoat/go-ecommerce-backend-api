package repositories

// type UserRepository struct{}

// func NewUserRepository() *UserRepository {
// 	return &UserRepository{}
// }

// func (ur *UserRepository) GetInfoUser() string {
// 	return "buiminhhoat"
// }

// INTERFACE_VERSION

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
}

// GetUserByEmail implements IUserRepository.
func (u *userRepository) GetUserByEmail(email string) bool {
	return true
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
