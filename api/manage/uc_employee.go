package manage

import "github.com/gogf/gf/v2/frame/g"

// AddEmployeeRoleReq
// @Description: 添加员工角色
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-02 00:17:48
type AddEmployeeRoleReq struct {
	g.Meta `path:"/employee/role/add" tags:"添加员工角色" method:"post" summary:"添加员工角色"`
	Name   string `p:"name" v:"required#员工角色名称不能为空"`
	Remark string `p:"remark"`
}
type AddEmployeeRoleRes struct{}

// GetEmployeeRoleReq
// @Description: 编辑员工角色
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-02 00:42:58
type GetEmployeeRoleReq struct {
	g.Meta `path:"/employee/role/edit" tags:"编辑员工角色" method:"get" summary:"编辑员工角色"`
	Id     uint16 `p:"id" v:"required|min:1|integer#ID不能为空|ID要大于等于1|ID只为正整数"`
}
type GetEmployeeRoleRes struct{}

// EditEmployeeRoleReq
// @Description: 编辑员工角色
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-02 00:43:05
type EditEmployeeRoleReq struct {
	g.Meta `path:"/employee/role/edit" tags:"编辑员工角色" method:"put" summary:"编辑员工角色"`
	Id     uint16 `p:"id" v:"required|min:1|integer#ID不能为空|ID要大于等于1|ID只为正整数"`
	Name   string `p:"name" v:"required#管理员角色名称不能为空"`
	Remark string `p:"remark"`
}
type EditEmployeeRoleRes struct{}

type ListEmployeeRoleReq struct {
	g.Meta `path:"/employee/role/show" tags:"员工角色列表" method:"get" summary:"员工角色列表"`
	Page   int `p:"page" v:"min:1|integer#页码要大于等于1|页码只为正整数"`
	Size   int `p:"size" v:"min:1|integer#显示条数要大于等于1|显示条数只为正整数"`
}
type ListEmployeeRoleRes struct{}
