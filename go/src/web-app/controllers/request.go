package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const ContextUseridKey = "userid"

var ErrorUserNotLogin = errors.New("用户未登录")

//获取当前登录的用户id
func GetCurrentUser(c *gin.Context) (userid int64, err error) {
	uid, ok := c.Get(ContextUseridKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userid, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
