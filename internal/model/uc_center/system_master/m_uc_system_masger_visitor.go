package system_master

// CreateVisitCategoryInput
// @Description: 添加访问标题
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-01 16:45:12
type CreateVisitCategoryInput struct {
	Title string
}

// ModifyVisitCategoryByIdInput
// @Description: 按ID修改访问标题
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-01 17:28:28
type ModifyVisitCategoryByIdInput struct {
	Id    uint16
	Title string
}

// ListVisitCategoryInput
// @Description: 访问标题列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-03 17:35:28
type ListVisitCategoryInput struct {
	Page int
	Size int
}

// ListVisitCategoryOutput
// @Description: 访问标题列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-01 17:28:25
type ListVisitCategoryOutput struct {
	Id        uint16
	Title     string
	CreatedAt string
	UpdatedAt string
}

// DeleteVisitCategoryInput
// @Description: 删除缓存key
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-03 20:38:04
type DeleteVisitCategoryInput struct {
	Id  string
	Key string
}

// CreateVisitorLogsInput
// @Description:访问信息写入
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-03 23:16:26
type CreateVisitorLogsInput struct {
	AccountId     uint64
	OsCategory    int8
	VisitCategory int16
	UnionId       uint64
	Description   string
	IPLong        interface{}
	IP            string
}

// ListVisitorLogsInput
// @Description: 访问日志列表
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-03-03 23:20:02
type ListVisitorLogsInput struct {
	AccountId     uint64
	OsCategory    int8
	VisitCategory int16
	Page          int
	Size          int
}

type ListVisitorLogsOutput struct {
	Id                uint64 `json:"id"`
	AccountId         uint64 `json:"account_id"`
	OsCategory        int8   `json:"os_category"`
	OsCategoryName    string `json:"os_category_name"`
	VisitCategory     int16  `json:"visit_category"`
	VisitCategoryName string `json:"visit_category_name"`
	UnionId           uint64 `json:"union_id"`
	Description       string `json:"description"`
	IP                string `json:"ip"`
	//IPLong            string `json:"ip_long"`
	CreatedAt string `json:"created_at"`
}
