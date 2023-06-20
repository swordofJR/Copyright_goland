package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/swordofJR/copyright_golang/dto"
	"github.com/swordofJR/copyright_golang/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(registerDTO dto.RegisterDTO) error {
	user := model.NewUser(registerDTO)
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
