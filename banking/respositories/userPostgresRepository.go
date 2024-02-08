package respositories

import (
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
	}
	result := r.db.Create(data)

	if result.Error != nil {
		log.Errorf("InsertUserData: %v", result.Error)
		return result.Error
	}
	log.Debugf("InsertUserData: %v", result.RowsAffected)
	return nil
}
