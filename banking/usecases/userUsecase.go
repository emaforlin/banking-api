package usecases

import (
	"html"
	"strings"

	"github.com/emaforlin/banking-api/banking/entities"
	"github.com/emaforlin/banking-api/banking/models"
	"github.com/emaforlin/banking-api/banking/repositories"
	"golang.org/x/crypto/bcrypt"
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
	hashPasswd, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	insertUserData := &entities.InsertUserDto{
		FullName: in.FullName,
		Username: html.EscapeString(strings.TrimSpace(in.Username)),
		Password: string(hashPasswd),
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
