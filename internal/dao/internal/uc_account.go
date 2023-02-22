// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UcAccountDao is the data access object for table uc_account.
type UcAccountDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns UcAccountColumns // columns contains all the column names of Table for convenient usage.
}

// UcAccountColumns defines and stores column names for table uc_account.
type UcAccountColumns struct {
	Id             string //
	IdentityCardId string // 关联用户身份证信息ID
	Username       string // 用户名
	PasswordHash   string // 用户密码
	Tel            string // 用户手机号码
	Status         string // 用户状态(0.启用1.停用)
	RealNameType   string // 实名状态1=未实名2=已上传未审核3=审核驳回4=实名成功
	CreatedAt      string //
	UpdatedAt      string //
}

// ucAccountColumns holds the columns for table uc_account.
var ucAccountColumns = UcAccountColumns{
	Id:             "id",
	IdentityCardId: "identity_card_id",
	Username:       "username",
	PasswordHash:   "password_hash",
	Tel:            "tel",
	Status:         "status",
	RealNameType:   "real_name_type",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewUcAccountDao creates and returns a new DAO object for table data access.
func NewUcAccountDao() *UcAccountDao {
	return &UcAccountDao{
		group:   "default",
		table:   "uc_account",
		columns: ucAccountColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UcAccountDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UcAccountDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UcAccountDao) Columns() UcAccountColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UcAccountDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UcAccountDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UcAccountDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}