package cmd

import (
	"context"
	"ice_flame/internal/controller"
	"ice_flame/internal/controller/test"

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
					test.TestDB.TestInsert,
					test.TestDB.TestInsertAndGetId,
					test.TestDB.TestGetStructById,
					test.TestDB.TestGetStructGCacheById,
					test.TestDB.TestGetStructRCacheById,
					test.TestDB.TestGetMapById,
					test.TestDB.TestGetMapGCacheById,
					test.TestDB.TestGetMapRCacheById,
					test.TestDB.TestGetStructByIds,
					test.TestDB.TestGetMapByIds,
				)
			})
			s.Run()
			return nil
		},
	}
)
