package system_master

import "github.com/gogf/gf/v2/os/gtime"

// CreateRoleInput
// @Description: 添加管理员角色
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-07 14:40:24
type CreateRoleInput struct {
	Name   string
	Remark string
}

// GetRoleByIdOutput
// @Description: 按ID获取管理员角色信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-16 14:29:23
type GetRoleByIdOutput struct {
	Id     uint16
	Name   string
	Remark string
}

// ModifyRoleByIdInput
// @Description: 按ID修改管理员角色
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-07 14:49:48
type ModifyRoleByIdInput struct {
	Id     uint16
	Name   string
	Remark string
}

// ListRoleInput
// @Description: 角色列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-07 15:35:59
type ListRoleInput struct {
	Page int
	Size int
}

// ListRoleOutput
// @Description: 角色列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-07 15:44:17
type ListRoleOutput struct {
	Id        uint16      `json:"id"`
	Name      string      `json:"name"`
	Remark    string      `json:"remark"`
	CreatedAt *gtime.Time `json:"created_at"`
	UpdatedAt *gtime.Time `json:"updated_at"`
}

type GetRoleAllOutput struct {
	Id   uint16 `json:"id"`
	Name string `json:"name"`
}

// CreateRoleRelationInput
// @Description: 绑定管理员角色信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-07 17:51:15
type CreateRoleRelationInput struct {
	AccountId uint64
	RoleId    uint16
}

// GetRoleRelationByIdOutput
// @Description: 按ID显示管理员角色绑定信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-16 11:36:05
type GetRoleRelationByIdOutput struct {
	Id        uint32
	AccountId uint64
	RoleId    uint16
}

// ModifyRoleRelationByIdInput
// @Description: 修改管理员角色绑定信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-07 17:52:18
type ModifyRoleRelationByIdInput struct {
	Id        uint32
	AccountId uint64
	RoleId    uint16
}

// ListRoleRelationInput
// @Description: 管理员角色信息绑定列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-07 17:53:31
type ListRoleRelationInput struct {
	Page int
	Size int
}

type ListRoleRelationOutput struct {
	Id        uint32      `json:"id"`
	AccountId uint64      `json:"account_id"`
	Name      string      `json:"name"`
	RoleId    uint16      `json:"role_id"`
	RoleName  string      `json:"role_name"`
	CreatedAt *gtime.Time `json:"created_at"`
	UpdatedAt *gtime.Time `json:"updated_at"`
}

// DeleteRoleRelationInput
// @Description: 删除管理员角色绑定信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-07 17:53:54
type DeleteRoleRelationInput struct {
	Id uint32
}

// CreateListPermissionInput
// @Description: 新建权限
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-09 16:36:31
type CreateListPermissionInput struct {
	Fid    uint32
	Name   string
	Module string
	Type   uint8
	Sort   uint32
	Remark string
}

// GetPermissionByIdOutput
// @Description: 按ID获取权限菜单信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-16 16:19:24
type GetPermissionByIdOutput struct {
	Id     uint32
	Fid    uint32
	Name   string
	Type   uint8
	Status uint8
	Sort   uint32
	Remark string
}

// ModifyPermissionByIdInput
// @Description: 按ID权限菜单信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-09 16:48:22
type ModifyPermissionByIdInput struct {
	Id     uint32
	Fid    uint32
	Name   string
	Status uint8
	Sort   uint32
	Remark string
}

// ListPermissionInput
// @Description: 权限菜单列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-09 16:48:54
type ListPermissionInput struct {
	Page int
	Size int
}

// ListPermissionOutput
// @Description: 权限菜单列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-15 14:35:26
type ListPermissionOutput struct {
	Id        uint32      `json:"id"`
	Fid       uint32      `json:"fid"`
	Name      string      `json:"name"`
	Status    uint8       `json:"status"`
	Remark    string      `json:"remark"`
	CreatedAt *gtime.Time `json:"created_at"`
	UpdatedAt *gtime.Time `json:"updated_at"`
}

// ModifyPermissionRelationInput
// @Description: 权限分配
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-13 16:20:05
type ModifyPermissionRelationInput struct {
	PermissionIds []uint16
	RoleId        uint16
}

// GetPermissionAllOutput
// @Description: 开启权限信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-15 14:35:02
type GetPermissionAllOutput struct {
	Id       uint32                    `json:"id"`
	Fid      uint32                    `json:"fid"`
	Name     string                    `json:"name"`
	Checked  uint8                     `json:"checked"`
	Children []*GetPermissionAllOutput `json:"children"`
}

type GetPermissionByRoleIdInput struct {
	RoleId uint16
}
