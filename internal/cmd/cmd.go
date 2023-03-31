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
						service.AuthMiddleware().MiddlewareVerifyPermission,
					)

					group.Bind(
						// 管理员用户管理
						manage.UcSystemMaster.CreateSystemMaster,
						manage.UcSystemMaster.EditPassword,
						manage.UcSystemMaster.ListSystemMaster,
						manage.UcSystemMaster.GetEditSystemMaster,
						manage.UcSystemMaster.EditSystemMaster,
						manage.UcSystemMaster.ResetPassword,
						manage.UcSystemMaster.EditStatus,
						manage.UcSystemMaster.UnlockSystemMaster,

						// 操作日志
						manage.UcSystemMasterVisitor.AddVisitCategory,
						manage.UcSystemMasterVisitor.GetEditVisitCategory,
						manage.UcSystemMasterVisitor.EditVisitCategory,
						manage.UcSystemMasterVisitor.ListVisitCategory,
						manage.UcSystemMasterVisitor.DeleteCacheVisitCategory,
						manage.UcSystemMasterVisitor.ListVisitorLogs,

						// 角色管理
						manage.UcSystemMasterAuth.AddRole,
						manage.UcSystemMasterAuth.GetEditRole,
						manage.UcSystemMasterAuth.EditRole,
						manage.UcSystemMasterAuth.ListRole,
						manage.UcSystemMasterAuth.GetListRole,
						// 角色关联
						manage.UcSystemMasterAuth.AddRoleRelation,
						manage.UcSystemMasterAuth.GetEditRoleRelation,
						manage.UcSystemMasterAuth.EditRoleRelation,
						manage.UcSystemMasterAuth.ListRoleRelation,
						manage.UcSystemMasterAuth.DeleteRoleRelation,

						// 权限管理
						manage.UcSystemMasterAuth.AddPermission,
						manage.UcSystemMasterAuth.GetEditPermission,
						manage.UcSystemMasterAuth.EditPermission,
						manage.UcSystemMasterAuth.ListPermission,

						// 权限分配
						manage.UcSystemMasterAuth.EditPermissionRelation,
						manage.UcSystemMasterAuth.ListPermissionRelation,

						// 导航菜单
						manage.UcSystemMasterAuth.ListMenu,

						// 权限排除模块
						manage.UcSystemMasterAuth.AddPermissionExclude,
						manage.UcSystemMasterAuth.GetEditPermissionExclude,
						manage.UcSystemMasterAuth.EditPermissionExclude,
						manage.UcSystemMasterAuth.ListPermissionExclude,
						manage.UcSystemMasterAuth.DeletePermissionExclude,

						// 文章管理
						// 单页管理
						manage.Article.AddSinglePage,
						manage.Article.GetSinglePage,
						manage.Article.EditSinglePage,
						manage.Article.ListSinglePage,
						manage.Article.DeleteSinglePage,

						// 频道管理
						manage.Article.AddArticleChannel,
						manage.Article.GetArticleChannel,
						manage.Article.EditArticleChannel,
						manage.Article.ListArticleChannel,
						manage.Article.DelArticleChannel,
						manage.Article.GetArticleChannelAll,

						// 分类管理
						manage.Article.AddArticleCategory,
						manage.Article.GetArticleCategory,
						manage.Article.EditArticleCategory,
						manage.Article.ListArticleCategory,
						manage.Article.DelArticleCategory,
						manage.Article.GetArticleCategoryAll,

						// 标签管理
						manage.Article.AddArticleTag,
						manage.Article.GetArticleTag,
						manage.Article.EditArticleTag,
						manage.Article.ListArticleTag,
						manage.Article.DelArticleTag,

						// 文章管理
						manage.Article.AddArticle,
						manage.Article.GetArticle,
						manage.Article.EditArticle,
						manage.Article.ListArticle,
						manage.Article.DelArticle,

						// 合伙人管理
						// 添加合伙人级别
						manage.UcPartner.AddPartnerLevel,
						manage.UcPartner.GetPartnerLevel,
						manage.UcPartner.EditPartnerLevel,
						manage.UcPartner.ListPartnerLevel,
					)
				},
			)

			s.Run()
			return nil
		},
	}
)
