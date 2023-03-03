package manage

import "github.com/gogf/gf/v2/frame/g"

// AddVisitCategoryReq
// @Description: 添加访问类型
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-03 18:10:49
type AddVisitCategoryReq struct {
	g.Meta `path:"/visit_category/add" tags:"添加访问类型" method:"post" summary:"访问类型"`
	Title  string `p:"title" v:"required#访问类型标题不能为空"`
}
type AddVisitCategoryRes struct {
}

// EditVisitCategoryReq
// @Description: 编辑访问类型
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-03 18:10:58
type EditVisitCategoryReq struct {
	g.Meta `path:"/visit_category/edit" tags:"编辑访问类型" method:"put" summary:"编辑访问类型"`
	Id     uint16 `p:"id" v:"required|min:1|integer#ID不能为空|ID要大于等于1|ID只为正整数"`
	Title  string `p:"title" v:"required#访问类型标题不能为空"`
}
type EditVisitCategoryRes struct {
}

// DeleteVisitCategoryReq
// @Description: 删除访问类型缓存
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-03 20:51:34
type DeleteVisitCategoryReq struct {
	g.Meta `path:"/visit_category/delete" tags:"编辑访问类型" method:"delete" summary:"编辑访问类型"`
}
type DeleteVisitCategoryRes struct {
}

// ListVisitCategoryReq
// @Description: 访问类型列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-03 18:11:11
type ListVisitCategoryReq struct {
	g.Meta `path:"/visit_category/show" tags:"访问类型列表" method:"get" summary:"访问类型列表"`
	Page   int `p:"page" v:"min:1|integer#页码要大于等于1|页码只为正整数"`
	Size   int `p:"size" v:"min:1|integer#显示条数要大于等于1|显示条数只为正整数"`
}
type ListVisitCategoryRes struct {
}
