// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"ice_flame/internal/dao/internal"
)

// internalUcSystemMasterRoleDao is internal type for wrapping internal DAO implements.
type internalUcSystemMasterRoleDao = *internal.UcSystemMasterRoleDao

// ucSystemMasterRoleDao is the data access object for table uc_system_master_role.
// You can define custom methods on it to extend its functionality as you wish.
type ucSystemMasterRoleDao struct {
	internalUcSystemMasterRoleDao
}

var (
	// UcSystemMasterRole is globally public accessible object for table uc_system_master_role operations.
	UcSystemMasterRole = ucSystemMasterRoleDao{
		internal.NewUcSystemMasterRoleDao(),
	}
)

// Fill with you ideas below.
