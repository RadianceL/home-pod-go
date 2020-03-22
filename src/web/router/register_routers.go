package routers

import (
	"app-test/src/web/controller"
	"github.com/gin-gonic/gin"

)

func RegisterRouterSys(app *gin.RouterGroup) {
	menu := controller.Result{}
	app.GET("/menu/list", menu.List)
}