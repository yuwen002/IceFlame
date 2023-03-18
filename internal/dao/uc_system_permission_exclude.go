// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"ice_flame/internal/dao/internal"
)

// internalUcSystemPermissionExcludeDao is internal type for wrapping internal DAO implements.
type internalUcSystemPermissionExcludeDao = *internal.UcSystemPermissionExcludeDao

// ucSystemPermissionExcludeDao is the data access object for table uc_system_permission_exclude.
// You can define custom methods on it to extend its functionality as you wish.
type ucSystemPermissionExcludeDao struct {
	internalUcSystemPermissionExcludeDao
}

var (
	// UcSystemPermissionExclude is globally public accessible object for table uc_system_permission_exclude operations.
	UcSystemPermissionExclude = ucSystemPermissionExcludeDao{
		internal.NewUcSystemPermissionExcludeDao(),
	}
)

// Fill with you ideas below.
