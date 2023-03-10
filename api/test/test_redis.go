package test

import "github.com/gogf/gf/v2/frame/g"

type RCExistsSetDataReq struct {
	g.Meta `path:"/test/exists_set_data" tags:"测试ExistsSetData" method:"get" summary:"测试ExistsSetData"`
	Id     int64 `p:"id" v:"required#ID不能为空"`
}
type RCExistsSetDataRes struct {
}

type RCCastHashDataReq struct {
	g.Meta `path:"/test/cast_hash_data" tags:"测试ExistsSetData" method:"get" summary:"测试ExistsSetData"`
	Id     int64 `p:"id" v:"required#ID不能为空"`
}
type RCCastHashDataRes struct {
}
