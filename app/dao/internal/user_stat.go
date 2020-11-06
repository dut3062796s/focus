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

// UserStatDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type UserStatDao struct {
	gmvc.M
	Table   string
	Columns userStatColumns
}

// UserStatColumns defines and stores column names for table gf_user_stat.
type userStatColumns struct {
	UserId     string // 用户ID    
    ZanCount   string // 被赞次数  
    CaiCount   string // 被踩次数  
    CreatedAt  string // 创建时间  
    UpdatedAt  string // 更新时间
}

var (
	// UserStat is globally public accessible object for table gf_user_stat operations.
	UserStat = &UserStatDao{
		M:     g.DB("default").Table("gf_user_stat").Safe(),
		Table: "gf_user_stat",
		Columns: userStatColumns{
			UserId:    "user_id",     
            ZanCount:  "zan_count",   
            CaiCount:  "cai_count",   
            CreatedAt: "created_at",  
            UpdatedAt: "updated_at",
		},
	}
)

// As sets an alias name for current table.
func (d *UserStatDao) As(as string) *UserStatDao {
	return &UserStatDao{M:d.M.As(as)}
}

// TX sets the transaction for current operation.
func (d *UserStatDao) TX(tx *gdb.TX) *UserStatDao {
	return &UserStatDao{M:d.M.TX(tx)}
}

// Master marks the following operation on master node.
func (d *UserStatDao) Master() *UserStatDao {
	return &UserStatDao{M:d.M.Master()}
}

// Slave marks the following operation on slave node.
// Note that it makes sense only if there's any slave node configured.
func (d *UserStatDao) Slave() *UserStatDao {
	return &UserStatDao{M:d.M.Slave()}
}

// LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *UserStatDao) LeftJoin(table ...string) *UserStatDao {
	return &UserStatDao{M:d.M.LeftJoin(table...)}
}

// RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *UserStatDao) RightJoin(table ...string) *UserStatDao {
	return &UserStatDao{M:d.M.RightJoin(table...)}
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *UserStatDao) InnerJoin(table ...string) *UserStatDao {
	return &UserStatDao{M:d.M.InnerJoin(table...)}
}

// Fields sets the operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *UserStatDao) Fields(fieldNamesOrMapStruct ...interface{}) *UserStatDao {
	return &UserStatDao{M:d.M.Fields(fieldNamesOrMapStruct...)}
}

// FieldsEx sets the excluded operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *UserStatDao) FieldsEx(fieldNamesOrMapStruct ...interface{}) *UserStatDao {
	return &UserStatDao{M:d.M.FieldsEx(fieldNamesOrMapStruct...)}
}

// Option sets the extra operation option for the model.
func (d *UserStatDao) Option(option int) *UserStatDao {
	return &UserStatDao{M:d.M.Option(option)}
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (d *UserStatDao) OmitEmpty() *UserStatDao {
	return &UserStatDao{M:d.M.OmitEmpty()}
}

// Filter marks filtering the fields which does not exist in the fields of the operated table.
func (d *UserStatDao) Filter() *UserStatDao {
	return &UserStatDao{M:d.M.Filter()}
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
func (d *UserStatDao) Where(where interface{}, args ...interface{}) *UserStatDao {
	return &UserStatDao{M:d.M.Where(where, args...)}
}

// WherePri does the same logic as M.Where except that if the parameter <where>
// is a single condition like int/string/float/slice, it treats the condition as the primary
// key value. That is, if primary key is "id" and given <where> parameter as "123", the
// WherePri function treats the condition as "id=123", but M.Where treats the condition
// as string "123".
func (d *UserStatDao) WherePri(where interface{}, args ...interface{}) *UserStatDao {
	return &UserStatDao{M:d.M.WherePri(where, args...)}
}

// And adds "AND" condition to the where statement.
func (d *UserStatDao) And(where interface{}, args ...interface{}) *UserStatDao {
	return &UserStatDao{M:d.M.And(where, args...)}
}

// Or adds "OR" condition to the where statement.
func (d *UserStatDao) Or(where interface{}, args ...interface{}) *UserStatDao {
	return &UserStatDao{M:d.M.Or(where, args...)}
}

// Group sets the "GROUP BY" statement for the model.
func (d *UserStatDao) Group(groupBy string) *UserStatDao {
	return &UserStatDao{M:d.M.Group(groupBy)}
}

// Order sets the "ORDER BY" statement for the model.
func (d *UserStatDao) Order(orderBy ...string) *UserStatDao {
	return &UserStatDao{M:d.M.Order(orderBy...)}
}

// Limit sets the "LIMIT" statement for the model.
// The parameter <limit> can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
func (d *UserStatDao) Limit(limit ...int) *UserStatDao {
	return &UserStatDao{M:d.M.Limit(limit...)}
}

// Offset sets the "OFFSET" statement for the model.
// It only makes sense for some databases like SQLServer, PostgreSQL, etc.
func (d *UserStatDao) Offset(offset int) *UserStatDao {
	return &UserStatDao{M:d.M.Offset(offset)}
}

// Page sets the paging number for the model.
// The parameter <page> is started from 1 for paging.
// Note that, it differs that the Limit function start from 0 for "LIMIT" statement.
func (d *UserStatDao) Page(page, limit int) *UserStatDao {
	return &UserStatDao{M:d.M.Page(page, limit)}
}

// Batch sets the batch operation number for the model.
func (d *UserStatDao) Batch(batch int) *UserStatDao {
	return &UserStatDao{M:d.M.Batch(batch)}
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
func (d *UserStatDao) Cache(duration time.Duration, name ...string) *UserStatDao {
	return &UserStatDao{M:d.M.Cache(duration, name...)}
}

// Data sets the operation data for the model.
// The parameter <data> can be type of string/map/gmap/slice/struct/*struct, etc.
// Eg:
// Data("uid=10000")
// Data("uid", 10000)
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
func (d *UserStatDao) Data(data ...interface{}) *UserStatDao {
	return &UserStatDao{M:d.M.Data(data...)}
}

// All does "SELECT FROM ..." statement for the model.
// It retrieves the records from table and returns the result as []*model.UserStat.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *UserStatDao) All(where ...interface{}) ([]*model.UserStat, error) {
	all, err := d.M.All(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.UserStat
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// One retrieves one record from table and returns the result as *model.UserStat.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *UserStatDao) One(where ...interface{}) (*model.UserStat, error) {
	one, err := d.M.One(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.UserStat
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindOne retrieves and returns a single Record by M.WherePri and M.One.
// Also see M.WherePri and M.One.
func (d *UserStatDao) FindOne(where ...interface{}) (*model.UserStat, error) {
	one, err := d.M.FindOne(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.UserStat
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindAll retrieves and returns Result by by M.WherePri and M.All.
// Also see M.WherePri and M.All.
func (d *UserStatDao) FindAll(where ...interface{}) ([]*model.UserStat, error) {
	all, err := d.M.FindAll(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.UserStat
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// Chunk iterates the table with given size and callback function.
func (d *UserStatDao) Chunk(limit int, callback func(entities []*model.UserStat, err error) bool) {
	d.M.Chunk(limit, func(result gdb.Result, err error) bool {
		var entities []*model.UserStat
		err = result.Structs(&entities)
		if err == sql.ErrNoRows {
			return false
		}
		return callback(entities, err)
	})
}

// LockUpdate sets the lock for update for current operation.
func (d *UserStatDao) LockUpdate() *UserStatDao {
	return &UserStatDao{M:d.M.LockUpdate()}
}

// LockShared sets the lock in share mode for current operation.
func (d *UserStatDao) LockShared() *UserStatDao {
	return &UserStatDao{M:d.M.LockShared()}
}

// Unscoped enables/disables the soft deleting feature.
func (d *UserStatDao) Unscoped() *UserStatDao {
	return &UserStatDao{M:d.M.Unscoped()}
}