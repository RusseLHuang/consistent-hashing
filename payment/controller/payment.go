package controller

import (
	"github.com/RusseLHuang/consistent-hashing/payment/dto"
	"github.com/RusseLHuang/consistent-hashing/user/entity"
	"github.com/RusseLHuang/consistent-hashing/user/repository"
	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	userRepository repository.UserRepository
}

func (pc PaymentController) Deposit() func(*gin.Context) {
	return func(c *gin.Context) {
		var requestBody dto.DepositRequestBody
		c.BindJSON(&requestBody)

		userEntity := entity.User{
			ID: requestBody.UserID,
		}

		pc.userRepository.AddBalance(&userEntity, requestBody.Amount)

		c.JSON(200, gin.H{
			"message": "success",
		})
	}
}

func (pc PaymentController) Transfer() func(*gin.Context) {
	return func(c *gin.Context) {
		var requestBody dto.TransferRequestBody
		c.BindJSON(&requestBody)

		fromUserEntity := entity.User{
			ID: requestBody.FromID,
		}

		toUserEntity := entity.User{
			ID: requestBody.ToID,
		}

		err := pc.userRepository.Transfer(
			&fromUserEntity,
			&toUserEntity,
			requestBody.Amount,
		)

		if err != nil {
			c.JSON(500, gin.H{
				"message": "failed",
				"error":   err,
			})

			return
		}

		c.JSON(200, gin.H{
			"message": "success",
		})
	}
}
