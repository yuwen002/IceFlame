package article

import (
	"context"
	"ice_flame/internal/dao"
	"ice_flame/internal/model/article"
	"ice_flame/internal/service"
	"ice_flame/utility"
)

var insArticle = sArticle{}

func Article() *sArticle {
	return &insArticle
}

func init() {
	service.RegisterArticle(Article())
}

type sArticle struct {
}

// CreateChannel
//
// @Title 新建频道
// @Description 新建频道
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-23 00:04:29
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sArticle) CreateChannel(ctx context.Context, in article.CreateChannelInput) (code int32, message string, err error) {
	return utility.DBInsert(dao.ArticleChannel.Ctx(ctx), utility.DBInsertInput{Data: in})
}

// GetChannelById
//
// @Title 按ID获取频道信息
// @Description 按ID获取频道信息
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-23 00:10:43
// @receiver s
// @param ctx
// @param id
// @return code
// @return message
// @return output
// @return err
func (s *sArticle) GetChannelById(ctx context.Context, id uint32) (code int32, message string, output *article.ChannelOutput, err error) {
	code, message, err = utility.DBGetStructById(dao.ArticleChannel.Ctx(ctx), utility.DBGetByIdInput{Where: id}, &output)
	return code, message, output, err
}

// ModifyChannelById
//
// @Title 按ID修改频道信息
// @Description 按ID修改频道信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 10:49:20
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sArticle) ModifyChannelById(ctx context.Context, in article.ModifyChannelInput) (code int32, message string, err error) {
	return utility.DBModifyById(dao.ArticleChannel.Ctx(ctx), utility.DBModifyByIdInput{
		Data:  in,
		Where: in.Id,
	})
}

// ListChannel
//
// @Title 频道信息列表
// @Description 频道信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 11:20:39
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return output
// @return err
func (s *sArticle) ListChannel(ctx context.Context, in article.ListChannelInput) (code int32, message string, output []*article.ChannelOutput, err error) {
	code, message, err = utility.DBGetAllStructByWhere(dao.ArticleChannel.Ctx(ctx), utility.DBGetAllByWhereInput{
		Order:    "id desc, sort desc",
		PageType: 1,
		DBPagination: utility.DBPagination{
			Page: in.Page,
			Size: in.Size,
		},
	}, &output)

	return code, message, output, err
}

// DelChannelById
//
// @Title 按ID删除频道信息
// @Description 按ID删除频道信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 11:44:52
// @receiver s
// @param ctx
// @param id
// @return code
// @return message
// @return err
func (s *sArticle) DelChannelById(ctx context.Context, id uint32) (code int32, message string, err error) {

	return utility.DBDelById(dao.ArticleChannel.Ctx(ctx), utility.DBDelByIdInput{Where: id})
}

// GetChannelAll
//
// @Title 获取所有频道信息
// @Description 获取所有频道信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-24 17:10:06
// @receiver s
// @param ctx
// @return code
// @return message
// @return output
// @return err
func (s *sArticle) GetChannelAll(ctx context.Context) (code int32, message string, output []*article.ChannelOutput, err error) {
	code, message, err = utility.DBGetAllStructByWhere(dao.ArticleChannel.Ctx(ctx), utility.DBGetAllByWhereInput{
		Where: "status=0",
		Order: "sort desc, id asc",
	}, &output)

	return code, message, output, err
}

// CreateCategory
//
// @Title 新建文章分类
// @Description 新建文章分类
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 17:23:37
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sArticle) CreateCategory(ctx context.Context, in article.CreateCategoryInput) (code int32, message string, err error) {
	return utility.DBInsert(dao.ArticleCategory.Ctx(ctx), utility.DBInsertInput{Data: in})
}

// GetCategoryById
//
// @Title 按ID获取文章分类
// @Description 按ID获取文章分类
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 17:24:14
// @receiver s
// @param ctx
// @param id
// @return code
// @return message
// @return err
func (s *sArticle) GetCategoryById(ctx context.Context, id uint32) (code int32, message string, output *article.ChannelOutput, err error) {
	code, message, err = utility.DBGetOneStructByWhere(dao.ArticleCategory.Ctx(ctx), utility.DBGetOneByWhereInput{Where: id}, &output)
	return code, message, output, err
}

// ModifyCategoryById
//
// @Title 按ID修改文章分类
// @Description 按ID修改文章分类
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 17:32:20
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sArticle) ModifyCategoryById(ctx context.Context, in article.ModifyCategoryInput) (code int32, message string, err error) {
	return utility.DBModifyById(dao.ArticleCategory.Ctx(ctx), utility.DBModifyByIdInput{
		Data:  in,
		Where: in.Id,
	})
}

// ListCategory
//
// @Title 文章分类列表
// @Description 文章分类列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 17:35:47
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return output
// @return err
func (s *sArticle) ListCategory(ctx context.Context, in article.ListCategoryInput) (code int32, message string, output []*article.CategoryOutput, err error) {
	code, message, err = utility.DBGetAllStructByWhere(dao.ArticleCategory.Ctx(ctx), utility.DBGetAllByWhereInput{
		Order:    "fid asc, sort desc",
		PageType: 1,
		DBPagination: utility.DBPagination{
			Page: in.Page,
			Size: in.Size,
		},
	}, &output)

	return code, message, output, err
}

// DelCategoryById
//
// @Title 按ID删除文章分类
// @Description 按ID删除文章分类
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 17:55:41
// @receiver s
// @param ctx
// @param id
// @return code
// @return message
// @return err
func (s *sArticle) DelCategoryById(ctx context.Context, id uint32) (code int32, message string, err error) {

	return utility.DBDelById(dao.ArticleCategory.Ctx(ctx), utility.DBDelByIdInput{Where: id})
}

// GetCategoryAll
//
// @Title 获取所有分类
// @Description 获取所有分类
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-24 17:32:21
// @receiver s
// @param ctx
// @return code
// @return message
// @return output
// @return err
func (s *sArticle) GetCategoryAll(ctx context.Context) (code int32, message string, output []*article.CategoryOutput, err error) {
	code, message, err = utility.DBGetAllStructByWhere(dao.ArticleCategory.Ctx(ctx), utility.DBGetAllByWhereInput{
		Where: "status=0",
		Order: "sort desc, id asc",
	}, &output)

	return code, message, output, err
}

// CreateArticle
//
// @Title 新建文章
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-25 10:40:09
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sArticle) CreateArticle(ctx context.Context, in article.CreateArticleInput) (code int32, message string, err error) {
	return utility.DBInsert(dao.Article.Ctx(ctx), utility.DBInsertInput{Data: in})
}

// GetArticleById
//
// @Title 按ID获取文章信息
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-25 10:49:59
// @receiver s
// @param ctx
// @param id
// @return code
// @return message
// @return output
// @return err
func (s *sArticle) GetArticleById(ctx context.Context, id uint32) (code int32, message string, output *article.GetArticleOutput, err error) {
	code, message, err = utility.DBGetStructById(dao.Article.Ctx(ctx), utility.DBGetByIdInput{Where: id}, &output)
	return code, message, output, err
}

// ModifyArticleById
//
// @Title
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-25 11:05:45
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sArticle) ModifyArticleById(ctx context.Context, in article.ModifyArticleInput) (code int32, message string, err error) {
	return utility.DBModifyById(dao.Article.Ctx(ctx), utility.DBModifyByIdInput{
		Data:  in,
		Where: in.Id,
	})
}

// ListArticle
//
// @Title 文章列表
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-25 11:19:24
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return output
// @return err
func (s *sArticle) ListArticle(ctx context.Context, in article.ListArticleInput) (code int32, message string, output []*article.GetArticleOutput, err error) {
	code, message, err = utility.DBGetAllStructByWhere(dao.Article.Ctx(ctx), utility.DBGetAllByWhereInput{
		Order:    "id",
		PageType: 1,
		DBPagination: utility.DBPagination{
			Page: in.Page,
			Size: in.Size,
		},
	}, &output)

	return code, message, output, err
}

// DelArticleById
//
// @Title 按ID删除文章
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-25 11:19:36
// @receiver s
// @param ctx
// @param id
// @return code
// @return message
// @return err
func (s *sArticle) DelArticleById(ctx context.Context, id uint32) (code int32, message string, err error) {
	return utility.DBDelById(dao.Article.Ctx(ctx), utility.DBDelByIdInput{Where: id})
}
