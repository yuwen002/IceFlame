// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"ice_flame/internal/model/uc_center/system_master"
)

type (
	IUcSystemMaster interface {
		ExistsUsername(ctx context.Context, username string) (code int32, message string, err error)
		ExistsTel(ctx context.Context, tel string) (code int32, message string, err error)
		AccountIdCastTel(ctx context.Context, accountId uint64) (code int32, message string, output interface{}, err error)
		UpdateAccountIdCastTel(ctx context.Context, accountId uint64, tel string) (code int32, message string, err error)
		GetLoginCountRKey(accountId string) string
		IncLoginCount(ctx context.Context, accountId string) (code int32, message string, err error)
		GetLoginCount(ctx context.Context, accountId string) (code int32, message string, err error)
		DelLoginCount(ctx context.Context, accountId string) (code int32, message string, err error)
		Register(ctx context.Context, in system_master.RegisterInput) (code int32, message string, err error)
		LoginTelPassword(ctx context.Context, in system_master.LoginTelPasswordInput) (code int32, message string, token string, err error)
		LoginUsernamePassword(ctx context.Context, in system_master.LoginUsernamePasswordInput) (code int32, message string, token string, err error)
		CreateSystemMaster(ctx context.Context, in system_master.CreateSystemMasterInput) (code int32, message string, err error)
		ModifyPasswordSelfById(ctx context.Context, in system_master.ModifyPasswordInput) (code int32, message string, err error)
		ListSystemMaster(ctx context.Context, in system_master.ListSystemMasterInput) (code int32, message string, out []system_master.ListSystemMasterOutput, err error)
		GetSystemMasterByAccountId(ctx context.Context, account uint64) (code int32, message string, output, err error)
		ModifySystemMasterByAccountId(ctx context.Context, in system_master.ModifySystemMasterByAccountIdInput) (code int32, message string, err error)
		ResetPasswordByAccountId(ctx context.Context, in system_master.ResetPasswordByAccountIdInput) (code int32, message string, err error)
		ModifyStatusByAccountId(ctx context.Context, in system_master.ModifyStatusByAccountIdInput) (code int32, message string, err error)
	}
	IUcSystemMasterAuth interface {
		CreateRole(ctx context.Context, in system_master.CreateRoleInput) (code int32, message string, err error)
		GetRoleById(ctx context.Context, id uint16) (code int32, message string, output system_master.GetRoleByIdOutput, err error)
		ModifyRoleById(ctx context.Context, in system_master.ModifyRoleByIdInput) (code int32, message string, err error)
		ListRole(ctx context.Context, in system_master.ListRoleInput) (code int32, message string, output []*system_master.ListRoleOutput, err error)
		GetRoleAll(ctx context.Context) (code int32, message string, output []*system_master.GetRoleAllOutput, err error)
		CreateRoleRelation(ctx context.Context, in system_master.CreateRoleRelationInput) (code int32, message string, err error)
		GetRoleRelationById(ctx context.Context, id uint32) (code int32, message string, output system_master.GetRoleRelationByIdOutput, err error)
		ModifyRoleRelationById(ctx context.Context, in system_master.ModifyRoleRelationByIdInput) (code int32, message string, err error)
		ListRoleRelation(ctx context.Context, in system_master.ListRoleRelationInput) (code int32, message string, output []*system_master.ListRoleRelationOutput, err error)
		DeleteRoleRelationById(ctx context.Context, in system_master.DeleteRoleRelationInput) (code int32, message string, err error)
		CreatePermission(ctx context.Context, in system_master.CreateListPermissionInput) (code int32, message string, err error)
		GetPermissionById(ctx context.Context, id uint32) (code int32, message string, output system_master.GetPermissionByIdOutput, err error)
		ModifyPermissionById(ctx context.Context, in system_master.ModifyPermissionByIdInput) (code int32, message string, err error)
		ListPermission(ctx context.Context, in system_master.ListPermissionInput) (code int32, message string, output []*system_master.ListPermissionOutput, err error)
		ModifyPermissionRelation(ctx context.Context, in system_master.ModifyPermissionRelationInput) (code int32, message string, err error)
		GetPermissionAll(ctx context.Context) (code int32, message string, output []*system_master.GetPermissionAllOutput, err error)
		GetPermissionByRoleId(ctx context.Context, in system_master.GetPermissionByRoleIdInput) (code int32, message string, output []*system_master.GetPermissionAllOutput, err error)
		GetMenuAll(ctx context.Context)
	}
	IUcSystemMasterVisitor interface {
		CreateVisitCategory(ctx context.Context, in system_master.CreateVisitCategoryInput) (code int32, message string, err error)
		ModifyVisitCategoryById(ctx context.Context, in system_master.ModifyVisitCategoryByIdInput) (code int32, message string, err error)
		ListVisitCategory(ctx context.Context, in system_master.ListVisitCategoryInput) (code int32, message string, output []*system_master.ListVisitCategoryOutput, err error)
		GetRCacheVisitCategory(ctx context.Context) (code int32, message string, output map[string]interface{}, err error)
		GetRCacheVisitCategoryById(id string) (code int32, message string, output map[string]interface{}, err error)
		DelRCacheVisitCategory() (code int32, message string, err error)
		DelRCacheVisitCategoryById(id string) (code int32, message string, err error)
		CreateVisitorLogs(ctx context.Context, in system_master.CreateVisitorLogsInput) (code int32, message string, err error)
		ListVisitorLogs(ctx context.Context, in system_master.ListVisitorLogsInput) (code int32, message string, output []*system_master.ListVisitorLogsOutput, err error)
	}
)

var (
	localUcSystemMaster        IUcSystemMaster
	localUcSystemMasterAuth    IUcSystemMasterAuth
	localUcSystemMasterVisitor IUcSystemMasterVisitor
)

func UcSystemMaster() IUcSystemMaster {
	if localUcSystemMaster == nil {
		panic("implement not found for interface IUcSystemMaster, forgot register?")
	}
	return localUcSystemMaster
}

func RegisterUcSystemMaster(i IUcSystemMaster) {
	localUcSystemMaster = i
}

func UcSystemMasterAuth() IUcSystemMasterAuth {
	if localUcSystemMasterAuth == nil {
		panic("implement not found for interface IUcSystemMasterAuth, forgot register?")
	}
	return localUcSystemMasterAuth
}

func RegisterUcSystemMasterAuth(i IUcSystemMasterAuth) {
	localUcSystemMasterAuth = i
}

func UcSystemMasterVisitor() IUcSystemMasterVisitor {
	if localUcSystemMasterVisitor == nil {
		panic("implement not found for interface IUcSystemMasterVisitor, forgot register?")
	}
	return localUcSystemMasterVisitor
}

func RegisterUcSystemMasterVisitor(i IUcSystemMasterVisitor) {
	localUcSystemMasterVisitor = i
}
