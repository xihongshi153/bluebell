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

type CommunityInfoReply struct {
	CId           string `form:"communityId"`
	CName         string `form:"communityName"`
	CIntroduction string `form:"communityIntroduction"`
}
type CreatePostParam struct {
	Title        string `json:"title" binding:"required"`
	Content      string `json:"content" binding:"required"`
	Author_id    int    `json:"author_id,string" binding:"required"`
	Community_id int    `json:"community_id,string" binding:"required"`
	Status       int    `json:"status" binding:"required,gte=1,lte=4"`
}
