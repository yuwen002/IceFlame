package model

// JwtClaimsInput
// @Description: token基本信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-22 10:51:59
type JwtClaimsInput struct {
	Id       uint64
	UserInfo map[string]interface{}
	Expire   int64
	Before   int64
	Issuer   string
	Subject  string
}
