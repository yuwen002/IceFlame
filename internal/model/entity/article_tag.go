// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ArticleTag is the golang structure for table article_tag.
type ArticleTag struct {
	Id        uint        `json:"id"         ` //
	Name      string      `json:"name"       ` // tag内容
	Remark    string      `json:"remark"     ` // 备注
	Status    int         `json:"status"     ` // 显示状态（0=显示，1=隐藏）
	CreatedAt *gtime.Time `json:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updated_at" ` //
}