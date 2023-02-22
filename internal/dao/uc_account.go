// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"ice_flame/internal/dao/internal"
)

// internalUcAccountDao is internal type for wrapping internal DAO implements.
type internalUcAccountDao = *internal.UcAccountDao

// ucAccountDao is the data access object for table uc_account.
// You can define custom methods on it to extend its functionality as you wish.
type ucAccountDao struct {
	internalUcAccountDao
}

var (
	// UcAccount is globally public accessible object for table uc_account operations.
	UcAccount = ucAccountDao{
		internal.NewUcAccountDao(),
	}
)

// Fill with you ideas below.