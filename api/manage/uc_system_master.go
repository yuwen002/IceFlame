package manage

import "github.com/gogf/gf/v2/frame/g"

// RegisterReq
// @Description: 管理员注册
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-23 11:37:31
type RegisterReq struct {
	g.Meta          `path:"/master/register" tags:"管理员注册" method:"post" summary:"管理员注册"`
	Tel             string `p:"tel" v:"required|phone#用户电话不能为空|请填写正确电话"`
	Password        string `p:"password" v:"required|length:5,32#密码不能为空|密码在5到32之间"`
	ConfirmPassword string `p:"confirm_password" v:"required|same:password#确认密码不能为空|两次输入密码不一致"`
	Name            string `p:"name" v:"required#姓名不能为空"`
}
type RegisterRes struct {
}

// LoginTelPasswordReq
// @Description: 管理员登入，登入方式电话密码登入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-23 14:33:03
type LoginTelPasswordReq struct {
	g.Meta   `path:"/master/login_tel_password" tags:"管理员登入" method:"post" summary:"管理员登入,登入方式电话和密码"`
	Tel      string `p:"tel" v:"required|phone#用户电话不能为空|请填写正确电话"`
	Password string `p:"password" v:"required|length:5,32#密码不能为空|密码在5到32之间"`
}
type LoginTelPasswordRes struct {
}

// LoginUsernamePasswordReq
// @Description: 管理员登入，登入方式用户名密码登入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-23 16:29:05
type LoginUsernamePasswordReq struct {
	g.Meta   `path:"/master/login_username_password" tags:"管理员登入" method:"post" summary:"管理员登入用户名和密码"`
	Username string `p:"tel" v:"required#用户名不能为空"`
	Password string `p:"password" v:"required|length:5,32#密码不能为空|密码在5到32之间"`
}
type LoginUsernamePasswordRes struct {
}
