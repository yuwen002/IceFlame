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
