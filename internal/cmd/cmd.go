package cmd

import (
	"context"
	"ice_flame/internal/controller"
	"ice_flame/internal/controller/manage"
	"ice_flame/internal/controller/test"
	"ice_flame/internal/service"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					controller.Hello,

					// 测试DB结构体方法
					test.DBTest.TestInsert,
					test.DBTest.TestInsertAndGetId,
					test.DBTest.TestGetStructById,
					test.DBTest.TestGetStructGCacheById,
					test.DBTest.TestGetStructRCacheById,
					test.DBTest.TestGetMapById,
					test.DBTest.TestGetMapGCacheById,
					test.DBTest.TestGetMapRCacheById,
					test.DBTest.TestGetStructByIds,
					test.DBTest.TestGetMapByIds,
					test.DBTest.TestGetOneStructByWhere,
					test.DBTest.TestGetOneMapByWhere,
					test.DBTest.TestDBGetAllStructByWhere,
					test.DBTest.TestDBGetAllMapByWhere,
					test.DBTest.TestDBModifyById,
					test.DBTest.TestDBModifyGCacheById,
					test.DBTest.TestDBModifyRCacheById,
					test.DBTest.TestDBModifyByWhere,
					test.DBTest.TestDBDelById,
					test.DBTest.TestDBDelByIds,
					test.DBTest.TestDBDelByWhere,
					test.DBTest.TestDBModifyIncById,
					test.DBTest.TestDBModifyDecById,
					test.DBTest.TestDBModifyIncByWhere,
					test.DBTest.TestDBModifyDecByWhere,

					// 测试Redis结构体方法
					test.RedisCacheTest.TestExistsSetData,
					test.RedisCacheTest.TestCastHashData,

					// Jwt测试
					test.JwtClaimsTest.TestJCSetToken,
					test.JwtClaimsTest.TestJCGetToken,
				)
			})
			// 后台管理分组
			s.Group("/manage",
				func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().MiddlewareHandlerResponse)
					group.Bind(
						manage.UcSystemMaster.Register,
						manage.UcSystemMaster.LoginTelPassword,
						manage.UcSystemMaster.LoginUsernamePassword,
					)

				},
				func(group *ghttp.RouterGroup) {
					group.Middleware(
						service.AuthMiddleware().MiddlewareAuthMaster,
						service.Middleware().MiddlewareHandlerResponse,
					)
					group.Bind(
						manage.UcSystemMaster.CreateSystemMaster,
						manage.UcSystemMaster.EditPassword,
						manage.UcSystemMaster.ListSystemMaster,
						manage.UcSystemMaster.EditSystemMaster,
					)
				},
			)

			s.Run()
			return nil
		},
	}
)
