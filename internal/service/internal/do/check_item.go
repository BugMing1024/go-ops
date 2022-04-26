// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CheckItem is the golang structure of table check_item for DAO operations like Where/Data.
type CheckItem struct {
	g.Meta      `orm:"table:check_item, do:true"`
	Id          interface{} //
	CheckItemId interface{} //
	Name        interface{} // 检查项名称
	Desc        interface{} // 检查项描述
	Type        interface{} //
	Content     interface{} // 检查项内容
	Creater     interface{} // 创建人
	Updater     interface{} // 更新人
	Created     *gtime.Time //
	Updated     *gtime.Time //
}