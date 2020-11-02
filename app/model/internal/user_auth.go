// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// UserAuth is the golang structure for table gf_user_auth.
type UserAuth struct {
    Id           uint        `orm:"id,primary"    json:"id"`            // UID                      
    UserId       uint        `orm:"user_id"       json:"user_id"`       // 用户ID                   
    IdentityType string      `orm:"identity_type" json:"identity_type"` // 第三方登录类型           
    Identifier   string      `orm:"identifier"    json:"identifier"`    // 第三方登录的唯一标识     
    Credential   string      `orm:"credential"    json:"credential"`    // 第三方登录token或者密钥  
    CreatedAt    *gtime.Time `orm:"created_at"    json:"created_at"`    // 创建时间                 
    UpdatedAt    *gtime.Time `orm:"updated_at"    json:"updated_at"`    // 更新时间                 
}