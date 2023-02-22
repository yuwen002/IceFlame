// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"ice_flame/internal/model/uc_center"
)

type (
	IUcSystemMaster interface {
		ExistsUsername(ctx context.Context, username string) (code int32, message string, err error)
		RegisterSystemMaster(ctx context.Context, in uc_center.RegisterSystemMasterInput) (code int32, message string, err error)
		LoginTelSystemMaster(ctx context.Context, in uc_center.LoginTelSystemMasterInput) (code int32, message string, token string, err error)
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
