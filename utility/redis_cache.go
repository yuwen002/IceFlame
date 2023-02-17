package utility

import (
	"context"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

// RedisExistsData
// @Description:
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-17 11:01:41
type RedisExistsData struct {
	Config string
	Key    string
	Value  interface{}
	Expire int64
}

// RedisCache
// @Description: Redis缓存
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-15 17:26:48
type RedisCache struct {
	Config string
	ctx    context.Context
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
	if rc.Config == "" {
		redis = g.Redis()
	} else {
		redis = g.Redis(rc.Config)
	}

	return redis
}

// ExistsSetData
//
// @Title 判断集合重元素是否存在
// @Description 判断集合重元素是否存在, 不存在将集合元素写入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-17 09:50:00
// @receiver rc
// @param data
// @param f
// @return code
//
//	0, "集合中数据已存在"
//	1, "失败，添加重复元素"，"添加元素成功"
//	2, "添加元素成功, key不存在或者不能为key设置过期时间"
//
// @return message
// @return err
func (rc *RedisCache) ExistsSetData(data RedisExistsData, f func(condition interface{}) (code int32, message string, err error)) (code int32, message string, err error) {
	// 初始化Redis
	redis := rc.InitRedis()
	res, err := redis.SIsMember(rc.ctx, data.Key, data.Value)
	if err != nil {
		return -1, "", err
	}

	// 数据存在，返回结果
	if res == 1 {
		return 0, "集合中数据已存在", nil
	}

	// 数据不存在，查询验证
	code, message, err = f(data.Value)
	if err != nil {
		return -1, "", err
	}

	// 查询数据不存在，直接返回
	if code == 1 {
		return code, message, nil
	}

	// 数据存在写入集合信息
	res, err = redis.SAdd(rc.ctx, data.Key, data.Value)
	if err != nil {
		return -1, "", err
	}

	// 添加元素失败
	if res == 0 {
		return 1, "失败，添加重复元素", nil
	}

	// 添加元素成功，验证过期时间是否开启
	if data.Expire > 0 {
		res, err = redis.Expire(rc.ctx, data.Key, data.Expire)
		if err != nil {
			return -1, "", err
		}

		if res == 0 {
			return 2, "添加元素成功, key不存在或者不能为key设置过期时间", nil
		}
	}

	return 0, "添加元素成功", nil
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
	var redis = RedisCache{Config: config[0]}
	return redis.InitRedis()
}

// ExistsSetData
//
// @Title 判断集合重元素是否存在
// @Description 判断集合重元素是否存在, 不存在将集合元素写入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-17 11:50:45
// @param data
// @param f
// @return code
// @return message
// @return err
func ExistsSetData(data RedisExistsData, f func(condition interface{}) (code int32, message string, err error)) (code int32, message string, err error) {
	var redis = RedisCache{Config: data.Config}
	return redis.ExistsSetData(data, f)
}
