package models

import "time"

type User struct {
	Id          int       `json:"id"`
	User_id     int       `json:"user_id,string"`
	UserName    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	Gender      int       `json:"gender"`
	Update_time time.Time `json:"update_time"`
	Create_time time.Time `json:"create_time"`
}
