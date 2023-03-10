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
	service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
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
	service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 4,
		Description:   "编辑管理员角色信息绑定",
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
	service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
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
	service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
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
		Type:   req.Type,
		Remark: req.Remark,
	})

	// 访问日志写入
	service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
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
		Id:     req.ID,
		Fid:    req.Fid,
		Name:   req.Name,
		Status: req.Status,
		Remark: req.Remark,
	})

	// 访问日志写入
	service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
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
	})

	return
}

// ListPermission
//
// @Title 权限信息列表
// @Description 权限信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-11 10:35:29
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMasterAuth) ListPermission(ctx context.Context, req *manage.ListPermissionReq) (res *manage.ListRoleRes, err error) {
	code, message, output, err := service.UcSystemMasterAuth().ListPermission(ctx, system_master.ListPermissionInput{
		Page: req.Page,
		Size: req.Size,
	})

	// 访问日志写入
	service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
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
		"data":    g.Map{"list": output},
	})

	return
}

func (c *cUcSystemMasterAuth) EditPermissionRelation(ctx context.Context, req *manage.EditPermissionRelationReq) (res *manage.EditPermissionRelationRes, err error) {
	code, message, err := service.UcSystemMasterAuth().ModifyPermissionRelation(ctx, system_master.ModifyPermissionRelationInput{
		PermissionIds: req.PermissionIds,
		RoleId:        req.RoleId,
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
