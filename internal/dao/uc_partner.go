// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"ice_flame/internal/dao/internal"
)

// internalUcPartnerDao is internal type for wrapping internal DAO implements.
type internalUcPartnerDao = *internal.UcPartnerDao

// ucPartnerDao is the data access object for table uc_partner.
// You can define custom methods on it to extend its functionality as you wish.
type ucPartnerDao struct {
	internalUcPartnerDao
}

var (
	// UcPartner is globally public accessible object for table uc_partner operations.
	UcPartner = ucPartnerDao{
		internal.NewUcPartnerDao(),
	}
)

// Fill with you ideas below.
