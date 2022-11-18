package logic

import (
	"bluebell/dao/models"
	"bluebell/dao/mysql"
	"bluebell/pkg/snowflake"
	"errors"
	"fmt"
)

func CreatePost(param models.CreatePostParam) (bool, error) {
	success, err := mysql.InsertPost(snowflake.GenID(), param)
	if err != nil {
		return false, err
	}
	if !success {
		return false, errors.New("服务器内部错误")
	}
	return true, err
}
func GetPostById(id int) (p models.Post, err error) {
	p, err = mysql.SelectPostById(id)
	return
}
func GetPostDetailById(id int) (pd models.PostDetail, err error) {
	// 1. 获取post
	p, err := mysql.SelectPostById(id)
	if err != nil {
		return pd, errors.New("postLogic.go GetPostDetailById mysql.SelectPostById(id) err: " + err.Error())
	}
	// 2. 获取 author_name
	fmt.Printf("%+v\n", p)
	author, err := mysql.SelectUserInfoById(p.Author_id)
	if err != nil {
		return pd, errors.New("postLogic.go GetPostDetailById mysql.SelectUserInfoById(p.Author_id) err: " + err.Error())
	}
	// 3. 获取 community
	com, err := mysql.SelectCommunityInfoById(p.Community_id)
	if err != nil {
		return pd, errors.New("postLogic.go GetPostDetailById SelectCommunityInfoById err: " + err.Error())
	}
	pd.AuthorName = author.UserName
	pd.Post = &p
	pd.Community = &com
	return
}
func GetPostPage(page, limit int) (data []models.Post, err error) {
	data, err = mysql.SelectPostPage(page, limit)
	return
}
