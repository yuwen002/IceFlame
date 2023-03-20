package article

import "github.com/gogf/gf/v2/os/gtime"

// CreateSinglePageCategoryInput
// @Description: 新建单页分类
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-21 00:10:31
type CreateSinglePageCategoryInput struct {
	Name   string
	Remark string
	Sort   uint32
}

// SinglePageCategoryOutput
// @Description: 单页分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-21 00:11:19
type SinglePageCategoryOutput struct {
	Id        uint16     `json:"id"`
	Name      string     `json:"name"`
	Remark    string     `json:"remark"`
	Sort      uint32     `json:"sort"`
	CreatedAt gtime.Time `json:"created_at"`
	UpdatedAt gtime.Time `json:"updated_at"`
}

type ModifySinglePageCategoryInput struct {
	Id     uint16
	Name   string
	Remark string
	Sort   uint32
}
