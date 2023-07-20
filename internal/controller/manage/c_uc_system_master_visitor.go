package manage

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"ice_flame/api/manage"
	"ice_flame/internal/model/uc_center/system_master"
	"ice_flame/internal/service"
)

var UcSystemMasterVisitor = cUcSystemMasterVisitor{}

type cUcSystemMasterVisitor struct {
}

// AddVisitCategory
//
// @Title 添加访问类型
// @Description 添加访问类型
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-03 17:18:27
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterVisitor) AddVisitCategory(ctx context.Context, req *manage.AddVisitCategoryReq) (res *manage.AddVisitCategoryRes, err error) {
	code, message, err := service.UcSystemMasterVisitor().CreateVisitCategory(ctx, system_master.CreateVisitCategoryInput{Title: req.Title})

	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 3,
		Description:   "新建访问日志分类",
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
	})

	return
}

// GetEditVisitCategory
//
// @Title 查看编辑访问日志分类
// @Description 查看编辑访问日志分类
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-18 15:46:49
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterVisitor) GetEditVisitCategory(ctx context.Context, req *manage.GetEditVisitCategoryReq) (res *manage.GetEditVisitCategoryRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 3,
		Description:   "查看编辑访问日志分类",
	})

	code, message, out, err := service.UcSystemMasterVisitor().GetVisitCategoryById(ctx, req.Id)

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
		"data":    g.Map{"info": out},
	})
	return
}

// EditVisitCategory
//
// @Title 编辑访问类型
// @Description 编辑访问类型
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-03 17:18:44
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterVisitor) EditVisitCategory(ctx context.Context, req *manage.EditVisitCategoryReq) (res *manage.EditVisitCategoryRes, err error) {
	code, message, err := service.UcSystemMasterVisitor().ModifyVisitCategoryById(ctx, system_master.ModifyVisitCategoryByIdInput{
		Id:    req.Id,
		Title: req.Title,
	})

	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 3,
		Description:   "提交访问日志分类",
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
	})

	return
}

// ListVisitCategory
//
// @Title 访问类型列表
// @Description 访问类型列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-03 18:07:40
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterVisitor) ListVisitCategory(ctx context.Context, req *manage.ListVisitCategoryReq) (res *manage.ListVisitCategoryRes, err error) {
	code, message, output, total, err := service.UcSystemMasterVisitor().ListVisitCategory(ctx, system_master.ListVisitCategoryInput{
		Page: req.Page,
		Size: req.Size,
	})

	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 3,
		Description:   "查看访问日志分类列表",
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
		"data":    g.Map{"list": output, "total": total},
	})

	return
}

// DeleteCacheVisitCategory
//
// @Title 删除缓存信息
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-03 21:04:03
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterVisitor) DeleteCacheVisitCategory(ctx context.Context, req *manage.DeleteVisitCategoryReq) (res *manage.DeleteVisitCategoryRes, err error) {
	code, message, err := service.UcSystemMasterVisitor().DelRCacheVisitCategory()

	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 3,
		Description:   "清除访问日志分类缓存",
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
	})
	return
}

// ListVisitorLogs
//
// @Title 管理员用户访问记录日志列表
// @Description 管理员用户访问记录日志列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-04 15:21:15
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterVisitor) ListVisitorLogs(ctx context.Context, req *manage.ListVisitorLogsReq) (res *manage.ListSystemMasterRes, err error) {
	code, message, output, err := service.UcSystemMasterVisitor().ListVisitorLogs(ctx, system_master.ListVisitorLogsInput{
		AccountId:     req.Id,
		OsCategory:    req.OsCategory,
		VisitCategory: req.VisitCategory,
		Page:          req.Page,
		Size:          req.Size,
	})

	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 3,
		Description:   "查看访问日志列表",
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
		"data":    g.Map{"list": output},
	})
	return
}
