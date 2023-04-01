package manage

import (
	"context"
	"ice_flame/api/manage"
	"ice_flame/internal/model/uc_center/system_master"
	"ice_flame/internal/service"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/frame/g"
)

var UcEmployee = cUcEmployee{}

type cUcEmployee struct {
}

// AddEmployeeRole
//
// @Title 添加员工角色
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-02 00:24:44
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcEmployee) AddEmployeeRole(ctx context.Context, req *manage.AddEmployeeRoleReq) (res *manage.AddEmployeeRoleRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 6,
		Description:   "添加员工角色",
	})

	code, message, err := service.UcEmployee().CreateEmployeeRole(ctx, system_master.CreateEmployeeRoleInput{
		Name:   req.Name,
		Remark: req.Remark,
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

// GetEmployeeRole
//
// @Title 查看员工角色
// @Description 查看员工角色
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-02 00:40:50
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcEmployee) GetEmployeeRole(ctx context.Context, req *manage.GetEmployeeRoleReq) (res *manage.GetEmployeeRoleRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 6,
		Description:   "查看员工角色",
	})

	code, message, out, err := service.UcEmployee().GetEmployeeRoleById(ctx, req.Id)

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

// EditEmployeeRole
//
// @Title 编辑员工角色
// @Description 编辑员工角色
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-02 00:47:09
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcEmployee) EditEmployeeRole(ctx context.Context, req *manage.EditEmployeeRoleReq) (res *manage.EditEmployeeRoleRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 6,
		Description:   "提交员工角色",
	})

	code, message, err := service.UcEmployee().ModifyEmployeeRoleById(ctx, system_master.ModifyRoleByIdInput{
		Id:     req.Id,
		Name:   req.Name,
		Remark: req.Remark,
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

// ListEmployeeRole
//
// @Title 员工角色列表
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-02 00:51:31
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcEmployee) ListEmployeeRole(ctx context.Context, req *manage.ListEmployeeRoleReq) (res *manage.ListEmployeeRoleRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 6,
		Description:   "员工角色列表",
	})

	code, message, out, err := service.UcEmployee().ListEmployeeRole(ctx, system_master.ListEmployeeRoleInput{
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
		"data":    g.Map{"list": out},
	})

	return
}
