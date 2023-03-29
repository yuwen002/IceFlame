// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UcEmployeeEwalletOrm is the golang structure for table uc_employee_ewallet_orm.
type UcEmployeeEwalletOrm struct {
	Id        uint        `json:"id"         ` //
	AccountId uint        `json:"account_id" ` // 关联登入ID
	UnionId   uint        `json:"union_id"   ` // 关联钱包充值、消费ID
	UnionType uint        `json:"union_type" ` // 关联类型（1.充值、0.消费）
	CreatedAt *gtime.Time `json:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updated_at" ` //
}