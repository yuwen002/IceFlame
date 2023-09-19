package article

import "github.com/gogf/gf/v2/os/gtime"

// CreateSinglePageInput
// @Description: 新建单页信息
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-21 00:10:31
type CreateSinglePageInput struct {
	Title       string
	Description string
	Keyword     string
	Content     string
	Thumbnail   string
}

// SinglePageOutput
// @Description: 获取单页信息
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-21 00:11:19
type SinglePageOutput struct {
	Id          uint32     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Keyword     string     `json:"keyword"`
	Content     string     `json:"content"`
	Thumbnail   string     `json:"thumbnail"`
	Click       uint32     `json:"click"`
	Status      int8       `json:"status"`
	CreatedAt   gtime.Time `json:"created_at"`
	UpdatedAt   gtime.Time `json:"updated_at"`
}

// ModifySinglePageInput
// @Description: 修改单页信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-21 11:26:17
type ModifySinglePageInput struct {
	Id          uint32
	Title       string
	Description string
	Keyword     string
	Content     string
	Thumbnail   string
	Click       uint32
	Status      int8
}

// ModifyStatusSinglePageInput
// @Description: 修改单页信息状态
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-09-19 11:48:41
type ModifyStatusSinglePageInput struct {
	Id     uint32
	Status int8
}

// ListSinglePageInput
// @Description:分类信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-21 11:26:36
type ListSinglePageInput struct {
	Page int
	Size int
}
