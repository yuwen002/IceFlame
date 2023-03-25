package manage

import "github.com/gogf/gf/v2/frame/g"

// AddSinglePageReq
// @Description:添加单页信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-22 11:53:20
type AddSinglePageReq struct {
	g.Meta      `path:"/article/single_page/add" tags:"添加单页信息" method:"post" summary:"添加单页信息"`
	Title       string `p:"title" v:"required#标题不能为空"`
	Description string `p:"description"`
	Keyword     string `p:"keyword"`
	Content     string `p:"content"`
	Thumbnail   string `p:"thumbnail"`
}
type AddSinglePageRes struct{}

// GetSinglePageReq
// @Description: 编辑单页信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-22 16:20:58
type GetSinglePageReq struct {
	g.Meta `path:"/article/single_page/edit" tags:"编辑单页信息" method:"get" summary:"编辑单页信息"`
	Id     uint32 `p:"id" v:"required|min:0|integer#父类ID不能为空|ID要大于等于0|ID只为整数"`
}
type GetSinglePageRes struct{}

// EditSinglePageReq
// @Description: 编辑单页信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-22 17:47:53
type EditSinglePageReq struct {
	g.Meta      `path:"/article/single_page/edit" tags:"编辑单页信息" method:"put" summary:"编辑单页信息"`
	Id          uint32 `p:"id" v:"required|min:0|integer#ID不能为空|ID要大于等于0|ID只为整数"`
	Title       string `p:"title" v:"required#标题不能为空"`
	Description string `p:"description"`
	Keyword     string `p:"keyword"`
	Content     string `p:"content"`
	Thumbnail   string `p:"thumbnail"`
	Click       uint32 `p:"click"`
	Status      int8   `p:"status" v:"required|in:0,1#频道状态不能为空|频道状态值错误"`
}
type EditSinglePageRes struct {
}

// ListSinglePageReq
// @Description: 单页信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-22 18:07:32
type ListSinglePageReq struct {
	g.Meta `path:"/article/single_page/show" tags:"单页信息列表" method:"get" summary:"单页信息列表"`
	Page   int `p:"page" v:"min:1|integer#页码要大于等于1|页码只为正整数"`
	Size   int `p:"size" v:"min:1|integer#显示条数要大于等于1|显示条数只为正整数"`
}
type ListSinglePageRes struct {
}

// DeleteSinglePageReq
// @Description: 删除单页信息
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-22 23:11:19
type DeleteSinglePageReq struct {
	g.Meta `path:"/article/single_page/delete" tags:"单页信息列表" method:"delete" summary:"单页信息列表"`
	Id     uint32 `p:"id" v:"required|min:0|integer#ID不能为空|ID要大于等于0|ID只为整数"`
}
type DeleteSinglePageRes struct {
}

// AddArticleChannelReq
// @Description: 添加频道信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 14:17:34
type AddArticleChannelReq struct {
	g.Meta `path:"/article/channel/add" tags:"添加频道信息" method:"post" summary:"添加频道信息"`
	Name   string `p:"name" v:"required#频道名称不能为空"`
	Remark string `p:"remark"`
	Sort   uint32 `p:"sort" v:"min:0|integer#排序要大于等于0|排序只为整数"`
}
type AddArticleChannelRes struct{}

// GetArticleChannelReq
// @Description: 编辑频道信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 14:42:30
type GetArticleChannelReq struct {
	g.Meta `path:"/article/channel/edit" tags:"编辑频道信息" method:"get" summary:"编辑频道信息"`
	Id     uint32 `p:"id" v:"required|min:0|integer#ID不能为空|ID要大于等于0|ID只为整数"`
}
type GetArticleChannelRes struct{}

// EditArticleChannelReq
// @Description: 编辑频道信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 16:00:13
type EditArticleChannelReq struct {
	g.Meta `path:"/article/channel/edit" tags:"编辑频道信息" method:"put" summary:"编辑频道信息"`
	Id     uint32 `p:"id" v:"required|min:0|integer#ID不能为空|ID要大于等于0|ID只为整数"`
	Name   string `p:"name" v:"required#频道名称不能为空"`
	Remark string `p:"remark"`
	Sort   uint32 `p:"sort" v:"min:0|integer#排序要大于等于0|排序只为整数"`
	Status int8   `p:"status" v:"required|in:0,1#频道状态不能为空|频道状态值错误"`
}
type EditArticleChannelRes struct{}

// ListArticleChannelReq
// @Description: 频道信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 16:00:47
type ListArticleChannelReq struct {
	g.Meta `path:"/article/channel/show" tags:"频道信息列表" method:"get" summary:"频道信息列表"`
	Page   int `p:"page" v:"min:1|integer#页码要大于等于1|页码只为正整数"`
	Size   int `p:"size" v:"min:1|integer#显示条数要大于等于1|显示条数只为正整数"`
}
type ListArticleChannelRes struct {
}

type DelArticleChannelReq struct {
	g.Meta `path:"/article/channel/delete" tags:"编辑频道信息" method:"delete" summary:"编辑频道信息"`
	Id     uint32 `p:"id" v:"required|min:0|integer#ID不能为空|ID要大于等于0|ID只为整数"`
}
type DelArticleChannelRes struct{}

// GetArticleChannelAllReq
// @Description: 获取所有频道信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-24 17:12:37
type GetArticleChannelAllReq struct {
	g.Meta `path:"/article/channel/get_all" tags:"所有频道信息列表" method:"get" summary:"所有频道信息列表"`
}
type GetArticleChannelAllRes struct{}

// AddArticleCategoryReq
// @Description: 添加分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 17:58:16
type AddArticleCategoryReq struct {
	g.Meta `path:"/article/category/add" tags:"添加分类信息" method:"post" summary:"添加分类信息"`
	Fid    uint32 `p:"id" v:"required|min:0|integer#父类ID不能为空|ID要大于等于0|ID只为整数"`
	Name   string `p:"title" v:"required#分类名称不能为空"`
	Remark string `p:"remark"`
	Sort   uint32 `p:"sort" v:"min:0|integer#排序要大于等于0|排序只为整数"`
}
type AddArticleCategoryRes struct{}

// GetArticleCategoryReq
// @Description: 获取分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 18:02:43
type GetArticleCategoryReq struct {
	g.Meta `path:"/article/category/edit" tags:"获取分类信息" method:"get" summary:"获取分类信息"`
	Id     uint32 `p:"id" v:"required|min:0|integer#ID不能为空|ID要大于等于0|ID只为整数"`
}
type GetArticleCategoryRes struct{}

// EditArticleCategoryReq
// @Description: 修改分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 18:04:15
type EditArticleCategoryReq struct {
	g.Meta `path:"/article/category/edit" tags:"修改分类信息" method:"put" summary:"编辑频道信息"`
	Id     uint32 `p:"id" v:"required|min:0|integer#ID不能为空|ID要大于等于0|ID只为整数"`
	Fid    uint32 `p:"fid" v:"required|min:0|integer#父类ID不能为空|ID要大于等于0|ID只为整数"`
	Name   string `p:"title" v:"required#频道名称不能为空"`
	Remark string `p:"remark"`
	Sort   uint32 `p:"sort" v:"min:0|integer#排序要大于等于0|排序只为整数"`
	Status int8   `p:"status" v:"required|in:0,1#频道状态不能为空|频道状态值错误"`
}
type EditArticleCategoryRes struct{}

// ListArticleCategoryReq
// @Description: 分类列表
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-23 23:14:28
type ListArticleCategoryReq struct {
	g.Meta `path:"/article/category/show" tags:"分类列表" method:"get" summary:"分类列表"`
	Page   int `p:"page" v:"min:1|integer#页码要大于等于1|页码只为正整数"`
	Size   int `p:"size" v:"min:1|integer#显示条数要大于等于1|显示条数只为正整数"`
}
type ListArticleCategoryRes struct{}

// DelArticleCategoryReq
// @Description: 删除分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-24 00:03:54
type DelArticleCategoryReq struct {
	g.Meta `path:"/article/category/edit" tags:"删除分类信息" method:"delete" summary:"删除分类信息"`
	Id     uint32 `p:"id" v:"required|min:0|integer#ID不能为空|ID要大于等于0|ID只为整数"`
}
type DelArticleCategoryRes struct{}

// GetArticleCategoryAllReq
// @Description: 所有分类信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-24 17:38:51
type GetArticleCategoryAllReq struct {
	g.Meta `path:"/article/category/get_all" tags:"所有分类信息列表" method:"get" summary:"所有分类信息列表"`
}
type GetArticleCategoryAllRes struct{}

// AddArticleReq
// @Description: 添加文章
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-25 13:48:43
type AddArticleReq struct {
	g.Meta      `path:"/article/add" tags:"添加文章" method:"post" summary:"添加文章"`
	CategoryId  uint32 `p:"category_id" v:"required|min:0|integer#分类ID不能为空|分类ID要大于等于0|分类ID只为整数"`
	ChannelId   uint32 `p:"channel_id" v:"required|min:0|integer#频道ID不能为空|频道ID要大于等于0|频道ID只为整数"`
	Title       string `p:"title" v:"required#标题不能为空"`
	Keyword     string `p:"keyword"`
	Description string `p:"description"`
	Link        string `p:"link"`
	Author      string `p:"author"`
	Tags        string `p:"tags"`
	PubDate     string `p:"pub_date"`
	Summary     string `p:"summary"`
	Content     string `p:"content"`
	Thumbnail   string `p:"thumbnail"`
	Click       uint32 `p:"click"`
}
type AddArticleRes struct{}

// GetArticleReq
// @Description:
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-25 14:36:27
type GetArticleReq struct {
	g.Meta `path:"/article/edit" tags:"编辑文章" method:"get" summary:"编辑文章"`
	Id     uint32 `p:"id" v:"required|min:0|integer#ID不能为空|ID要大于等于0|ID只为整数"`
}
type GetArticleRes struct{}

// EditArticleReq
// @Description: 提交编辑文章
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-25 16:05:09
type EditArticleReq struct {
	g.Meta      `path:"/article/edit" tags:"编辑文章" method:"put" summary:"编辑文章"`
	Id          uint32 `p:"id" v:"required|min:0|integer#ID不能为空|ID要大于等于0|ID只为整数"`
	CategoryId  uint32 `p:"category_id" v:"required|min:0|integer#分类ID不能为空|分类ID要大于等于0|分类ID只为整数"`
	ChannelId   uint32 `p:"channel_id" v:"required|min:0|integer#频道ID不能为空|频道ID要大于等于0|频道ID只为整数"`
	Title       string `p:"title" v:"required#标题不能为空"`
	Keyword     string `p:"keyword"`
	Description string `p:"description"`
	Link        string `p:"link"`
	Author      string `p:"author"`
	Tags        string `p:"tags"`
	PubDate     string `p:"pub_date"`
	Summary     string `p:"summary"`
	Content     string `p:"content"`
	Thumbnail   string `p:"thumbnail"`
	Click       uint32 `p:"click"`
	Status      int8   `p:"status"`
}
type EditArticleRes struct{}

// ListArticleReq
// @Description: 文章信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-25 16:49:32
type ListArticleReq struct {
	g.Meta `path:"/article/show" tags:"文章信息列表" method:"get" summary:"文章信息列表"`
	Page   int `p:"page" v:"min:1|integer#页码要大于等于1|页码只为正整数"`
	Size   int `p:"size" v:"min:1|integer#显示条数要大于等于1|显示条数只为正整数"`
}
type ListArticleRes struct{}

// DelArticleReq
// @Description: 删除文章
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-25 17:26:15
type DelArticleReq struct {
	g.Meta `path:"/article/delete" tags:"删除文章" method:"delete" summary:"删除文章"`
	Id     uint32 `p:"id" v:"required|min:0|integer#ID不能为空|ID要大于等于0|ID只为整数"`
}
type DelArticleRes struct{}
