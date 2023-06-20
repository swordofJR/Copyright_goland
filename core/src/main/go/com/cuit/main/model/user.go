package model

import (
	"time"

	"github.com/swordofJR/copyright_golang/dto"
)

type Role struct {
	Name        string
	Permissions []Permission
}
type Permission struct {
	Resource  Resource
	Operation Operation
}

type ResourceConst string

const (
	ResourceUser     ResourceConst = "user"
	ResourceArticle  ResourceConst = "article"
	ResourceComment  ResourceConst = "comment"
	ResourceCategory ResourceConst = "category"
)

type Operation string

const (
	OperationCreate Operation = "create"
	OperationRead   Operation = "read"
	OperationUpdate Operation = "update"
	OperationDelete Operation = "delete"
)

type User struct {
	ID         int64 `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Username   string
	Password   string
	Email      string
	Phone      string
	Sex        int
	Deleted    int `gorm:"column:deleted"`
	CreateTime time.Time
	UpdateTime time.Time
	Roles      []Role
}

func NewUser(registerDTO dto.RegisterDTO) *User {
	return &User{
		Username:   registerDTO.Username,
		Password:   registerDTO.Password,
		Email:      registerDTO.Email,
		Phone:      registerDTO.Phone,
		Sex:        registerDTO.Sex,
		Deleted:    0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
}

func (u *User) TableName() string {
	return "user"
}
