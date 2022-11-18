package models

import "time"

type Post struct {
	Id           int       `json:"id,omitempty"`
	Post_id      int       `json:"post_id,omitempty,string"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Author_id    int       `json:"author_id,string"`
	Community_id int       `json:"community_id,string"`
	Status       int       `json:"status"`
	Create_time  time.Time `json:"create_time,omitempty"`
	Update_time  time.Time `json:"update_time,omitempty"`
}

// 当点开一个帖子所需要的参数
type PostDetail struct {
	AuthorName string `json:"author_name"`
	*Post
	*Community `json:"community"`
}
