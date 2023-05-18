package manage

import (
	"context"
	"ice_flame/api/manage"
	"ice_flame/internal/model/uc_center/system_master"
	"ice_flame/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var UcSystemMaster = cUcSystemMaster{}

type cUcSystemMaster struct {
}

// Register
//
// @Title 管理员用户注册
// @Description 管理员用户注册
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-22 17:32:27
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMaster) Register(ctx context.Context, req *manage.RegisterReq) (res *manage.RegisterRes, err error) {
	code, message, err := service.UcSystemMaster().Register(ctx, system_master.RegisterInput{
		Password: req.Password,
		Tel:      req.Tel,
		Name:     req.Name,
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

// LoginTelPassword
//
// @Title 管理员用户登入
// @Description 管理员用户登入，登入方式电话号码和密码
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-23 11:40:42
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMaster) LoginTelPassword(ctx context.Context, req *manage.LoginTelPasswordReq) (res *manage.LoginTelPasswordRes, err error) {
	code, message, token, err := service.UcSystemMaster().LoginTelPassword(ctx, system_master.LoginTelPasswordInput{
		Tel:      req.Tel,
		Password: req.Password,
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
		"data":    g.Map{"token": token},
	})

	return
}

// LoginUsernamePassword
//
// @Title 管理员用户登入
// @Description 管理员用户登入，登入方式用户名和密码
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-23 16:30:03
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMaster) LoginUsernamePassword(ctx context.Context, req *manage.LoginUsernamePasswordReq) (res *manage.LoginUsernamePasswordRes, err error) {
	code, message, token, err := service.UcSystemMaster().LoginUsernamePassword(ctx, system_master.LoginUsernamePasswordInput{
		Username: req.Username,
		Password: req.Password,
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
		"data":    g.Map{"token": token},
	})

	return
}

// CreateSystemMaster
//
// @Title 管理员新建管理用户
// @Description 管理员新建管理用户
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-23 17:06:28
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMaster) CreateSystemMaster(ctx context.Context, req *manage.CreateSystemMasterReq) (res *manage.CreateSystemMasterRes, err error) {
	code, message, err := service.UcSystemMaster().CreateSystemMaster(ctx, system_master.CreateSystemMasterInput{
		Username: req.Username,
		Password: req.Password,
		Tel:      req.Tel,
		Name:     req.Name,
	})

	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 2,
		Description:   "新建管理员用户",
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

// EditPassword
//
// @Title 修改管理员密码
// @Description 修改管理员密码
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-24 17:06:04
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMaster) EditPassword(ctx context.Context, req *manage.EditPasswordReq) (res *manage.EditPasswordRes, err error) {
	code, message, err := service.UcSystemMaster().ModifyPasswordSelfById(ctx, system_master.ModifyPasswordInput{
		OldPassword: req.OldPassword,
		NewPassword: req.OldPassword,
	})

	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 2,
		Description:   "管理员修改自己密码",
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

// ListSystemMaster
//
// @Title 管理员用户显示列表
// @Description 管理员用户显示列表
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-02-25 23:23:54
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMaster) ListSystemMaster(ctx context.Context, req *manage.ListSystemMasterReq) (res *manage.ListSystemMasterRes, err error) {
	// 判断页码
	if req.Page == 0 {
		req.Page = 1
	}

	// 判断每页显示条数
	if req.Size == 0 {
		req.Size = 10
	}
	code, message, out, total, err := service.UcSystemMaster().ListSystemMaster(ctx, system_master.ListSystemMasterInput{
		Size: req.Size,
		Page: req.Page,
	})

	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 2,
		Description:   "查看管理员列表",
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
		"data":    g.Map{"list": out, "total": total},
	})
	return
}

// GetEditSystemMaster
//
// @Title 查看编辑管理员信息
// @Description 查看编辑管理员信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-18 15:31:18
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMaster) GetEditSystemMaster(ctx context.Context, req *manage.GetEditSystemMasterReq) (res *manage.GetEditSystemMasterRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 2,
		Description:   "查看编辑管理员信息",
	})

	code, message, out, err := service.UcSystemMaster().GetSystemMasterByAccountId(ctx, req.AccountId)

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

// EditSystemMaster
//
// @Title 编辑用户信息
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-28 14:55:53
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMaster) EditSystemMaster(ctx context.Context, req *manage.EditSystemMasterReq) (res *manage.EditSystemMasterRes, err error) {
	code, message, err := service.UcSystemMaster().ModifySystemMasterByAccountId(ctx, system_master.ModifySystemMasterByAccountIdInput{
		AccountId: req.AccountId,
		Tel:       req.Tel,
		Name:      req.Name,
	})

	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 2,
		Description:   "提交编辑管理员信息",
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

// ResetPassword
//
// @Title 重置管理员密码
// @Description 重置管理员密码
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-28 15:43:40
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMaster) ResetPassword(ctx context.Context, req *manage.ResetPasswordSystemMasterReq) (res *manage.ResetPasswordSystemMasterRes, err error) {
	code, message, err := service.UcSystemMaster().ResetPasswordByAccountId(ctx, system_master.ResetPasswordByAccountIdInput{
		AccountId: req.AccountId,
		Password:  req.Password,
	})

	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 2,
		Description:   "重置管理员密码",
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

// EditStatus
//
// @Title 修改管理用户状态
// @Description 修改管理用户状态
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-28 15:59:46
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMaster) EditStatus(ctx context.Context, req *manage.EditStatusSystemMasterReq) (res *manage.EditStatusSystemMasterRes, err error) {
	code, message, err := service.UcSystemMaster().ModifyStatusByAccountId(ctx, system_master.ModifyStatusByAccountIdInput{
		AccountId: req.AccountId,
		Status:    req.Status,
	})

	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 2,
		Description:   "重置管理员密码",
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

// UnlockSystemMaster
//
// @Title 解锁用户
// @Description 按account ID 解锁用户
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-01 10:51:56
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMaster) UnlockSystemMaster(ctx context.Context, req *manage.UnlockSystemMasterReq) (res *manage.UnlockSystemMasterRes, err error) {
	code, message, err := service.UcSystemMaster().DelLoginCount(ctx, gconv.String(req.AccountId))

	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 2,
		Description:   "解锁管理员登入限制",
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

// GetSystemMasterInfo
//
// @Title 获取用户详细信息
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-04-12 14:38:00
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcSystemMaster) GetSystemMasterInfo(ctx context.Context, req *manage.GetSystemMasterInfoReq) (res *manage.GetSystemMasterInfoRes, err error) {
	code, message, out, err := service.UcSystemMaster().GetSystemMasterByAccountId(ctx, gconv.Uint64(ctx.Value("master_id")))

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
