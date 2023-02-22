package uc_center

// RegisterSystemMasterInput
// @Description: 创建新用户
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-22 16:06:35
type RegisterSystemMasterInput struct {
	Username string
	Password string
	Tel      string
	Name     string
}

// LoginTelSystemMasterInput
// @Description: 用户电话密码登入
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-02-22 23:42:24
type LoginTelSystemMasterInput struct {
	Tel      string
	Password string
}
