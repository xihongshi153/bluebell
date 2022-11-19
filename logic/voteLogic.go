package logic

import (
	"bluebell/dao/mysql"

	"go.uber.org/zap"
)

func MakeVoteLogic(post_id, user_id int) (err error) {
	tx, err := mysql.GetTx()
	if err != nil {
		zap.S().Error("mysql.GetTx() fail", " err: ", err.Error())
		return
	}

	err = mysql.InsertVotePostIdUserId(user_id, post_id, tx)
	if err != nil {
		tx.Rollback()
		zap.S().Error("mysql.InsertVotePostIdUserId(user_id, post_id, tx)fail", " err: ", err.Error())
		return
	}
	err = mysql.IncreasePostVoteNum(post_id, tx)
	if err != nil {
		tx.Rollback()
		zap.S().Error("mysql.IncreasePostVoteNum(post_id, tx) fail", " err: ", err.Error())
		return
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		zap.S().Error("tx.Commit() fail", " err: ", err.Error())
	}
	return
}
func DeleteVoteLogic(post_id, user_id int) (err error) {
	tx, err := mysql.GetTx()
	if err != nil {
		zap.S().Error("mysql.GetTx() fail", " #err: ", err.Error())
		return
	}
	err = mysql.DeleteVotePostIdUserId(user_id, post_id, tx)
	if err != nil {
		tx.Rollback()
		zap.S().Error("mysql.DeleteVotePostIdUserId(user_id, post_id, tx) fail", " #err: ", err.Error())
		return
	}
	err = mysql.DecreasePostVoteNum(post_id, tx)
	if err != nil {
		tx.Rollback()
		zap.S().Error("mysql.DecreasePostVoteNum(post_id, tx) fail", " #err: ", err.Error())
		return
	}
	tx.Commit()
	if err != nil {
		tx.Rollback()
		zap.S().Error("tx.Commit() fail", " err: ", err.Error())
	}
	return
}
