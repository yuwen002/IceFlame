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

// CreateCategory
//
// @Title 新建单页分类
// @Description 新建单页分类
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-21 00:05:59
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sSinglePage) CreateCategory(ctx context.Context, in article.CreateSinglePageCategoryInput) (code int32, message string, err error) {
	return utility.DBInsert(dao.SinglePageCategory.Ctx(ctx), utility.DBInsertInput{Data: in})
}

// GetCategoryById
//
// @Title 按ID获取单页分类
// @Description 按ID获取单页分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-21 00:13:28
// @receiver s
// @param ctx
// @param id
// @return code
// @return message
// @return output
// @return err
func (s *sSinglePage) GetCategoryById(ctx context.Context, id uint16) (code int32, message string, output *article.SinglePageCategoryOutput, err error) {
	code, message, err = utility.DBGetStructById(dao.SinglePageCategory.Ctx(ctx), utility.DBGetByIdInput{Where: id}, &output)
	return code, message, output, err
}

// ModifyCategoryById
//
// @Title 按ID修改单页分类
// @Description 按ID修改单页分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-21 00:20:16
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sSinglePage) ModifyCategoryById(ctx context.Context, in article.ModifySinglePageCategoryInput) (code int32, message string, err error) {
	return utility.DBModifyById(dao.SinglePageCategory.Ctx(ctx), utility.DBModifyByIdInput{
		Data:  in,
		Where: in.Id,
	})
}

func (s *sSinglePage) ListCategoryById(ctx context.Context) {
}

func (s *sSinglePage) DeleteCategoryById(ctx context.Context, id uint16) {
}

func (s *sSinglePage) GetCategoryAll(ctx context.Context) {
}
