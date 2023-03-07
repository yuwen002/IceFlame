package manage

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"ice_flame/api/manage"
	"ice_flame/internal/model/uc_center/system_master"
	"ice_flame/internal/service"
)

var UcSystemMasterAuth = cUcSystemMasterAuth{}

type cUcSystemMasterAuth struct {
}

// AddRole
//
// @Title 添加管理员角色信息
// @Description 添加管理员角色信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-07 16:32:22
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterAuth) AddRole(ctx context.Context, req *manage.AddRoleReq) (res *manage.AddRoleRes, err error) {
	code, message, err := service.UcSystemMasterAuth().CreateRole(ctx, system_master.CreateRoleInput{
		Name:   req.Name,
		Remark: req.Remark,
	})

	// 访问日志写入
	service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 4,
		Description:   "添加管理员角色信息",
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

// EditRole
//
// @Title 编辑管理员角色信息
// @Description 编辑管理员角色信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-07 16:37:00
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterAuth) EditRole(ctx context.Context, req *manage.EditRoleReq) (res *manage.EditRoleRes, err error) {
	code, message, err := service.UcSystemMasterAuth().ModifyRoleById(ctx, system_master.ModifyRoleByIdInput{
		Id:     req.Id,
		Name:   req.Name,
		Remark: req.Remark,
	})

	// 访问日志写入
	service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 4,
		Description:   "编辑管理员角色信息",
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

// ListRole
//
// @Title 管理员角色信息列表
// @Description 管理员角色信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-07 16:51:49
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterAuth) ListRole(ctx context.Context, req *manage.ListRoleReq) (res *manage.ListRoleRes, err error) {
	code, message, output, err := service.UcSystemMasterAuth().ListRole(ctx, system_master.ListRoleInput{
		Page: req.Page,
		Size: req.Size,
	})

	// 访问日志写入
	service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 4,
		Description:   "查看管理员角色信息列表",
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
		"data": g.Map{
			"list": output,
		},
	})

	return
}
