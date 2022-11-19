package redis

import (
	"bluebell/setting"
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
)

var rdb *redis.Client

var (
	PostPrefix   = "post:"
	VoteUserSet  = "vote:userVote"    //  value值为 post_id+"_"+user_id
	VotePostZSet = "vote:postVoteNum" //  value值为  post_id  对应分数为 点赞数
)

func Init(cfg *setting.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	_, err = rdb.Ping(context.TODO()).Result()
	return
}
func Close() {
	_ = rdb.Close()
}
