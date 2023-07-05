package manage

import (
	"context"
	"ice_flame/api/manage"
	"ice_flame/internal/model/uc_center/system_master"
	"ice_flame/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
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
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
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

// GetEditRole
//
// @Title 按ID获取角色编辑信息
// @Description 按ID获取角色编辑信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-16 14:46:29
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterAuth) GetEditRole(ctx context.Context, req *manage.GetEditRoleReq) (res *manage.GetEditRoleRes, err error) {
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 4,
		Description:   "查看编辑管理员角色信息",
	})

	code, message, output, err := service.UcSystemMasterAuth().GetRoleById(ctx, req.Id)
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
		"data":    output,
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
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 4,
		Description:   "提交编辑管理员角色信息",
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
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
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

// GetListRole
//
// @Title 获取所有管理员角色信息
// @Description 获取所有管理员角色信息 select借口
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-15 18:20:48
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterAuth) GetListRole(ctx context.Context, req *manage.GetListRoleReq) (res *manage.GetListRoleRes, err error) {
	code, message, output, err := service.UcSystemMasterAuth().GetRoleAll(ctx)

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

// AddRoleRelation
//
// @Title 添加管理员绑定角色
// @Description 添加管理员绑定角色
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-09 10:37:22
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterAuth) AddRoleRelation(ctx context.Context, req *manage.AddRoleRelationReq) (res *manage.AddRoleRelationRes, err error) {
	code, message, err := service.UcSystemMasterAuth().CreateRoleRelation(ctx, system_master.CreateRoleRelationInput{
		AccountId: req.AccountId,
		RoleId:    req.RoleId,
	})

	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 4,
		Description:   "添加管理员角色信息绑定",
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

// GetEditRoleRelation
//
// @Title 编辑管理员绑定角色
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-16 14:22:24
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterAuth) GetEditRoleRelation(ctx context.Context, req *manage.GetEditRoleRelationReq) (res *manage.GetEditRoleRelationRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 4,
		Description:   "查看编辑管理员角色信息绑定",
	})

	code, message, output, err := service.UcSystemMasterAuth().GetRoleRelationById(ctx, req.Id)
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
		"data":    output,
	})

	return
}

// EditRoleRelation
//
// @Title 编辑管理员绑定角色
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-09 11:11:02
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterAuth) EditRoleRelation(ctx context.Context, req *manage.EditRoleRelationReq) (res *manage.EditRoleRelationRes, err error) {
	code, message, err := service.UcSystemMasterAuth().ModifyRoleRelationById(ctx, system_master.ModifyRoleRelationByIdInput{
		Id:        req.Id,
		AccountId: req.AccountId,
		RoleId:    req.RoleId,
	})

	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 4,
		Description:   "提交编辑管理员角色信息绑定",
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

// ListRoleRelation
//
// @Title 管理员绑定角色列表
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-09 11:11:15
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterAuth) ListRoleRelation(ctx context.Context, req *manage.ListRoleRelationReq) (res *manage.ListRoleRelationRes, err error) {
	code, message, out, err := service.UcSystemMasterAuth().ListRoleRelation(ctx, system_master.ListRoleRelationInput{
		Page: req.Page,
		Size: req.Size,
	})

	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 4,
		Description:   "查看管理员角色信息绑定列表",
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

// DeleteRoleRelation
//
// @Title 删除管理员角色信息绑定
// @Description 删除管理员角色信息绑定
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-09 11:53:40
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterAuth) DeleteRoleRelation(ctx context.Context, req *manage.DeleteRoleRelationReq) (res *manage.DeleteRoleRelationRes, err error) {
	code, message, err := service.UcSystemMasterAuth().DeleteRoleRelationById(ctx, system_master.DeleteRoleRelationInput{Id: req.Id})

	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 4,
		Description:   "删除管理员角色信息绑定",
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

// AddPermission
//
// @Title 添加权限信息
// @Description 添加权限信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-09 17:45:18
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterAuth) AddPermission(ctx context.Context, req *manage.AddPermissionReq) (res *manage.AddRoleRes, err error) {
	code, message, err := service.UcSystemMasterAuth().CreatePermission(ctx, system_master.CreateListPermissionInput{
		Fid:    req.Fid,
		Name:   req.Name,
		Module: req.Module,
		Uri:    req.Uri,
		Type:   req.Type,
		Sort:   req.Sort,
		Remark: req.Remark,
	})

	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 4,
		Description:   "添加菜单信息",
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

// GetEditPermission
//
// @Title 编辑权限信息
// @Description 编辑权限信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-16 16:42:20
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterAuth) GetEditPermission(ctx context.Context, req *manage.GetEditPermissionReq) (res *manage.GetEditPermissionRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 4,
		Description:   "查看编辑菜单信息",
	})

	code, message, output, err := service.UcSystemMasterAuth().GetPermissionById(ctx, req.Id)

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
		"data":    output,
	})

	return
}

// EditPermission
//
// @Title 编辑权限信息
// @Description 编辑权限信息
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-09 22:00:41
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterAuth) EditPermission(ctx context.Context, req *manage.EditPermissionReq) (res *manage.EditPermissionRes, err error) {
	code, message, err := service.UcSystemMasterAuth().ModifyPermissionById(ctx, system_master.ModifyPermissionByIdInput{
		Id:     req.Id,
		Fid:    req.Fid,
		Name:   req.Name,
		Module: req.Module,
		Uri:    req.Uri,
		Status: req.Status,
		Sort:   req.Sort,
		Remark: req.Remark,
	})

	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 4,
		Description:   "提交编辑菜单信息",
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

func (c *cUcSystemMasterAuth) EditStatusPermission(ctx context.Context, req *manage.EditPermissionStatusReq) (res *manage.EditPermissionStatusRes, err error) {
	code, message, err := service.UcSystemMasterAuth().ModifyStatusPermissionById(ctx, system_master.ModifyStatusPermissionByIdInput{
		Id:     req.Id,
		Status: req.Status,
	})

	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 4,
		Description:   "提交编辑菜单状态信息",
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

// ListPermission
//
// @Title 查看权限信息列表
// @Description 权限信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-11 10:35:29
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterAuth) ListPermission(ctx context.Context, req *manage.ListPermissionReq) (res *manage.ListRoleRes, err error) {
	code, message, output, total, err := service.UcSystemMasterAuth().ListPermission(ctx, system_master.ListPermissionInput{
		Page: req.Page,
		Size: req.Size,
	})

	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 4,
		Description:   "编辑菜单信息",
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
			"list":  output,
			"total": total,
		},
	})

	return
}

// GetFirstPermissionList
//
// @Title 一级权限信息列表
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-06-30 11:15:31
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterAuth) GetFirstPermissionList(ctx context.Context, req *manage.GetFirstPermissionListReq) (res *manage.GetFirstPermissionListRes, err error) {
	code, message, output, err := service.UcSystemMasterAuth().ListFirstPermission(ctx)

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

// EditPermissionRelation
//
// @Title 编辑权限信息列表
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-18 14:37:29
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterAuth) EditPermissionRelation(ctx context.Context, req *manage.EditPermissionRelationReq) (res *manage.EditPermissionRelationRes, err error) {
	code, message, err := service.UcSystemMasterAuth().ModifyPermissionRelation(ctx, system_master.ModifyPermissionRelationInput{
		PermissionIds: req.PermissionIds,
		RoleId:        req.RoleId,
	})

	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 4,
		Description:   "编辑分配权限信息",
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

// ListPermissionRelation
//
// @Title 获取权限信息
// @Description 获取权限信息,并分配以选中权限
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-15 17:29:12
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterAuth) ListPermissionRelation(ctx context.Context, req *manage.ListPermissionRelationReq) (res *manage.ListPermissionRelationRes, err error) {
	code, message, output, err := service.UcSystemMasterAuth().GetPermissionByRoleId(ctx, req.RoleId)

	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 4,
		Description:   "查看分配权限信息",
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

// ListMenu
//
// @Title
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-17 15:19:22
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterAuth) ListMenu(ctx context.Context, req *manage.ListMenuReq) (res *manage.ListMenuRes, err error) {
	code, message, output, err := service.UcSystemMasterAuth().GetMasterMenu(ctx, gconv.Uint64(ctx.Value("master_id")))

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

// AddPermissionExclude
//
// @Title 添加权限排除模块
// @Description 添加权限排除模块
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-18 23:20:09
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterAuth) AddPermissionExclude(ctx context.Context, req *manage.AddPermissionExcludeReq) (res *manage.AddPermissionExcludeRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 4,
		Description:   "添加权限排除信息",
	})

	code, message, err := service.UcSystemMasterAuth().CreatePermissionExclude(ctx, system_master.CreatePermissionExcludeInput{
		Name:   req.Name,
		Module: req.Module,
		Uri:    req.Uri,
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

// GetEditPermissionExclude
//
// @Title 查看编辑排除权限
// @Description 查看编辑排除权限
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-18 23:45:33
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterAuth) GetEditPermissionExclude(ctx context.Context, req *manage.GetEditPermissionExcludeReq) (res *manage.GetEditPermissionExcludeRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 4,
		Description:   "查看编辑排除权限信息",
	})

	code, message, out, err := service.UcSystemMasterAuth().GetPermissionExcludeById(ctx, req.Id)

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

// EditPermissionExclude
//
// @Title 修改编辑排除权限
// @Description 修改编辑排除权限
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-18 23:46:03
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterAuth) EditPermissionExclude(ctx context.Context, req *manage.EditPermissionExcludeReq) (res *manage.EditPermissionExcludeRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 4,
		Description:   "修改编辑排除权限信息",
	})

	code, message, err := service.UcSystemMasterAuth().ModifyPermissionExcludeById(ctx, system_master.ModifyPermissionExcludeByIdInput{
		Id:     req.Id,
		Name:   req.Name,
		Module: req.Module,
		Uri:    req.Uri,
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

// ListPermissionExclude
//
// @Title 查看排除权限信息列表
// @Description 查看排除权限信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-18 23:55:44
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterAuth) ListPermissionExclude(ctx context.Context, req *manage.ListPermissionExcludeReq) (res *manage.ListPermissionExcludeRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 4,
		Description:   "查看排除权限信息列表",
	})

	code, message, out, err := service.UcSystemMasterAuth().ListPermissionExclude(ctx, system_master.ListPermissionExcludeInput{
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

// DeletePermissionExclude
//
// @Title 删除排除权限信息列表
// @Description 删除排除权限信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-19 00:00:16
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterAuth) DeletePermissionExclude(ctx context.Context, req *manage.DeletePermissionExcludeReq) (res *manage.DeletePermissionExcludeRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 4,
		Description:   "删除排除权限信息列表",
	})

	code, message, err := service.UcSystemMasterAuth().DeletePermissionExcludeById(ctx, req.Id)

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
