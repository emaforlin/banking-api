package usecases

import (
	"github.com/emaforlin/banking-api/banking/entities"
	"github.com/emaforlin/banking-api/banking/models"
	"github.com/emaforlin/banking-api/banking/repositories"
)

type UserUsecase interface {
	RegisterUser(in *models.AddUserData) error
	// GetUserByName(name string) (out *models.ShowUserData)
}

type userUsecaseImpl struct {
	userRepository repositories.UserRepository
}

func NewUserUsecaseImpl(userRepo repositories.UserRepository) UserUsecase {
	return &userUsecaseImpl{
		userRepository: userRepo,
	}
}

func (u *userUsecaseImpl) RegisterUser(in *models.AddUserData) error {
	insertUserData := &entities.InsertUserDto{
		FullName: in.FullName,
	}
	if err := u.userRepository.InsertUserData(insertUserData); err != nil {
		return err
	}
	return nil
}
