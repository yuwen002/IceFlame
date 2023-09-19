package article

import (
	"context"
	"ice_flame/internal/dao"
	"ice_flame/internal/model/article"
	"ice_flame/internal/service"
	"ice_flame/utility"
)

var insSinglePage = sSinglePage{}

func SinglePage() *sSinglePage {
	return &insSinglePage
}

func init() {
	service.RegisterSinglePage(SinglePage())
}

type sSinglePage struct {
}

// Create
//
// @Title 新建单页信息
// @Description 新建单页信息
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-21 00:05:59
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sSinglePage) Create(ctx context.Context, in article.CreateSinglePageInput) (code int32, message string, err error) {
	return utility.DBInsert(dao.SinglePage.Ctx(ctx), utility.DBInsertInput{Data: in})
}

// GetById
//
// @Title 按ID获取单页信息
// @Description 按ID获取单页信息
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-21 00:13:28
// @receiver s
// @param ctx
// @param id
// @return code
// @return message
// @return output
// @return err
func (s *sSinglePage) GetById(ctx context.Context, id uint32) (code int32, message string, output *article.SinglePageOutput, err error) {
	code, message, err = utility.DBGetStructById(dao.SinglePage.Ctx(ctx), utility.DBGetByIdInput{Where: id}, &output)
	return code, message, output, err
}

// ModifyById
//
// @Title 按ID修改单页信息
// @Description 按ID修改单页分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-21 00:20:16
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sSinglePage) ModifyById(ctx context.Context, in article.ModifySinglePageInput) (code int32, message string, err error) {
	return utility.DBModifyById(dao.SinglePage.Ctx(ctx), utility.DBModifyByIdInput{
		Data:  in,
		Where: in.Id,
	})
}

// ModifyStatusById
//
// @Title 按ID修改单页状态信息
// @Description 按ID修改单页状态信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-09-19 14:07:25
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sSinglePage) ModifyStatusById(ctx context.Context, in article.ModifyStatusSinglePageInput) (code int32, message string, err error) {
	return utility.DBModifyById(dao.SinglePage.Ctx(ctx), utility.DBModifyByIdInput{
		Data:  in,
		Where: in.Id,
	})
}

// List
//
// @Title 单页信息列表
// @Description 单页信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-21 13:45:08
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return output
// @return err
func (s *sSinglePage) List(ctx context.Context, in article.ListSinglePageInput) (code int32, message string, output []*article.SinglePageOutput, err error) {
	code, message, err = utility.DBGetAllStructByWhere(dao.SinglePage.Ctx(ctx), utility.DBGetAllByWhereInput{
		Order:    "id asc",
		PageType: 1,
		DBPagination: utility.DBPagination{
			Page: in.Page,
			Size: in.Size,
		},
	}, &output)

	return code, message, output, err
}

// Delete
//
// @Title 按ID删除单页信息
// @Description 按ID删除单页信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-21 14:15:30
// @receiver s
// @param ctx
// @param id
// @return code
// @return message
// @return err
func (s *sSinglePage) Delete(ctx context.Context, id uint32) (code int32, message string, err error) {
	return utility.DBDelByWhere(dao.SinglePage.Ctx(ctx), utility.DBDelByWhereInput{Where: id})
}
