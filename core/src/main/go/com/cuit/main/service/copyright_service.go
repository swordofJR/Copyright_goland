package service

import (
	"fmt"

	"github.com/swordofJR/copyright_golang/dto"
	"github.com/swordofJR/copyright_golang/repository"
)

type CopyrightService struct {
	repository repository.CopyrightRepository
}

func NewCopyrightService(repository repository.CopyrightRepository) *CopyrightService {
	return &CopyrightService{
		repository: repository,
	}
}

func (s *CopyrightService) Create(createDTO dto.CreateCopyrightDTO) error {
	err := s.repository.Create(createDTO)
	if err != nil {
		return fmt.Errorf("failed to create copyright: %s", err)
	}
	return nil
}
