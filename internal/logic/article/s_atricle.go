package article

import (
	"context"
	"fmt"
	"ice_flame/internal/dao"
	"ice_flame/internal/model/article"
	"ice_flame/internal/service"
	"ice_flame/utility"
	"strings"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/frame/g"
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
	// 查找是否关联文章信息
	code, message, _, err = utility.DBGetMapById(dao.Article.Ctx(ctx), utility.DBGetByIdInput{Where: id})
	if err != nil {
		return -1, "", err
	}
	if code != 1 {
		return 1, "频道关联文章，不能删除相关频道", nil
	}
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
	// 查找是否关联文章信息
	code, message, _, err = utility.DBGetMapById(dao.Article.Ctx(ctx), utility.DBGetByIdInput{Where: id})
	if err != nil {
		return -1, "", err
	}
	if code != 1 {
		return 1, "分类关联文章，不能删除相关分类", nil
	}

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

// CreateTag
//
// @Title 新建标签
// @Description 新建标签
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-26 00:18:53
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sArticle) CreateTag(ctx context.Context, in article.CreateTagInput) (code int32, message string, err error) {
	return utility.DBInsert(dao.ArticleTag.Ctx(ctx), utility.DBInsertInput{Data: in})
}

// GetTagById
//
// @Title 显示标签
// @Description 显示标签
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-26 00:33:52
// @receiver s
// @param ctx
// @param id
// @return code
// @return message
// @return out
// @return err
func (s *sArticle) GetTagById(ctx context.Context, id uint32) (code int32, message string, out *article.TagOutput, err error) {
	code, message, err = utility.DBGetStructById(dao.ArticleTag.Ctx(ctx), utility.DBGetByIdInput{Where: id}, &out)
	return code, message, out, err
}

// ModifyTagById
//
// @Title 按ID修改标签
// @Description 按ID修改标签
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-26 00:56:56
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return err
func (s *sArticle) ModifyTagById(ctx context.Context, in article.ModifyTagInput) (code int32, message string, err error) {
	return utility.DBModifyById(dao.ArticleTag.Ctx(ctx), utility.DBModifyByIdInput{
		Data:  in,
		Where: in.Id,
	})
}

// ListTag
//
// @Title 标签列表
// @Description 标签列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-27 16:49:22
// @receiver s
// @param ctx
// @param in
// @return code
// @return message
// @return out
// @return err
func (s *sArticle) ListTag(ctx context.Context, in article.ListTagInput) (code int32, message string, out []*article.TagOutput, err error) {
	code, message, err = utility.DBGetAllStructByWhere(dao.ArticleTag.Ctx(ctx), utility.DBGetAllByWhereInput{
		Order:    "sort desc, id desc",
		PageType: 1,
		DBPagination: utility.DBPagination{
			Page: in.Page,
			Size: in.Size,
		},
	}, &out)

	return code, message, out, err
}

// DelTagById
//
// @Title 按ID删除标签
// @Description 按ID删除标签
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-26 01:05:22
// @receiver s
// @param ctx
// @param id
// @return code
// @return message
// @return err
func (s *sArticle) DelTagById(ctx context.Context, id uint32) (code int32, message string, err error) {
	// 查找是否关联文章信息
	code, message, _, err = utility.DBGetMapById(dao.ArticleTagOrm.Ctx(ctx), utility.DBGetByIdInput{
		Where: "tag_id=?",
		Args:  id,
	})
	if err != nil {
		return -1, "", err
	}
	if code != 1 {
		return 1, "标签关联文章，不能删除相关标签", nil
	}

	return utility.DBDelById(dao.ArticleTag.Ctx(ctx), utility.DBDelByIdInput{Where: id})
}

// GetTagAll
//
// @Title 所有显示标签
// @Description 所有显示标签
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-26 01:04:30
// @receiver s
// @param ctx
// @return code
// @return message
// @return out
// @return err
func (s *sArticle) GetTagAll(ctx context.Context) (code int32, message string, out []*article.TagOutput, err error) {
	code, message, err = utility.DBGetAllStructByWhere(dao.ArticleTag.Ctx(ctx), utility.DBGetAllByWhereInput{
		Where: "status=0",
		Order: "sort desc, id desc",
	}, &out)

	return code, message, out, err
}

// GetMapTags
//
// @Title 获取所有标签map
// @Description
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-28 23:57:28
// @receiver s
// @param ctx
// @param ids
// @return code
// @return message
// @return out
// @return err
func (s *sArticle) GetMapTags(ctx context.Context, ids []uint32) (code int32, message string, out map[string]string, err error) {
	code, message, output, err := utility.DBGetAllMapByWhere(dao.ArticleTag.Ctx(ctx), utility.DBGetAllByWhereInput{Where: "status = 0"})
	if code != 0 {
		return code, message, nil, err
	}

	tagMap := utility.MapStrStr(utility.MapsFromColumns(output, "id", "name"))

	return 0, "转换成功", tagMap, nil
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
	var tagIds []string
	// 给标签加标记
	if in.Tags != "" {
		// 提取标签Id
		tagIds = strings.Split(in.Tags, ",")
		in.Tags = "," + in.Tags + ","
	}
	code, message, id, err := utility.DBInsertAndGetId(dao.Article.Ctx(ctx), utility.DBInsertInput{Data: in})
	if code != 0 {
		return code, message, err
	}

	// 添加ORM标签关联表
	if len(tagIds) > 0 {
		insertData := make([]map[string]interface{}, len(tagIds))
		// 写入数据
		for index := range tagIds {
			insertData[index] = g.Map{
				"tag_id":     tagIds[index],
				"article_id": id,
			}
		}
		_, _, _ = utility.DBInsert(dao.ArticleTagOrm.Ctx(ctx), utility.DBInsertInput{Data: insertData})
	}

	return 1, "发布成功", nil
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
	// 处理标签ID
	if output.Tags != "" {
		output.Tags = strings.Trim(output.Tags, ",")

	}

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
	if in.Tags != "" {
		// 提取标签Id
		tagIds := strings.Split(in.Tags, ",")
		length := len(tagIds)
		// 给标签加标记
		if length > 0 {
			in.Tags = "," + in.Tags + ","
		}

		code, _, tagOrm, err := utility.DBGetAllMapByWhere(dao.ArticleTagOrm.Ctx(ctx), utility.DBGetAllByWhereInput{
			Where: "article_id=?",
			Args:  in.Id,
		})
		if err != nil {
			return -1, "", err
		}

		if code == 1 {
			insertData := make([]map[string]interface{}, length)
			// 添加ORM标签关联表
			for index := range tagIds {
				insertData[index] = g.Map{
					"tag_id":     tagIds[index],
					"article_id": in.Id,
				}
			}
			// 数据写入
			_, _, _ = utility.DBInsert(dao.ArticleTagOrm.Ctx(ctx), utility.DBInsertInput{Data: insertData})
		}

		// 数据转换成sting数组
		fmt.Println(tagOrm)
		oldTagIds := utility.ArrayColumnCast[interface{}, string](tagOrm, "tag_id", func(v interface{}) string {
			return gconv.String(v)
		})

		// 需要删除的数据
		delTagIds := utility.ArrayDiff[string](oldTagIds, tagIds)
		if len(delTagIds) > 0 {
			code, message, err = utility.DBDelByWhere(dao.ArticleTagOrm.Ctx(ctx), utility.DBDelByWhereInput{
				Where: "tag_id in (?) and article_id =?",
				Args:  g.Slice{delTagIds, in.Id},
			})
			if code != 0 {
				return code, message, err
			}
		}

		// 需要添加的数据
		insertTagIds := utility.ArrayDiff[string](tagIds, oldTagIds)
		if len(insertTagIds) != 0 {
			insertData := make([]map[string]interface{}, len(insertTagIds))
			// 添加ORM标签关联表
			for index := range insertTagIds {
				insertData[index] = g.Map{
					"tag_id":     insertTagIds[index],
					"article_id": in.Id,
				}
			}
			_, _, _ = utility.DBInsert(dao.ArticleTagOrm.Ctx(ctx), utility.DBInsertInput{Data: insertData})
		}
	}

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
	_, _, err = utility.DBDelByWhere(dao.ArticleTagOrm.Ctx(ctx), utility.DBDelByWhereInput{
		Where: "article_id = ?",
		Args:  id,
	})
	if err != nil {
		return -1, "", err
	}

	return utility.DBDelById(dao.Article.Ctx(ctx), utility.DBDelByIdInput{Where: id})
}
