package logic

import (
	"bluebell/dao/models"
	"bluebell/dao/mysql"
	"errors"
)

func SelectCommunityInfo() ([]models.CommunityInfoReply, error) {
	reply, err := mysql.SelectCommunityInfo()
	return reply, err
}
func SelectCommunityIntroBycId(cid int) (string, error) {
	intro, err := mysql.SelectCommunityIntroBycId(cid)
	if err != nil {
		return "", err
	}
	if len(intro) == 0 {
		return "", errors.New(" 介绍信息为\"\" ")
	}
	return intro, err
}
