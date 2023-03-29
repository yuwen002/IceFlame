// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UcEmployeeEwallet is the golang structure for table uc_employee_ewallet.
type UcEmployeeEwallet struct {
	Id             uint        `json:"id"              ` //
	AccountId      uint        `json:"account_id"      ` // 关联员工ID
	Money          uint        `json:"money"           ` // 账户余额（单位：里）（冻结金额）
	Balance        uint        `json:"balance"         ` // 上次入账余额（单位：里）
	TotalMoney     int         `json:"total_money"     ` // 账户累计金额（单位：里）
	UnfreezeAmount int         `json:"unfreeze_amount" ` // 解冻金额
	HashEwallet    string      `json:"hash_ewallet"    ` // hash钱包
	EntryId        uint        `json:"entry_id"        ` // 入账ID
	WithdrawId     uint        `json:"withdraw_id"     ` // 提现ID
	WxOpenid       string      `json:"wx_openid"       ` // 微信提现openid
	CreatedAt      *gtime.Time `json:"created_at"      ` //
	UpdatedAt      *gtime.Time `json:"updated_at"      ` //
}