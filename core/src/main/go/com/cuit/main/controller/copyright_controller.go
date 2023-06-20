package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/swordofJR/copyright_golang/dto"
	"github.com/swordofJR/copyright_golang/model"
	"github.com/swordofJR/copyright_golang/service"
	"github.com/swordofJR/copyright_golang/utils"
	"net/http"
)

type CopyrightController struct {
	service service.CopyrightService
}

func NewCopyrightController(service service.CopyrightService) *CopyrightController {
	return &CopyrightController{
		service: service,
	}
}

func (cc *CopyrightController) Create(c *gin.Context) {
	var createDTO dto.CreateCopyrightDTO
	err := c.ShouldBindJSON(&createDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := utils.GetUserIdFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	copyright := model.NewCopyright(createDTO)
	copyright.OwnerId = userId

	err = cc.service.Create(copyright)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "created"})
}
