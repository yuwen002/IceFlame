package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"ice_flame/internal/consts"
	"ice_flame/utility"
)

var insAuthMiddleware = sAuthMiddleware{}

func AuthMiddleware() *sAuthMiddleware {
	return &insAuthMiddleware
}

type sAuthMiddleware struct{}

// MiddlewareAuthMaster
//
// @Title 管理员登入验证中间件
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-22 14:16:59
// @receiver s
// @param r
func (s *sAuthMiddleware) MiddlewareAuthMaster(r *ghttp.Request) {
	token := r.Get("token")
	claims, err := utility.ParseToken(token.String(), consts.MasterSecretKey)
	if err != nil {
		r.Response.WriteJsonExit(g.Map{
			"code":    -1,
			"message": err.Error(),
		})

		r.SetCtxVar("master_id", claims.Id)
		r.SetCtxVar("user_info", claims.UserInfo)
		r.Middleware.Next()
	}
}
