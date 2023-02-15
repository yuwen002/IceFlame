package utility

import (
	"context"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

type RedisExistsData struct {
	RedisKey   string
	RedisValue interface{}
	Message    string
}

// RedisCache
// @Description: Redis缓存
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-15 17:26:48
type RedisCache struct {
	RedisConfig string

	ctx context.Context
}

// InitRedis
//
// @Title 初始化Redis
// @Description 初始化Redis
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-15 17:31:04
// @receiver rc
// @return *gredis.Redis
func (rc *RedisCache) InitRedis() *gredis.Redis {
	// 判断Redis配置链接
	var redis *gredis.Redis
	if rc.RedisConfig == "" {
		redis = g.Redis()
	} else {
		redis = g.Redis(rc.RedisConfig)
	}

	return redis
}

func (rc *RedisCache) ExistsData(data RedisExistsData) (code int32, message string, err error) {
	redis := rc.InitRedis()
	res, err := redis.SIsMember(rc.ctx, data.RedisKey, data.RedisValue)
	if err != nil {
		return -1, "", err
	}

	if res == 1 {
		return 0, data.Message, nil
	}
}

// InitRedis
//
// @Title 初始化Redis函数
// @Description 初始化Redis函数
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-15 17:52:28
// @param config
// @return *gredis.Redis
func InitRedis(config ...string) *gredis.Redis {
	var redis = RedisCache{RedisConfig: config[0]}
	return redis.InitRedis()
}
