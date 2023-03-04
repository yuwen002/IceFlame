// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UcSystemMasterVisitCategory is the golang structure of table uc_system_master_visit_category for DAO operations like Where/Data.
type UcSystemMasterVisitCategory struct {
	g.Meta    `orm:"table:uc_system_master_visit_category, do:true"`
	Id        interface{} // 自增ID
	Title     interface{} // 访问类型标题
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}