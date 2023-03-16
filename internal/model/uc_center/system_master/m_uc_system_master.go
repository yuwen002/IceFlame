package system_master

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// RegisterInput
// @Description: 创建新用户
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-22 16:06:35
type RegisterInput struct {
	Password string
	Tel      string
	Name     string
}

// LoginTelPasswordInput
// @Description: 用户电话密码登入
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-02-22 23:42:24
type LoginTelPasswordInput struct {
	Tel      string
	Password string
}

// LoginUsernamePasswordInput
// @Description:用户名密码登入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-23 16:20:59
type LoginUsernamePasswordInput struct {
	Username string
	Password string
}

// CreateSystemMasterInput
// @Description: 后台管理员新建用户
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-23 16:56:11
type CreateSystemMasterInput struct {
	Username string
	Password string
	Tel      string
	Name     string
}

// ModifyPasswordInput
// @Description: 修改用户密码
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-24 13:56:52
type ModifyPasswordInput struct {
	OldPassword string
	NewPassword string
}

// ListSystemMasterInput
// @Description: 管理员数据列表
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-02-24 23:46:23
type ListSystemMasterInput struct {
	Page int
	Size int
}

// ListSystemMasterOutput
// @Description:
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-02-26 01:43:35
type ListSystemMasterOutput struct {
	Id           uint64     `json:"id"`
	AccountId    uint64     `json:"account_id"`
	Username     string     `json:"username"`
	Name         string     `json:"name"`
	Tel          string     `json:"tel"`
	Status       uint8      `json:"status"`
	RealNameType string     `json:"real_name_type"`
	CreatedAt    gtime.Time `json:"created_at"`
	UpdatedAt    gtime.Time `json:"updated_at"`
}

// UcSystemMaster
// @Description: 管理员用户表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-25 17:47:49
type UcSystemMaster struct {
	gmeta.Meta `orm:"table:uc_system_master"`
	Id         uint64     `json:"id"`
	AccountId  uint64     `json:"account_id"`
	Name       string     `json:"name"`
	Tel        string     `json:"tel"`
	UcAccount  *UcAccount `orm:"with:id=account_id"`
}

// UcAccount
// @Description: 用户中心总表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-25 17:48:08
type UcAccount struct {
	gmeta.Meta   `orm:"table:uc_account"`
	Id           uint64     `json:"id"`
	Username     string     `json:"username"`
	RealNameType string     `json:"real_name_type"`
	Status       uint8      `json:"status"`
	CreatedAt    gtime.Time `json:"created_at"`
	UpdatedAt    gtime.Time `json:"updated_at"`
}

type GetSystemMasterByAccountIdOutput struct {
}

// ModifySystemMasterByAccountIdInput
// @Description:修改管理员信息
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-02-26 23:53:36
type ModifySystemMasterByAccountIdInput struct {
	AccountId uint64
	Tel       string
	Name      string
	Status    int8
}

// ResetPasswordByAccountIdInput
// @Description: 重置管理员密码
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-28 14:13:00
type ResetPasswordByAccountIdInput struct {
	AccountId uint64
	Password  string
}

// ModifyStatusByAccountIdInput
// @Description: 更改管理用户状态
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-28 15:48:38
type ModifyStatusByAccountIdInput struct {
	AccountId uint64
	Status    string
}

// IncLoginCountInput
// @Description: 登入限制
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-02-28 21:36:45
type IncLoginCountInput struct {
	Username string
	Tel      string
}
