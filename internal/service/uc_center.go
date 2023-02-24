// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	system_master "ice_flame/internal/model/uc_center"
)

type (
	IUcSystemMaster interface {
		ExistsUsername(ctx context.Context, username string) (code int32, message string, err error)
		ExistsTel(ctx context.Context, tel string) (code int32, message string, err error)
		Register(ctx context.Context, in system_master.RegisterInput) (code int32, message string, err error)
		LoginTelPassword(ctx context.Context, in system_master.LoginTelPasswordInput) (code int32, message string, token string, err error)
		LoginUsernamePassword(ctx context.Context, in system_master.LoginUsernamePasswordInput) (code int32, message string, token string, err error)
		CreateSystemMaster(ctx context.Context, in system_master.CreateSystemMasterInput) (code int32, message string, err error)
		ModifyPasswordSelfById(ctx context.Context, in system_master.ModifyPasswordInput) (code int32, message string, err error)
	}
)

var (
	localUcSystemMaster IUcSystemMaster
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
