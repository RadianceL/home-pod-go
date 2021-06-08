package controller

import (
	"github.com/gin-gonic/gin"
	models "go-gin-template/src/pkg"
	"go-gin-template/src/web/entity"
	"go-gin-template/src/web/entity/response"
)

type Result struct{}

func (Result) List(c *gin.Context) {
	email := entity.Email{ID: 10, UserID: 1111, Email: "931305033@qq.com", Subscribed: true}
	_ = models.Create(email)
	response.ResSuccessMsg(c)
}
