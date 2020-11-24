package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"web-app/Logger"
	"web-app/controllers"
)

func SetupRoute() *gin.Engine {

	r := gin.New()
	//注册业务路由
	//
	r.POST("/signup", controllers.Signuphandler)
	r.POST("/login", controllers.Loginhandler)
	r.Use(Logger.GinLogger(), Logger.GinRecovery(true))
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, viper.GetString("app.mode"))

	})

	return r

}
