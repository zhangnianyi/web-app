package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"web-app/logic"
	"web-app/models"
)

func Signuphandler(c *gin.Context) {
	var p models.ParamSignup
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误
		zap.L().Error("Should BindJSON error", zap.Error(err))
		//判断error 是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"mes": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	fmt.Println(p)
	//获取参数
	//业务处理
	//将p传递到创建数据库的阶段
	err := logic.Signup(&p)
	if err != nil {
		fmt.Println("logic.Signup faild ", err)
		c.JSON(http.StatusOK, gin.H{
			"message": "注册失败",
		})
		return
	}
	//返回响应
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}

func Loginhandler(c *gin.Context) {
	var p models.ParamLogin
	//获取请求参数及参数校验
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Should BindJSON  for login error ", zap.Error(err))
		//判断error 是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"mes": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	//业务逻辑处理
	if err := logic.Login(&p); err != nil {
		zap.L().Error("logic.Login(&p)", zap.String("username", p.Username), zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"mess": "user login faild",
		})
		return
	}
	//返回相应
	c.JSON(http.StatusOK, gin.H{
		"mes": "登陆成功",
	})
}
