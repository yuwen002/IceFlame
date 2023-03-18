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
type AddVisitCategoryRes struct{}

// GetEditVisitCategoryReq
// @Description: 编辑访问类型
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-18 16:36:00
type GetEditVisitCategoryReq struct {
	g.Meta `path:"/visit_category/edit" tags:"编辑访问类型" method:"get" summary:"编辑访问类型"`
	Id     uint16 `p:"id" v:"required|min:1|integer#ID不能为空|ID要大于等于1|ID只为正整数"`
}
type GetEditVisitCategoryRes struct{}

// EditVisitCategoryReq
// @Description: 编辑访问类型
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-03 18:10:58
type EditVisitCategoryReq struct {
	g.Meta `path:"/visit_category/edit" tags:"编辑访问类型" method:"put" summary:"编辑访问类型"`
	Id     uint16 `p:"id" v:"required|min:1|integer#ID不能为空|ID要大于等于1|ID只为正整数"`
	Title  string `p:"title" v:"required#访问类型标题不能为空"`
}
type EditVisitCategoryRes struct{}

// DeleteVisitCategoryReq
// @Description: 删除访问类型缓存
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-03 20:51:34
type DeleteVisitCategoryReq struct {
	g.Meta `path:"/visit_category/delete" tags:"删除访问类型缓存" method:"delete" summary:"编辑访问类型"`
}
type DeleteVisitCategoryRes struct{}

// ListVisitCategoryReq
// @Description: 访问类型列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-03 18:11:11
type ListVisitCategoryReq struct {
	g.Meta `path:"/visit_category/show" tags:"访问类型列表" method:"get" summary:"访问类型列表"`
	Page   int `p:"page" v:"min:1|integer#页码要大于等于1|页码只为正整数"`
	Size   int `p:"size" v:"min:1|integer#显示条数要大于等于1|显示条数只为正整数"`
}
type ListVisitCategoryRes struct{}

// ListVisitorLogsReq
// @Description:
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-04 14:14:53
type ListVisitorLogsReq struct {
	g.Meta        `path:"/visitor_logs/show" tags:"访问类型日志列表" method:"get" summary:"访问类型列表"`
	Id            uint64 `p:"account_id" v:"min:1|integer#ID要大于等于1|ID只为正整数"`
	OsCategory    int8   `p:"os_category" v:"in:1,2,3#系统访问类型错误"`
	VisitCategory int16  `p:"visit_category"  v:"min:1|integer#访问类型要大于等于1|访问类型只为正整数"`
	Page          int    `p:"page" v:"min:1|integer#页码要大于等于1|页码只为正整数"`
	Size          int    `p:"size" v:"min:1|integer#显示条数要大于等于1|显示条数只为正整数"`
}
type ListVisitorLogsRes struct{}
