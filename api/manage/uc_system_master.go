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
	Username string `p:"username" v:"required#用户名不能为空"`
	Password string `p:"password" v:"required|length:5,32#密码不能为空|密码在5到32之间"`
}
type LoginUsernamePasswordRes struct {
}

// CreateSystemMasterReq
// @Description: 管理员新建管理用户
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-23 17:04:22
type CreateSystemMasterReq struct {
	g.Meta          `path:"/master/create_system_master" tags:"管理员新建管理用户" method:"post" summary:"管理员新建管理用户"`
	Username        string `p:"username" v:"required#用户名不能为空"`
	Tel             string `p:"tel" v:"required|phone#用户电话不能为空|请填写正确电话"`
	Password        string `p:"password" v:"required|length:5,32#密码不能为空|密码在5到32之间"`
	ConfirmPassword string `p:"confirm_password" v:"required|same:password#确认密码不能为空|两次输入密码不一致"`
	Name            string `p:"name" v:"required#姓名不能为空"`
}
type CreateSystemMasterRes struct {
}

// EditPasswordReq
// @Description:修改关联
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-24 16:24:25
type EditPasswordReq struct {
	g.Meta             `path:"/master/edit_password" tags:"管理员新建管理用户" method:"put" summary:"管理员新建管理用户"`
	OldPassword        string `p:"old_password" v:"required|length:5,32#密码不能为空|密码在5到32之间"`
	NewPassword        string `p:"new_password" v:"required|length:5,32#密码不能为空|密码在5到32之间"`
	ConfirmNewPassword string `p:"confirm_new_password" v:"required|same:new_password#确认密码不能为空|两次输入密码不一致"`
}
type EditPasswordRes struct {
}

// ListSystemMasterReq
// @Description: 管理员列表
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-02-25 23:15:07
type ListSystemMasterReq struct {
	g.Meta `path:"/master/show_system_master" tags:"管理员用户列表" method:"get" summary:"管理员用户列表"`
	Page   int `p:"page" v:"min:1|integer#页码要大于等于1|页码只为正整数"`
	Size   int `p:"size" v:"min:1|integer#显示条数要大于等于1|显示条数只为正整数"`
}
type ListSystemMasterRes struct {
}

// EditSystemMasterReq
// @Description:编辑管理员列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-27 10:23:48
type EditSystemMasterReq struct {
	g.Meta    `path:"/master/edit_system_master" tags:"管理员用户列表" method:"put" summary:"管理员用户列表"`
	AccountId uint64 `p:"account_id" v:"required|min:1|integer#ID不能为空|ID要大于等于1|ID只为正整数"`
	Tel       string `p:"tel" v:"required|phone#用户电话不能为空|请填写正确电话"`
	Name      string `p:"name" v:"required#姓名不能为空"`
	Status    int8   `p:"status" v:"required#用户状态不能为空"`
}
type EditSystemMasterRes struct {
}

// ResetPasswordSystemMasterReq
// @Description: 管理员重置密码
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-28 15:13:21
type ResetPasswordSystemMasterReq struct {
	g.Meta          `path:"/master/reset_password" tags:"重置管理员密码" method:"put" summary:"重置管理员密码"`
	AccountId       uint64 `p:"account_id" v:"required|min:1|integer#ID不能为空|ID要大于等于1|ID只为正整数"`
	Password        string `p:"password" v:"length:5,32#密码在5到32之间"`
	ConfirmPassword string `p:"confirm_password" v:"same:password#两次输入密码不一致"`
}
type ResetPasswordSystemMasterRes struct {
}

// EditStatusSystemMasterReq
// @Description: 管理员修改管理员登入状态
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-28 15:57:42
type EditStatusSystemMasterReq struct {
	g.Meta    `path:"/master/edit_status" tags:"重置管理员密码" method:"put" summary:"重置管理员密码"`
	AccountId uint64 `p:"account_id" v:"required|min:1|integer#ID不能为空|ID要大于等于1|ID只为正整数"`
	Status    string `p:"status" v:"in:required|0,1#用户状态不能为空|用户状态值错误"`
}
type EditStatusSystemMasterRes struct {
}

type UnlockSystemMasterReq struct {
	g.Meta    `path:"/master/unlock_system_master" tags:"重置管理员密码" method:"put" summary:"重置管理员密码"`
	AccountId uint64 `p:"account_id" v:"required|min:1|integer#ID不能为空|ID要大于等于1|ID只为正整数"`
}
type UnlockSystemMasterRes struct {
}
