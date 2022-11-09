package models

// 定义一些 param model 方便调用
type RegisterParam struct {
	UserName   string `form:"username" binding:"required"`
	Passwrod   string `form:"password"  binding:"required"`
	Email      string `form:"email" binding:"required"`
	Repassword string `form:"repassword"  binding:"required,eqfield=Password"`
}

type LoginParam struct {
	UserName string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
