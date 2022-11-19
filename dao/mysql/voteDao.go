package mysql

import (
	"database/sql"
	"fmt"
)

func InsertVotePostIdUserId(user_id, post_id int, tx *sql.Tx) (err error) {
	// redis_userVoteInfo 添加一条信息
	sqlStr := "insert into redis_userVoteInfo (post_id_user_id) values (?)"
	_, err = tx.Exec(sqlStr, fmt.Sprintf("%d_%d", post_id, user_id))
	return
}
func IncreasePostVoteNum(post_id int, tx *sql.Tx) (err error) {
	// redis_postVoteInfo 对应 +1
	sqlStr := "update redis_postVoteInfo set vote_num = vote_num + 1 where post_id=?"
	_, err = tx.Exec(sqlStr, post_id)
	return
}

func DeleteVotePostIdUserId(user_id, post_id int, tx *sql.Tx) (err error) {
	// redis_userVoteInfo 删除一条信息
	sqlStr := "delete from redis_userVoteInfo where post_id_user_id = ? ;"
	_, err = tx.Exec(sqlStr, fmt.Sprintf("%d_%d", post_id, user_id))
	return
}
func DecreasePostVoteNum(post_id int, tx *sql.Tx) (err error) {
	// redis_postVoteInfo 对应 -1
	sqlStr := "update redis_postVoteInfo set vote_num = vote_num - 1 where post_id=?;"
	_, err = tx.Exec(sqlStr, post_id)
	return
}
