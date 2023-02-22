// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		MiddlewareHandlerResponse(r *ghttp.Request)
	}
	IAuthMiddleware interface {
		MiddlewareAuthMaster(r *ghttp.Request)
	}
)

var (
	localMiddleware     IMiddleware
	localAuthMiddleware IAuthMiddleware
)

func AuthMiddleware() IAuthMiddleware {
	if localAuthMiddleware == nil {
		panic("implement not found for interface IAuthMiddleware, forgot register?")
	}
	return localAuthMiddleware
}

func RegisterAuthMiddleware(i IAuthMiddleware) {
	localAuthMiddleware = i
}

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}