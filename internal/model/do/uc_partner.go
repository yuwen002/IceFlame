// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UcPartner is the golang structure of table uc_partner for DAO operations like Where/Data.
type UcPartner struct {
	g.Meta        `orm:"table:uc_partner, do:true"`
	Id            interface{} //
	AccountId     interface{} // 关联登入用户表ID
	Fid           interface{} // 上级id
	TopId         interface{} // 顶级id
	LevelId       interface{} // 级别id
	Depth         interface{} // 层级深度
	IsAgent       interface{} // 是否是代理商:1=是,0=否
	Path          interface{} // 层级路径
	FullPath      interface{} //
	JoinStoreNum  interface{} // 入驻商家数
	TeamNum       interface{} // 团队人数
	PromotionType interface{} // 晋级方式:1系统自动,0手动
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
}
