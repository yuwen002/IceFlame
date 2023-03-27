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

// AddArticleChannel
//
// @Title 添加频道信息
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 14:17:28
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cArticle) AddArticleChannel(ctx context.Context, req *manage.AddArticleChannelReq) (res *manage.AddArticleChannelRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 5,
		Description:   "添加频道信息",
	})

	code, message, err := service.Article().CreateChannel(ctx, article.CreateChannelInput{
		Name:   req.Name,
		Remark: req.Remark,
		Sort:   req.Sort,
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

// GetArticleChannel
//
// @Title 获取频道信息
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 14:56:57
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cArticle) GetArticleChannel(ctx context.Context, req *manage.GetArticleChannelReq) (res *manage.GetArticleChannelRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 5,
		Description:   "查看编辑频道信息",
	})

	code, message, out, err := service.Article().GetChannelById(ctx, req.Id)

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

// EditArticleChannel
//
// @Title 编辑频道信息
// @Description 编辑频道信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 15:44:06
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cArticle) EditArticleChannel(ctx context.Context, req *manage.EditArticleChannelReq) (res *manage.EditArticleChannelRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 5,
		Description:   "提交编辑频道信息",
	})

	code, message, err := service.Article().ModifyChannelById(ctx, article.ModifyChannelInput{
		Id:     req.Id,
		Name:   req.Name,
		Remark: req.Remark,
		Sort:   req.Sort,
		Status: req.Status,
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

// ListArticleChannel
//
// @Title 频道信息列表
// @Description 频道信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 16:03:07
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cArticle) ListArticleChannel(ctx context.Context, req *manage.ListArticleChannelReq) (res *manage.ListArticleChannelRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 5,
		Description:   "查看频道信息列表",
	})

	code, message, out, err := service.Article().ListChannel(ctx, article.ListChannelInput{
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

// DelArticleChannel
//
// @Title 删除频道信息
// @Description 删除频道信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 16:33:32
// @receiver c
// @param ctx
// @param req
// @return res
func (c *cArticle) DelArticleChannel(ctx context.Context, req *manage.DelArticleChannelReq) (res *manage.DelArticleChannelRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 5,
		Description:   "删除频道信息",
	})

	code, message, err := service.Article().DelChannelById(ctx, req.Id)

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

// GetArticleChannelAll
//
// @Title 获取所有频道信息列表
// @Description 获取所有频道信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-24 17:15:56
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cArticle) GetArticleChannelAll(ctx context.Context, req *manage.GetArticleChannelAllReq) (res *manage.GetArticleChannelAllRes, err error) {
	code, message, out, err := service.Article().GetChannelAll(ctx)

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

// AddArticleCategory
//
// @Title 添加文章分类
// @Description 添加文章分类
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 18:05:16
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cArticle) AddArticleCategory(ctx context.Context, req *manage.AddArticleCategoryReq) (res *manage.AddArticleCategoryRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 5,
		Description:   "添加分类信息",
	})

	code, message, err := service.Article().CreateCategory(ctx, article.CreateCategoryInput{
		Fid:    req.Fid,
		Name:   req.Name,
		Remark: req.Remark,
		Sort:   req.Sort,
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

// GetArticleCategory
//
// @Title 查看编辑分类信息
// @Description 查看编辑分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 18:11:16
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cArticle) GetArticleCategory(ctx context.Context, req *manage.GetArticleCategoryReq) (res *manage.GetArticleCategoryRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 5,
		Description:   "查看编辑分类信息",
	})

	code, message, out, err := service.Article().GetCategoryById(ctx, req.Id)

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

// EditArticleCategory
//
// @Title 提交编辑分类信息
// @Description 提交编辑分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 18:12:49
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cArticle) EditArticleCategory(ctx context.Context, req *manage.EditArticleCategoryReq) (res *manage.EditArticleCategoryRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 5,
		Description:   "提交编辑分类信息",
	})

	code, message, err := service.Article().ModifyCategoryById(ctx, article.ModifyCategoryInput{
		Id:     req.Id,
		Fid:    req.Fid,
		Name:   req.Name,
		Remark: req.Remark,
		Sort:   req.Sort,
		Status: req.Status,
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

// ListArticleCategory
//
// @Title 查看分类信息列表
// @Description 查看分类信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-23 23:17:16
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cArticle) ListArticleCategory(ctx context.Context, req *manage.ListArticleCategoryReq) (res *manage.ListArticleCategoryRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 5,
		Description:   "查看分类信息列表",
	})

	code, message, out, err := service.Article().ListCategory(ctx, article.ListCategoryInput{
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

// DelArticleCategory
//
// @Title 删除分类信息
// @Description 删除分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-24 00:07:26
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cArticle) DelArticleCategory(ctx context.Context, req *manage.DelArticleCategoryReq) (res *manage.DelArticleCategoryRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 5,
		Description:   "删除分类信息",
	})

	code, message, err := service.Article().DelCategoryById(ctx, req.Id)

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

// GetArticleCategoryAll
//
// @Title 所有分类信息列表
// @Description 所有分类信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-24 17:39:58
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cArticle) GetArticleCategoryAll(ctx context.Context, req *manage.GetArticleCategoryAllReq) (res *manage.GetArticleCategoryAllRes, err error) {
	code, message, out, err := service.Article().GetCategoryAll(ctx)

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

// AddArticleTag
//
// @Title 添加标签
// @Description 添加标签
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-26 01:27:42
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cArticle) AddArticleTag(ctx context.Context, req *manage.AddArticleTagReq) (res *manage.AddArticleTagRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 5,
		Description:   "添加标签信息",
	})

	code, message, err := service.Article().CreateTag(ctx, article.CreateTagInput{
		Name:   req.Name,
		Remark: req.Remark,
		Sort:   req.Sort,
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

// GetArticleTag
//
// @Title 查看标签
// @Description 查看标签
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-26 01:27:34
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cArticle) GetArticleTag(ctx context.Context, req *manage.GetArticleTagReq) (res *manage.GetArticleTagRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 5,
		Description:   "查看编辑标签信息",
	})

	code, message, out, err := service.Article().GetTagById(ctx, req.Id)

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

func (c *cArticle) EditArticleTag(ctx context.Context, req *manage.EditArticleTagReq) (res *manage.EditArticleTagRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 5,
		Description:   "提交编辑标签信息",
	})

	code, message, err := service.Article().ModifyTagById(ctx, article.ModifyTagInput{
		Id:     req.Id,
		Name:   req.Name,
		Remark: req.Remark,
		Sort:   req.Sort,
		Status: req.Status,
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

// ListArticleTag
//
// @Title 查看标签列表
// @Description 查看标签列表
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-26 01:38:23
// @receiver c
// @param ctx
// @param req
// @return res
func (c *cArticle) ListArticleTag(ctx context.Context, req *manage.ListArticleTagReq) (res *manage.ListArticleTagRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 5,
		Description:   "查看标签列表",
	})

	code, message, out, err := service.Article().ListTag(ctx, article.ListTagInput{
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

// DelArticleTag
//
// @Title 删除标签
// @Description 删除标签
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-26 01:40:40
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cArticle) DelArticleTag(ctx context.Context, req *manage.DelArticleTagReq) (res *manage.DelArticleTagRes, err error) {
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 5,
		Description:   "删除标签",
	})

	code, message, err := service.Article().DelTagById(ctx, req.Id)

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

// GetArticleTagAll
//
// @Title 所有标签信息
// @Description 所有标签信息
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-26 01:43:20
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cArticle) GetArticleTagAll(ctx context.Context, req *manage.GetArticleTagAllReq) (res *manage.GetArticleTagAllRes, err error) {
	code, message, out, err := service.Article().GetTagAll(ctx)
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

// AddArticle
//
// @Title 添加文章
// @Description 添加文章
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-25 13:26:06
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cArticle) AddArticle(ctx context.Context, req *manage.AddArticleReq) (res *manage.AddArticleRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 5,
		Description:   "添加文章",
	})

	code, message, err := service.Article().CreateArticle(ctx, article.CreateArticleInput{
		CategoryId:  req.CategoryId,
		ChannelId:   req.ChannelId,
		Title:       req.Title,
		Keyword:     req.Keyword,
		Description: req.Description,
		Link:        req.Link,
		Author:      req.Author,
		Tags:        req.Tags,
		PubDate:     req.PubDate,
		Summary:     req.Summary,
		Content:     req.Content,
		Thumbnail:   req.Thumbnail,
		Click:       req.Click,
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

// GetArticle
//
// @Title 查看编辑文章
// @Description 查看编辑文章
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-25 14:50:28
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cArticle) GetArticle(ctx context.Context, req *manage.GetArticleReq) (res *manage.GetArticleRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 5,
		Description:   "查看编辑文章",
	})

	code, message, out, err := service.Article().GetArticleById(ctx, req.Id)

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

// EditArticle
//
// @Title 提交编辑文章
// @Description 提交编辑文章
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-25 16:37:58
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cArticle) EditArticle(ctx context.Context, req *manage.EditArticleReq) (res *manage.EditArticleRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 5,
		Description:   "提交编辑文章",
	})

	code, message, err := service.Article().ModifyArticleById(ctx, article.ModifyArticleInput{
		Id:          req.Id,
		CategoryId:  req.CategoryId,
		ChannelId:   req.ChannelId,
		Title:       req.Title,
		Keyword:     req.Keyword,
		Description: req.Description,
		Link:        req.Link,
		Author:      req.Author,
		Tags:        req.Tags,
		PubDate:     req.PubDate,
		Summary:     req.Summary,
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

// ListArticle
//
// @Title 文章信息列表
// @Description 文章信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-25 16:50:02
// @receiver c
// @param ctx
// @param req
// @return res
func (c *cArticle) ListArticle(ctx context.Context, req *manage.ListArticleReq) (res *manage.ListArticleRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 5,
		Description:   "查看文章列表",
	})

	code, message, out, err := service.Article().ListArticle(ctx, article.ListArticleInput{
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

// DelArticle
//
// @Title 删除文章
// @Description 删除文章
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-25 17:30:14
// @receiver c
// @param ctx
// @param req
// @return res
// @return err
func (c *cArticle) DelArticle(ctx context.Context, req *manage.DelArticleReq) (res *manage.DelArticleRes, err error) {
	// 访问日志写入
	_, _, _ = service.UcSystemMasterVisitor().CreateVisitorLogs(ctx, system_master.CreateVisitorLogsInput{
		AccountId:     gconv.Uint64(ctx.Value("master_id")),
		VisitCategory: 5,
		Description:   "删除文章",
	})

	code, message, err := service.Article().DelArticleById(ctx, req.Id)

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
