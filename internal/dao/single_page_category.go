// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"ice_flame/internal/dao/internal"
)

// internalSinglePageCategoryDao is internal type for wrapping internal DAO implements.
type internalSinglePageCategoryDao = *internal.SinglePageCategoryDao

// singlePageCategoryDao is the data access object for table single_page_category.
// You can define custom methods on it to extend its functionality as you wish.
type singlePageCategoryDao struct {
	internalSinglePageCategoryDao
}

var (
	// SinglePageCategory is globally public accessible object for table single_page_category operations.
	SinglePageCategory = singlePageCategoryDao{
		internal.NewSinglePageCategoryDao(),
	}
)

// Fill with you ideas below.
