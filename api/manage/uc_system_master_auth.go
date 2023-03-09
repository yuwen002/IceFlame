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
	g.Meta `path:"/master/auth/show_role" tags:"管理员角色列表" method:"get" summary:"管理员角色列表"`
	Page   int `p:"page" v:"min:1|integer#页码要大于等于1|页码只为正整数"`
	Size   int `p:"size" v:"min:1|integer#显示条数要大于等于1|显示条数只为正整数"`
}
type ListRoleRes struct{}

// AddRoleRelationReq
// @Description: 添加管理员绑定角色信息
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-08 00:19:37
type AddRoleRelationReq struct {
	g.Meta    `path:"/master/auth/add_role_relation" tags:"加管理员绑定角色信息" method:"post" summary:"加管理员绑定角色信息"`
	AccountId uint64 `p:"account_id" v:"required|min:1|integer#管理员ID不能为空|管理员ID要大于等于1|管理员ID只为正整数"`
	RoleId    uint16 `p:"role_id" v:"required|min:1|integer#角色ID不能为空|角色ID要大于等于1|ID只为正整数"`
}
type AddRoleRelationRes struct{}

// EditRoleRelationReq
// @Description: 编辑管理员绑定角色信息
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-08 00:20:03
type EditRoleRelationReq struct {
	g.Meta    `path:"/master/auth/edit_role_relation" tags:"编辑管理员绑定角色信息" method:"put" summary:"编辑管理员绑定角色信息"`
	Id        uint32 `p:"id" v:"required|min:1|integer#ID不能为空|ID要大于等于1|ID只为正整数"`
	AccountId uint64 `p:"account_id" v:"required|min:1|integer#管理员ID不能为空|管理员ID要大于等于1|管理员ID只为正整数"`
	RoleId    uint16 `p:"role_id" v:"required|min:1|integer#角色ID不能为空|角色ID要大于等于1|ID只为正整数"`
}
type EditRoleRelationRes struct{}

// ListRoleRelationReq
// @Description:管理员绑定角色信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-08 00:20:15
type ListRoleRelationReq struct {
	g.Meta `path:"/master/auth/show_role_relation" tags:"管理员绑定角色信息列表" method:"get" summary:"管理员绑定角色信息"`
	Page   int `p:"page" v:"min:1|integer#页码要大于等于1|页码只为正整数"`
	Size   int `p:"size" v:"min:1|integer#显示条数要大于等于1|显示条数只为正整数"`
}
type ListRoleRelationRes struct{}

// DeleteRoleRelationReq
// @Description: 删除管理员绑定角色信息
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-08 00:20:31
type DeleteRoleRelationReq struct {
	g.Meta `path:"/master/auth/delete_role_relation" tags:"删除管理员绑定角色信息" method:"delete" summary:"删除管理员绑定角色信息"`
	Id     uint32 `p:"id" v:"required|min:1|integer#ID不能为空|ID要大于等于1|ID只为正整数"`
}
type DeleteRoleRelationRes struct{}