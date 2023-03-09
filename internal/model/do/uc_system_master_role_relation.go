// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UcSystemMasterRoleRelation is the golang structure of table uc_system_master_role_relation for DAO operations like Where/Data.
type UcSystemMasterRoleRelation struct {
	g.Meta    `orm:"table:uc_system_master_role_relation, do:true"`
	Id        interface{} //
	AccountId interface{} // 管理员ID
	RoleId    interface{} // 身份ID
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}