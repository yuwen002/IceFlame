package utility

import (
	"context"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

// RedisExistsData
// @Description: 集合数据
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-17 11:01:41
type RedisExistsData struct {
	Config string      // Redis配置信息
	Key    string      // Redis键值
	Value  interface{} // Redis值
	Expire int64       // Redis过期时间
}

// RedisCastData
// @Description: hash数据转换
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-18 11:23:26
type RedisCastData struct {
	Config string // Redis配置信息
	Key    string // Redis键值
	Field  string // Redis字段
	Expire int64  // Redis过期时间
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
	if code != 0 {
		return code, message, err
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
			return 2, "添加元素成功, 不能为key设置过期时间", nil
		}
	}

	return 0, "添加元素成功", nil
}

// CastHashData
//
// @Title 数据转换
// @Description 数据转换，例如：店铺ID转换商家ID
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-18 11:49:39
// @receiver rc
// @param data
// @param f
// @return code
// @return message
// @return output
// @return err
func (rc *RedisCache) CastHashData(data RedisCastData, f func(condition interface{}) (code int32, message string, output interface{}, err error)) (code int32, message string, output interface{}, err error) {
	// 初始化Redis
	redis := rc.InitRedis()
	res, err := redis.HGet(rc.ctx, data.Key, data.Field)
	if err != nil {
		return -1, "", nil, err
	}

	// 数据正确，取出成功
	if !res.IsEmpty() {
		return 0, "取出转换数据成功", res, nil
	}

	// 无数据从新读取数据
	code, message, output, err = f(data.Field)
	if code != 0 {
		return code, message, nil, err
	}

	field := map[string]interface{}{data.Field: output}
	result, err := redis.HSet(rc.ctx, data.Key, field)
	if err != nil {
		return -1, "", nil, err
	}

	// 添加元素成功，验证过期时间是否开启
	if data.Expire > 0 {
		result, err = redis.Expire(rc.ctx, data.Key, data.Expire)
		return -1, "", nil, err

		if result == 0 {
			return 2, "添加元素成功, 不能为key设置过期时间", output, nil
		}
	}

	if result == 1 {
		return 0, "取出转换数据成功(新字段)", output, nil
	}

	return 0, "取出转换数据成功(覆盖字段值)", output, nil
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

// RCExistsSetData
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
//
//	0, "集合中数据已存在"
//	1, "失败，添加重复元素"，"添加元素成功"
//	2, "添加元素成功, key不存在或者不能为key设置过期时间"
func RCExistsSetData(data RedisExistsData, f func(condition interface{}) (code int32, message string, err error)) (code int32, message string, err error) {
	var redis = RedisCache{Config: data.Config}
	return redis.ExistsSetData(data, f)
}

// RCCastHashData
//
// @Title 数据转换
// @Description 数据转换，例如：店铺ID转换商家ID
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-18 15:09:58
// @param data
// @param f
// @return code
// @return message
// @return output
// @return err
func RCCastHashData(data RedisCastData, f func(condition interface{}) (code int32, message string, output interface{}, err error)) (code int32, message string, output interface{}, err error) {
	var redis = RedisCache{Config: data.Config}
	return redis.CastHashData(data, f)
}
