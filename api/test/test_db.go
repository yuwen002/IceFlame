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

type DBGetStructRCacheByIdReq struct {
	g.Meta `path:"/test/get_struct_rcache_by_id" tags:"测试GetStructRCacheById" method:"post" summary:"测试GetStructRCacheById，按ID返回数据,redis缓存"`
	Id     int64 `p:"id" v:"required#ID不能为空"`
}
type DBGetStructRCacheByIdRes struct {
}

type DBGetMapByIdReq struct {
	g.Meta `path:"/test/get_map_by_id" tags:"测试GetMapById" method:"post" summary:"测试GetMapById，按ID返回数据"`
	Id     int64 `p:"id" v:"required#ID不能为空"`
}
type DBGetMapByIdRes struct {
}

type DBGetMapGCacheByIdReq struct {
	g.Meta `path:"/test/get_map_gcache_by_id" tags:"测试GetMapGCacheById" method:"post" summary:"测试GetMapGCacheById，按ID返回数据"`
	Id     int64 `p:"id" v:"required#ID不能为空"`
}
type DBGetMapGCacheByIdRes struct {
}

type DBGetMapRCacheByIdReq struct {
	g.Meta `path:"/test/get_map_rcache_by_id" tags:"测试GetMapRCacheById" method:"post" summary:"测试GetMapRCacheById，按ID返回数据"`
	Id     int64 `p:"id" v:"required#ID不能为空"`
}
type DBGetMapRCacheByIdRes struct {
}

type DBGetStructByIdsReq struct {
	g.Meta `path:"/test/get_struct_by_ids" tags:"测试GetStructByIds" method:"post" summary:"测试GetStructByIds，按Ids返回数据"`
	Ids    string `p:"ids" v:"required#ID不能为空"`
}
type DBGetStructByIdsRes struct {
}

type DBGetMapByIdsReq struct {
	g.Meta `path:"/test/get_map_by_ids" tags:"测试GetMapByIds" method:"post" summary:"测试GetMapByIds，按Ids返回数据"`
	Ids    string `p:"ids" v:"required#ID不能为空"`
}
type DBGetMapByIdsRes struct {
}

type DBGetOneStructByWhereReq struct {
	g.Meta `path:"/test/get_one_struct_by_where" tags:"测试GetMapByIds" method:"post" summary:"测试GetMapByIds，按条件返回数据"`
	Title  string `p:"title" v:"required#标题不能为空"`
}
type DBGetOneStructByWhereRes struct {
}

type DBGetOneMapByWhereReq struct {
	g.Meta `path:"/test/get_one_map_by_where" tags:"测试GetMapByIds" method:"post" summary:"测试GetMapByIds，按条件返回数据"`
	Title  string `p:"title" v:"required#标题不能为空"`
}
type DBGetOneMapByWhereRes struct {
}

type DBGetAllStructByWhereReq struct {
	g.Meta   `path:"/test/get_all_struct_by_where" tags:"测试GetAllStructByWhere" method:"post" summary:"测试GetAllStructByWhere，按条件返回数据"`
	Classify string `p:"classify" v:"required#分类不能为空"`
}
type DBGetAllStructByWhereRes struct {
}

type DBGetAllMapByWhereReq struct {
	g.Meta   `path:"/test/get_all_map_by_where" tags:"测试GetAllMapByWhere" method:"post" summary:"测试GetAllMapByWhere，按条件返回数据"`
	Classify string `p:"classify" v:"required#分类不能为空"`
}
type DBGetAllMapByWhereRes struct {
}

type DBModifyByIdReq struct {
	g.Meta   `path:"/test/modify_by_id" tags:"测试ModifyById" method:"put" summary:"测试ModifyGCacheById，按Id修改数据"`
	Id       string `p:"id" v:"required#ID不能为空"`
	TestData string `p:"test_data"`
	Title    string `p:"title"`
	Classify string `p:"classify"`
}
type DBModifyByIdRes struct {
}

type DBModifyGCacheByIdReq struct {
	g.Meta   `path:"/test/modify_gcache_by_id" tags:"测试ModifyGCacheById" method:"put" summary:"测试ModifyGCacheById，按Id修改数据"`
	Id       string `p:"id" v:"required#ID不能为空"`
	TestData string `p:"test_data"`
	Title    string `p:"title"`
	Classify string `p:"classify"`
}
type DBModifyGCacheByIdRes struct {
}

type DBModifyRCacheByIdReq struct {
	g.Meta   `path:"/test/modify_rcache_by_id" tags:"测试ModifyGCacheById" method:"put" summary:"测试ModifyGCacheById，按Id修改数据"`
	Id       string `p:"id" v:"required#ID不能为空"`
	TestData string `p:"test_data"`
	Title    string `p:"title"`
	Classify string `p:"classify"`
}
type DBModifyRCacheByIdRes struct {
}

type DBModifyByWhereReq struct {
	g.Meta   `path:"/test/modify_by_where" tags:"测试ModifyByWhere" method:"put" summary:"测试ModifyByWhere，按条件修改数据"`
	Id       string `p:"id" v:"required#ID不能为空"`
	TestData string `p:"test_data"`
	Title    string `p:"title"`
	Classify string `p:"classify"`
}
type DBModifyByWhereRes struct {
}

type DBDelByIdReq struct {
	g.Meta `path:"/test/del_by_id" tags:"测试DelById" method:"delete" summary:"测试DelById，按条件删除数据"`
	Id     int64 `p:"id" v:"required#ID不能为空"`
}
type DBDelByIdRes struct {
}
