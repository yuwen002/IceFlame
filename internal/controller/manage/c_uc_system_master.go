package manage

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"ice_flame/api/manage"
	"ice_flame/internal/model/uc_center"
	"ice_flame/internal/service"
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

// ModifyPassword
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
func (c *cUcSystemMaster) ModifyPassword(ctx context.Context, req *manage.ModifyPasswordReq) (res *manage.ModifyPasswordRes, err error) {
	code, message, err := service.UcSystemMaster().ModifyPasswordSelfById(ctx, system_master.ModifyPasswordInput{
		OldPassword: req.OldPassword,
		NewPassword: req.OldPassword,
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
	code, message, out, err := service.UcSystemMaster().ListSystemMaster(ctx, system_master.ListSystemMasterInput{
		Size: req.Size,
		Page: req.Page,
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
