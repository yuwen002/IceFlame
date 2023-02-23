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
	code, message, err := service.UcSystemMaster().Register(ctx, uc_center.RegisterInput{
		Username: req.Tel,
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
	code, message, token, err := service.UcSystemMaster().LoginTelPassword(ctx, uc_center.LoginTelPasswordInput{
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
	code, message, token, err := service.UcSystemMaster().LoginUsernamePassword(ctx, uc_center.LoginUsernamePasswordInput{
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
