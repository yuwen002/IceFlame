package test

import "github.com/gogf/gf/v2/frame/g"

type DBInsertReq struct {
	g.Meta `path:"/test/insert" tags:"测试Insert" method:"post" summary:"测试Insert"`
	Data   string `p:"data" v:"required#写入数据不能为空"`
}

type DBInsertRes struct {
}

type DBInsertAndGetIdReq struct {
	g.Meta `path:"/test/insert_and_get_id" tags:"测试InsertAndGetId" method:"post" summary:"测试Insert并获取插入ID"`
	Data   string `p:"data" v:"required#写入数据不能为空"`
}
type DBInsertAndGetIdRes struct {
}

type DBGetStructByIdReq struct {
	g.Meta `path:"/test/get_struct_by_id" tags:"测试GetStructById" method:"post" summary:"测试GetStructById，按ID返回数据"`
	Id     int64 `p:"id" v:"required#ID不能为空"`
}

type DBGetStructByIdRes struct {
}

type DBGetStructGCacheByIdReq struct {
	g.Meta `path:"/test/get_struct_gcache_by_id" tags:"测试GetStructGCacheById" method:"post" summary:"测试GetStructGCacheById，按ID返回数据,gcache缓存"`
	Id     int64 `p:"id" v:"required#ID不能为空"`
}

type DBGetStructGCacheByIdRes struct {
}
