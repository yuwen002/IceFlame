// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// DbTest is the golang structure of table db_test for DAO operations like Where/Data.
type DbTest struct {
	g.Meta   `orm:"table:db_test, do:true"`
	Id       interface{} // 自增ID
	TestData interface{} // 数据
	Title    interface{} // 标题
}
