// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"ice_flame/internal/dao/internal"
)

// internalUcEmployeeExtDao is internal type for wrapping internal DAO implements.
type internalUcEmployeeExtDao = *internal.UcEmployeeExtDao

// ucEmployeeExtDao is the data access object for table uc_employee_ext.
// You can define custom methods on it to extend its functionality as you wish.
type ucEmployeeExtDao struct {
	internalUcEmployeeExtDao
}

var (
	// UcEmployeeExt is globally public accessible object for table uc_employee_ext operations.
	UcEmployeeExt = ucEmployeeExtDao{
		internal.NewUcEmployeeExtDao(),
	}
)

// Fill with you ideas below.
