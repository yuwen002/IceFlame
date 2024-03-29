// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Article is the golang structure of table article for DAO operations like Where/Data.
type Article struct {
	g.Meta      `orm:"table:article, do:true"`
	Id          interface{} //
	CategoryId  interface{} // 分类ID
	ChannelId   interface{} // 栏目ID
	Title       interface{} // 标题
	Keyword     interface{} // 关键字
	Description interface{} // 描述
	Link        interface{} // 文章链接
	Author      interface{} // 作者
	Tags        interface{} // Tag标签
	PubDate     *gtime.Time // 发布时间
	Summary     interface{} // 摘要
	Content     interface{} // 内容
	Thumbnail   interface{} // 缩略图
	Click       interface{} // 点击量
	Status      interface{} // 显示状态（0=显示，1=隐藏）
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
