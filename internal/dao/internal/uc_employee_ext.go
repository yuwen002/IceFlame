// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UcEmployeeExtDao is the data access object for table uc_employee_ext.
type UcEmployeeExtDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns UcEmployeeExtColumns // columns contains all the column names of Table for convenient usage.
}

// UcEmployeeExtColumns defines and stores column names for table uc_employee_ext.
type UcEmployeeExtColumns struct {
	Id               string //
	AccountId        string // 用户id
	EmployeeSex      string // 1=男2=女
	EmployeeAvatar   string // 头像
	EmployeeNickname string // 昵称
	CreatedAt        string //
	UpdatedAt        string //
}

// ucEmployeeExtColumns holds the columns for table uc_employee_ext.
var ucEmployeeExtColumns = UcEmployeeExtColumns{
	Id:               "id",
	AccountId:        "account_id",
	EmployeeSex:      "employee_sex",
	EmployeeAvatar:   "employee_avatar",
	EmployeeNickname: "employee_nickname",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

// NewUcEmployeeExtDao creates and returns a new DAO object for table data access.
func NewUcEmployeeExtDao() *UcEmployeeExtDao {
	return &UcEmployeeExtDao{
		group:   "default",
		table:   "uc_employee_ext",
		columns: ucEmployeeExtColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UcEmployeeExtDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UcEmployeeExtDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UcEmployeeExtDao) Columns() UcEmployeeExtColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UcEmployeeExtDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UcEmployeeExtDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UcEmployeeExtDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}