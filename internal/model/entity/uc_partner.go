// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UcPartner is the golang structure for table uc_partner.
type UcPartner struct {
	Id            uint        `json:"id"             ` //
	AccountId     uint        `json:"account_id"     ` // 关联登入用户表ID
	Fid           uint        `json:"fid"            ` // 上级id
	TopId         uint        `json:"top_id"         ` // 顶级id
	LevelId       uint        `json:"level_id"       ` // 级别id
	Depth         uint        `json:"depth"          ` // 层级深度
	IsAgent       uint        `json:"is_agent"       ` // 是否是代理商:1=是,0=否
	Path          string      `json:"path"           ` // 层级路径
	FullPath      string      `json:"full_path"      ` //
	JoinStoreNum  uint        `json:"join_store_num" ` // 入驻商家数
	TeamNum       uint        `json:"team_num"       ` // 团队人数
	PromotionType uint        `json:"promotion_type" ` // 晋级方式:1系统自动,0手动
	CreatedAt     *gtime.Time `json:"created_at"     ` //
	UpdatedAt     *gtime.Time `json:"updated_at"     ` //
}
