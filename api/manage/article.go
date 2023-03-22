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
	Id          uint32 `p:"id" v:"required|min:0|integer#父类ID不能为空|ID要大于等于0|ID只为整数"`
	Title       string `p:"title" v:"required#标题不能为空"`
	Description string `p:"description"`
	Keyword     string `p:"keyword"`
	Content     string `p:"content"`
	Thumbnail   string `p:"thumbnail"`
	Click       uint32 `p:"click"`
	Status      int8   `p:"status"`
}
type EditSinglePageRes struct {
}

// ListSinglePageReq
// @Description: 单页信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-22 18:07:32
type ListSinglePageReq struct {
	g.Meta `path:"/article/single_page/show" tags:"单页信息列表" method:"put" summary:"单页信息列表"`
	Page   int `p:"page" v:"min:1|integer#页码要大于等于1|页码只为正整数"`
	Size   int `p:"size" v:"min:1|integer#显示条数要大于等于1|显示条数只为正整数"`
}
type ListSinglePageRes struct {
}
