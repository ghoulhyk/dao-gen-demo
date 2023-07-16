// Code generated with the command 'go generate', DO NOT EDIT.

package deleter

import (
	"demo/dao/base"
	"demo/dao/daoerrors"
	"demo/dao/whereCond"

	"demo/dao/databaseDef"
	"demo/dao/util"
)

// region deleterParams

type userDeleterParams struct {
	whereCondList []base.IWhereCond
}

func (receiver *userDeleterParams) TableInfo() base.TableInfo {
	return base.TableInfo{
		DatabaseInfo: util.ToPtr(databaseDef.Db1()),
		TableName:    "user",
	}
}

func (receiver *userDeleterParams) GetWhereCondList() *[]base.IWhereCond {
	return &receiver.whereCondList
}

// endregion

type UserWrapper struct {
	clientHolder base.IClientHolder
	params       *userDeleterParams
}

func (receiver *UserWrapper) Init(client base.IClient) *UserWrapper {
	clientHolder := &base.ClientHolder{}
	clientHolder.Init(client)
	receiver.clientHolder = clientHolder
	receiver.params = &userDeleterParams{}
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

func (receiver *UserWrapper) Delete() (deleteRows uint64, err error) {
	return receiver.clientHolder.GetClient().Delete(receiver.params)
}

// 发生错误直接panic
// 删除数量为0 不会 panic
func (receiver *UserWrapper) MustDelete() (deleteRows uint64) {
	deleteRows, err := receiver.Delete()
	if err != nil {
		if receiver.clientHolder.GetClient().IsInTransaction() {
			receiver.clientHolder.GetClient().Rollback()
		}
		panic(daoerrors.DeleteErr("user", err))
	}
	return
}

// Deprecated: 使用 MustDelete
func (receiver *UserWrapper) DeleteOrPanic() (deleteRows uint64) {
	return receiver.MustDelete()
}

func (receiver *UserWrapper) DeleteById(id uint32) (deleted bool, err error) {
	localVar := UserWrapper{}

	localVar.Init(receiver.clientHolder.GetClient())

	rows, err := localVar.AndWhere(func(cond *whereCond.User) {
		cond.And().Id().Equ(id)
	}).Delete()
	if err != nil {
		return
	}
	return rows == 1, nil
}

// 发生错误直接panic
// 删除数量为0 不会 panic
func (receiver *UserWrapper) MustDeleteById(id uint32) (deleted bool) {
	result, err := receiver.DeleteById(id)
	if err != nil {
		panic(daoerrors.DeleteErr("user", err))
	}
	return result
}

// Deprecated: 使用 MustDeleteById
func (receiver *UserWrapper) DeleteByIdOrPanic(id uint32) (deleted bool) {
	return receiver.MustDeleteById(id)
}

func (receiver *UserWrapper) DeleteByIds(ids ...uint32) (deleteRows uint64, err error) {
	localVar := UserWrapper{}

	localVar.Init(receiver.clientHolder.GetClient())

	rows, err := localVar.AndWhere(func(cond *whereCond.User) {
		cond.And().Id().In(ids...)
	}).Delete()
	if err != nil {
		return
	}
	return rows, nil
}

// 发生错误直接panic
// 删除数量为0 不会 panic
func (receiver *UserWrapper) MustDeleteByIds(ids ...uint32) (deleteRows uint64) {
	result, err := receiver.DeleteByIds(ids...)
	if err != nil {
		panic(daoerrors.DeleteErr("user", err))
	}
	return result
}

// Deprecated: 使用 MustDeleteByIds
func (receiver *UserWrapper) DeleteByIdsOrPanic(ids ...uint32) (deleteRows uint64) {
	return receiver.MustDeleteByIds(ids...)
}
