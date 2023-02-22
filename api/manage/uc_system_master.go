package manage

import "github.com/gogf/gf/v2/frame/g"

type RegisterSystemMasterReq struct {
	g.Meta          `path:"/master/register" tags:"管理员注册" method:"post" summary:"管理员注册"`
	Tel             string `p:"tel" v:"required|phone#用户电话不能为空|请填写正确电话"`
	Password        string `p:"password" v:"required|length:5,32#密码不能为空|密码在5到32之间"`
	ConfirmPassword string `p:"confirm_password" v:"required|same:password#确认密码不能为空|两次输入密码不一致"`
	Name            string `p:"name" v:"required#姓名不能为空"`
}
type RegisterSystemMasterRes struct {
}
