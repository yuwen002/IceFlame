package manage

import (
	"context"
	"ice_flame/api/manage"
	"ice_flame/internal/model/article"
	"ice_flame/internal/model/uc_center/system_master"
	"ice_flame/internal/service"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/frame/g"
)

var Article = cArticle{}

type cArticle struct {
}

// AddSinglePage
//
// @Title 添加单页信息
// @Description 添加单页信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-22 16:04:21
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cArticle) AddSinglePage(ctx context.Context, req *manage.AddSinglePageReq) (res *manage.AddSinglePageRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 5,
		Description:   "添加单页信息",
	})

	code, message, err := service.SinglePage().Create(ctx, article.CreateSinglePageInput{
		Title:       req.Title,
		Description: req.Description,
		Keyword:     req.Keyword,
		Content:     req.Content,
		Thumbnail:   req.Thumbnail,
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

// GetSinglePage
//
// @Title 获取单页信息
// @Description 获取单页信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-22 16:26:17
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cArticle) GetSinglePage(ctx context.Context, req *manage.GetSinglePageReq) (res *manage.GetSinglePageRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 5,
		Description:   "查看编辑单页信息",
	})

	code, message, out, err := service.SinglePage().GetById(ctx, req.Id)

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

// EditSinglePage
//
// @Title 修改单页信息
// @Description 修改单页信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-22 18:01:17
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cArticle) EditSinglePage(ctx context.Context, req *manage.EditSinglePageReq) (res *manage.EditSinglePageRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 5,
		Description:   "提交编辑单页信息",
	})

	code, message, err := service.SinglePage().ModifyById(ctx, article.ModifySinglePageInput{
		Id:          req.Id,
		Title:       req.Title,
		Description: req.Description,
		Keyword:     req.Keyword,
		Content:     req.Content,
		Thumbnail:   req.Thumbnail,
		Click:       req.Click,
		Status:      req.Status,
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

// ListSinglePage
//
// @Title 单页信息列表
// @Description 单页信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-22 18:10:19
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cArticle) ListSinglePage(ctx context.Context, req *manage.ListSinglePageReq) (res *manage.ListSinglePageRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 5,
		Description:   "查看单页信息列表",
	})

	code, message, out, err := service.SinglePage().List(ctx, article.ListSinglePageInput{
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

// DeleteSinglePage
//
// @Title 删除单页信息
// @Description 删除单页信息
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-22 23:31:00
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cArticle) DeleteSinglePage(ctx context.Context, req *manage.DeleteSinglePageReq) (res *manage.DeleteSinglePageRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 5,
		Description:   "删除编辑单页信息",
	})

	code, message, err := service.SinglePage().Delete(ctx, req.Id)

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
