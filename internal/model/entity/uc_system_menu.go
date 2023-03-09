// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UcSystemMenu is the golang structure for table uc_system_menu.
type UcSystemMenu struct {
	Id        uint        `json:"id"         ` //
	Fid       uint        `json:"fid"        ` // 父级ID
	Name      string      `json:"name"       ` // 菜单名称
	Status    uint        `json:"status"     ` // 菜单状态（0=启用，停用）
	Remark    string      `json:"remark"     ` // 备注信息
	CreatedAt *gtime.Time `json:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updated_at" ` //
}
