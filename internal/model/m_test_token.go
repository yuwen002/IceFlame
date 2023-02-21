package model

type JwtClaimsInput struct {
	Id       uint64
	UserInfo map[string]interface{}
	Expire   int64
	Before   int64
	Issuer   string
	Subject  string
}
