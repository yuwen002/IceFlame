package test

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"ice_flame/internal/dao"
	"ice_flame/internal/model"
	"ice_flame/internal/service"
	"ice_flame/utility"
)

var insDBTest = sDBTest{}

func NewDBTest() *sDBTest {
	return &insDBTest
}

func init() {
	service.RegisterDBTest(NewDBTest())
}

// sDBTest
// @Description: 数据库测试
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-02 15:29:14
type sDBTest struct {
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
func (s *sDBTest) TestInsert(ctx context.Context, in model.TestDBInsertInput) (code int32, message string, err error) {
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
func (s *sDBTest) TestInsertAndGetId(ctx context.Context, in model.TestDBInsertInput) (code int32, message string, lastInsertId int64, err error) {
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
func (s *sDBTest) TestGetStructById(ctx context.Context, in model.TestGetByIdInput, output interface{}) (code int32, message string, err error) {
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
func (s *sDBTest) TestGetStructGCacheById(ctx context.Context, in model.TestGetByIdInput, output interface{}) (code int32, message string, err error) {
	return utility.DBGetStructGCacheById(dao.DbTest.Ctx(ctx), utility.DBGetByIdInput{
		Where:       in.Id,
		CacheKey:    dao.DbTest.Table() + ":id-" + gconv.String(in.Id),
		RedisConfig: "test",
		Expire:      300,
	}, output)
}

// TestGetStructRCacheById
//
// @Title 按ID查询数据
// @Description  按ID查询数据，redis缓存,返回struct
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-03 16:55:09
// @receiver s
// @param ctx
// @param in
// @param output
func (s *sDBTest) TestGetStructRCacheById(ctx context.Context, in model.TestGetByIdInput, output interface{}) (code int32, message string, err error) {
	return utility.DBGetStructRCacheById(dao.DbTest.Ctx(ctx), utility.DBGetByIdInput{
		Where:       in.Id,
		CacheKey:    dao.DbTest.Table() + ":id-" + gconv.String(in.Id),
		RedisConfig: "test",
		Expire:      300,
	}, output)
}

// TestGetMapById
//
// @Title 按ID查询数据
// @Description  按ID查询数据，返回map
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-03 17:07:38
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return output
// @return err
func (s *sDBTest) TestGetMapById(ctx context.Context, in model.TestGetByIdInput) (code int32, message string, output interface{}, err error) {
	return utility.DBGetMapById(dao.DbTest.Ctx(ctx), utility.DBGetByIdInput{Where: in.Id})
}

// TestGetMapGCacheById
//
// @Title 按ID查询数据
// @Description  按ID查询数据，gcache缓存，返回map
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-03 17:22:00
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return output
// @return err
func (s *sDBTest) TestGetMapGCacheById(ctx context.Context, in model.TestGetByIdInput) (code int32, message string, output interface{}, err error) {
	return utility.DBGetMapGCacheById(dao.DbTest.Ctx(ctx), utility.DBGetByIdInput{
		Where:       "id=?",
		Args:        in.Id,
		CacheKey:    dao.DbTest.Table() + ":id-" + gconv.String(in.Id),
		RedisConfig: "test",
		Expire:      300,
	})
}

// TestGetMapRCacheById
//
// @Title 按ID查询数据
// @Description  按ID查询数据，redis缓存，返回map
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-03 17:38:49
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return output
// @return err
func (s *sDBTest) TestGetMapRCacheById(ctx context.Context, in model.TestGetByIdInput) (code int32, message string, output interface{}, err error) {
	return utility.DBGetMapRCacheById(dao.DbTest.Ctx(ctx), utility.DBGetByIdInput{
		Where:       "id=?",
		Args:        in.Id,
		CacheKey:    dao.DbTest.Table() + ":id-" + gconv.String(in.Id),
		RedisConfig: "test",
		Expire:      300,
	})
}

// TestGetStructByIds
//
// @Title 按ID查询多条数据
// @Description 按ID查询多条数据，返回struct
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-03 18:23:02
// @receiver s
// @param ctx
// @param in
// @param output
// @return code
// @return message
// @return err
func (s *sDBTest) TestGetStructByIds(ctx context.Context, in model.TestGetByIdsInput, output interface{}) (code int32, message string, err error) {
	return utility.DBGetStructByIds(dao.DbTest.Ctx(ctx), utility.DBGetByIdsInput{
		Column: "id",
		In:     in.Ids,
		Order:  "id desc",
	}, output)
}

// TestGetMapByIds
//
// @Title 按ID查询多条数据
// @Description 按ID查询多条数据，返回map
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-03 18:53:35
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return output
// @return err
func (s *sDBTest) TestGetMapByIds(ctx context.Context, in model.TestGetByIdsInput) (code int32, message string, output interface{}, err error) {
	return utility.DBGetMapByIds(dao.DbTest.Ctx(ctx), utility.DBGetByIdsInput{
		Column: "id",
		In:     in.Ids,
		Order:  "id desc",
	})
}

// TestGetOneStructByWhere
//
// @Title 条件查询单条信息
// @Description 条件查询单条信息，返回struct
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-02-05 23:30:06
// @receiver s
// @param ctx
// @param in
// @param output
// @return code
// @return message
// @return err
func (s *sDBTest) TestGetOneStructByWhere(ctx context.Context, in model.TestGetOneByWhereInput, output interface{}) (code int32, message string, err error) {
	return utility.DBGetOneStructByWhere(dao.DbTest.Ctx(ctx), utility.DBGetOneByWhereInput{
		Where: in.Where,
		Args:  in.Args,
		Order: "id desc",
	}, output)
}

// TestGetOneMapByWhere
//
// @Title 条件查询单条信息
// @Description 条件查询单条信息，返回map
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-06 15:08:16
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return output
// @return err
func (s *sDBTest) TestGetOneMapByWhere(ctx context.Context, in model.TestGetOneByWhereInput) (code int32, message string, output interface{}, err error) {
	return utility.DBGetOneMapByWhere(dao.DbTest.Ctx(ctx), utility.DBGetOneByWhereInput{
		Where: in.Where,
		Args:  in.Args,
	})
}

// TestDBGetAllStructByWhere
//
// @Title 条件查询多条信息
// @Description 条件查询多条信息，返回struct
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-06 17:14:25
// @receiver s
// @param ctx
// @param in
// @param output
// @return code
// @return message
// @return err
func (s *sDBTest) TestDBGetAllStructByWhere(ctx context.Context, in model.TestGetAllByWhereInput, output interface{}) (code int32, message string, err error) {
	var condition utility.DBGetAllByWhereInput
	condition.Where = in.Where
	condition.Args = in.Args
	condition.PageType = 1
	condition.Length = in.Limit
	return utility.DBGetAllStructByWhere(dao.DbTest.Ctx(ctx), condition, output)
}

// TestDBGetAllMapByWhere
//
// @Title 条件查询多条信息
// @Description 条件查询多条信息，返回map
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-07 11:14:21
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return output
// @return err
func (s *sDBTest) TestDBGetAllMapByWhere(ctx context.Context, in model.TestGetAllByWhereInput) (code int32, message string, output interface{}, err error) {
	return utility.DBGetAllMapByWhere(dao.DbTest.Ctx(ctx), utility.DBGetAllByWhereInput{
		Where:    in.Where,
		Args:     in.Args,
		Order:    "id desc",
		PageType: 1,
		DBLimit:  utility.DBLimit{Offset: in.Limit},
	})
}

// TestDBModifyById
//
// @Title 按ID修改信息
// @Description 按ID修改信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-08 10:14:06
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sDBTest) TestDBModifyById(ctx context.Context, in model.TestModifyByIdInput) (code int32, message string, err error) {
	return utility.DBModifyById(dao.DbTest.Ctx(ctx), utility.DBModifyByIdInput{
		Data:  in.Data,
		Where: in.Where,
		Args:  in.Args,
	})
}

// TestDBModifyGCacheById
//
// @Title 按ID修改信息
// @Description 按ID修改信息,删除gcache缓存
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-08 11:17:45
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sDBTest) TestDBModifyGCacheById(ctx context.Context, in model.TestModifyByIdInput) (code int32, message string, err error) {
	return utility.DBModifyGCacheById(dao.DbTest.Ctx(ctx), utility.DBModifyByIdInput{
		Data:        in.Data,
		Where:       in.Where,
		Args:        in.Args,
		CacheKey:    dao.DbTest.Table() + ":id-" + in.Where,
		RedisConfig: "test",
	})
}

// TestDBModifyRCacheById
//
// @Title 按ID修改信息
// @Description 按ID修改信息,删除redis缓存
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-08 11:44:24
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sDBTest) TestDBModifyRCacheById(ctx context.Context, in model.TestModifyByIdInput) (code int32, message string, err error) {
	return utility.DBModifyRCacheById(dao.DbTest.Ctx(ctx), utility.DBModifyByIdInput{
		Data:        in.Data,
		Where:       in.Where,
		Args:        in.Args,
		CacheKey:    dao.DbTest.Table() + ":id-" + in.Where,
		RedisConfig: "test",
	})
}

// TestDBModifyByWhere
//
// @Title 按条件修改信息
// @Description 按条件修改信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-08 14:21:10
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sDBTest) TestDBModifyByWhere(ctx context.Context, in model.TestModifyByWhereInput) (code int32, message string, err error) {
	return utility.DBModifyByWhere(dao.DbTest.Ctx(ctx), utility.DBModifyByWhereInput{
		Data:     in.Data,
		Where:    in.Where,
		Args:     in.Args,
		PageType: 1,
		DBLimit:  utility.DBLimit{Offset: in.Limit},
	})
}

// TestDBDelById
//
// @Title 按ID删除数据
// @Description 按ID删除数据
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-08 14:39:24
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sDBTest) TestDBDelById(ctx context.Context, in model.TestDelByIdInput) (code int32, message string, err error) {
	return utility.DBDelById(dao.DbTest.Ctx(ctx), utility.DBDelByIdInput{Where: in.Id})
}

// TestDBDelByIds
//
// @Title 按ID删除多条数据
// @Description 按ID删除多条数据
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-08 16:45:10
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sDBTest) TestDBDelByIds(ctx context.Context, in model.TestDelByIdsInput) (code int32, message string, err error) {
	return utility.DBDelByIds(dao.DbTest.Ctx(ctx), utility.DBDelByIdsInput{
		Column: "id",
		In:     in.Ids,
	})
}

// TestDBDelByWhere
//
// @Title 按条件删除数据
// @Description 按条件删除数据
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-08 16:04:40
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sDBTest) TestDBDelByWhere(ctx context.Context, in model.TestDelByWhereInput) (code int32, message string, err error) {
	return utility.DBDelByWhere(dao.DbTest.Ctx(ctx), utility.DBDelByWhereInput{
		Where:    in.Where,
		Args:     in.Args,
		PageType: 1,
		DBLimit:  utility.DBLimit{Offset: in.Limit},
	})
}

// TestDBModifyIncById
//
// @Title 按ID自增
// @Description 按ID自增
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-11 14:53:06
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sDBTest) TestDBModifyIncById(ctx context.Context, in model.TestModifyIncByIdInput) (code int32, message string, err error) {
	return utility.DBModifyIncById(dao.DbTest.Ctx(ctx), utility.DBModifyIncByIdInput{
		Field:  in.Column,
		Where:  in.Id,
		Amount: in.Amount,
	})
}

// TestDBModifyDecById
//
// @Title 按ID自减
// @Description 按ID自减
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-11 16:14:48
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sDBTest) TestDBModifyDecById(ctx context.Context, in model.TestModifyDecByIdInput) (code int32, message string, err error) {
	return utility.DBModifyDecById(dao.DbTest.Ctx(ctx), utility.DBModifyDecByIdInput{
		Field:  in.Column,
		Where:  in.Id,
		Amount: in.Amount,
	})
}

// TestDBModifyIncByWhere
//
// @Title
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-11 16:38:56
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sDBTest) TestDBModifyIncByWhere(ctx context.Context, in model.TestModifyIncByWhereInput) (code int32, message string, err error) {
	return utility.DBModifyIncByWhere(dao.DbTest.Ctx(ctx), utility.DBModifyIncByWhereInput{
		Field:    in.Column,
		Where:    in.Where,
		Args:     in.Args,
		Amount:   in.Amount,
		PageType: 1,
		DBLimit:  utility.DBLimit{Offset: 2},
	})
}

// TestDBModifyDecByWhere
//
// @Title
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-11 16:38:59
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sDBTest) TestDBModifyDecByWhere(ctx context.Context, in model.TestModifyDecByWhereInput) (code int32, message string, err error) {
	return utility.DBModifyDecByWhere(dao.DbTest.Ctx(ctx), utility.DBModifyDecByWhereInput{
		Field:    in.Column,
		Where:    in.Where,
		Args:     in.Args,
		Amount:   in.Amount,
		PageType: 1,
		DBLimit:  utility.DBLimit{Offset: 2},
	})
}
