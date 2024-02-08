package usecases

import (
	"github.com/emaforlin/banking-api/banking/entities"
	"github.com/emaforlin/banking-api/banking/models"
	"github.com/emaforlin/banking-api/banking/repositories"
)

type UserUsecase interface {
	RegisterUser(in *models.AddUserData) error
	GetUserByName(name string) (*models.ShowUserData, error)
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
		Username: in.Username,
		Password: in.Password,
	}

	if err := u.userRepository.InsertUserData(insertUserData); err != nil {
		return err
	}
	return nil
}

func (u *userUsecaseImpl) GetUserByName(name string) (*models.ShowUserData, error) {
	obtainedUserData := &entities.User{}
	if err := u.userRepository.RetrieveUserData(name, obtainedUserData); err != nil {
		return nil, err
	}
	return &models.ShowUserData{
		FullName: obtainedUserData.FullName,
		Funds:    obtainedUserData.Funds,
	}, nil

}
