package logic

import (
	"fmt"
	"web-app/dao/mysql"
	"web-app/models"
	"web-app/pkg/jwt"
	"web-app/pkg/snowflake"
)

func Signup(p *models.ParamSignup) (err error) {
	//判断用户是否存在
	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		//这里是数据库查询错误
		return err
	}

	//生成uid
	Userid := snowflake.GenID()
	//在这里构造一个用户的实例
	user := &models.User{
		UserID:   Userid,
		Username: p.Username,
		Password: p.Password,
	}
	//保存到数据库
	//密码要加密
	err = mysql.InserUser(user)
	if err != nil {
		fmt.Println("mysql.InserUser faild", err)
		return
	}

	return
}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	//下面的传递的user是一个指针  所以就能拿到user的id了
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	return jwt.GenToken(user.UserID, user.Username)

}
