// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"focus/app/dao/internal"
)

// userRoleAuthDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type userRoleAuthDao struct {
	*internal.UserRoleAuthDao
}

var (
	// UserRoleAuth is globally public accessible object for table {TplTableName} operations.
	UserRoleAuth = &userRoleAuthDao{
		internal.UserRoleAuth,
	}
)

// Fill with you ideas below.