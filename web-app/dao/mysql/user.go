package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"web-app/models"
)

const secret = "liwenzhou.com"

func CheckUserExist(username string) (err error) {
	sqlstr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlstr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已经存在")
	}
	return

}

func InserUser(user *models.User) (err error) {
	//对密码进行加密
	user.Password = encryPassword(user.Password)
	sqlstr := `insert into  user(user_id,username,password) values(?,?,?)`
	_, err = db.Exec(sqlstr, user.UserID, user.Username, user.Password)
	return
}
func encryPassword(opassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(opassword)))

}
func Login(user *models.User) (err error) {
	opassword := user.Password
	sqlstr := `select user_id,username,password from user where username=?`
	err = db.Get(user, sqlstr, user.Username)
	if err == sql.ErrNoRows {
		fmt.Println("你要登陆的用户不存在 请先去注册")
		return
	}
	if err != nil {
		fmt.Println("查询数据库失败")
		return
	}
	//判断密码
	password := encryPassword(opassword)
	if password != user.Password {
		return errors.New("用户登录的密码错误")
	}
	return
}
