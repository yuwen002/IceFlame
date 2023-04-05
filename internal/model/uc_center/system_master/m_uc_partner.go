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

// ListPartnerLevelInput
// @Description: 合伙人等级列表
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-02 01:23:48
type ListPartnerLevelInput struct {
	Page int
	Size int
}

// CreatePartnerInput
// @Description: 新建合伙人
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-02 01:24:04
type CreatePartnerInput struct {
	Fid           uint64
	LevelId       uint16
	RoleId        uint16
	Password      string
	Name          string
	Tel           string
	PromotionType uint8
}

// ModifyPartnerInput
// @Description: 修改合伙人
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-05 23:39:50
type ModifyPartnerInput struct {
	AccountId     uint64
	LevelId       uint16
	Name          string
	Tel           string
	PromotionType uint8
}

// ListPartnerInput
// @Description: 合伙人列表
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-04-05 23:40:55
type ListPartnerInput struct {
	Page int
	Size int
}

type ListPartnerOutput struct {
	AccountId     uint64      `json:"account_id"`
	LevelId       uint16      `json:"level_id"`
	LevelName     string      `json:"level_name"`
	Team1st       uint32      `json:"team_1st"`
	Team2nd       uint32      `json:"team_2nd"`
	Name          string      `json:"name"`
	Tel           string      `json:"tel"`
	InviteCode    string      `json:"invite_code"`
	PromotionType uint8       `json:"promotion_type"`
	Status        uint8       `json:"status"`
	CreatedAt     *gtime.Time `json:"created_at"`
	UpdatedAt     *gtime.Time `json:"updated_at"`
}

// PartnerOutput
// @Description: 合伙人详细信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-04-03 17:41:24
type PartnerOutput struct {
	Fid           uint64      `json:"fid"`
	FName         string      `json:"f_name"`
	FTel          string      `json:"f_tel"`
	AccountId     uint64      `json:"account_id"`
	LevelId       uint16      `json:"level_id"`
	LevelName     string      `json:"level_name"`
	Team1st       uint32      `json:"team_1st"`
	Team2nd       uint32      `json:"team_2nd"`
	Name          string      `json:"name"`
	Tel           string      `json:"tel"`
	InviteCode    string      `json:"invite_code"`
	PromotionType uint8       `json:"promotion_type"`
	Status        uint8       `json:"status"`
	CreatedAt     *gtime.Time `json:"created_at"`
	UpdatedAt     *gtime.Time `json:"updated_at"`
}
