package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseDate struct {
	Code Rescode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponceError(c *gin.Context, code Rescode) {
	rd := &ResponseDate{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}
func ResponceErrorWithMsg(c *gin.Context, code Rescode, msg interface{}) {
	rd := &ResponseDate{
		Code: CodeSuccess,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}

func ResponceSuccess(c *gin.Context, data interface{}) {
	rd := &ResponseDate{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, rd)
}
