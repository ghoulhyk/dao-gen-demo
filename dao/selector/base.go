// Code generated with the command 'go generate', DO NOT EDIT.

package selector

import (
	"demo/dao/base"
	"demo/dao/daoerrors"
)

type _count struct {
	clientHolder base.IClientHolder
	params       base.ISelectorParams
}

func (receiver *_count) Count() (uint64, error) {
	return receiver.clientHolder.GetClient().Count(receiver.params)
}

// 发生错误直接panic
// Count为0 不会 panic
func (receiver *_count) MustCount() uint64 {
	cnt, err := receiver.Count()
	if err != nil {
		if receiver.clientHolder.GetClient().IsInTransaction() {
			receiver.clientHolder.GetClient().Rollback()
		}
		panic(daoerrors.SelectErr(receiver.params.TableInfo().TableName, err))
	}
	return cnt
}

// Deprecated: 使用 MustCount
func (receiver *_count) CountOrPanic() uint64 {
	return receiver.MustCount()
}

type _exist struct {
	_count
}

func (receiver *_exist) Exist() (bool, error) {
	cnt, err := receiver.Count()
	if err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func (receiver *_exist) ExistIgnoreErr() bool {
	exist, _ := receiver.Exist()
	return exist
}
