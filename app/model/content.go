// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"focus/app/model/internal"
	"github.com/gogf/gf/os/gtime"
)

// Content is the golang structure for table gf_content.
type Content internal.Content

const (
	ContentListDefaultSize = 10
	ContentListMaxSize     = 50
	ContentSortDefault     = 0
	ContentSortActive      = 1
	ContentSortHot         = 2
	ContentTypeArticle     = "article"
	ContentTypeAsk         = "ask"
	ContentTypeTopic       = "topic"
)

// ==========================================================================================
// 通用数据结构
// ==========================================================================================
// 主要用于列表展示
type ContentListItem struct {
	Id         uint        `json:"id"`          // 自增ID
	CategoryId uint        `json:"category_id"` // 栏目ID
	UserId     uint        `json:"user_id"`     // 用户ID
	Title      string      `json:"title"`       // 标题
	Content    string      `json:"content"`     // 内容
	Sort       uint        `json:"sort"`        // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶
	Brief      string      `json:"brief"`       // 摘要
	Thumb      string      `json:"thumb"`       // 缩略图
	Tags       string      `json:"tags"`        // 标签名称列表，以JSON存储
	Referer    string      `json:"referer"`     // 内容来源，例如github/gitee
	Status     uint        `json:"status"`      // 状态 0: 正常, 1: 禁用
	ViewCount  uint        `json:"view_count"`  // 浏览数量
	ZanCount   uint        `json:"zan_count"`   // 赞
	CaiCount   uint        `json:"cai_count"`   // 踩
	CreatedAt  *gtime.Time `json:"created_at"`  // 创建时间
	UpdatedAt  *gtime.Time `json:"updated_at"`  // 修改时间
}

// 绑定到Content列表中的栏目信息
type ContentListCategoryItem struct {
	Id          uint   `json:"id"`           // 分类ID，自增主键
	Name        string `json:"name"`         // 分类名称
	Thumb       string `json:"thumb"`        // 封面图
	ContentType string `json:"content_type"` // 内容类型：content, ask, article, reply

}

// 绑定到Content列表中的用户信息
type ContentListUserItem struct {
	Id       uint   `json:"id"`       // UID
	Nickname string `json:"nickname"` // 昵称
	Avatar   string `json:"avatar"`   // 头像地址
}

// Content详情
type ContentDetail struct {
	Content Content
	User    User
}

// ==========================================================================================
// API
// ==========================================================================================
// API查看内容详情
type ContentApiDetailReq struct {
	Id uint `v:"min:1#请选择查看的内容"`
}

// API创建/修改内容基类
type ContentApiCreateUpdateBase struct {
	ContentServiceCreateUpdateBase
	CategoryId uint   `v:"min:1#请输入栏目ID"`    // 栏目ID
	Title      string `v:"required#请输入内容标题"` // 标题
	Content    string `v:"required#请输入内容内容"` // 内容
}

// API创建内容
type ContentApiCreateReq struct {
	ContentApiCreateUpdateBase
}

// API修改内容
type ContentApiUpdateReq struct {
	ContentApiCreateUpdateBase
	Id uint `v:"min:1#请选择需要修改的内容"` // 修改时ID不能为空
}

// API删除内容
type ContentApiDeleteReq struct {
	Id uint `v:"min:1#请选择需要删除的内容"` // 删除时ID不能为空
}

// ==========================================================================================
// Service
// ==========================================================================================
// Service查询列表
type ContentServiceGetListReq struct {
	Type       string // 内容模型
	CategoryId uint   `p:"cate"`                    // 栏目ID
	Page       int    `d:"1"  v:"min:0#分页号码错误"`     // 分页号码
	Size       int    `d:"10" v:"max:50#分页数量最大50条"` // 分页数量，最大50
	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// Service查询列表结果
type ContentServiceGetListRes struct {
	List  []*ContentServiceGetListResItem `json:"list"`  // 列表
	Page  int                             `json:"page"`  // 分页码
	Size  int                             `json:"size"`  // 分页数量
	Total int                             `json:"total"` // 数据总数
}

type ContentServiceGetListResItem struct {
	Content  *ContentListItem         `json:"content"`
	Category *ContentListCategoryItem `json:"category"`
	User     *ContentListUserItem     `json:"user"`
}

// Service查询详情结果
type ContentServiceGetDetailRes struct {
	Content *Content `json:"content"`
	User    *User    `json:"user"`
}

// Service创建/修改内容基类
type ContentServiceCreateUpdateBase struct {
	Type       string   // 内容模型
	CategoryId uint     // 栏目ID
	Title      string   // 标题
	Content    string   // 内容
	Brief      string   // 摘要
	Thumb      string   // 缩略图
	Tags       []string // 标签名称列表，以JSON存储
	Referer    string   // 内容来源，例如github/gitee
}

// Service创建内容
type ContentServiceCreateReq struct {
	ContentServiceCreateUpdateBase
	UserId uint
}

// Service修改内容
type ContentServiceUpdateReq struct {
	ContentServiceCreateUpdateBase
	Id uint
}
