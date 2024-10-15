package globals

import (
	"fmt"
	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
	"log"
)

var (
	R      *redis.Client
	R_LOCK *redislock.Client
)

func InitRedis() {
	if CONF == nil {
		log.Fatalf("config must be present")
	}
	r := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", CONF.Redis.Host, CONF.Redis.Port),
		Password: CONF.Redis.Password,
		DB:       CONF.Redis.Db,
		DisableIndentity: true,
	})
	R = r
	R_LOCK = redislock.New(r)
	
}
