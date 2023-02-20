package test

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"ice_flame/api/test"
	"ice_flame/internal/model"
	"ice_flame/internal/service"
)

var RedisCacheTest = cRedisCacheTest{}

type cRedisCacheTest struct {
}

// TestExistsSetData
//
// @Title 测试判断集合重元素是否存在
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-20 14:20:40
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cRedisCacheTest) TestExistsSetData(ctx context.Context, req *test.RCExistsSetDataReq) (res *test.RCExistsSetDataRes, err error) {
	code, message, err := service.RedisCacheTest().TestExistsSetData(ctx, model.TestExistsSetDataInput{Id: req.Id})

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

// TestCastHashData
//
// @Title 测试数据转换
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-20 14:21:13
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cRedisCacheTest) TestCastHashData(ctx context.Context, req *test.RCCastHashDataReq) (res *test.RCCastHashDataRes, err error) {
	code, message, output, err := service.RedisCacheTest().TestCastHashData(ctx, model.TestCastHashDataInput{Id: req.Id})
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
		"data":    g.Map{"union_id": output},
	})

	return
}
