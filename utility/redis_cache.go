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
	Config string      // Redis配置信息
	Key    string      // Redis键值
	Field  string      // Redis字段
	Data   interface{} // 更新数据字段
	Expire int64       // Redis过期时间
}

// RedisHashIdData
// @Description: hash ID 数据
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-02 10:07:24
type RedisHashIdData struct {
	Config string      // Redis配置信息
	Key    string      // Redis键值
	Id     string      // Redis字段
	Data   interface{} // 更新数据字段
	Expire int64       // Redis过期时间
}

type RedisDelKey struct {
	Config string // Redis配置信息
	Key    string // Redis键值
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

// DelKey
//
// @Title 删除缓存key
// @Description 删除缓存key
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-03 14:13:04
// @receiver rc
// @param key
// @return code
// @return message
// @return err
func (rc *RedisCache) DelKey(data RedisDelKey) (code int32, message string, err error) {
	// 初始化Redis
	redis := rc.InitRedis()
	res, err := redis.Del(rc.ctx, data.Key)
	if err != nil {
		return -1, "", err
	}

	if res == 0 {
		return 1, "删除成功", nil
	}

	return 0, "删除成功", nil
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
//	2, "失败，添加重复元素"，"添加元素成功"
//	3, "添加元素成功, key不存在或者不能为key设置过期时间"
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
		return 2, "失败，添加重复元素", nil
	}

	// 添加元素成功，验证过期时间是否开启
	if data.Expire > 0 {
		res, err = redis.Expire(rc.ctx, data.Key, data.Expire)
		if err != nil {
			return -1, "", err
		}

		if res == 0 {
			return 3, "添加元素成功, 不能为key设置过期时间", nil
		}
	}

	return 0, "添加元素成功,集合中数据已存在", nil
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
		if err != nil {
			return -1, "", nil, err
		}

		if result == 0 {
			return 2, "添加元素成功, 不能为key设置过期时间", output, nil
		}
	}

	if result == 1 {
		return 0, "取出转换数据成功(新字段)", output, nil
	}

	return 0, "取出转换数据成功(覆盖字段值)", output, nil
}

// CastHashData
//
// @Title 更新数据转换
// @Description 更新数据转换，例如：店铺ID转换商家ID
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-18 11:49:39
// @receiver rc
// @param data
// @param f
// @return code
// @return message
// @return output
// @return err

func (rc *RedisCache) UpdateCastHashData(data RedisCastData) (code int32, message string, err error) {
	// 初始化Redis
	redis := rc.InitRedis()
	// 更新hash字段
	field := map[string]interface{}{data.Field: data.Data}
	result, err := redis.HSet(rc.ctx, data.Key, field)
	if err != nil {
		return -1, "", err
	}

	if result == 1 {
		return 1, "取出转换数据成功(新字段)", nil
	}

	return 0, "取出转换数据成功(覆盖字段值)", nil
}

// SetHashId
//
// @Title 按ID添加缓存
// @Description 按ID添加缓存信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-02 10:33:57
// @receiver rc
// @param data
// @return code
// @return message
// @return err
func (rc *RedisCache) SetHashId(data RedisHashIdData) (code int32, message string, err error) {
	// 初始化Redis
	redis := rc.InitRedis()
	// 写入或更新ID信息
	field := map[string]interface{}{data.Id: data.Data}
	res, err := redis.HSet(rc.ctx, data.Key, field)
	if err != nil {
		return -1, "", err
	}

	// 添加元素成功，验证过期时间是否开启
	if data.Expire > 0 {
		result, err := redis.Expire(rc.ctx, data.Key, data.Expire)
		if err != nil {
			return -1, "", err
		}
		if result == 0 {
			return 1, "缓存设置成功, 不能为key设置过期时间", nil
		}
	}

	if res == 1 {
		message = "(新字段)"
	} else {
		message = "(覆盖字段值)"
	}

	return 0, "缓存设置成功" + message, nil
}

// GetHashId
//
// @Title 按ID获取缓存
// @Description 按ID获取缓存信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-02 13:53:32
// @receiver rc
// @param data
// @return code
// @return message
// @return output
// @return err
func (rc *RedisCache) GetHashId(data RedisHashIdData) (code int32, message string, output *g.Var, err error) {
	// 初始化Redis
	redis := rc.InitRedis()
	res, err := redis.HGet(rc.ctx, data.Key, data.Id)
	if err != nil {
		return -1, "", nil, err
	}
	if res.IsEmpty() {
		return 1, "缓存数据不存在", nil, nil
	}

	return 0, "取出缓存数据成功", output, nil
}

// GetMapHashId
//
// @Title 按ID获取缓存
// @Description 按ID获取缓存信息，返回map
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-02 13:54:36
// @receiver rc
// @param data
// @return code
// @return message
// @return output
// @return err
func (rc *RedisCache) GetMapHashId(data RedisHashIdData) (code int32, message string, output map[string]interface{}, err error) {
	code, message, out, err := rc.GetHashId(data)
	if code != 0 {
		return code, message, nil, err
	}

	return 0, "取出缓存数据成功", out.Map(), nil
}

// GeStructHashId
//
// @Title 按ID获取缓存
// @Description 按ID获取缓存信息，返回struct
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-02 13:06:24
// @receiver rc
// @param data
// @param output
// @return code
// @return message
// @return err
func (rc *RedisCache) GeStructHashId(data RedisHashIdData, output interface{}) (code int32, message string, err error) {
	code, message, out, err := rc.GetHashId(data)
	if code != 0 {
		return code, message, err
	}

	err = out.Struct(output)
	if err != nil {
		return -1, "", err
	}

	return 0, "取出缓存数据成功", nil
}

// GetHashAll
//
// @Title 获取所有字段和值
// @Description 获取所有字段和值
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-02 15:58:57
// @receiver rc
// @param data
// @return code
// @return message
// @return output
// @return err
func (rc *RedisCache) GetHashAll(data RedisHashIdData) (code int32, message string, output *g.Var, err error) {
	// 初始化数据
	redis := rc.InitRedis()
	res, err := redis.HGetAll(rc.ctx, data.Key)
	if err != nil {
		return -1, "", nil, err
	}

	if res.IsEmpty() {
		return 1, "缓存数据不存在", nil, nil
	}

	return 0, "取出缓存数据成功", res, nil
}

// GetMapHashAll
//
// @Title 获取所有字段和值
// @Description 获取所有字段和值,返回map数组
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-02 16:05:22
// @receiver rc
// @param data
// @return code
// @return message
// @return output
// @return err
func (rc *RedisCache) GetMapHashAll(data RedisHashIdData) (code int32, message string, output map[string]interface{}, err error) {
	code, message, out, err := rc.GetHashAll(data)
	if code != 0 {
		return code, message, nil, err
	}
	return code, message, out.Map(), nil
}

// GetStructsHashAll
//
// @Title 获取所有字段和值
// @Description 获取所有字段和值,返回struct数组
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-02 16:15:53
// @receiver rc
// @param data
// @param output
// @return code
// @return message
// @return err
func (rc *RedisCache) GetStructsHashAll(data RedisHashIdData, output interface{}) (code int32, message string, err error) {
	code, message, out, err := rc.GetHashAll(data)
	if code != 0 {
		return code, message, err
	}

	err = out.Structs(output)
	if err != nil {
		return -1, "", err
	}

	return 0, message, nil
}

// DelHashId
//
// @Title 按ID删除缓存
// @Description 按ID删除缓存信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-02 11:14:38
// @receiver rc
// @param data
// @return code
// @return message
// @return err
func (rc *RedisCache) DelHashId(data RedisHashIdData) (code int32, message string, err error) {
	// 初始化Redis
	redis := rc.InitRedis()
	// 按ID删除缓存信息
	res, err := redis.HDel(rc.ctx, data.Key, data.Id)
	if err != nil {
		return -1, "", err
	}

	if res == 1 {
		return 1, "删除数据成功", nil
	}

	return 1, "删除ID不存在", nil
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

// RCDelKey
//
// @Title 删除缓存key
// @Description 删除缓存key
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-03 14:45:33
// @param data
// @return code
// @return message
// @return err
func RCDelKey(data RedisDelKey) (code int32, message string, err error) {
	var redis = RedisCache{Config: data.Config}
	return redis.DelKey(data)
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

// RCUpdateCastHashData
//
// @Title 更新数据转换
// @Description 更新数据转换，例如：店铺ID转换商家IDn
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-02-27 23:45:47
// @param data
// @return code
// @return message
// @return err
func RCUpdateCastHashData(data RedisCastData) (code int32, message string, err error) {
	var redis = RedisCache{Config: data.Config}
	return redis.UpdateCastHashData(data)
}

// RCSetHashId
//
// @Title 按ID添加缓存
// @Description 按ID添加缓存信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-02 14:10:52
// @param data
func RCSetHashId(data RedisHashIdData) (code int32, message string, err error) {
	var redis = RedisCache{Config: data.Config}
	return redis.SetHashId(data)
}

// RCGetHashId
//
// @Title 按ID获取缓存
// @Description 按ID获取缓存信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-02 14:38:35
// @param data
// @return code
// @return message
// @return output
// @return err
func RCGetHashId(data RedisHashIdData) (code int32, message string, output *g.Var, err error) {
	var redis = RedisCache{Config: data.Config}
	return redis.GetHashId(data)
}

// RCGetMapHashId
//
// @Title 按ID获取缓存
// @Description 按ID获取缓存信息，返回map
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-02 14:14:57
// @param data
// @return code
// @return message
// @return output
// @return err
func RCGetMapHashId(data RedisHashIdData) (code int32, message string, output map[string]interface{}, err error) {
	var redis = RedisCache{Config: data.Config}
	return redis.GetMapHashId(data)
}

// RCGeStructHashId
//
// @Title 按ID获取缓存
// @Description 按ID获取缓存信息，返回struct
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-02 14:19:02
// @param data
// @param output
// @return code
// @return message
// @return err
func RCGeStructHashId(data RedisHashIdData, output interface{}) (code int32, message string, err error) {
	var redis = RedisCache{Config: data.Config}
	return redis.GeStructHashId(data, output)
}

// RCDelHashId
//
// @Title 按ID删除缓存
// @Description 按ID删除缓存信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-02 14:40:26
// @param data
// @return code
// @return message
// @return err
func RCDelHashId(data RedisHashIdData) (code int32, message string, err error) {
	var redis = RedisCache{Config: data.Config}
	return redis.DelHashId(data)
}

// RCGetHashAll
//
// @Title 获取所有字段和值
// @Description 获取所有字段和值
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-02 16:51:54
// @param data
// @return code
// @return message
// @return output
// @return err
func RCGetHashAll(data RedisHashIdData) (code int32, message string, output *g.Var, err error) {
	var redis = RedisCache{Config: data.Config}
	return redis.GetHashAll(data)
}

// RCGetMapsHashAll
//
// @Title 获取所有字段和值
// @Description 获取所有字段和值,返回map数组
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-02 16:35:25
// @param data
// @return code
// @return message
// @return output
// @return err
func RCGetMapsHashAll(data RedisHashIdData) (code int32, message string, output map[string]interface{}, err error) {
	var redis = RedisCache{Config: data.Config}
	return redis.GetMapHashAll(data)
}

// RCGetStructsHashAll
//
// @Title 获取所有字段和值
// @Description 获取所有字段和值,返回struct数组
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-02 16:53:08
// @param data
// @param output
// @return code
// @return message
// @return err
func RCGetStructsHashAll(data RedisHashIdData, output interface{}) (code int32, message string, err error) {
	var redis = RedisCache{Config: data.Config}
	return redis.GetStructsHashAll(data, output)
}
