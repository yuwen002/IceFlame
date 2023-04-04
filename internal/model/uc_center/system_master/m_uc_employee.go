package system_master

import "github.com/gogf/gf/v2/os/gtime"

// CreateEmployeeRoleInput
// @Description: 新建员工角色
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-01 23:53:33
type CreateEmployeeRoleInput struct {
	Name   string
	Remark string
}

// GetEmployeeRoleOutput
// @Description: 显示员工角色
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-01 23:56:52
type GetEmployeeRoleOutput struct {
	Id        uint16      `json:"id"`
	Name      string      `json:"name"`
	Remark    string      `json:"remark"`
	CreatedAt *gtime.Time `json:"created_at"`
	UpdatedAt *gtime.Time `json:"updated_at"`
}

// ModifyEmployeeRole
// @Description: 修改员工角色
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-02 00:07:33
type ModifyEmployeeRole struct {
	Id     uint16
	Name   string
	Remark string
}

// ListEmployeeRoleInput
// @Description: 员工角色列表
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-02 00:08:05
type ListEmployeeRoleInput struct {
	Page int
	Size int
}

// CreateEmployeeRoleRelationInput
// @Description: 新建员工角色绑定员工ID
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-02 01:16:05
type CreateEmployeeRoleRelationInput struct {
	AccountId uint64
	RoleId    uint16
}

// CreateEmployeeInput
// @Description: 新建员工
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-02 01:14:53
type CreateEmployeeInput struct {
	Password string
	Name     string
	Tel      string
}

// ModifyEmployeeInput
// @Description: 修改员工信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-04-04 17:41:21
type ModifyEmployeeInput struct {
	AccountId uint64
	Name      string
	Tel       string
}

type EmployeeOutput struct {
	Password     string
	Name         string
	Tel          string
	WxUnionId    string
	WxOpenId     string
	InviteCode   string
	InviteQrcode string
	RealNameType uint8
	Gender       uint8
	Avatar       string
	Nickname     string
}
