package manage

import "github.com/gogf/gf/v2/frame/g"

// AddPartnerLevelReq
// @Description: 添加合伙人级别
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-30 17:33:47
type AddPartnerLevelReq struct {
	g.Meta `path:"/partner/level/add" tags:"添加合伙人级别" method:"post" summary:"添加合伙人级别"`
	Name   string `p:"name" v:"required#级别名称不能为空"`
	Sort   uint16 `p:"sort" v:"min:0|integer#ID要大于等于0|ID只为整数"`
	Remark string `p:"remark"`
}
type AddPartnerLevelRes struct{}

// GetPartnerLevelReq
// @Description: 编辑合伙人级别
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-30 17:55:55
type GetPartnerLevelReq struct {
	g.Meta `path:"/partner/level/edit" tags:"编辑合伙人级别" method:"get" summary:"编辑合伙人级别"`
	Id     uint16 `p:"id" v:"required|min:0|integer#ID不能为空|ID要大于等于0|ID只为整数"`
}
type GetPartnerLevelRes struct{}

// EditPartnerLevelReq
// @Description: 编辑合伙人级别
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-30 18:16:33
type EditPartnerLevelReq struct {
	g.Meta `path:"/partner/level/edit" tags:"编辑合伙人级别" method:"put" summary:"编辑合伙人级别"`
	Id     uint16 `p:"id" v:"required|min:0|integer#ID不能为空|ID要大于等于0|ID只为整数"`
	Name   string `p:"name" v:"required#级别名称不能为空"`
	Sort   uint16 `p:"sort" v:"min:0|integer#ID要大于等于0|ID只为整数"`
	Remark string `p:"remark"`
}
type EditPartnerLevelRes struct{}

// ListPartnerLevelReq
// @Description: 合伙人级别列表
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-31 00:03:09
type ListPartnerLevelReq struct {
	g.Meta `path:"/partner/level/show" tags:"合伙人级别列表" method:"get" summary:"合伙人级别列表"`
	Page   int `p:"page" v:"min:1|integer#页码要大于等于1|页码只为正整数"`
	Size   int `p:"size" v:"min:1|integer#显示条数要大于等于1|显示条数只为正整数"`
}
type ListPartnerLevelRes struct{}

// AddPartnerReq
// @Description: 添加合伙人
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-04-03 16:02:07
type AddPartnerReq struct {
	g.Meta        `path:"/partner/add" tags:"添加合伙人" method:"post" summary:"添加合伙人"`
	Fid           uint64 `p:"fid" v:"required|min:0|integer#父类ID不能为空|ID要大于等于0|ID只为整数"`
	Name          string `p:"name" v:"required#合伙人名称不能为空"`
	Tel           string `p:"tel" v:"required|phone#用户电话不能为空|请填写正确电话"`
	LevelId       uint16 `p:"level_id" v:"required|min:0|integer#级别ID不能为空|ID要大于等于0|ID只为整数"`
	Password      string `p:"password" v:"length:5,32#密码在5到32之间"`
	PromotionType uint8  `p:"promotion_type" v:"min:0|integer|in:0,1#晋级方式ID要大于等于0|晋级方式ID只为整数|晋级方式传入数据错误"`
}
type AddPartnerRes struct {
}
