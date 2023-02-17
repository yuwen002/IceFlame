package test

import "github.com/gogf/gf/v2/frame/g"

type RCExistsSetDataReq struct {
	g.Meta `path:"/test/exists_set_data" tags:"测试ExistsSetData" method:"put" summary:"测试ExistsSetData"`
	Id     int64 `p:"id" v:"required#ID不能为空"`
}
type RCExistsSetDataRes struct {
}
