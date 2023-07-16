// Code generated with the command 'go generate', DO NOT EDIT.

package updater

import (
	"demo/dao/base"
	"demo/dao/baseCond"
	"demo/dao/daoerrors"
	"demo/dao/whereCond"
	"fmt"

	"demo/dao/databaseDef"
	"demo/dao/util"

	"time"
)

// region updaterParams

type userUpdaterParams struct {
	whereCondList []base.IWhereCond
	limit         int

	// region UpdateValItem

	id         *baseCond.UpdateValItem
	name       *baseCond.UpdateValItem // 姓名
	gender     *baseCond.UpdateValItem
	createTime *baseCond.UpdateValItem

	// endregion
}

func (receiver *userUpdaterParams) TableInfo() base.TableInfo {
	return base.TableInfo{
		DatabaseInfo: util.ToPtr(databaseDef.Db1()),
		TableName:    "user",
	}
}

func (receiver *userUpdaterParams) GetWhereCondList() *[]base.IWhereCond {
	return &receiver.whereCondList
}

func (receiver *userUpdaterParams) GetUpdateValItemList() *[]base.IUpdateValItem {
	var result []base.IUpdateValItem

	if receiver.id != nil {
		result = append(result, receiver.id)
	}

	if receiver.name != nil {
		result = append(result, receiver.name)
	}

	if receiver.gender != nil {
		result = append(result, receiver.gender)
	}

	if receiver.createTime != nil {
		result = append(result, receiver.createTime)
	}

	return &result
}

func (receiver *userUpdaterParams) GetLimit() int {
	return receiver.limit
}

// endregion

type UserWrapper struct {
	clientHolder base.IClientHolder
	params       *userUpdaterParams
}

func (receiver *UserWrapper) Init(client base.IClient) *UserWrapper {
	clientHolder := &base.ClientHolder{}
	clientHolder.Init(client)
	receiver.clientHolder = clientHolder
	receiver.params = &userUpdaterParams{}
	receiver.params.limit = -1
	return receiver
}

func (receiver *UserWrapper) Where(fun func(cond *whereCond.User)) *UserWrapper {
	return receiver.AndWhere(fun)
}

func (receiver *UserWrapper) AndWhere(fun func(cond *whereCond.User)) *UserWrapper {
	cond := whereCond.New_User(base.Op_AND, util.ToPtr(receiver.params.TableInfo()))
	fun(cond)
	receiver.params.whereCondList = append(receiver.params.whereCondList, cond)
	return receiver
}

func (receiver *UserWrapper) OrWhere(fun func(cond *whereCond.User)) *UserWrapper {
	cond := whereCond.New_User(base.Op_OR, util.ToPtr(receiver.params.TableInfo()))
	fun(cond)
	receiver.params.whereCondList = append(receiver.params.whereCondList, cond)
	return receiver
}

func (receiver *UserWrapper) Limit(limit uint) *UserWrapper {
	receiver.params.limit = int(limit)
	return receiver
}

func (receiver *UserWrapper) Update() (updateRows uint64, err error) {
	return receiver.clientHolder.GetClient().Update(receiver.params)
}

// 发生错误直接panic
// 更新数量为0 不会 panic
func (receiver *UserWrapper) MustUpdate() (updateRows uint64) {
	updateRows, err := receiver.Update()
	if err != nil {
		if receiver.clientHolder.GetClient().IsInTransaction() {
			receiver.clientHolder.GetClient().Rollback()
		}
		panic(daoerrors.UpdateErr("user", err))
	}

	return
}

// Deprecated: 使用 MustUpdate
func (receiver *UserWrapper) UpdateOrPanic() (updateRows uint64) {
	return receiver.MustUpdate()
}

// region val setter

// region id

func (receiver *UserWrapper) SetId(val uint32) *UserWrapper {
	return receiver.SetIdPtr(&val)
}

func (receiver *UserWrapper) SetIdPtr(ptr *uint32) *UserWrapper {
	receiver.getId().SetPtr(ptr)
	return receiver
}

func (receiver *UserWrapper) SetIdExpr(sql string) *UserWrapper {
	receiver.getId().SetRawSql(sql)
	return receiver
}

// SetIdIncr 自增,默认为 1
func (receiver *UserWrapper) SetIdIncr(steps ...interface{}) *UserWrapper {
	var step interface{} = 1

	if len(steps) > 0 {
		switch steps[0].(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
			step = steps[0]
		default:
			panic("SetIdIncr 参数类型不支持")
		}
	}

	return receiver.SetIdExpr(fmt.Sprintf("id + %v", step))
}

// SetIdDecr 自减,默认为 1
func (receiver *UserWrapper) SetIdDecr(steps ...interface{}) *UserWrapper {
	var step interface{} = 1

	if len(steps) > 0 {
		switch steps[0].(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
			step = steps[0]
		default:
			panic("SetIdDecr 参数类型不支持")
		}
	}

	return receiver.SetIdExpr(fmt.Sprintf("id - %v", step))
}

func (receiver *UserWrapper) getId() *baseCond.UpdateValItem {
	if receiver.params.id == nil {
		receiver.params.id = baseCond.NewUpdateValItem("id")
	}
	return receiver.params.id
}

// endregion

// region name

func (receiver *UserWrapper) SetName(val string) *UserWrapper {
	return receiver.SetNamePtr(&val)
}

func (receiver *UserWrapper) SetNamePtr(ptr *string) *UserWrapper {
	receiver.getName().SetPtr(ptr)
	return receiver
}

func (receiver *UserWrapper) SetNameExpr(sql string) *UserWrapper {
	receiver.getName().SetRawSql(sql)
	return receiver
}

func (receiver *UserWrapper) getName() *baseCond.UpdateValItem {
	if receiver.params.name == nil {
		receiver.params.name = baseCond.NewUpdateValItem("name")
	}
	return receiver.params.name
}

// endregion

// region gender

func (receiver *UserWrapper) SetGender(val uint8) *UserWrapper {
	return receiver.SetGenderPtr(&val)
}

func (receiver *UserWrapper) SetGenderPtr(ptr *uint8) *UserWrapper {
	receiver.getGender().SetPtr(ptr)
	return receiver
}

func (receiver *UserWrapper) SetGenderExpr(sql string) *UserWrapper {
	receiver.getGender().SetRawSql(sql)
	return receiver
}

// SetGenderIncr 自增,默认为 1
func (receiver *UserWrapper) SetGenderIncr(steps ...interface{}) *UserWrapper {
	var step interface{} = 1

	if len(steps) > 0 {
		switch steps[0].(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
			step = steps[0]
		default:
			panic("SetGenderIncr 参数类型不支持")
		}
	}

	return receiver.SetGenderExpr(fmt.Sprintf("gender + %v", step))
}

// SetGenderDecr 自减,默认为 1
func (receiver *UserWrapper) SetGenderDecr(steps ...interface{}) *UserWrapper {
	var step interface{} = 1

	if len(steps) > 0 {
		switch steps[0].(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
			step = steps[0]
		default:
			panic("SetGenderDecr 参数类型不支持")
		}
	}

	return receiver.SetGenderExpr(fmt.Sprintf("gender - %v", step))
}

func (receiver *UserWrapper) getGender() *baseCond.UpdateValItem {
	if receiver.params.gender == nil {
		receiver.params.gender = baseCond.NewUpdateValItem("gender")
	}
	return receiver.params.gender
}

// endregion

// region createTime

func (receiver *UserWrapper) SetCreateTime(val time.Time) *UserWrapper {
	return receiver.SetCreateTimePtr(&val)
}

func (receiver *UserWrapper) SetCreateTimePtr(ptr *time.Time) *UserWrapper {
	receiver.getCreateTime().SetPtr(ptr)
	return receiver
}

func (receiver *UserWrapper) SetCreateTimeExpr(sql string) *UserWrapper {
	receiver.getCreateTime().SetRawSql(sql)
	return receiver
}

// SetCreateTimeIncr 自增,默认为 1
func (receiver *UserWrapper) SetCreateTimeIncr(steps ...interface{}) *UserWrapper {
	var step interface{} = 1

	if len(steps) > 0 {
		switch steps[0].(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
			step = steps[0]
		default:
			panic("SetCreateTimeIncr 参数类型不支持")
		}
	}

	return receiver.SetCreateTimeExpr(fmt.Sprintf("create_time + %v", step))
}

// SetCreateTimeDecr 自减,默认为 1
func (receiver *UserWrapper) SetCreateTimeDecr(steps ...interface{}) *UserWrapper {
	var step interface{} = 1

	if len(steps) > 0 {
		switch steps[0].(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
			step = steps[0]
		default:
			panic("SetCreateTimeDecr 参数类型不支持")
		}
	}

	return receiver.SetCreateTimeExpr(fmt.Sprintf("create_time - %v", step))
}

func (receiver *UserWrapper) getCreateTime() *baseCond.UpdateValItem {
	if receiver.params.createTime == nil {
		receiver.params.createTime = baseCond.NewUpdateValItem("create_time")
	}
	return receiver.params.createTime
}

// endregion

// endregion
