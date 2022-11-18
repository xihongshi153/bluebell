package mysql

import (
	"bluebell/dao/models"

	"go.uber.org/zap"
)

func InsertPost(id int64, post models.CreatePostParam) (bool, error) {
	sqlStr := "insert into post (post_id, title, content, author_id, community_id,status) values (?,?,?,?,?,?);"
	_, err := db.Exec(sqlStr, id, post.Title, post.Content, post.Author_id, post.Community_id, post.Status)
	if err != nil {
		return false, err
	}
	return true, nil
}
func SelectPostById(id int) (post models.Post, err error) {
	sqlStr := "select title, content, author_id, community_id, status, create_time, update_time from post where post_id = ?"
	row := db.QueryRow(sqlStr, id)
	err = row.Scan(&post.Title, &post.Content, &post.Author_id, &post.Community_id, &post.Status, &post.Create_time, &post.Update_time)
	post.Post_id = id
	if err != nil {
		return
	}
	return
}

func SelectPostPage(page, limit int) (data []models.Post, err error) {
	sqlStr := "select post_id,title,content,author_id,community_id,status from post order by create_time limit ?,?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		zap.S().Error("stmt, err := db.Prepare(sqlStr)", err.Error())
		return
	}
	rows, err := stmt.Query(page, limit)
	if err != nil {
		zap.S().Error("rows,err:stmt.Query(page, limit)", err.Error())
		return
	}
	for rows.Next() {
		item := models.Post{}
		err = rows.Scan(&item.Post_id, &item.Title, &item.Content, &item.Author_id, &item.Community_id, &item.Status)
		if err != nil {
			zap.S().Error("err = rows.Scan(&item.Post_id, &item.Title, &item.Content, &item.Author_id, &item.Community_id, &item.Status) ", err.Error())
			return
		}
		data = append(data, item)
	}
	return
}
