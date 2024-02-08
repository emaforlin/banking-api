package usecases

import (
	"github.com/emaforlin/banking-api/banking/entities"
	"github.com/emaforlin/banking-api/banking/models"
	"github.com/emaforlin/banking-api/banking/respositories"
)

type UserUsecase interface {
	UserDataProcessing(in *models.AddUserData) error
}

type userUsecaseImpl struct {
	userRepository respositories.UserRepository
}

func NewUserUsecaseImpl(userRepo respositories.UserRepository) UserUsecase {
	return &userUsecaseImpl{
		userRepository: userRepo,
	}
}

func (u *userUsecaseImpl) UserDataProcessing(in *models.AddUserData) error {
	insertUserData := &entities.InsertUserDto{
		FullName: in.FullName,
	}
	if err := u.userRepository.InsertUserData(insertUserData); err != nil {
		return err
	}
	return nil
}
