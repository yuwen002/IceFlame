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

// GetEditRoleReq
// @Description: 获取角色编辑信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-16 14:38:08
type GetEditRoleReq struct {
	g.Meta `path:"/master/auth/edit_role" tags:"编辑管理员角色" method:"get" summary:"编辑管理员角色"`
	Id     uint16 `p:"id" v:"required|min:1|integer#ID不能为空|ID要大于等于1|ID只为正整数"`
}
type GetEditRoleRes struct{}

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
type GetListRoleRes struct{}

// GetListRoleReq
// @Description: 所有管理员角色列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-15 18:19:25
type GetListRoleReq struct {
	g.Meta `path:"/master/auth/get_role" tags:"管理员角色列表" method:"get" summary:"管理员角色列表"`
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

// GetEditRoleRelationReq
// @Description: 编辑管理员绑定角色信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-16 11:53:42
type GetEditRoleRelationReq struct {
	g.Meta `path:"/master/auth/edit_role_relation" tags:"编辑管理员绑定角色信息" method:"get" summary:"编辑管理员绑定角色信息"`
	Id     uint32 `p:"id" v:"required|min:1|integer#ID不能为空|ID要大于等于1|ID只为正整数"`
}
type GetEditRoleRelationRes struct {
}

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

// AddPermissionReq
// @Description: 添加权限信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-09 17:44:16
type AddPermissionReq struct {
	g.Meta `path:"/master/auth/add_permission" tags:"添加权限" method:"post" summary:"添加权限"`
	Fid    uint32 `p:"fid" v:"required|min:0|integer#父类ID不能为空|ID要大于等于0|ID只为整数"`
	Name   string `p:"name" v:"required#权限名称不能为空"`
	Module string `p:"module"`
	Type   uint8  `p:"type" v:"in:1,2#请输入正确的菜单分"`
	Sort   uint32 `p:"sort" v:"min:0|integer#ID要大于等于0|ID只为整数"`
	Remark string `p:"remark"`
}
type AddPermissionRes struct{}

// GetEditPermissionReq
// @Description: 按ID获取权限信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-16 16:28:12
type GetEditPermissionReq struct {
	g.Meta `path:"/master/auth/edit_permission" tags:"修改权限" method:"get" summary:"修改权限"`
	Id     uint32 `p:"id" v:"required|min:0|integer#ID不能为空|ID要大于等于0|ID只为整数"`
}
type GetEditPermissionRes struct{}

// EditPermissionReq
// @Description: 修改权限信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-10 14:22:05
type EditPermissionReq struct {
	g.Meta `path:"/master/auth/edit_permission" tags:"修改权限" method:"put" summary:"修改权限"`
	ID     uint32 `p:"id" v:"required|min:0|integer#ID不能为空|ID要大于等于0|ID只为整数"`
	Fid    uint32 `p:"fid" v:"required|min:0|integer#父类ID不能为空|ID要大于等于0|ID只为整数"`
	Name   string `p:"name" v:"required#权限名称不能为空"`
	Module string `p:"module"`
	Type   uint8  `p:"type" v:"in:1,2#请输入正确的菜单分"`
	Status uint8  `p:"status" v:"required|in:0,1#权限状态不能为空|权限状态输入不正确"`
	Sort   uint32 `p:"fid" v:"min:0|integer#ID要大于等于0|ID只为整数"`
	Remark string `p:"remark"`
}
type EditPermissionRes struct{}

// ListPermissionReq
// @Description: 权限信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-13 00:15:38
type ListPermissionReq struct {
	g.Meta `path:"/master/auth/show_permission" tags:"权限列表" method:"get" summary:"权限列表"`
	Page   int `p:"page" v:"min:1|integer#页码要大于等于1|页码只为正整数"`
	Size   int `p:"size" v:"min:1|integer#显示条数要大于等于1|显示条数只为正整数"`
}
type ListPermissionRes struct{}

// EditPermissionRelationReq
// @Description: 按角色ID权限信息修改
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-15 17:14:04
type EditPermissionRelationReq struct {
	g.Meta        `path:"/master/auth/edit_permission_relation" tags:"权限信息修改" method:"put" summary:"权限信息修改"`
	PermissionIds []uint16 `p:"permission_ids" v:"required#权限ID不能为空"`
	RoleId        uint16   `p:"role_id" v:"required#角色ID不能为空"`
}
type EditPermissionRelationRes struct{}

// ListPermissionRelationReq
// @Description: 按角色ID权限信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-15 17:15:39
type ListPermissionRelationReq struct {
	g.Meta `path:"/master/auth/show_permission_relation" tags:"权限信息修改" method:"get" summary:"权限信息修改"`
	RoleId uint16 `p:"role_id" v:"required#角色ID不能为空"`
}
type ListPermissionRelationRes struct{}

// ListMenuReq
// @Description:获取菜单导航列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-17 15:13:18
type ListMenuReq struct {
	g.Meta `path:"/show_menu" tags:"权限信息修改" method:"get" summary:"权限信息修改"`
}
type ListMenuRes struct{}
