// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// Role is the golang structure for table gf_role.
type Role struct {
    Id        uint        `orm:"id,primary" json:"id"`         // 角色ID                                                  
    Name      string      `orm:"name"       json:"name"`       // 角色名称                                                
    Sort      uint        `orm:"sort"       json:"sort"`       // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶  
    Brief     string      `orm:"brief"      json:"brief"`      // 角色描述                                                
    CreatedAt *gtime.Time `orm:"created_at" json:"created_at"` // 创建时间                                                
}