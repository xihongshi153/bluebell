package snowflake

import (
	"testing"
)

func TestSnowflake(t *testing.T) {
	err := Init("2003-11-09", 1)
	if err != nil {
		t.Error("init fail err: ", err)
		return
	}
	for i := 0; i < 10; i++ {

		t.Log(GenID())
	}
}
