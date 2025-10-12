// internal/repository/user_repository.go
package repository

import (
	"github.com/saku-730/web-specimen/backend/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("mail_address = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
