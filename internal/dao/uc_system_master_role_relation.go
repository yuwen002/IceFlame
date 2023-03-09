// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"ice_flame/internal/dao/internal"
)

// internalUcSystemMasterRoleRelationDao is internal type for wrapping internal DAO implements.
type internalUcSystemMasterRoleRelationDao = *internal.UcSystemMasterRoleRelationDao

// ucSystemMasterRoleRelationDao is the data access object for table uc_system_master_role_relation.
// You can define custom methods on it to extend its functionality as you wish.
type ucSystemMasterRoleRelationDao struct {
	internalUcSystemMasterRoleRelationDao
}

var (
	// UcSystemMasterRoleRelation is globally public accessible object for table uc_system_master_role_relation operations.
	UcSystemMasterRoleRelation = ucSystemMasterRoleRelationDao{
		internal.NewUcSystemMasterRoleRelationDao(),
	}
)

// Fill with you ideas below.