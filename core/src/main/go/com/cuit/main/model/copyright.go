package model

import (
	"time"

	"github.com/swordofJR/copyright_golang/dto"
)

type Copyright struct {
	ID            int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	CopyrightName string
	OwnerId       int64
	Label         string
	DigitalMark   string
	Hash          string
	Deleted       int `gorm:"column:deleted"`
	CreateTime    time.Time
	UpdateTime    time.Time
}

func NewCopyright(createDTO dto.CreateCopyrightDTO) *Copyright {
	return &Copyright{
		CopyrightName: createDTO.CopyrightName,
		Label:         createDTO.Label,
		DigitalMark:   createDTO.DigitalMark,
		Hash:          createDTO.Hash,
		Deleted:       0,
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	}
}

func (c *Copyright) TableName() string {
	return "copyright"
}
