package respositories

import "github.com/emaforlin/banking-api/banking/entities"

type UserRepository interface {
	InsertUserData(in *entities.InsertUserDto) error
}
