// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UcSystemMasterDao is the data access object for table uc_system_master.
type UcSystemMasterDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns UcSystemMasterColumns // columns contains all the column names of Table for convenient usage.
}

// UcSystemMasterColumns defines and stores column names for table uc_system_master.
type UcSystemMasterColumns struct {
	Id        string // 自增ID
	AccountId string // account id
	Name      string //
	Tel       string //
	CreatedAt string //
	UpdatedAt string //
}

// ucSystemMasterColumns holds the columns for table uc_system_master.
var ucSystemMasterColumns = UcSystemMasterColumns{
	Id:        "id",
	AccountId: "account_id",
	Name:      "name",
	Tel:       "tel",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewUcSystemMasterDao creates and returns a new DAO object for table data access.
func NewUcSystemMasterDao() *UcSystemMasterDao {
	return &UcSystemMasterDao{
		group:   "default",
		table:   "uc_system_master",
		columns: ucSystemMasterColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UcSystemMasterDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UcSystemMasterDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UcSystemMasterDao) Columns() UcSystemMasterColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UcSystemMasterDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UcSystemMasterDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UcSystemMasterDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
