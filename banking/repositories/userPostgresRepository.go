package repositories

import (
	"fmt"

	"github.com/emaforlin/banking-api/banking/entities"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type userPostgresRepository struct {
	db *gorm.DB
}

func NewUserPostgresRepository(db *gorm.DB) UserRepository {
	return &userPostgresRepository{db: db}
}

func (r *userPostgresRepository) InsertUserData(in *entities.InsertUserDto) error {
	data := &entities.User{
		FullName: in.FullName,
		Username: in.Username,
		Password: in.Password,
	}
	out := &entities.User{}
	r.db.Where("username = ?", in.Username).First(&out)
	if out.Username == in.Username {
		log.Errorf("InsertUserData: username already exists")
		return fmt.Errorf("InsertUserData: username already exists")
	}

	result := r.db.Create(data)
	if result.Error != nil {
		log.Errorf("InsertUserData: %v", result.Error)
		return result.Error
	}
	log.Debugf("InsertUserData: %v", result.RowsAffected)
	return result.Error
}

func (r *userPostgresRepository) RetrieveUserData(name string, out *entities.User) (err error) {
	result := r.db.Where("full_name = ?", name).First(out)
	err = result.Error
	if err != nil {
		log.Errorf("RetrieveUserData: %v", err)
		return
	}
	log.Debugf("RetrieveUserData: %v", result.RowsAffected)
	return
}
