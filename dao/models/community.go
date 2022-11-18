package models

import "time"

type Community struct {
	Id             int       `json:"id,string"`
	Community_id   int       `json:"community_id,string"`
	Community_name string    `json:"community_name"`
	Introduction   string    `json:"introduction"`
	Create_time    time.Time `json:"create_time"`
	Update_time    time.Time `json:"update_time"`
}
