package repositories

import (
	"github.com/buiminhhoat/go-ecommerce-backend-api/global"
	"github.com/buiminhhoat/go-ecommerce-backend-api/internal/database"
)

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
	sqlc *database.Queries
}

// GetUserByEmail implements IUserRepository.
func (ur *userRepository) GetUserByEmail(email string) bool {
	// row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&model.GoCrmUser{}).RowsAffected

	user, err := ur.sqlc.GetUserByEmailSQLC(ctx, email)

	if err != nil {
		return false
	}
	return user.UsrID != NumberNull
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		sqlc: database.New(global.Mdbc),
	}
}
