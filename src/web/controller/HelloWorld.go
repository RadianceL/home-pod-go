package controller

import (
	models "app-test/src/pkg"
	"app-test/src/web/entity"
	"app-test/src/web/entity/response"
	"github.com/gin-gonic/gin"
)

type Result struct{}

func (Result) List(c *gin.Context){
	email := entity.Email{ID: 10, UserID: 1111, Email: "931305033@qq.com", Subscribed: true}
	_ = models.Create(email)
	response.ResSuccessMsg(c)
}
