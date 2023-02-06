// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"ice_flame/internal/model"
)

type (
	ITestDB interface {
		TestInsert(ctx context.Context, in model.TestDBInsertInput) (code int32, message string, err error)
		TestInsertAndGetId(ctx context.Context, in model.TestDBInsertInput) (code int32, message string, lastInsertId int64, err error)
		TestGetStructById(ctx context.Context, in model.TestGetByIdInput, output interface{}) (code int32, message string, err error)
		TestGetStructGCacheById(ctx context.Context, in model.TestGetByIdInput, output interface{}) (code int32, message string, err error)
		TestGetStructRCacheById(ctx context.Context, in model.TestGetByIdInput, output interface{}) (code int32, message string, err error)
		TestGetMapById(ctx context.Context, in model.TestGetByIdInput) (code int32, message string, output interface{}, err error)
		TestGetMapGCacheById(ctx context.Context, in model.TestGetByIdInput) (code int32, message string, output interface{}, err error)
		TestGetMapRCacheById(ctx context.Context, in model.TestGetByIdInput) (code int32, message string, output interface{}, err error)
		TestGetStructByIds(ctx context.Context, in model.TestGetByIdsInput, output interface{}) (code int32, message string, err error)
		TestGetMapByIds(ctx context.Context, in model.TestGetByIdsInput) (code int32, message string, output interface{}, err error)
		TestGetOneStructByWhere(ctx context.Context, in model.TestGetOneByWhereInput, output interface{}) (code int32, message string, err error)
		TestGetOneMapByWhere(ctx context.Context, in model.TestGetOneByWhereInput) (code int32, message string, output interface{}, err error)
		TestDBGetAllStructByWhereInput(ctx context.Context, in model.TestGetAllByWhereInput, output []interface{}) (code int32, message string, err error)
	}
)

var (
	localTestDB ITestDB
)

func TestDB() ITestDB {
	if localTestDB == nil {
		panic("implement not found for interface ITestDB, forgot register?")
	}
	return localTestDB
}

func RegisterTestDB(i ITestDB) {
	localTestDB = i
}
