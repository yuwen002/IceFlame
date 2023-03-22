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
func (s *sArticle) ModifyChannelById() {
}
func (s *sArticle) ListChannel() {
}
func (s *sArticle) DeleteChannelById() {
}
