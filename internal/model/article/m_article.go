package article

import "github.com/gogf/gf/v2/os/gtime"

// CreateChannelInput
// @Description: 新建频道
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-23 00:00:53
type CreateChannelInput struct {
	Name   string
	Remark string
	Sort   uint32
}

// ChannelOutput
// @Description: 频道信息显示
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-23 00:08:34
type ChannelOutput struct {
	Id        uint32     `json:"id"`
	Name      string     `json:"name"`
	Remark    string     `json:"remark"`
	Sort      uint32     `json:"sort"`
	Status    int8       `json:"status"`
	CreatedAt gtime.Time `json:"created_at"`
	UpdatedAt gtime.Time `json:"updated_at"`
}
