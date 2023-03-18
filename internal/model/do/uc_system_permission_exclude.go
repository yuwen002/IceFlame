// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UcSystemPermissionExclude is the golang structure of table uc_system_permission_exclude for DAO operations like Where/Data.
type UcSystemPermissionExclude struct {
	g.Meta    `orm:"table:uc_system_permission_exclude, do:true"`
	Id        interface{} //
	Name      interface{} // 菜单名称
	Module    interface{} // 对应的程序模块
	Uri       interface{} // 对应的URI地址
	Remark    interface{} // 备注信息
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
