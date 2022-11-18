package mysql

import (
	"bluebell/dao/models"
	"bluebell/pkg/snowflake"

	"go.uber.org/zap"
)

func InsertUser(userid int64, username, password, email string) (err error) {
	sqlStr := "insert into user(user_id,username,password,email) values (?,?,?,?)"
	_, err = db.Exec(sqlStr, snowflake.GenID(), username, password, email)
	if err != nil {
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
func SelectUserInfoById(id int) (u models.User, err error) {
	sqlStr := "select username,email,gender from user where user_id = ?"
	row := db.QueryRow(sqlStr, id)
	err = row.Scan(&u.UserName, &u.Email, &u.Gender)
	if err != nil {
		zap.L().Info("userDao.go SelectUserInfoById fail err: ", zap.Error(err))
	}
	return
}
