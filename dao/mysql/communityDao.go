package mysql

import (
	"bluebell/dao/models"
)

func SelectCommunityInfo() ([]models.CommunityInfoReply, error) {
	sqlStr := "select community_id ,community_name ,introduction from community;"
	row, err := db.Query(sqlStr)
	reply := make([]models.CommunityInfoReply, 0)
	if err != nil {
		return reply, err
	}
	for row.Next() {
		info := models.CommunityInfoReply{}
		row.Scan(&info.CId, &info.CName, &info.CIntroduction)
		reply = append(reply, info)
	}
	return reply, nil
}

func SelectCommunityIntroBycId(cid int) (string, error) {
	sqlStr := "select introduction from community where community_id = ?;"
	row := db.QueryRow(sqlStr, cid)
	var intro string = ""
	err := row.Scan(&intro)
	return intro, err
}
func SelectCommunityInfoById(id int) (com models.Community, err error) {
	sqlStr := "select community_name,introduction from community where community_id = ?;"
	row := db.QueryRow(sqlStr, id)
	err = row.Scan(&com.Community_name, &com.Introduction)
	return
}
