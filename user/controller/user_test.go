package controller

import (
	"testing"

	"github.com/gin-gonic/gin"
)

type GinContextMock struct {
	*gin.Context
}

func (g *GinContextMock) BindJSON(obj interface{}) error {
	var jsonObj map[string]string
	jsonObj["Name"] = "Russel"
	obj = jsonObj

	return nil
}

func TestUserController(t *testing.T) {
	uc := UserController{}

	uc.CreateUser()
}
