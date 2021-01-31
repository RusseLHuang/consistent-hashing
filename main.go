package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/RusseLHuang/consistent-hashing/consistent"
	"github.com/RusseLHuang/consistent-hashing/database"
	pc "github.com/RusseLHuang/consistent-hashing/payment/controller"
	"github.com/RusseLHuang/consistent-hashing/user/controller"
	"github.com/RusseLHuang/consistent-hashing/user/repository"
)

func main() {
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	dbKeyList := database.FetchDBKeyList()
	dbConnectionMap := database.InitDB()
	consistentHandler := consistent.NewConsistent(dbKeyList)

	userRepository := repository.UserRepository{}.CreateRepository(
		dbConnectionMap,
		consistentHandler,
	)

	userController := controller.NewUserController(userRepository)

	r := gin.Default()
	r.GET("/balance", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/deposit", pc.PaymentController{}.Deposit())
	r.POST("/transfer", pc.PaymentController{}.Transfer())
	r.POST("/user", userController.CreateUser())
	r.Run()
}
