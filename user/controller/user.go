package controller

import (
	"time"

	"github.com/RusseLHuang/consistent-hashing/user/dto"
	"github.com/RusseLHuang/consistent-hashing/user/entity"
	"github.com/RusseLHuang/consistent-hashing/user/repository"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository *repository.UserRepository
}

func NewUserController(
	userRepository *repository.UserRepository,
) UserController {
	return UserController{
		userRepository: userRepository,
	}
}

func (uc UserController) CreateUser() func(*gin.Context) {
	return func(c *gin.Context) {
		var requestBody dto.CreateUserRequestBody
		c.BindJSON(&requestBody)

		userEntity := entity.User{
			Name:       requestBody.Name,
			NationalID: requestBody.NationalID,
			CreatedAt:  time.Now(),
			DeletedAt:  time.Now(),
		}

		uc.userRepository.Save(&userEntity)

		c.JSON(200, gin.H{
			"message": "success",
		})
	}
}
