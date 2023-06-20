package utils

import (
	"github.com/swordofJR/copyright_golang/model"
)

type Resource string

const (
	ResourceUser     Resource = "user"
	ResourceArticle  Resource = "article"
	ResourceComment  Resource = "comment"
	ResourceCategory Resource = "category"
)

type Operation string

const (
	OperationCreate Operation = "create"
	OperationRead   Operation = "read"
	OperationUpdate Operation = "update"
	OperationDelete Operation = "delete"
)

func GetResourcesByAuthorities(authorities []string) ([]model.Resource, error) {
	authorityRepo := GetGlobalAuthorityRepository()
	resources, err := authorityRepo.FindResourcesByAuthorities(authorities)
	if err != nil {
		return nil, err
	}
	return resources, nil
}

func HasResource(resources []model.Resource, resource model.Resource) bool {
	for _, r := range resources {
		if r.Name == resource.Name && r.Type == resource.Type && r.Value == resource.Value {
			return true
		}
	}
	return false
}
