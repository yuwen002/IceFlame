package test

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"ice_flame/api/test"
	"ice_flame/internal/model"
	"ice_flame/internal/service"
)

var JwtClaimsTest = cJwtClaimsTest{}

type cJwtClaimsTest struct{}

// TestJCSetToken
//
// @Title 测试生成token
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-22 10:54:04
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cJwtClaimsTest) TestJCSetToken(ctx context.Context, req *test.JCSetTokenReq) (res *test.JCSetTokenRes, err error) {
	code, message, token, err := service.JwtTokenTest().TestSetToken(ctx, model.JwtClaimsInput{
		Id: 1,
		UserInfo: g.Map{
			"username": "yuwen002",
			"nickname": "冰之焰",
		},
		Expire:  3600,
		Before:  0,
		Issuer:  "UserIssue",
		Subject: "用户",
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
		"token":   token,
	})

	return
}

// TestJCGetToken
//
// @Title 测试解析token
// @Description 测试解析token
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-22 10:59:50
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cJwtClaimsTest) TestJCGetToken(ctx context.Context, req *test.JCGetTokenReq) (res *test.JCSetTokenRes, err error) {
	code, message, claims, err := service.JwtTokenTest().TestGetToken(ctx, req.Token)
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
		"claims":  claims,
	})

	return
}
