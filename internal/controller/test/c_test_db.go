package test

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"ice_flame/api/test"
	"ice_flame/internal/model"
	"ice_flame/internal/service"
)

var TestDB = cTestDB{}

type cTestDB struct {
}

// TestInsert
//
// @Title 数据写入
// @Description 数据写入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-03 10:35:59
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cTestDB) TestInsert(ctx context.Context, req *test.DBInsertReq) (res *test.DBInsertRes, err error) {
	code, message, err := service.TestDB().TestInsert(ctx, model.TestDBInsertInput{TestData: req.Data})

	var json = g.RequestFromCtx(ctx).Response
	if err != nil {
		json.WriteJsonExit(g.Map{
			"code":    code,
			"message": err.Error(),
		})
	}

	json.WriteJsonExit(g.Map{
		"code":    code,
		"message": message,
	})

	return
}

// TestInsertAndGetId
//
// @Title 数据写入
// @Description 数据写入并返回写入ID
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-03 10:40:33
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cTestDB) TestInsertAndGetId(ctx context.Context, req *test.DBInsertAndGetIdReq) (res *test.DBInsertAndGetIdRes, err error) {
	code, message, id, err := service.TestDB().TestInsertAndGetId(ctx, model.TestDBInsertInput{TestData: req.Data})

	var json = g.RequestFromCtx(ctx).Response
	if err != nil {
		json.WriteJsonExit(g.Map{
			"code":    code,
			"message": err.Error(),
		})
	}

	json.WriteJsonExit(g.Map{
		"code":    code,
		"message": message,
		"data":    g.Map{"id": id},
	})

	return
}

// TestGetStructById
//
// @Title 按ID查询信息
// @Description 按ID查询信息，返回类型struct
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-03 11:54:06
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cTestDB) TestGetStructById(ctx context.Context, req *test.DBGetStructByIdReq) (res *test.DBGetStructByIdRes, err error) {
	var output *model.TestGetByIdOutput
	code, message, err := service.TestDB().TestGetStructById(ctx, model.TestGetByIdInput{Id: req.Id}, &output)
	var json = g.RequestFromCtx(ctx).Response
	if err != nil {
		json.WriteJsonExit(g.Map{
			"code":    code,
			"message": err.Error(),
		})
	}

	json.WriteJsonExit(g.Map{
		"code":    code,
		"message": message,
		"data":    output,
	})

	return
}

// TestGetStructGCacheById
//
// @Title 按ID查询信息
// @Description 按ID查询信息，带gcache缓存，返回类型struct
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-03 12:40:40
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cTestDB) TestGetStructGCacheById(ctx context.Context, req *test.DBGetStructGCacheByIdReq) (res *test.DBGetStructGCacheByIdRes, err error) {
	var output *model.TestGetByIdOutput
	code, message, err := service.TestDB().TestGetStructGCacheById(ctx, model.TestGetByIdInput{Id: req.Id}, &output)
	var json = g.RequestFromCtx(ctx).Response
	if err != nil {
		json.WriteJsonExit(g.Map{
			"code":    code,
			"message": err.Error(),
		})
	}

	json.WriteJsonExit(g.Map{
		"code":    code,
		"message": message,
		"data":    output,
	})

	return
}
