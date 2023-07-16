// Code generated with the command 'go generate', DO NOT EDIT.

package whereCond

import (
	"demo/dao/base"
	"demo/dao/baseCond"
	"fmt"
	"xorm.io/builder"

	"time"
)

func New_User(op base.Op, _tableInfo *base.TableInfo) *User {
	return &User{
		_op:        op,
		_tableInfo: _tableInfo,
	}
}

func New_User_Inner(op base.Op, _tableInfo *base.TableInfo) *User_Inner {
	return &User_Inner{
		_op:        op,
		_tableInfo: _tableInfo,
	}
}

// region wrapper

type User struct {
	_op        base.Op
	_tableInfo *base.TableInfo
	_conds     []base.IWhereCond
}

func (receiver *User) And() *User_Inner {
	cond := New_User_Inner(base.Op_AND, receiver._tableInfo)
	receiver._conds = append(receiver._conds, cond)
	return cond
}

func (receiver *User) Or() *User_Inner {
	cond := New_User_Inner(base.Op_OR, receiver._tableInfo)
	receiver._conds = append(receiver._conds, cond)
	return cond
}

// region 原生sql

func (receiver *User) AndRaw(sql string, args ...any) {
	receiver._conds = append(receiver._conds, baseCond.NewRawWhereCondItem[any](base.Op_AND, sql, args...))
}

func (receiver *User) OrRaw(sql string, args ...any) {
	receiver._conds = append(receiver._conds, baseCond.NewRawWhereCondItem[any](base.Op_OR, sql, args...))
}

// endregion

// region 嵌套

func (receiver *User) AndNest(fun func(cond *User)) {
	cond := New_User(base.Op_AND, receiver._tableInfo)
	fun(cond)
	receiver._conds = append(receiver._conds, cond)
}
func (receiver *User) OrNest(fun func(cond *User)) {
	cond := New_User(base.Op_OR, receiver._tableInfo)
	fun(cond)
	receiver._conds = append(receiver._conds, cond)
}

// endregion

func (receiver *User) GetOp() base.Op {
	return receiver._op
}

func (receiver *User) GetWhereBuilder() builder.Cond {
	cond := builder.NewCond()
	for _, whereCond := range receiver._conds {
		if whereCond.GetOp() == base.Op_OR {
			cond = cond.Or(whereCond.GetWhereBuilder())
		} else {
			cond = cond.And(whereCond.GetWhereBuilder())
		}
	}
	return cond
}

// endregion

type User_Inner struct {
	_op        base.Op
	_tableInfo *base.TableInfo
	_raw       *baseCond.RawWhereCondItem

	id         *baseCond.WhereCondItem[uint32]
	name       *baseCond.WhereCondItem[string] // 姓名
	gender     *baseCond.WhereCondItem[uint8]
	createTime *baseCond.WhereCondItem[time.Time]
}

// region 数据库字段

func (receiver *User_Inner) Id() *baseCond.WhereCondItem[uint32] {
	receiver.id = baseCond.NewWhereCondItem[uint32](base.Op_AND, "id")
	return receiver.id
}

// 姓名
func (receiver *User_Inner) Name() *baseCond.WhereCondItem[string] {
	receiver.name = baseCond.NewWhereCondItem[string](base.Op_AND, "name")
	return receiver.name
}

func (receiver *User_Inner) Gender() *baseCond.WhereCondItem[uint8] {
	receiver.gender = baseCond.NewWhereCondItem[uint8](base.Op_AND, "gender")
	return receiver.gender
}

func (receiver *User_Inner) CreateTime() *baseCond.WhereCondItem[time.Time] {
	receiver.createTime = baseCond.NewWhereCondItem[time.Time](base.Op_AND, "create_time")
	return receiver.createTime
}

// endregion

func (receiver *User_Inner) tableName() string {
	tableName := ""
	if receiver._tableInfo != nil && receiver._tableInfo.Alias != "" {
		tableName = receiver._tableInfo.Alias
		tableName = fmt.Sprintf("`%s`.", tableName)
	}
	return tableName
}

func (receiver *User_Inner) GetOp() base.Op {
	return receiver._op
}

func (receiver *User_Inner) GetWhereBuilder() builder.Cond {
	cond := builder.NewCond()

	if receiver.id != nil {
		if receiver.id.GetOp() == base.Op_OR {
			cond = cond.Or(receiver.id.GetWhereBuilder())
		} else {
			cond = cond.And(receiver.id.GetWhereBuilder())
		}
	}

	if receiver.name != nil {
		if receiver.name.GetOp() == base.Op_OR {
			cond = cond.Or(receiver.name.GetWhereBuilder())
		} else {
			cond = cond.And(receiver.name.GetWhereBuilder())
		}
	}

	if receiver.gender != nil {
		if receiver.gender.GetOp() == base.Op_OR {
			cond = cond.Or(receiver.gender.GetWhereBuilder())
		} else {
			cond = cond.And(receiver.gender.GetWhereBuilder())
		}
	}

	if receiver.createTime != nil {
		if receiver.createTime.GetOp() == base.Op_OR {
			cond = cond.Or(receiver.createTime.GetWhereBuilder())
		} else {
			cond = cond.And(receiver.createTime.GetWhereBuilder())
		}
	}

	if receiver._raw != nil {
		if receiver._raw.GetOp() == base.Op_OR {
			cond = cond.Or(receiver._raw.GetWhereBuilder())
		} else {
			cond = cond.And(receiver._raw.GetWhereBuilder())
		}
	}

	return cond
}
