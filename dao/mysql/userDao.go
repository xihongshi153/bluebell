package mysql

import (
	"bluebell/pkg/snowflake"
	"fmt"

	"go.uber.org/zap"
)

func InsertUser(userid int64, username, password, email string) (err error) {
	sqlStr := "insert into user(user_id,username,password,email) values (?,?,?,?)"
	_, err = db.Exec(sqlStr, snowflake.GenID(), username, password, email)
	fmt.Println(err)
	if err != nil {
		zap.L().Debug(fmt.Sprint("dao user.go Register insert fail err: ", err.Error()))
		return
	}
	return
}
func FindUserWithUserNamePassword(username, password string) (cnt int, err error) {
	sqlStr := "select count(*) from user where username = ? and password = ?;"
	row := db.QueryRow(sqlStr, username, password)
	err = row.Scan(&cnt)
	if err != nil {
		zap.L().Info("FindUserWithUserNamePassword fail err: ", zap.Error(err))
	}
	return
}
