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
