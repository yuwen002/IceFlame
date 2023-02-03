package test

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"ice_flame/internal/dao"
	"ice_flame/internal/model"
	"ice_flame/internal/service"
	"ice_flame/utility"
)

var insTestDB = sTestDB{}

func NewTestDB() *sTestDB {
	return &insTestDB
}

func init() {
	service.RegisterTestDB(NewTestDB())
}

// sTestDB
// @Description: 数据库测试
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-02 15:29:14
type sTestDB struct {
}

// TestInsert
//
// @Title 测试数据写入
// @Description 测试数据写入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-02 15:29:23
// @receiver s
// @param ctx
// @param data
// @return code
// @return message
// @return err
func (s *sTestDB) TestInsert(ctx context.Context, in model.TestDBInsertInput) (code int32, message string, err error) {
	return utility.DBInsert(dao.DbTest.Ctx(ctx), utility.DBInsertInput{Data: in})
}

// TestInsertAndGetId
//
// @Title 测试数据写入并返回ID
// @Description 测试数据写入并返回ID
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-02 16:43:37
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return lastInsertId
// @return err
func (s *sTestDB) TestInsertAndGetId(ctx context.Context, in model.TestDBInsertInput) (code int32, message string, lastInsertId int64, err error) {
	return utility.DBInsertAndGetId(dao.DbTest.Ctx(ctx), utility.DBInsertInput{Data: in})
}

// TestGetStructById
//
// @Title 按ID查询数据
// @Description 按ID查询数据，返回struct
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-02 16:54:22
// @receiver s
// @param ctx
// @param in
// @param out
// @return code
// @return message
// @return output
// @return err
func (s *sTestDB) TestGetStructById(ctx context.Context, in model.TestGetByIdInput, output interface{}) (code int32, message string, err error) {
	code, message, err = utility.DBGetStructById(dao.DbTest.Ctx(ctx), utility.DBGetByIdInput{Where: in.Id}, output)
	return code, message, err
}

// TestGetStructGCacheById
//
// @Title 按ID查询数据
// @Description 按ID查询数据，gcache缓存,返回struct
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-03 14:01:49
// @receiver s
// @param ctx
// @param in
// @param out
// @return code
// @return message
// @return err
func (s *sTestDB) TestGetStructGCacheById(ctx context.Context, in model.TestGetByIdInput, output interface{}) (code int32, message string, err error) {
	return utility.DBGetStructGCacheById(dao.DbTest.Ctx(ctx), utility.DBGetByIdInput{
		Where:       in.Id,
		CacheKey:    dao.DbTest.Table() + ":id-" + gconv.String(in.Id),
		RedisConfig: "test",
		Expire:      300,
	}, output)
}
