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
func (c *cUcSystemMaster) Register(ctx context.Context, req *manage.RegisterSystemMasterReq) (res *manage.RegisterSystemMasterRes, err error) {
	code, message, err := service.UcSystemMaster().CreateSystemMaster(ctx, uc_center.CreateSystemMasterInput{
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
