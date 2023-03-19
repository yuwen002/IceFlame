package middleware

import (
	"context"
	"ice_flame/internal/consts"
	"ice_flame/internal/service"
	"ice_flame/utility"
	"regexp"
	"strings"

	"github.com/gogf/gf/v2/util/gconv"

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

type sAuthMiddleware struct {
	ctx context.Context
}

// extractHandler
//
// @Title 解析Handler
// @Description 传入"controller/manage.(*cUcSystemMasterAuth).ListPermission-fm" 解析package, struct, method
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

	r.SetCtxVar("master_id", claims.Id)
	r.SetCtxVar("user_info", claims.UserInfo)
	r.Middleware.Next()
}

// MiddlewareVerifyPermission
//
// @Title 管理员权限验证中间件
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-17 17:59:59
// @receiver s
// @param r
func (s *sAuthMiddleware) MiddlewareVerifyPermission(r *ghttp.Request) {
	userInfo := r.GetCtxVar("user_info")
	if gconv.Uint8(userInfo.Map()["supper_master"]) == 1 {
		r.Middleware.Next()
	}

	// 访问模块
	module := s.concatModule(s.extractHandler(r.GetServeHandler().Handler.Name))
	// 排除权限列表
	code, message, exclude, err := service.UcSystemMasterAuth().GetPermissionExcludeAll(s.ctx)
	if err != nil {
		r.Response.WriteJsonExit(g.Map{
			"code":    -1,
			"message": err.Error(),
		})
	}

	if code == 0 {
		// 查找排除模块
		existsPermissionExclude := false
		for index := range exclude {
			if exclude[index].Module == module {
				existsPermissionExclude = true
				break
			}
		}

		// 排除权限判断，直接跳转
		if existsPermissionExclude {
			r.Middleware.Next()
		}
	}

	// 查询管理员所分配的角色关联的权限
	code, message, rolePermission, err := service.UcSystemMasterAuth().GetRolePermissionByAccountId(s.ctx, gconv.Uint64(r.GetCtxVar("master_id")))
	if err != nil {
		r.Response.WriteJsonExit(g.Map{
			"code":    -1,
			"message": err.Error(),
		})
	}

	if code != 0 {
		r.Response.WriteJsonExit(g.Map{
			"code":    -1,
			"message": message,
		})
	}

	// 查询所有可以使用的权限
	code, message, permission, err := service.UcSystemMasterAuth().GetPermissionAll(s.ctx)
	if err != nil {
		r.Response.WriteJsonExit(g.Map{
			"code":    -1,
			"message": err.Error(),
		})
	}
	if code != 0 {
		r.Response.WriteJsonExit(g.Map{
			"code":    -1,
			"message": message,
		})
	}

	// 获取访问模块信息，controller调用的方法
	var permissionId uint32
	for index := range permission {
		if permission[index].Module == module {
			// 匹配权限ID
			permissionId = permission[index].Id
			break
		}
	}

	// 权限模块不存在
	if permissionId == 0 {
		r.Response.WriteJsonExit(g.Map{
			"code":    -1,
			"message": "访问权限错误",
		})
	}

	// 没有访问权限
	existsPermission := utility.InArray(permissionId, rolePermission)
	if existsPermission == false {
		r.Response.WriteJsonExit(g.Map{
			"code":    -1,
			"message": "访问权限错误",
		})
	}

	r.Middleware.Next()
}
