package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"web-app/Logger"
	"web-app/controllers"
	"web-app/middlewares"
)

func SetupRoute() *gin.Engine {

	r := gin.New()
	//注册业务路由
	//
	r.POST("/signup", controllers.Signuphandler)
	r.POST("/login", controllers.Loginhandler)
	r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		//如果是登录的用户 就继续执行 就是判断请求头中是否有jwt  是否有有效的token
		//ISlogin := true
		//c.Request.Header.Get("Authorization")
		//if ISlogin {
		c.String(http.StatusOK, "PONG")
		//} else {
		//	c.String(http.StatusOK, "请登录后在进行操作")
		//}

	})
	r.Use(Logger.GinLogger(), Logger.GinRecovery(true))
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, viper.GetString("app.mode"))

	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})

	})

	return r
}

// JWTAuthMiddleware 基于JWT的认证中间件
