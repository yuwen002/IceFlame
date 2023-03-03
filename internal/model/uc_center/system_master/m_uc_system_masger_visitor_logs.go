package system_master

// AddVisitCategoryInput
// @Description: 添加访问标题
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-03-01 16:45:12
type AddVisitCategoryInput struct {
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

type AddVisitorLogsInput struct {
	AccountId  uint64
	OsCategory int8
}
