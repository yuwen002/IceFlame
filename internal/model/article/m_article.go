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

// ModifyChannelInput
// @Description: 修改频道信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 11:10:52
type ModifyChannelInput struct {
	Id     uint32
	Name   string
	Remark string
	Sort   uint32
	Status int8
}

// ListChannelInput
// @Description: 频道信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 14:18:02
type ListChannelInput struct {
	Page int
	Size int
}

// CreateCategoryInput
// @Description: 新建文章分类
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 17:22:20
type CreateCategoryInput struct {
	Fid    uint32
	Name   string
	Remark string
	Sort   uint32
}

// CategoryOutput
// @Description: 文章分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 17:26:51
type CategoryOutput struct {
	Id        uint32     `json:"id"`
	Fid       uint32     `json:"fid"`
	Name      string     `json:"name"`
	Remark    string     `json:"remark"`
	Sort      uint32     `json:"sort"`
	Status    int8       `json:"status"`
	CreatedAt gtime.Time `json:"created_at"`
	UpdatedAt gtime.Time `json:"updated_at"`
}

// ModifyCategoryInput
// @Description: 修改文章分类信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 17:30:45
type ModifyCategoryInput struct {
	Id     uint32
	Fid    uint32
	Name   string
	Remark string
	Sort   uint32
	Status int8
}

// ListCategoryInput
// @Description: 文章分类信息列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-23 17:33:07
type ListCategoryInput struct {
	Page int
	Size int
}

// CreateTagInput
// @Description: 新建标签
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-26 00:16:25
type CreateTagInput struct {
	Name   string
	Remark string
	Sort   uint32
}

// TagOutput
// @Description: 标签显示
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-26 00:30:34
type TagOutput struct {
	Id        uint32     `json:"id"`
	Name      string     `json:"name"`
	Remark    string     `json:"remark"`
	Sort      uint32     `json:"sort"`
	Status    int8       `json:"status"`
	CreatedAt gtime.Time `json:"created_at"`
	UpdatedAt gtime.Time `json:"updated_at"`
}

// ModifyTagInput
// @Description: 修改标签
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-26 00:38:57
type ModifyTagInput struct {
	Id     uint32
	Name   string
	Remark string
	Sort   uint32
	Status int8
}

// ListTagInput
// @Description: 标签列表
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-26 00:49:20
type ListTagInput struct {
	Page int
	Size int
}

// CreateArticleInput
// @Description: 新建文章
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-25 00:37:55
type CreateArticleInput struct {
	CategoryId  uint32
	ChannelId   uint32
	Title       string
	Keyword     string
	Description string
	Link        string
	Author      string
	Tags        string
	PubDate     string
	Summary     string
	Content     string
	Thumbnail   string
	Click       uint32
}

// GetArticleOutput
// @Description: 获取文章
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-25 00:39:13
type GetArticleOutput struct {
	Id          uint32     `json:"id"`
	CategoryId  uint32     `json:"category_id"`
	ChannelId   uint32     `json:"channel_id"`
	Title       string     `json:"title"`
	Keyword     string     `json:"keyword"`
	Description string     `json:"description"`
	Link        string     `json:"link"`
	Author      string     `json:"author"`
	Tags        string     `json:"tags"`
	TagNames    string     `json:"tag_names"`
	PubDate     string     `json:"pub_date"`
	Summary     string     `json:"summary"`
	Content     string     `json:"content"`
	Thumbnail   string     `json:"thumbnail"`
	Click       uint32     `json:"click"`
	Status      int8       `json:"status"`
	CreatedAt   gtime.Time `json:"created_at"`
	UpdatedAt   gtime.Time `json:"updated_at"`
}

// ModifyArticleInput
// @Description: 修改文章
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-25 00:39:25
type ModifyArticleInput struct {
	Id          uint32
	CategoryId  uint32
	ChannelId   uint32
	Title       string
	Keyword     string
	Description string
	Link        string
	Author      string
	Tags        string
	PubDate     string
	Summary     string
	Content     string
	Thumbnail   string
	Click       uint32
	Status      int8
}

// ListArticleInput
// @Description: 文章列表
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-25 00:39:37
type ListArticleInput struct {
	Page int
	Size int
}
