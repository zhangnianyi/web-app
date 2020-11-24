package controllers

type Rescode int64

const (
	CodeSuccess Rescode = 1000 + iota
	CodeInvalidparam
	CodeUserExist
	CodeUserNotExist
	CodeInvailPassword
	CodeServiceBusy
)

var codMsgmap = map[Rescode]string{
	CodeSuccess:        "success",
	CodeInvalidparam:   "请求参数错误",
	CodeUserExist:      "用户已存在",
	CodeUserNotExist:   "用户不存在",
	CodeInvailPassword: "密码错误",
	CodeServiceBusy:    "服务器内部忙",
}

func (c Rescode) Msg() string {
	msg, ok := codMsgmap[c]
	if !ok {
		msg = codMsgmap[CodeServiceBusy]
	}
	return msg
}
