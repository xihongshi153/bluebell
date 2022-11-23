package jwt

import "testing"

func TestJWTgerateAndParse(t *testing.T) {
	username := "tomato"
	userid := "1111111111"
	token, err := GenerateJwt(username, userid)
	if err != nil {
		t.Errorf("jwt generate fail err:%s", err.Error())
		return
	}
	calims, err := Parsejwt(token)
	if err != nil {
		t.Errorf("jwt parse fail err:%s", err.Error())
	}
	if calims.Username != username {
		t.Errorf("数据解析错误 username不相等")
		return
	}
	if calims.UserId != userid {
		t.Errorf("数据解析错误 userid不相等")
		return
	}
}
