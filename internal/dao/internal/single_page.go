// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SinglePageDao is the data access object for table single_page.
type SinglePageDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns SinglePageColumns // columns contains all the column names of Table for convenient usage.
}

// SinglePageColumns defines and stores column names for table single_page.
type SinglePageColumns struct {
	Id          string //
	Title       string // 页面标题
	Description string // 页面描述
	Keyword     string // 页面关键字
	Content     string // 页面内容
	Thumbnail   string // 缩略图
	Click       string // 点击量
	Status      string // 显示状态（0=显示，1=隐藏）
	CreatedAt   string //
	UpdatedAt   string //
}

// singlePageColumns holds the columns for table single_page.
var singlePageColumns = SinglePageColumns{
	Id:          "id",
	Title:       "title",
	Description: "description",
	Keyword:     "keyword",
	Content:     "content",
	Thumbnail:   "thumbnail",
	Click:       "click",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewSinglePageDao creates and returns a new DAO object for table data access.
func NewSinglePageDao() *SinglePageDao {
	return &SinglePageDao{
		group:   "default",
		table:   "single_page",
		columns: singlePageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SinglePageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SinglePageDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SinglePageDao) Columns() SinglePageColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SinglePageDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SinglePageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SinglePageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}