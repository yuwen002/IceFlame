package uc_center

// RegisterInput
// @Description: 创建新用户
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-22 16:06:35
type RegisterInput struct {
	Username string
	Password string
	Tel      string
	Name     string
}

// LoginTelPasswordInput
// @Description: 用户电话密码登入
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-02-22 23:42:24
type LoginTelPasswordInput struct {
	Tel      string
	Password string
}

// LoginUsernamePasswordInput
// @Description:用户名密码登入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-23 16:20:59
type LoginUsernamePasswordInput struct {
	Username string
	Password string
}
