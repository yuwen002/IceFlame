// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SinglePageCategoryDao is the data access object for table single_page_category.
type SinglePageCategoryDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns SinglePageCategoryColumns // columns contains all the column names of Table for convenient usage.
}

// SinglePageCategoryColumns defines and stores column names for table single_page_category.
type SinglePageCategoryColumns struct {
	Id        string //
	Name      string // 分类名称
	Remark    string // 备注信息
	Sort      string // 排序顺序
	CreatedAt string //
	UpdatedAt string //
}

// singlePageCategoryColumns holds the columns for table single_page_category.
var singlePageCategoryColumns = SinglePageCategoryColumns{
	Id:        "id",
	Name:      "name",
	Remark:    "remark",
	Sort:      "sort",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewSinglePageCategoryDao creates and returns a new DAO object for table data access.
func NewSinglePageCategoryDao() *SinglePageCategoryDao {
	return &SinglePageCategoryDao{
		group:   "default",
		table:   "single_page_category",
		columns: singlePageCategoryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SinglePageCategoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SinglePageCategoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SinglePageCategoryDao) Columns() SinglePageCategoryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SinglePageCategoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SinglePageCategoryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SinglePageCategoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
