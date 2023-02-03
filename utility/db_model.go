package utility

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"time"
)

// RedisData
// @Description: Redis缓存配置
// @Author liuxingyu <yuwen002@163.com>
// @TestData 2023-01-27 00:46:01
type RedisData struct {
	Key    string
	Data   interface{}
	Expire int64
}

// DBLimit
// @Description: 分页
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-29 15:24:58
type DBLimit struct {
	Offset int
	Length int
}

// DBPagination
// @Description: 分页
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-30 12:54:04
type DBPagination struct {
	Page int
	Size int
}

// DBInsertInput
// @Description: 数据库写入数据
// @Author liuxingyu <yuwen002@163.com>
// @TestData 2023-01-28 23:29:48
type DBInsertInput struct {
	Data interface{} // 数据库数据插入
}

// DBGetByIdInput
// @Description: 按ID获取信息
// @Author liuxingyu <yuwen002@163.com>
// @TestData 2023-01-28 23:29:58
type DBGetByIdInput struct {
	Field       string      // 查询字段
	Where       interface{} // 查询条件
	Args        interface{} // 查询参数
	CacheKey    string      // 缓存KEY
	RedisConfig string      // 缓存配置信息
	Expire      int64       // 缓存过期时间
}

// DBGetByIdsInput
// @Description: 按ID获取多条信息
// @Author liuxingyu <yuwen002@163.com>
// @TestData 2023-01-31 23:50:39
type DBGetByIdsInput struct {
	Field  string      // 查询字段
	Column string      // In主键字段
	In     interface{} // In条件信息
	Order  string      // 排序
}

// DBGetOneByWhereInput
// @Description: 条件获取单条信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-29 15:03:16
type DBGetOneByWhereInput struct {
	Field string      // 查询字段
	Where interface{} // 查询条件
	Args  interface{} // 查询参数
	Order string      // 排序
}

// DBGetAllByWhereInput
// @Description: 条件获取多条信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-01 17:42:21
type DBGetAllByWhereInput struct {
	Field        string      // 查询字段
	Where        interface{} // 查询条件
	Args         interface{} // 查询参数
	Order        string      // 排序
	PageType     int8        // 分页类型
	DBLimit                  // 分页，偏移量
	DBPagination             // 分页，页码
}

// DBModifyByIdInput
// @Description: 按ID修改单条数据
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-01 17:42:36
type DBModifyByIdInput struct {
	Data        interface{} // 更新数据
	Where       interface{} // 查询条件
	Args        interface{} // 查询参数
	OmitEmpty   bool        // 是否过滤空字段
	CacheKey    string      // 缓存KEY
	RedisConfig string      // 缓存配置信息
}

// DBModifyByWhereInput
// @Description: 按条件修改数据
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-01 17:43:20
type DBModifyByWhereInput struct {
	Data         interface{} // 更新数据
	Where        interface{} // 查询条件
	Args         interface{} // 查询参数
	Order        string      // 排序
	OmitEmpty    bool        // 是否过滤空字段
	PageType     int8        // 分页类型
	DBLimit                  // 分页，偏移量
	DBPagination             // 分页，页码
}

// DBDelByIdInput
// @Description: 按ID删除数据
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-01 17:43:51
type DBDelByIdInput struct {
	Where       interface{} // 查询条件
	Args        interface{} // 查询参数
	Cache       bool        // 缓存类型
	CacheKey    string      // 缓存KEY
	RedisConfig string      // 缓存配置信息
}

// DB
// @Description: 数据库常用操作结构体
// @Author liuxingyu <yuwen002@163.com>
// @TestData 2023-01-28 23:28:34
type DB struct {
	ctx         context.Context
	M           *gdb.Model
	RedisConfig string
}

// InitRedis
//
// @Title: Redis初始化连接
// @Description: Redis初始化连接
// @Author liuxingyu <yuwen002@163.com>
// @TestData 2023-01-28 23:43:51
// @receiver db
// @return *gredis.Redis
func (db *DB) InitRedis() *gredis.Redis {
	// 判断Redis配置链接
	var redis *gredis.Redis
	if db.RedisConfig == "" {
		redis = g.Redis()
	} else {
		redis = g.Redis(db.RedisConfig)
	}

	return redis
}

// SetRedisCache
//
// @Title: Redis缓存写入
// @Description:  Redis缓存写入
// @Author liuxingyu <yuwen002@163.com>
// @TestData 2023-01-28 23:45:33
// @receiver db
// @param cache
// @return code
// @return message
// @return err
func (db *DB) SetRedisCache(cache RedisData) (code int32, message string, err error) {
	// Redis初始化
	redis := db.InitRedis()
	// 数据写入
	res, err := redis.Set(db.ctx, cache.Key, cache.Data)
	if err != nil {
		return -1, "", err
	}

	// 失败返回
	if res.String() != "OK" {
		return 1, "Redis数据写入失败", nil
	}

	// 过期时间开启
	if cache.Expire > 0 {
		_, _ = redis.Do(db.ctx, "EXPIRE", cache.Key, cache.Expire)
	}

	return 0, "数据写入成功", nil
}

// GetStructRedisCache
//
// @Title: 缓存取出
// @Description: Redis缓存获取，返回struct
// @Author liuxingyu <yuwen002@163.com>
// @TestData 2023-01-28 23:54:21
// @receiver db
// @param cache
// @param output
// @return code
// @return message
// @return err
func (db *DB) GetStructRedisCache(cache RedisData, output interface{}) (code int32, message string, err error) {
	// Redis初始化
	redis := db.InitRedis()
	// 数据获取
	res, err := redis.Get(db.ctx, cache.Key)
	if err != nil {
		return -1, "", err
	}

	if res.IsNil() {
		return 1, "缓存信息不存在", nil
	}

	err = res.Struct(output)
	if err != nil {
		return -1, "", err
	}

	return 0, "缓存数据取出成功", nil
}

// GetMapRedisCache
//
// @Title 缓存取出
// @Description Redis缓存获取，返回map
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-29 10:56:37
// @receiver db
// @param cache
// @return code
// @return message
// @return output
// @return err
func (db *DB) GetMapRedisCache(cache RedisData) (code int32, message string, output interface{}, err error) {
	// Redis初始化
	redis := db.InitRedis()
	// 数据获取
	res, err := redis.Get(db.ctx, cache.Key)
	if err != nil {
		return -1, "", nil, err
	}

	if res.IsNil() {
		return 1, "缓存信息不存在", nil, nil
	}

	return 0, "缓存数据取出成功", res.Map(), nil
}

// DelRedisCache
//
// @Title 删除缓存
// @Description 删除缓存信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-31 16:58:40
// @receiver db
// @param cache
// @return code
// @return message
// @return err
func (db *DB) DelRedisCache(cache RedisData) (code int32, message string, err error) {
	// Redis初始化
	redis := db.InitRedis()
	res, err := redis.Del(db.ctx, cache.Key)
	if err != nil {
		return -1, "", err
	}

	if res == 1 {
		return 0, "删除缓存信息成功", nil
	}

	return 1, "缓存信息key不存在", nil
}

// Insert
//
// @Title 数据写入
// @Description 数据库插入方法，无返回值
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-28 14:21:25
// @receiver db
// @param condition
// @return code
// @return message
// @return err
func (db *DB) Insert(condition DBInsertInput) (code int32, message string, err error) {
	// 数据写入
	res, err := db.M.OmitNil().OmitEmpty().Data(condition.Data).Insert()
	if err != nil {
		return -1, "", err
	}

	insertId, err := res.LastInsertId()
	if err != nil {
		return -1, "", err
	}

	if insertId > 0 {
		return 0, "数据写入成功", nil
	}

	return 1, "数据写入失败", nil
}

// InsertAndGetId
//
// @Title: 数据写入
// @Description: 数据库插入方法，返回写入ID
// @Author Lxy <yuwen002@163.com>
// @Date 2023-01-28 14:42:45
// @receiver db
// @param condition
// @return code
// @return message
// @return lastInsertId
// @return err
func (db *DB) InsertAndGetId(condition DBInsertInput) (code int32, message string, lastInsertId int64, err error) {
	// 数据写入
	lastInsertId, err = db.M.OmitNil().OmitEmpty().InsertAndGetId(condition.Data)
	if err != nil {
		return -1, "", 0, err
	}

	if lastInsertId > 0 {
		return 0, "数据写入成功", lastInsertId, nil
	}

	return 1, "数据写入失败", 0, nil
}

// ConditionGetById
//
// @Title: 数据写入条件判断
// @Description: 数据写入条件判断，查询字段判断
// @Author liuxingyu <yuwen002@163.com>
// @TestData 2023-01-28 23:33:08
// @receiver db
// @param condition
// @return *gdb.Model
func (db *DB) ConditionGetById(condition DBGetByIdInput) *gdb.Model {
	// 查询字段
	if condition.Field != "" {
		db.M = db.M.Fields(condition.Field)
	}

	// 查询条件参数判断
	if condition.Args == nil {
		db.M = db.M.Where("id=?", condition.Where)
	} else {
		db.M = db.M.Where(condition.Where, condition.Args)
	}

	return db.M
}

// GetStructById
//
// @Title: 按ID（主键）查询数据
// @Description: 按ID（主键）查询数据，返回类型struct
// @Author liuxingyu <yuwen002@163.com>
// @TestData 2023-01-28 23:33:53
// @receiver db
// @param condition
// @param output
// @return code
// @return message
// @return err
func (db *DB) GetStructById(condition DBGetByIdInput, output interface{}) (code int32, message string, err error) {
	// 查询数据
	err = db.ConditionGetById(condition).Scan(output)
	if err != nil {
		return -1, "", err
	}

	if output == nil {
		return 1, "无查询信息", nil
	}

	return 0, "数据查询成功", nil
}

// GetStructGCacheById
//
// @Title: 按ID获取数据，带缓存
// @Description: 按ID获取数据，gcache缓存，返回struct
// @Author liuxingyu <yuwen002@163.com>
// @TestData 2023-01-28 23:52:46
// @receiver db
// @param condition
// @param output
// @return code
// @return message
// @return err
func (db *DB) GetStructGCacheById(condition DBGetByIdInput, output interface{}) (code int32, message string, err error) {
	// Redis 缓存适配
	redisCache := gcache.NewAdapterRedis(db.InitRedis())
	g.DB().GetCache().SetAdapter(redisCache)

	// 缓存查询
	err = db.ConditionGetById(condition).Cache(gdb.CacheOption{
		Duration: time.Second * time.Duration(condition.Expire),
		Name:     condition.CacheKey,
		Force:    false,
	}).Scan(output)

	if err != nil {
		return -1, "", err
	}

	if output == nil {
		return 1, "查询数据不存在", nil
	}

	return 0, "查询数据成功", nil
}

// GetStructRCacheById
//
// @Title: 按ID获取数据，带缓存
// @Description: 按ID获取数据，Redis缓存，返回struct
// @Author liuxingyu <yuwen002@163.com>
// @TestData 2023-01-29 00:04:18
// @receiver db
// @param condition
// @param output
// @return code
// @return message
// @return err
func (db *DB) GetStructRCacheById(condition DBGetByIdInput, output interface{}) (code int32, message string, err error) {
	// 查询缓存数据
	code, message, err = db.GetStructRedisCache(RedisData{Key: condition.CacheKey}, output)
	if err != nil {
		return -1, "", err
	}

	if code == 0 {
		return 0, "查询数据成功", nil
	}

	// 查询数据
	code, message, err = db.GetStructById(condition, output)
	if err != nil {
		return -1, "", err
	}

	if code == 1 {
		return code, message, nil
	}

	// 缓存信息写入
	code, _, err = db.SetRedisCache(RedisData{
		Key:    condition.CacheKey,
		Data:   output,
		Expire: condition.Expire,
	})
	if err != nil {
		return -1, "", err
	}
	if code == 1 {
		return 2, "查询数据成功, 写入缓存失败", nil
	}

	return 0, "查询数据成功", nil
}

// GetMapById
//
// @Title 按ID获取数据
// @Description 按ID获取数据，返回map
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-29 09:51:04
// @receiver db
// @param condition
// @return code
// @return message
// @return output
// @return err
func (db *DB) GetMapById(condition DBGetByIdInput) (code int32, message string, output interface{}, err error) {
	// 数据查询
	res, err := db.ConditionGetById(condition).One()
	if err != nil {
		return -1, "", nil, err
	}

	if res.IsEmpty() {
		return 1, "查询数据不存在", nil, nil
	}

	return 0, "数据查询成功", res.Map(), nil
}

// GetMapGCacheById
//
// @Title 按ID获取数据
// @Description 按ID获取数据，gcache缓存，返回map
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-29 10:45:49
// @receiver db
// @param condition
// @return code
// @return message
// @return output
// @return err
func (db *DB) GetMapGCacheById(condition DBGetByIdInput) (code int32, message string, output interface{}, err error) {
	// Redis 缓存适配
	redisCache := gcache.NewAdapterRedis(db.InitRedis())
	g.DB().GetCache().SetAdapter(redisCache)

	// 缓存查询数据
	res, err := db.ConditionGetById(condition).Cache(gdb.CacheOption{
		Duration: time.Second * time.Duration(condition.Expire),
		Name:     condition.CacheKey,
		Force:    false,
	}).One()

	if err != nil {
		return -1, "", nil, err
	}

	if res.IsEmpty() {
		return 1, "查询数据不存在", nil, nil
	}

	return 0, "数据查询成功", res.Map(), nil
}

// GetMapRCacheById
//
// @Title 按ID获取数据
// @Description 按ID获取数据，gcache缓存，返回map
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-29 13:50:24
// @receiver db
// @param condition
// @return code
// @return message
// @return output
// @return err
func (db *DB) GetMapRCacheById(condition DBGetByIdInput) (code int32, message string, output interface{}, err error) {
	// 缓存数据查询
	code, message, output, err = db.GetMapRedisCache(RedisData{Key: condition.CacheKey})
	if err != nil {
		return -1, "", nil, err
	}

	if code == 0 {
		return 0, "查询数据成功", output, nil
	}

	// 无缓存数据，查询数据库
	code, message, output, err = db.GetMapById(condition)
	if err != nil {
		return -1, "", nil, err
	}

	if code == 1 {
		return 1, message, nil, nil
	}

	// 查询缓存信息写入
	code, message, err = db.SetRedisCache(RedisData{
		Key:    condition.CacheKey,
		Data:   output,
		Expire: condition.Expire,
	})
	if err != nil {
		return -1, "", nil, err
	}

	if code == 1 {
		return 2, "查询数据成功，缓存信息写入失败", output, nil
	}

	return 0, "查询数据成功", output, nil
}

// ConditionGetByIds
//
// @Title 按ID（主键）查询多条数据条件
// @Description 按ID（主键）查询多条数据,返回struct数组
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-01 17:20:16
// @receiver db
// @param condition
// @return *gdb.Model
func (db *DB) ConditionGetByIds(condition DBGetByIdsInput) *gdb.Model {
	// 查询字段
	if condition.Field != "" {
		db.M = db.M.Fields(condition.Field)
	}

	if condition.Order != "" {
		db.M = db.M.Order(condition.Order)
	}

	return db.M
}

// GetStructByIds
//
// @Title 按ID（主键）查询多条数据
// @Description 按ID（主键）查询多条数据,返回struct数组
// @Author liuxingyu <yuwen002@163.com>
// @TestData 2023-01-31 23:58:53
// @receiver db
// @param condition
// @param output
// @return code
// @return message
// @return err
func (db *DB) GetStructByIds(condition DBGetByIdsInput, output interface{}) (code int32, message string, err error) {

	err = db.ConditionGetByIds(condition).WhereIn(condition.Column, condition.In).Scan(output)
	if err != nil {
		return -1, "", err
	}

	if output == nil {
		return 1, "无查询信息", nil
	}

	return 0, "数据查询成功", nil
}

// GetMapByIds
//
// @Title 按ID（主键）查询多条数据
// @Description 按ID（主键）查询多条数据,返回map数组
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-01 15:50:49
// @receiver db
// @param condition
// @return code
// @return message
// @return output
// @return err
func (db *DB) GetMapByIds(condition DBGetByIdsInput) (code int32, message string, output interface{}, err error) {
	// 数据查询
	res, err := db.ConditionGetByIds(condition).WhereIn(condition.Column, condition.In).All()
	if err != nil {
		return -1, "", nil, err
	}

	if res.IsEmpty() {
		return 1, "查询数据不存在", nil, err
	}

	return 0, "数据查询成功", res.List(), nil
}

// GetOneByWhere
//
// @Title 条件查询判断
// @Description 条件查询判断，单条数据
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-29 15:09:58
// @receiver db
// @param condition
// @return *gdb.Model
func (db *DB) GetOneByWhere(condition DBGetOneByWhereInput) *gdb.Model {
	if condition.Field != "" {
		db.M = db.M.Fields(condition.Field)
	}

	if condition.Args == nil {
		db.M = db.M.Where(condition.Where)
	} else {
		db.M = db.M.Where(condition.Where, condition.Args)
	}

	if condition.Order == "" {
		db.M = db.M.Order("id desc")
	} else {
		db.M = db.M.Order(condition.Order)
	}

	return db.M
}

// GetOneStructByWhere
//
// @Title 按条件获取数据
// @Description 按条件获取单条数据，返回struct
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-29 15:09:34
// @receiver db
// @param condition
// @param output
// @return code
// @return message
// @return err
func (db *DB) GetOneStructByWhere(condition DBGetOneByWhereInput, output interface{}) (code int32, message string, err error) {
	// 查询数据
	err = db.GetOneByWhere(condition).Scan(output)
	if err != nil {
		return -1, "", err
	}

	if output == nil {
		return 1, "查询数据不存在", nil
	}

	return 0, "查询数据成功", nil
}

// GetOneMapByWhere
//
// @Title 按条件获取数据
// @Description 按条件获取单条数据，返回map
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-29 15:13:18
// @receiver db
// @param condition
// @return code
// @return message
// @return output
// @return err
func (db *DB) GetOneMapByWhere(condition DBGetOneByWhereInput) (code int32, message string, output interface{}, err error) {
	// 查询数据
	res, err := db.GetOneByWhere(condition).One()
	if err != nil {
		return -1, "", nil, err
	}

	if res.IsEmpty() {
		return 1, "查询数据不存在", nil, nil
	}

	return 0, "数据查询成功", res.Map(), nil
}

// ConditionGetAllByWhere
//
// @Title 按条件获取列表
// @Description 判断查询条件，分页信息
// @Author liuxingyu <yuwen002@163.com>
// @TestData 2023-01-29 23:37:47
// @receiver db
// @param condition
// @return *gdb.Model
func (db *DB) ConditionGetAllByWhere(condition DBGetAllByWhereInput) *gdb.Model {
	// 查询字段判断
	if condition.Field != "" {
		db.M = db.M.Fields(condition.Field)
	}

	// 条件判断
	if condition.Args == nil {
		if condition.Where != nil {
			db.M = db.M.Where(condition.Where)
		}
	} else {
		db.M = db.M.Where(condition.Where, condition.Args)
	}

	// 排序判断
	if condition.Order != "" {
		db.M = db.M.Order(condition.Order)
	}

	// 分页类型判断
	switch condition.PageType {
	case 1:
		// 判断是否需要分页
		if condition.Page > 0 {
			if condition.Size == 0 {
				// 分页默认长度
				condition.Size = 10
			}
			db.M = db.M.Page(condition.Page, condition.Size)
		}
	case 2:
		// 判断分页
		if condition.Length > 0 {
			db.M.Limit(condition.Offset, condition.Length)
		} else if condition.Offset > 0 {
			db.M.Limit(condition.Offset)
		}
	}

	return db.M
}

// GetAllStructByWhere
//
// @Title 按条件获取列表
// @Description 按条件获取列表，返回struct
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-30 17:45:48
// @receiver db
// @param condition
// @param output
// @return code
// @return message
// @return err
func (db *DB) GetAllStructByWhere(condition DBGetAllByWhereInput, output []interface{}) (code int32, message string, err error) {
	// 查询数据
	err = db.ConditionGetAllByWhere(condition).Scan(output)
	if err != nil {
		return -1, "", err
	}

	if len(output) == 0 {
		return 1, "查询数据不存在", nil
	}

	return 0, "查询数据成功", nil
}

// GetAllMapByWhere
//
// @Title 按条件获取列表
// @Description 按条件获取列表，返回list<[]Map>
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-31 10:04:44
// @receiver db
// @param condition
// @return code
// @return message
// @return output
// @return err
func (db *DB) GetAllMapByWhere(condition DBGetAllByWhereInput) (code int32, message string, output gdb.List, err error) {
	res, err := db.ConditionGetAllByWhere(condition).All()
	if err != nil {
		return -1, "", nil, err
	}
	if res.IsEmpty() {
		return 1, "查询数据不存在", nil, nil
	}

	return 0, "查询数据成功", res.List(), nil
}

// ConditionModifyById
//
// @Title 条件按ID修改数据
// @Description 按ID修改数据，条件判断
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-31 15:22:42
// @receiver db
// @param condition
// @return *gdb.Model
func (db *DB) ConditionModifyById(condition DBModifyByIdInput) *gdb.Model {
	// 判断条件
	if condition.Args == nil {
		db.M = db.M.Where("id=?", condition.Where)
	} else {
		db.M = db.M.Where(condition.Where, condition.Args)
	}

	// 过滤空字段是否开启
	if condition.OmitEmpty {
		db.M = db.M.OmitEmpty()
	}

	return db.M
}

// ModifyById
//
// @Title 按ID修改数据
// @Description 按ID修改数据
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-31 14:31:49
// @receiver db
// @param condition
// @return code
// @return message
// @return err
func (db *DB) ModifyById(condition DBModifyByIdInput) (code int32, message string, err error) {
	// 更新数据
	res, err := db.ConditionModifyById(condition).Data(condition.Data).OmitNil().Update()
	if err != nil {
		return -1, "", err
	}

	// 受影响行数
	affected, err := res.RowsAffected()
	if err != nil {
		return -1, "", err
	}

	if affected == 0 {
		return 1, "无变更数据", nil
	}

	return 0, "更新数据成功", nil
}

// ModifyGCacheById
//
// @Title 按ID修改数据
// @Description 按ID修改数据,清除gcache缓存
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-31 15:57:29
// @receiver db
// @param condition
// @return code
// @return message
// @return err
func (db *DB) ModifyGCacheById(condition DBModifyByIdInput) (code int32, message string, err error) {
	// Redis 缓存适配
	redisCache := gcache.NewAdapterRedis(db.InitRedis())
	g.DB().GetCache().SetAdapter(redisCache)

	// 更新数据
	res, err := db.ConditionModifyById(condition).Data(condition.Data).Cache(gdb.CacheOption{
		Duration: -1,
		Name:     condition.CacheKey,
		Force:    false,
	}).Update()
	// 受影响行数
	affected, err := res.RowsAffected()
	if err != nil {
		return -1, "", err
	}

	if affected == 0 {
		return 1, "无变更数据", nil
	}

	return 0, "更新数据成功", nil
}

// ModifyRCacheById
//
// @Title 按ID修改数据
// @Description 按ID修改数据,清除Redis缓存
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-31 17:20:40
// @receiver db
// @param condition
// @return code
// @return message
// @return err
func (db *DB) ModifyRCacheById(condition DBModifyByIdInput) (code int32, message string, err error) {
	res, err := db.ConditionModifyById(condition).Data(condition.Data).Update()
	if err != nil {
		return -1, "", err
	}

	// 受影响行数
	affected, err := res.RowsAffected()
	if err != nil {
		return -1, "", err
	}

	if affected == 0 {
		return 1, "无变更数据", nil
	}

	// 删除缓存信息
	_, _, err = db.DelRedisCache(RedisData{Key: condition.CacheKey})

	return 0, "更新数据成功", nil
}

// ModifyByWhere
//
// @Title 条件修改多条数据
// @Description 条件修改多条数据
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-31 17:57:32
// @receiver db
// @param condition
// @return code
// @return message
// @return err
func (db *DB) ModifyByWhere(condition DBModifyByWhereInput) (code int32, message string, err error) {
	// 条件判断
	if condition.Args == nil {
		if condition.Where != nil {
			db.M = db.M.Where(condition.Where)
		}
	} else {
		db.M = db.M.Where(condition.Where, condition.Args)
	}

	// 排序判断
	if condition.Order != "" {
		db.M = db.M.Order(condition.Order)
	}

	// 分页类型判断
	switch condition.PageType {
	case 1:
		// 判断是否需要分页
		if condition.Page > 0 {
			if condition.Size == 0 {
				// 分页默认长度
				condition.Size = 10
			}
			db.M = db.M.Page(condition.Page, condition.Size)
		}
	case 2:
		// 判断分页
		if condition.Length > 0 {
			db.M.Limit(condition.Offset, condition.Length)
		} else if condition.Offset > 0 {
			db.M.Limit(condition.Offset)
		}
	}
	// 过滤空字段是否开启
	if condition.OmitEmpty {
		db.M = db.M.OmitEmpty()
	}

	res, err := db.M.Data(condition.Data).Update()

	// 受影响行数
	affected, err := res.RowsAffected()
	if err != nil {
		return -1, "", err
	}

	if affected == 0 {
		return 1, "无变更数据", nil
	}

	return 0, "更新数据成功", nil
}

// DelById
//
// @Title 按ID删除数据
// @Description 按ID删除数据 开启cache缓存删除功能
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-01 17:52:41
// @receiver db
// @param condition
// @return code
// @return message
// @return err
func (db *DB) DelById(condition DBDelByIdInput) (code int32, message string, err error) {
	// 条件判断
	// 查询条件参数判断
	if condition.Args == nil {
		db.M = db.M.Where("id=?", condition.Where)
	} else {
		db.M = db.M.Where(condition.Where, condition.Args)
	}

	if condition.Cache == true {
		// Redis 缓存适配
		redisCache := gcache.NewAdapterRedis(db.InitRedis())
		g.DB().GetCache().SetAdapter(redisCache)

		db.M = db.M.Cache(gdb.CacheOption{
			Duration: -1,
			Name:     condition.CacheKey,
			Force:    false,
		})
	}

	res, err := db.M.Delete()
	if err != nil {
		return -1, "", err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return -1, "", err
	}

	if affected == 0 {
		return 1, "删除数据不存在", nil
	}

	if condition.Cache == true {
		_, _, err = db.DelRedisCache(RedisData{Key: condition.CacheKey})
		if err != nil {
			return -1, "", err
		}
	}

	return 0, "删除数据成功", nil
}

// DBInsert
//
// @Title: 数据写入
// @Description: 数据库插入方法，无返回值
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-28 14:37:05
// @param m
// @param condition
// @return code
// @return message
// @return err
func DBInsert(m *gdb.Model, condition DBInsertInput) (code int32, message string, err error) {
	db := DB{M: m}
	return db.Insert(condition)
}

// DBInsertAndGetId
//
// @Title: 数据写入
// @Description: 数据库插入方法，返回写入ID
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-28 15:03:44
// @param m
// @param condition
// @return code
// @return message
// @return lastInsertId
// @return err
func DBInsertAndGetId(m *gdb.Model, condition DBInsertInput) (code int32, message string, lastInsertId int64, err error) {
	db := DB{M: m}
	return db.InsertAndGetId(condition)
}

// DBGetStructById
//
// @Title: 按ID（主键）查询数据
// @Description: 按ID（主键）查询数据，返回类型struct
// @Author liuxingyu <yuwen002@163.com>
// @TestData 2023-01-28 23:38:37
// @param m
// @param condition
// @param output
// @return code
// @return message
// @return err
func DBGetStructById(m *gdb.Model, condition DBGetByIdInput, output interface{}) (code int32, message string, err error) {
	db := DB{M: m}
	return db.GetStructById(condition, output)
}

// DBGetStructGCacheById
//
// @Title: 按ID（主键）查询数据
// @Description: 按ID（主键）查询数据，返回类型struct
// @Author liuxingyu <yuwen002@163.com>
// @TestData 2023-01-29 00:08:28
// @param m
// @param condition
// @param output
// @return code
// @return message
// @return err
func DBGetStructGCacheById(m *gdb.Model, condition DBGetByIdInput, output interface{}) (code int32, message string, err error) {
	db := DB{M: m, RedisConfig: condition.RedisConfig}
	return db.GetStructGCacheById(condition, output)
}

// DBGetStructRCacheById
//
// @Title: 按ID获取数据，带缓存
// @Description: 按ID获取数据，Redis缓存，返回struct
// @Author liuxingyu
// @TestData 2023-01-29 00:08:54
// @param m
// @param condition
// @param output
// @return code
// @return message
// @return err
func DBGetStructRCacheById(m *gdb.Model, condition DBGetByIdInput, output interface{}) (code int32, message string, err error) {
	db := DB{M: m, RedisConfig: condition.RedisConfig}
	return db.GetStructRCacheById(condition, output)
}

// DBGetMapById
//
// @Title 按ID获取数据
// @Description  按ID获取数据，返回map
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-29 14:32:34
// @param m
// @param condition
// @return code
// @return message
// @return output
// @return err
func DBGetMapById(m *gdb.Model, condition DBGetByIdInput) (code int32, message string, output interface{}, err error) {
	db := DB{M: m}
	return db.GetMapById(condition)
}

// DBGetMapGCacheById
//
// @Title 按ID获取数据
// @Description 按ID获取数据，gcache缓存，返回map
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-29 14:33:49
// @param m
// @param condition
// @return code
// @return message
// @return output
// @return err
func DBGetMapGCacheById(m *gdb.Model, condition DBGetByIdInput) (code int32, message string, output interface{}, err error) {
	db := DB{M: m, RedisConfig: condition.RedisConfig}
	return db.GetMapGCacheById(condition)
}

// DBGetMapRCacheById
//
// @Title 按ID获取数据
// @Description 按ID获取数据，Redis缓存，返回map
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-29 14:43:23
// @param m
// @param condition
// @return code
// @return message
// @return output
// @return err
func DBGetMapRCacheById(m *gdb.Model, condition DBGetByIdInput) (code int32, message string, output interface{}, err error) {
	db := DB{M: m, RedisConfig: condition.RedisConfig}
	return db.GetMapRCacheById(condition)
}

// DBGetStructByIds
//
// @Title 按ID（主键）查询多条数据
// @Description 按ID（主键）查询多条数据,返回struct数组
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-01 17:27:07
// @param m
// @param condition
// @param output
// @return code
// @return message
// @return err
func DBGetStructByIds(m *gdb.Model, condition DBGetByIdsInput, output interface{}) (code int32, message string, err error) {
	db := DB{M: m}
	return db.GetStructByIds(condition, output)
}

// DBGetMapByIds
//
// @Title 按ID（主键）查询多条数据
// @Description 按ID（主键）查询多条数据,返回map数组
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-01 17:27:10
// @param m
// @param condition
// @return code
// @return message
// @return output
// @return err
func DBGetMapByIds(m *gdb.Model, condition DBGetByIdsInput) (code int32, message string, output interface{}, err error) {
	db := DB{M: m}
	return db.GetMapByIds(condition)
}

// DBGetOneStructByWhere
//
// @Title 按条件获取数据
// @Description 按条件获取单条数据，返回struct
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-31 10:15:03
// @param m
// @param condition
// @param output
// @return code
// @return message
// @return err
func DBGetOneStructByWhere(m *gdb.Model, condition DBGetOneByWhereInput, output interface{}) (code int32, message string, err error) {
	db := DB{M: m}
	return db.GetOneStructByWhere(condition, output)
}

// DBGetOneMapByWhere
//
// @Title 按条件获取数据
// @Description 按条件获取单条数据，返回map
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-31 10:18:15
// @param m
// @param condition
// @return code
// @return message
// @return output
// @return err
func DBGetOneMapByWhere(m *gdb.Model, condition DBGetOneByWhereInput) (code int32, message string, output interface{}, err error) {
	db := DB{M: m}
	return db.GetOneMapByWhere(condition)
}

// DBGetAllStructByWhere
//
// @Title 按条件获取列表
// @Description 按条件获取列表，返回struct
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-31 10:36:57
// @param m
// @param condition
// @param output
// @return code
// @return message
// @return err
func DBGetAllStructByWhere(m *gdb.Model, condition DBGetAllByWhereInput, output []interface{}) (code int32, message string, err error) {
	db := DB{M: m}
	return db.GetAllStructByWhere(condition, output)
}

// DBGetAllMapByWhere
//
// @Title 按条件获取列表
// @Description 按条件获取列表，返回list<[]Map>
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-31 10:38:39
// @param m
// @param condition
// @return code
// @return message
// @return output
// @return err
func DBGetAllMapByWhere(m *gdb.Model, condition DBGetAllByWhereInput) (code int32, message string, output gdb.List, err error) {
	db := DB{M: m}
	return db.GetAllMapByWhere(condition)
}

// DBModifyById
//
// @Title 按ID修改数据
// @Description  按ID修改数据
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-31 14:44:09
// @param m
// @param condition
// @return code
// @return message
// @return err
func DBModifyById(m *gdb.Model, condition DBModifyByIdInput) (code int32, message string, err error) {
	db := DB{M: m}
	return db.ModifyById(condition)
}

// DBModifyGCacheById
//
// @Title 按ID修改数据
// @Description 按ID修改数据,清除gcache缓存
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-31 17:23:19
// @param m
// @param condition
// @return code
// @return message
// @return err
func DBModifyGCacheById(m *gdb.Model, condition DBModifyByIdInput) (code int32, message string, err error) {
	db := DB{M: m, RedisConfig: condition.RedisConfig}
	return db.ModifyGCacheById(condition)
}

// DBModifyRCacheById
//
// @Title 按ID修改数据
// @Description 按ID修改数据,清除Redis缓存
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-31 17:23:59
// @param m
// @param condition
// @return code
// @return message
// @return err
func DBModifyRCacheById(m *gdb.Model, condition DBModifyByIdInput) (code int32, message string, err error) {
	db := DB{M: m, RedisConfig: condition.RedisConfig}
	return db.ModifyRCacheById(condition)
}

// DBModifyByWhere
//
// @Title 条件修改多条数据
// @Description 条件修改多条数据
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-01-31 18:04:25
// @param m
// @param condition
// @return code
// @return message
// @return err
func DBModifyByWhere(m *gdb.Model, condition DBModifyByWhereInput) (code int32, message string, err error) {
	db := DB{M: m}
	return db.ModifyByWhere(condition)
}

// DBDelById
//
// @Title 按ID删除数据
// @Description 按ID删除数据 开启cache缓存删除功能
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-02 10:15:07
// @param m
// @param condition
// @return code
// @return message
// @return err
func DBDelById(m *gdb.Model, condition DBDelByIdInput) (code int32, message string, err error) {
	db := DB{M: m, RedisConfig: condition.RedisConfig}
	return db.DelById(condition)
}
