package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/swordofJR/copyright_golang/dto"
	"github.com/swordofJR/copyright_golang/model"
)

type CopyrightRepository struct {
	db *gorm.DB
}

func NewCopyrightRepository(db *gorm.DB) *CopyrightRepository {
	return &CopyrightRepository{
		db: db,
	}
}

func (r *CopyrightRepository) Create(createDTO dto.CreateCopyrightDTO) error {
	copyright := model.NewCopyright(createDTO)
	return r.db.Create(copyright).Error
}
