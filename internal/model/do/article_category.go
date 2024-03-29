// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ArticleCategory is the golang structure of table article_category for DAO operations like Where/Data.
type ArticleCategory struct {
	g.Meta    `orm:"table:article_category, do:true"`
	Id        interface{} //
	Fid       interface{} // 父级ID
	Name      interface{} // 分类名称
	Remark    interface{} // 备注信息
	Sort      interface{} // 排序顺序
	Status    interface{} // 显示状态（0=显示，1=隐藏）
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
