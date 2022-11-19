package logic

import (
	"bluebell/dao/mysql"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
	"errors"
	"fmt"
)

func Register(username string, password string, email string) (err error) {
	// 各种复杂的逻辑 如 是否重复 是否要在redis中操作
	err = mysql.InsertUser(snowflake.GenID(), username, password, email)
	return
}
func Login(username, password string) (bool, string, error) {
	cnt, err := mysql.FindUserWithUserNamePassword(username, password)
	if err != nil {
		return false, "", err
	}
	if cnt == 0 {
		return false, "", errors.New("can not find user")
	}
	if cnt == 1 {
		// 签发 token
		userId, err := mysql.SelectUserIdByUserName(username)
		if err != nil {
			return false, "", errors.New("can not find user_id by this username:" + username)
		}
		token, _ := jwt.GenerateJwt(username, fmt.Sprintf("%d", userId))
		return true, token, nil
	}
	return false, "", errors.New(" user repeat, two or more same username and same password")
}
