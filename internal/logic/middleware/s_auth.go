package middleware

import (
	"ice_flame/internal/consts"
	"ice_flame/internal/service"
	"ice_flame/utility"
	"regexp"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var insAuthMiddleware = sAuthMiddleware{}

func AuthMiddleware() *sAuthMiddleware {
	return &insAuthMiddleware
}

func init() {
	service.RegisterAuthMiddleware(AuthMiddleware())
}

type sAuthMiddleware struct{}

// extractHandler
//
// @Title 解析Handler
// @Description 传入"controller/manage.(*cUcSystemMasterAuth).ListMenu-fm" 解析package, struct, method
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-12 02:00:20
// @receiver s
// @param str
// @return packageName
// @return structName
// @return methodName
func (s *sAuthMiddleware) extractHandler(str string) (packageName, structName, methodName string) {
	re := regexp.MustCompile(`controller/(.+)\.\(\*(.+)\)\.(.+)-fm`)
	match := re.FindStringSubmatch(str)
	if len(match) == 4 {
		return match[1], match[2], match[3]
	}
	return "", "", ""
}

// concatModule
//
// @Title 拼接访问模型
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-12 02:16:45
// @receiver s
// @param args
// @return string
func (s *sAuthMiddleware) concatModule(args ...string) string {
	if len(args) <= 0 {
		return ""
	}

	return strings.Join(args, ".")
}

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
	// token 未传入
	if token.String() == "" {
		r.Response.WriteJsonExit(g.Map{
			"code":    -1,
			"message": "token不能为空",
		})
	}
	claims, err := utility.ParseToken(token.String(), consts.MasterSecretKey)
	if err != nil {
		r.Response.WriteJsonExit(g.Map{
			"code":    -1,
			"message": err.Error(),
		})
	}

	claims.UserInfo["os_category"] = utility.GetOsCategory(r)

	r.SetCtxVar("master_id", claims.Id)
	r.SetCtxVar("user_info", claims.UserInfo)
	r.Middleware.Next()
}

func (s *sAuthMiddleware) MiddlewareVerifyPermission(r *ghttp.Request) {
	module := s.concatModule(s.extractHandler(r.GetServeHandler().Handler.Name))
}
