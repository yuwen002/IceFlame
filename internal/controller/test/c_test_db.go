package test

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"ice_flame/api/test"
	"ice_flame/internal/model"
	"ice_flame/internal/service"
	"strings"
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

// TestGetStructRCacheById
//
// @Title 按ID查询信息
// @Description 按ID查询信息，带redis缓存，返回类型struct
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-03 17:08:15
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c cTestDB) TestGetStructRCacheById(ctx context.Context, req *test.DBGetStructRCacheByIdReq) (res *test.DBGetStructRCacheByIdRes, err error) {
	var output *model.TestGetByIdOutput
	code, message, err := service.TestDB().TestGetStructRCacheById(ctx, model.TestGetByIdInput{Id: req.Id}, &output)
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

// TestGetMapById
//
// @Title 按ID查询信息
// @Description 按ID查询信息，返回类型map
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-03 17:14:45
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cTestDB) TestGetMapById(ctx context.Context, req *test.DBGetMapByIdReq) (res *test.DBGetMapByIdRes, err error) {
	code, message, output, err := service.TestDB().TestGetMapById(ctx, model.TestGetByIdInput{Id: req.Id})
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

// TestGetMapGCacheById
//
// @Title 按ID查询信息
// @Description 按ID查询信息，带gcache缓存，返回类型map
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-03 17:28:33
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cTestDB) TestGetMapGCacheById(ctx context.Context, req *test.DBGetMapGCacheByIdReq) (res *test.DBGetMapGCacheByIdRes, err error) {
	code, message, output, err := service.TestDB().TestGetMapGCacheById(ctx, model.TestGetByIdInput{Id: req.Id})
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

// TestGetMapRCacheById
//
// @Title 按ID查询信息
// @Description 按ID查询信息，带redis缓存，返回类型map
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-03 17:40:58
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cTestDB) TestGetMapRCacheById(ctx context.Context, req *test.DBGetMapRCacheByIdReq) (res *test.DBGetMapRCacheByIdRes, err error) {
	code, message, output, err := service.TestDB().TestGetMapRCacheById(ctx, model.TestGetByIdInput{Id: req.Id})
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

// TestGetStructByIds
//
// @Title 按ID查询多条信息
// @Description 按ID查询信息，返回类型struct数组
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-03 18:38:39
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cTestDB) TestGetStructByIds(ctx context.Context, req *test.DBGetStructByIdsReq) (res *test.DBGetStructByIdsRes, err error) {
	var output []*model.TestGetByIdOutput
	code, message, err := service.TestDB().TestGetStructByIds(ctx, model.TestGetByIdsInput{Ids: strings.Split(req.Ids, ",")}, &output)
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

// TestGetMapByIds
//
// @Title 按ID查询多条信息
// @Description 按ID查询信息，返回类型map数组
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-03 18:54:37
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cTestDB) TestGetMapByIds(ctx context.Context, req *test.DBGetMapByIdsReq) (res *test.DBGetMapByIdsRes, err error) {
	code, message, output, err := service.TestDB().TestGetMapByIds(ctx, model.TestGetByIdsInput{Ids: strings.Split(req.Ids, ",")})
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
