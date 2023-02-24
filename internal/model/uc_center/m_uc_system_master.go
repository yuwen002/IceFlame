package system_master

// RegisterInput
// @Description: 创建新用户
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-22 16:06:35
type RegisterInput struct {
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

// CreateSystemMasterInput
// @Description: 后台管理员新建用户
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-23 16:56:11
type CreateSystemMasterInput struct {
	Username string
	Password string
	Tel      string
	Name     string
}

// ModifyPasswordInput
// @Description: 修改用户密码
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-24 13:56:52
type ModifyPasswordInput struct {
	OldPassword string
	NewPassword string
}
