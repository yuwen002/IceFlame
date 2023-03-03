package manage

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
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
	code, message, err := service.UcSystemMasterVisitorLogs().AddVisitCategory(ctx, system_master.AddVisitCategoryInput{Title: req.Title})
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
	code, message, err := service.UcSystemMasterVisitorLogs().ModifyVisitCategoryById(ctx, system_master.ModifyVisitCategoryByIdInput{
		Id:    req.Id,
		Title: req.Title,
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
	code, message, output, err := service.UcSystemMasterVisitorLogs().ListVisitCategory(ctx, system_master.ListVisitCategoryInput{
		Page: req.Page,
		Size: req.Size,
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
