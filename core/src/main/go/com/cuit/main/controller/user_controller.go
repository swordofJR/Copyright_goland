package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swordofJR/copyright_golang/dto"
	"github.com/swordofJR/copyright_golang/service"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	var registerDTO dto.RegisterDTO
	err := c.ShouldBindJSON(&registerDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uc.service.Register(registerDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "created"})
}

func (uc *UserController) Login(c *gin.Context) {
	var loginDTO dto.LoginDTO
	err := c.ShouldBindJSON(&loginDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := uc.service.Login(loginDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
