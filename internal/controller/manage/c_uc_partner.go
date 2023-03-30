package manage

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"ice_flame/api/manage"
	"ice_flame/internal/model/uc_center/system_master"
	"ice_flame/internal/service"
)

var UcPartner = cUcPartner{}

type cUcPartner struct {
}

// AddPartnerLevel
//
// @Title 添加合伙人级别
// @Description 添加合伙人级别
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-30 17:34:20
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcPartner) AddPartnerLevel(ctx context.Context, req *manage.AddPartnerLevelReq) (res *manage.AddPartnerLevelRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 6,
		Description:   "添加合伙人级别",
	})

	code, message, err := service.UcPartner().CreatePartnerLevel(ctx, system_master.CreatePartnerLevelInput{
		Name:   req.Name,
		Sort:   req.Sort,
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

// GetPartnerLevel
//
// @Title 编辑合伙人级别
// @Description 编辑合伙人级别
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-30 17:58:33
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcPartner) GetPartnerLevel(ctx context.Context, req *manage.GetPartnerLevelReq) (res *manage.GetPartnerLevelRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 6,
		Description:   "查看合伙人级别",
	})

	code, message, out, err := service.UcPartner().GetPartnerLevelById(ctx, req.Id)

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

// EditPartnerLevel
//
// @Title 编辑合伙人级别
// @Description 编辑合伙人级别
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-30 18:18:04
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cUcPartner) EditPartnerLevel(ctx context.Context, req *manage.EditPartnerLevelReq) (res *manage.EditPartnerLevelRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 6,
		Description:   "提交合伙人级别",
	})

	code, message, err := service.UcPartner().ModifyPartnerLevelById(ctx, system_master.ModifyPartnerLevelInput{
		Id:     req.Id,
		Name:   req.Name,
		Sort:   req.Sort,
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
