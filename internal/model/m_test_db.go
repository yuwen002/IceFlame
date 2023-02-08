package model

// TestDBInsertInput
// @Description: 数据库插入测试
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-03 12:55:27
type TestDBInsertInput struct {
	TestData string
}

// TestGetByIdInput
// @Description: 按ID查询信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-03 12:55:46
type TestGetByIdInput struct {
	Id int64
}

// TestGetByIdOutput
// @Description: 按ID查询信息结果
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-03 12:55:58
type TestGetByIdOutput struct {
	Id       int64  `json:"id"`
	TestData string `json:"test_data"`
}

// TestGetByIdsInput
// @Description: 按ID查询多条信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-03 18:19:26
type TestGetByIdsInput struct {
	Ids interface{}
}

// TestGetOneByWhereInput
// @Description:条件查询单条信息
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-02-05 23:29:27
type TestGetOneByWhereInput struct {
	Where string
	Args  string
}

// TestGetOneByWhereOutput
// @Description:条件查询单条信息
// @Author liuxingyu <yuwen002@163.com>
// @Data 2023-02-05 23:29:37
type TestGetOneByWhereOutput struct {
	Id       int64  `json:"id"`
	TestData string `json:"test_data"`
	Title    string `json:"title"`
}

// TestGetAllByWhereInput
// @Description:条件查询多条信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-07 11:09:42
type TestGetAllByWhereInput struct {
	Where string
	Args  string
	Limit int
}

// TestGetAllByWhereOutput
// @Description: 条件查询单条信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-07 11:09:56
type TestGetAllByWhereOutput struct {
	Id       int64  `json:"id"`
	TestData string `json:"test_data"`
	Title    string `json:"title"`
}

// TestModifyByIdInput
// @Description: 按ID修改数据
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-08 14:05:11
type TestModifyByIdInput struct {
	Data  interface{}
	Where string
	Args  string
}

// TestModifyByWhereInput
// @Description: 按条件修改数据
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-08 14:08:18
type TestModifyByWhereInput struct {
	Data  interface{}
	Where string
	Args  string
	Limit int
}

// TestDelByIdInput
// @Description: 按ID删除数据
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-08 15:59:44
type TestDelByIdInput struct {
	Id int64
}

// TestDelByIdsInput
// @Description: 按ID删除多条数据
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-08 16:43:23
type TestDelByIdsInput struct {
	Ids interface{}
}

// TestDelByWhereInput
// @Description: 按条件删除数据
// @Author liuxingyu <yuwen002@163.com>
// @Date 2023-02-08 16:00:29
type TestDelByWhereInput struct {
	Where string
	Args  string
	Limit int
}
