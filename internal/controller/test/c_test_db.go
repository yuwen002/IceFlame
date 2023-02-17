package test

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"ice_flame/api/test"
	"ice_flame/internal/model"
	"ice_flame/internal/service"
	"strings"
)

var DBTest = cDBTest{}

type cDBTest struct {
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
func (c *cDBTest) TestInsert(ctx context.Context, req *test.DBInsertReq) (res *test.DBInsertRes, err error) {
	code, message, err := service.DBTest().TestInsert(ctx, model.TestDBInsertInput{TestData: req.Data})

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
func (c *cDBTest) TestInsertAndGetId(ctx context.Context, req *test.DBInsertAndGetIdReq) (res *test.DBInsertAndGetIdRes, err error) {
	code, message, id, err := service.DBTest().TestInsertAndGetId(ctx, model.TestDBInsertInput{TestData: req.Data})

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
func (c *cDBTest) TestGetStructById(ctx context.Context, req *test.DBGetStructByIdReq) (res *test.DBGetStructByIdRes, err error) {
	var output *model.TestGetByIdOutput
	code, message, err := service.DBTest().TestGetStructById(ctx, model.TestGetByIdInput{Id: req.Id}, &output)
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
func (c *cDBTest) TestGetStructGCacheById(ctx context.Context, req *test.DBGetStructGCacheByIdReq) (res *test.DBGetStructGCacheByIdRes, err error) {
	var output *model.TestGetByIdOutput
	code, message, err := service.DBTest().TestGetStructGCacheById(ctx, model.TestGetByIdInput{Id: req.Id}, &output)
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
func (c *cDBTest) TestGetStructRCacheById(ctx context.Context, req *test.DBGetStructRCacheByIdReq) (res *test.DBGetStructRCacheByIdRes, err error) {
	var output *model.TestGetByIdOutput
	code, message, err := service.DBTest().TestGetStructRCacheById(ctx, model.TestGetByIdInput{Id: req.Id}, &output)
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
func (c *cDBTest) TestGetMapById(ctx context.Context, req *test.DBGetMapByIdReq) (res *test.DBGetMapByIdRes, err error) {
	code, message, output, err := service.DBTest().TestGetMapById(ctx, model.TestGetByIdInput{Id: req.Id})
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
func (c *cDBTest) TestGetMapGCacheById(ctx context.Context, req *test.DBGetMapGCacheByIdReq) (res *test.DBGetMapGCacheByIdRes, err error) {
	code, message, output, err := service.DBTest().TestGetMapGCacheById(ctx, model.TestGetByIdInput{Id: req.Id})
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
func (c *cDBTest) TestGetMapRCacheById(ctx context.Context, req *test.DBGetMapRCacheByIdReq) (res *test.DBGetMapRCacheByIdRes, err error) {
	code, message, output, err := service.DBTest().TestGetMapRCacheById(ctx, model.TestGetByIdInput{Id: req.Id})
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
func (c *cDBTest) TestGetStructByIds(ctx context.Context, req *test.DBGetStructByIdsReq) (res *test.DBGetStructByIdsRes, err error) {
	var output []*model.TestGetByIdOutput
	code, message, err := service.DBTest().TestGetStructByIds(ctx, model.TestGetByIdsInput{Ids: strings.Split(req.Ids, ",")}, &output)
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
func (c *cDBTest) TestGetMapByIds(ctx context.Context, req *test.DBGetMapByIdsReq) (res *test.DBGetMapByIdsRes, err error) {
	code, message, output, err := service.DBTest().TestGetMapByIds(ctx, model.TestGetByIdsInput{Ids: strings.Split(req.Ids, ",")})
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

// TestGetOneStructByWhere
//
// @Title 条件查询单条信息
// @Description 测试条件查询单条信息，返回struct
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-02-05 23:30:39
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cDBTest) TestGetOneStructByWhere(ctx context.Context, req *test.DBGetOneStructByWhereReq) (res *test.DBGetOneStructByWhereRes, err error) {
	var output *model.TestGetOneByWhereOutput
	code, message, err := service.DBTest().TestGetOneStructByWhere(ctx, model.TestGetOneByWhereInput{
		Where: "title=?",
		Args:  req.Title,
	}, &output)

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

// TestGetOneMapByWhere
//
// @Title 条件查询单条信息
// @Description 测试条件查询单条信息，返回map
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-06 12:16:16
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cDBTest) TestGetOneMapByWhere(ctx context.Context, req *test.DBGetOneMapByWhereReq) (res *test.DBGetOneMapByWhereRes, err error) {
	code, message, output, err := service.DBTest().TestGetOneMapByWhere(ctx, model.TestGetOneByWhereInput{
		Where: "title=?",
		Args:  req.Title,
	})

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

// TestDBGetAllStructByWhere
//
// @Title 条件查询多条信息
// @Description 测试条件查询多条信息，返回struct
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-07 11:10:21
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cDBTest) TestDBGetAllStructByWhere(ctx context.Context, req *test.DBGetAllStructByWhereReq) (res *test.DBGetAllStructByWhereRes, err error) {
	var output []*model.TestGetAllByWhereOutput
	code, message, err := service.DBTest().TestDBGetAllStructByWhere(ctx, model.TestGetAllByWhereInput{
		Where: "classify=?",
		Args:  req.Classify,
		Limit: 0,
	}, &output)

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

// TestDBGetAllMapByWhere
//
// @Title 条件查询多条信息
// @Description 测试条件查询多条信息，返回map
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-07 11:20:18
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cDBTest) TestDBGetAllMapByWhere(ctx context.Context, req *test.DBGetAllMapByWhereReq) (res *test.DBGetAllMapByWhereRes, err error) {
	code, message, output, err := service.DBTest().TestDBGetAllMapByWhere(ctx, model.TestGetAllByWhereInput{
		Where: "classify=?",
		Args:  req.Classify,
		Limit: 10,
	})

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

// TestDBModifyById
//
// @Title 按ID修改数据
// @Description 按ID修改数据
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-08 11:08:28
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cDBTest) TestDBModifyById(ctx context.Context, req *test.DBModifyByIdReq) (res *test.DBModifyByIdRes, err error) {
	code, message, err := service.DBTest().TestDBModifyById(ctx, model.TestModifyByIdInput{
		Data: g.Map{
			"test_data": req.TestData,
			"title":     req.Title,
			"classify":  req.Classify,
		},
		Where: req.Id,
	})

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

// TestDBModifyGCacheById
//
// @Title 按ID修改数据
// @Description 按ID修改数据,删除gcache缓存
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-08 11:21:18
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cDBTest) TestDBModifyGCacheById(ctx context.Context, req *test.DBModifyGCacheByIdReq) (res *test.DBModifyGCacheByIdRes, err error) {
	code, message, err := service.DBTest().TestDBModifyGCacheById(ctx, model.TestModifyByIdInput{
		Data: g.Map{
			"test_data": req.TestData,
			"title":     req.Title,
			"classify":  req.Classify,
		},
		Where: req.Id,
	})

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

// TestDBModifyRCacheById
//
// @Title 按ID修改数据
// @Description 按ID修改数据,删除Redis缓存
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-08 13:36:09
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cDBTest) TestDBModifyRCacheById(ctx context.Context, req *test.DBModifyRCacheByIdReq) (res *test.DBModifyRCacheByIdRes, err error) {
	code, message, err := service.DBTest().TestDBModifyRCacheById(ctx, model.TestModifyByIdInput{
		Data: g.Map{
			"test_data": req.TestData,
			"title":     req.Title,
			"classify":  req.Classify,
		},
		Where: req.Id,
		Args:  "",
	})

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

// TestDBModifyByWhere
//
// @Title 按条件修改数据
// @Description 按条件修改数据,删除Redis缓存
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-08 14:28:10
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cDBTest) TestDBModifyByWhere(ctx context.Context, req *test.DBModifyByWhereReq) (res *test.DBModifyByWhereRes, err error) {
	code, message, err := service.DBTest().TestDBModifyByWhere(ctx, model.TestModifyByWhereInput{
		Data: g.Map{
			"test_data": req.TestData,
			"title":     req.Title,
			"classify":  req.Classify,
		},
		Where: "id=?",
		Args:  req.Id,
		Limit: 1,
	})

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

// TestDBDelById
//
// @Title 按ID删除信息
// @Description 按ID删除信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-08 15:07:04
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cDBTest) TestDBDelById(ctx context.Context, req *test.DBDelByIdReq) (res *test.DBDelByIdRes, err error) {
	code, message, err := service.DBTest().TestDBDelById(ctx, model.TestDelByIdInput{Id: req.Id})

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

// TestDBDelByIds
//
// @Title 按ID删除多条信息
// @Description 按ID删除多条信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-08 16:47:51
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cDBTest) TestDBDelByIds(ctx context.Context, req *test.DBDelByIdsReq) (res *test.DBDelByIdsRes, err error) {
	code, message, err := service.DBTest().TestDBDelByIds(ctx, model.TestDelByIdsInput{Ids: req.Ids})

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

// TestDBDelByWhere
//
// @Title 按条件删除信息
// @Description 按条件删除信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-08 16:53:39
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cDBTest) TestDBDelByWhere(ctx context.Context, req *test.DBDelByWhereReq) (res *test.DBDelByWhereRes, err error) {
	code, message, err := service.DBTest().TestDBDelByWhere(ctx, model.TestDelByWhereInput{
		Where: "classify=?",
		Args:  req.Classify,
	})

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

// TestDBModifyIncById
//
// @Title 按ID自增
// @Description 按ID自增
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-11 15:47:10
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cDBTest) TestDBModifyIncById(ctx context.Context, req *test.DBModifyIncByIdReq) (res *test.DBModifyIncByIdRes, err error) {
	code, message, err := service.DBTest().TestDBModifyIncById(ctx, model.TestModifyIncByIdInput{
		Id:     req.Id,
		Column: "num",
		Amount: 1,
	})

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

// TestDBModifyDecById
//
// @Title 按ID自减
// @Description 按ID自减
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-11 16:18:09
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cDBTest) TestDBModifyDecById(ctx context.Context, req *test.DBModifyDecByIdReq) (res *test.DBModifyDecByIdRes, err error) {
	code, message, err := service.DBTest().TestDBModifyDecById(ctx, model.TestModifyDecByIdInput{
		Id:     req.Id,
		Column: "num",
		Amount: 1,
	})

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

// TestDBModifyIncByWhere
//
// @Title 按条件自增
// @Description 按条件自增
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-11 17:27:23
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cDBTest) TestDBModifyIncByWhere(ctx context.Context, req *test.DBModifyIncByWhereReq) (res *test.DBModifyIncByWhereRes, err error) {
	code, message, err := service.DBTest().TestDBModifyIncByWhere(ctx, model.TestModifyIncByWhereInput{
		Where:  "classify=?",
		Args:   req.Classify,
		Column: "num",
		Amount: 1,
	})

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

// TestDBModifyDecByWhere
//
// @Title 按条件自减
// @Description按条件自减
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-11 17:27:40
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cDBTest) TestDBModifyDecByWhere(ctx context.Context, req *test.DBModifyDecByWhereReq) (res *test.DBModifyDecByWhereRes, err error) {
	code, message, err := service.DBTest().TestDBModifyDecByWhere(ctx, model.TestModifyDecByWhereInput{
		Where:  "classify=?",
		Args:   req.Classify,
		Column: "num",
		Amount: 1,
	})

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
