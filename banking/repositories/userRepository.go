package repositories

import "github.com/emaforlin/banking-api/banking/entities"

type UserRepository interface {
	InsertUserData(in *entities.InsertUserDto) error
	RetrieveUserData(name string, out *entities.User) error
}
