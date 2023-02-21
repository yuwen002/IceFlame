package test

import "github.com/gogf/gf/v2/frame/g"

type JCSetTokenReq struct {
	g.Meta `path:"/test/jwt/set_token" tags:"测试jwt" method:"get" summary:"测试jwt"`
}
type JCSetTokenRes struct {
}

type JCGetTokenReq struct {
	g.Meta `path:"/test/jwt/get_token" tags:"测试jwt" method:"get" summary:"测试jwt"`
	Token  string
}
type JCGetTokenRes struct {
}
