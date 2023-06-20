package model

import (
	"time"
)

type Resource struct {
	ID         int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name       string
	Type       string
	Value      string
	Deleted    int `gorm:"column:deleted"`
	CreateTime time.Time
	UpdateTime time.Time
}
