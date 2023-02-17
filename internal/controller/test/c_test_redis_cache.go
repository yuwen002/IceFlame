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
