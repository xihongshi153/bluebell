package mysql

import (
	"bluebell/setting"
	"database/sql"
	"fmt"
)

var db *sql.DB

func Init(cfg *setting.MySQLConfig) (err error) {
	//dsn := "root:root@tcp(127.0.0.1:3306)/blubell"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName,
	)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	// 尝试与数据连接
	err = db.Ping()
	if err != nil {
		return
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns) // 最大连接数
	db.SetMaxIdleConns(cfg.MaxIdleConns) // 最多休眠连接
	return
}

func Close() {
	_ = db.Close()
}
