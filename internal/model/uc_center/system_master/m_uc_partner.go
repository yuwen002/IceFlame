package system_master

import "github.com/gogf/gf/v2/os/gtime"

// CreatePartnerLevelInput
// @Description: 新建合伙人等级
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-30 15:10:17
type CreatePartnerLevelInput struct {
	Name   string
	Sort   uint16
	Remark string
}

// PartnerLevelOutput
// @Description: 合伙人等级信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-30 15:59:47
type PartnerLevelOutput struct {
	Id        uint16     `json:"id"`
	Name      string     `json:"name"`
	Sort      uint16     `json:"sort"`
	Remark    string     `json:"remark"`
	CreatedAt gtime.Time `json:"created_at"`
	UpdatedAt gtime.Time `json:"updated_at"`
}

// ModifyPartnerLevelInput
// @Description: 按ID修改合伙人等级信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-30 16:32:38
type ModifyPartnerLevelInput struct {
	Id     uint16
	Name   string
	Sort   uint16
	Remark string
}

type ListPartnerLevelInput struct {
	Page int
	Size int
}

type CreatePartnerInput struct {
	Fid           uint64
	LevelId       uint16
	Password      string
	Name          string
	Tel           string
	PromotionType uint8
}
