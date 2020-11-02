// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package internal

import (
	"database/sql"
	"focus/app/model"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
	"time"
)

// UserRoleAuthDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type UserRoleAuthDao struct {
	gmvc.M
	Table   string
	Columns userRoleAuthColumns
}

// UserRoleAuthColumns defines and stores column names for table gf_user_role_auth.
type userRoleAuthColumns struct {
	Id         string // 自增ID                                  
    RoleId     string // 角色ID                                  
    Name       string // 权限名称                                
    Key        string // 权限键名(用于程序)                      
    Value      string // 权限键值，部分自定义权限可能有键值存在  
    Brief      string // 自定义权限描述                          
    CreatedAt  string // 创建时间
}

var (
	// UserRoleAuth is globally public accessible object for table gf_user_role_auth operations.
	UserRoleAuth = &UserRoleAuthDao{
		M:     g.DB("default").Table("gf_user_role_auth").Safe(),
		Table: "gf_user_role_auth",
		Columns: userRoleAuthColumns{
			Id:        "id",          
            RoleId:    "role_id",     
            Name:      "name",        
            Key:       "key",         
            Value:     "value",       
            Brief:     "brief",       
            CreatedAt: "created_at",
		},
	}
)

// As sets an alias name for current table.
func (dao *UserRoleAuthDao) As(as string) *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.As(as)}
}

// TX sets the transaction for current operation.
func (dao *UserRoleAuthDao) TX(tx *gdb.TX) *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.TX(tx)}
}

// Master marks the following operation on master node.
func (dao *UserRoleAuthDao) Master() *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.Master()}
}

// Slave marks the following operation on slave node.
// Note that it makes sense only if there's any slave node configured.
func (dao *UserRoleAuthDao) Slave() *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.Slave()}
}

// LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
func (dao *UserRoleAuthDao) LeftJoin(table ...string) *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.LeftJoin(table...)}
}

// RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
func (dao *UserRoleAuthDao) RightJoin(table ...string) *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.RightJoin(table...)}
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
func (dao *UserRoleAuthDao) InnerJoin(table ...string) *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.InnerJoin(table...)}
}

// Fields sets the operation fields of the model, multiple fields joined using char ','.
func (dao *UserRoleAuthDao) Fields(fields ...string) *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.Fields(fields...)}
}

// FieldsEx sets the excluded operation fields of the model, multiple fields joined using char ','.
func (dao *UserRoleAuthDao) FieldsEx(fields ...string) *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.FieldsEx(fields...)}
}

// Option sets the extra operation option for the model.
func (dao *UserRoleAuthDao) Option(option int) *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.Option(option)}
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (dao *UserRoleAuthDao) OmitEmpty() *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.OmitEmpty()}
}

// Filter marks filtering the fields which does not exist in the fields of the operated table.
func (dao *UserRoleAuthDao) Filter() *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.Filter()}
}

// Where sets the condition statement for the model. The parameter <where> can be type of
// string/map/gmap/slice/struct/*struct, etc. Note that, if it's called more than one times,
// multiple conditions will be joined into where statement using "AND".
// Eg:
// Where("uid=10000")
// Where("uid", 10000)
// Where("money>? AND name like ?", 99999, "vip_%")
// Where("uid", 1).Where("name", "john")
// Where("status IN (?)", g.Slice{1,2,3})
// Where("age IN(?,?)", 18, 50)
// Where(User{ Id : 1, UserName : "john"})
func (dao *UserRoleAuthDao) Where(where interface{}, args ...interface{}) *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.Where(where, args...)}
}

// WherePri does the same logic as M.Where except that if the parameter <where>
// is a single condition like int/string/float/slice, it treats the condition as the primary
// key value. That is, if primary key is "id" and given <where> parameter as "123", the
// WherePri function treats the condition as "id=123", but M.Where treats the condition
// as string "123".
func (dao *UserRoleAuthDao) WherePri(where interface{}, args ...interface{}) *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.WherePri(where, args...)}
}

// And adds "AND" condition to the where statement.
func (dao *UserRoleAuthDao) And(where interface{}, args ...interface{}) *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.And(where, args...)}
}

// Or adds "OR" condition to the where statement.
func (dao *UserRoleAuthDao) Or(where interface{}, args ...interface{}) *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.Or(where, args...)}
}

// Group sets the "GROUP BY" statement for the model.
func (dao *UserRoleAuthDao) Group(groupBy string) *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.Group(groupBy)}
}

// Order sets the "ORDER BY" statement for the model.
func (dao *UserRoleAuthDao) Order(orderBy ...string) *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.Order(orderBy...)}
}

// Limit sets the "LIMIT" statement for the model.
// The parameter <limit> can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
func (dao *UserRoleAuthDao) Limit(limit ...int) *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.Limit(limit...)}
}

// Offset sets the "OFFSET" statement for the model.
// It only makes sense for some databases like SQLServer, PostgreSQL, etc.
func (dao *UserRoleAuthDao) Offset(offset int) *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.Offset(offset)}
}

// Page sets the paging number for the model.
// The parameter <page> is started from 1 for paging.
// Note that, it differs that the Limit function start from 0 for "LIMIT" statement.
func (dao *UserRoleAuthDao) Page(page, limit int) *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.Page(page, limit)}
}

// Batch sets the batch operation number for the model.
func (dao *UserRoleAuthDao) Batch(batch int) *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.Batch(batch)}
}

// Cache sets the cache feature for the model. It caches the result of the sql, which means
// if there's another same sql request, it just reads and returns the result from cache, it
// but not committed and executed into the database.
//
// If the parameter <duration> < 0, which means it clear the cache with given <name>.
// If the parameter <duration> = 0, which means it never expires.
// If the parameter <duration> > 0, which means it expires after <duration>.
//
// The optional parameter <name> is used to bind a name to the cache, which means you can later
// control the cache like changing the <duration> or clearing the cache with specified <name>.
//
// Note that, the cache feature is disabled if the model is operating on a transaction.
func (dao *UserRoleAuthDao) Cache(duration time.Duration, name ...string) *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.Cache(duration, name...)}
}

// Data sets the operation data for the model.
// The parameter <data> can be type of string/map/gmap/slice/struct/*struct, etc.
// Eg:
// Data("uid=10000")
// Data("uid", 10000)
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
func (dao *UserRoleAuthDao) Data(data ...interface{}) *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.Data(data...)}
}

// All does "SELECT FROM ..." statement for the model.
// It retrieves the records from table and returns the result as []*model.UserRoleAuth.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (dao *UserRoleAuthDao) All(where ...interface{}) ([]*model.UserRoleAuth, error) {
	all, err := dao.M.All(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.UserRoleAuth
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// One retrieves one record from table and returns the result as *model.UserRoleAuth.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (dao *UserRoleAuthDao) One(where ...interface{}) (*model.UserRoleAuth, error) {
	one, err := dao.M.One(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.UserRoleAuth
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindOne retrieves and returns a single Record by M.WherePri and M.One.
// Also see M.WherePri and M.One.
func (dao *UserRoleAuthDao) FindOne(where ...interface{}) (*model.UserRoleAuth, error) {
	one, err := dao.M.FindOne(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.UserRoleAuth
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindAll retrieves and returns Result by by M.WherePri and M.All.
// Also see M.WherePri and M.All.
func (dao *UserRoleAuthDao) FindAll(where ...interface{}) ([]*model.UserRoleAuth, error) {
	all, err := dao.M.FindAll(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.UserRoleAuth
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// Chunk iterates the table with given size and callback function.
func (dao *UserRoleAuthDao) Chunk(limit int, callback func(entities []*model.UserRoleAuth, err error) bool) {
	dao.M.Chunk(limit, func(result gdb.Result, err error) bool {
		var entities []*model.UserRoleAuth
		err = result.Structs(&entities)
		if err == sql.ErrNoRows {
			return false
		}
		return callback(entities, err)
	})
}

// LockUpdate sets the lock for update for current operation.
func (dao *UserRoleAuthDao) LockUpdate() *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.LockUpdate()}
}

// LockShared sets the lock in share mode for current operation.
func (dao *UserRoleAuthDao) LockShared() *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.LockShared()}
}

// Unscoped enables/disables the soft deleting feature.
func (dao *UserRoleAuthDao) Unscoped() *UserRoleAuthDao {
	return &UserRoleAuthDao{M:dao.M.Unscoped()}
}