// internal/repository/user_defaults_repository.go
package repository

import (
	"github.com/saku-730/web-specimen/backend/internal/entity"
//	"github.com/saku-730/web-specimen/backend/internal/model"
	"gorm.io/gorm"
)

type UserDefaultsRepository interface {
	FindDefaultsByUserID(userID int) (*entity.UserDefault, error)
}

type userDefaultsRepository struct {
	db *gorm.DB
}

func NewUserDefaultsRepository(db *gorm.DB) UserDefaultsRepository {
	return &userDefaultsRepository{db: db}
}

func (r *userDefaultsRepository) FindDefaultsByUserID(userID int) (*entity.UserDefault, error) {
	var defaults entity.UserDefault
	if err := r.db.Where("user_id= ?",userID).First(&defaults).Error; err != nil {
		return nil, err
	}
	return &defaults, nil
}
