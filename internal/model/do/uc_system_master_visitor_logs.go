// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UcSystemMasterVisitorLogs is the golang structure of table uc_system_master_visitor_logs for DAO operations like Where/Data.
type UcSystemMasterVisitorLogs struct {
	g.Meta        `orm:"table:uc_system_master_visitor_logs, do:true"`
	Id            interface{} // 自增ID
	AccountId     interface{} // 关联用户ID
	OsCategory    interface{} // 系统访问类型（1=web端，2=android端，3=IOS端）
	VisitCategory interface{} // 访问类型
	UnionId       interface{} // 关联ID
	Description   interface{} // 访问信息描述
	IpLong        []byte      //
	Ip            interface{} //
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
}
