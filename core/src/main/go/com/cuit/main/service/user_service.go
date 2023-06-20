package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/swordofJR/copyright_golang/dto"
	"github.com/swordofJR/copyright_golang/model"
	"github.com/swordofJR/copyright_golang/repository"
	"github.com/swordofJR/copyright_golang/utils"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) Register(registerDTO dto.RegisterDTO) error {
	err := s.repository.Create(registerDTO)
	if err != nil {
		return fmt.Errorf("failed to register user: %s", err)
	}
	return nil
}

func (s *UserService) Login(loginDTO dto.LoginDTO) (string, error) {
	user, err := s.repository.FindByUsername(loginDTO.Username)
	if err != nil {
		return "", fmt.Errorf("invalid username or password")
	}
	if user.Password != loginDTO.Password {
		return "", fmt.Errorf("invalid username or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      utils.TimeAfter(time.Hour * 24),
	})

	signedToken, err := token.SignedString([]byte(utils.GetEnv("