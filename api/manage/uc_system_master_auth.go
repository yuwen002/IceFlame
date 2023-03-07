package manage

import "github.com/gogf/gf/v2/frame/g"

// AddRoleReq
// @Description: 添加管理员角色
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-07 16:15:34
type AddRoleReq struct {
	g.Meta `path:"/master/auth/add_role" tags:"添加管理员角色" method:"post" summary:"添加管理员角色"`
	Name   string `p:"name" v:"required#管理员角色名称不能为空"`
	Remark string `p:"remark"`
}
type AddRoleRes struct{}

// EditRoleReq
// @Description: 编辑管理员角色
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-07 16:15:52
type EditRoleReq struct {
	g.Meta `path:"/master/auth/edit_role" tags:"编辑管理员角色" method:"put" summary:"编辑管理员角色"`
	Id     uint16 `p:"id" v:"required|min:1|integer#ID不能为空|ID要大于等于1|ID只为正整数"`
	Name   string `p:"name" v:"required#管理员角色名称不能为空"`
	Remark string `p:"remark"`
}
type EditRoleRes struct{}

// ListRoleReq
// @Description: 管理员角色列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-07 16:15:57
type ListRoleReq struct {
	g.Meta `path:"/master/auth/show_role" tags:"管理员角色列表" method:"get" summary:"编辑管理员角色"`
	Page   int `p:"page" v:"min:1|integer#页码要大于等于1|页码只为正整数"`
	Size   int `p:"size" v:"min:1|integer#显示条数要大于等于1|显示条数只为正整数"`
}
type ListRoleRes struct{}
