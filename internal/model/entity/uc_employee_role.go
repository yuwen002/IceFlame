// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UcEmployeeRole is the golang structure for table uc_employee_role.
type UcEmployeeRole struct {
	Id        uint        `json:"id"         ` //
	AccountId uint        `json:"account_id" ` //
	RoleId    int         `json:"role_id"    ` //
	CreatedAt *gtime.Time `json:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updated_at" ` //
}
