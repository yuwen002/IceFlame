// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UcSystemPermissionExcludeDao is the data access object for table uc_system_permission_exclude.
type UcSystemPermissionExcludeDao struct {
	table   string                           // table is the underlying table name of the DAO.
	group   string                           // group is the database configuration group name of current DAO.
	columns UcSystemPermissionExcludeColumns // columns contains all the column names of Table for convenient usage.
}

// UcSystemPermissionExcludeColumns defines and stores column names for table uc_system_permission_exclude.
type UcSystemPermissionExcludeColumns struct {
	Id        string //
	Name      string // 菜单名称
	Module    string // 对应的程序模块
	Uri       string // 对应的URI地址
	Remark    string // 备注信息
	CreatedAt string //
	UpdatedAt string //
}

// ucSystemPermissionExcludeColumns holds the columns for table uc_system_permission_exclude.
var ucSystemPermissionExcludeColumns = UcSystemPermissionExcludeColumns{
	Id:        "id",
	Name:      "name",
	Module:    "module",
	Uri:       "uri",
	Remark:    "remark",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewUcSystemPermissionExcludeDao creates and returns a new DAO object for table data access.
func NewUcSystemPermissionExcludeDao() *UcSystemPermissionExcludeDao {
	return &UcSystemPermissionExcludeDao{
		group:   "default",
		table:   "uc_system_permission_exclude",
		columns: ucSystemPermissionExcludeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UcSystemPermissionExcludeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UcSystemPermissionExcludeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UcSystemPermissionExcludeDao) Columns() UcSystemPermissionExcludeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UcSystemPermissionExcludeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UcSystemPermissionExcludeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UcSystemPermissionExcludeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
