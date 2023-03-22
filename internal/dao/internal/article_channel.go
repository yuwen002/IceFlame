// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ArticleChannelDao is the data access object for table article_channel.
type ArticleChannelDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns ArticleChannelColumns // columns contains all the column names of Table for convenient usage.
}

// ArticleChannelColumns defines and stores column names for table article_channel.
type ArticleChannelColumns struct {
	Id        string //
	Name      string // 栏目名称
	Remark    string // 备注信息
	Sort      string // 排序顺序
	Status    string // 显示状态（0=显示，1=隐藏）
	CreatedAt string //
	UpdatedAt string //
}

// articleChannelColumns holds the columns for table article_channel.
var articleChannelColumns = ArticleChannelColumns{
	Id:        "id",
	Name:      "name",
	Remark:    "remark",
	Sort:      "sort",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewArticleChannelDao creates and returns a new DAO object for table data access.
func NewArticleChannelDao() *ArticleChannelDao {
	return &ArticleChannelDao{
		group:   "default",
		table:   "article_channel",
		columns: articleChannelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ArticleChannelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ArticleChannelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ArticleChannelDao) Columns() ArticleChannelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ArticleChannelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ArticleChannelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ArticleChannelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
