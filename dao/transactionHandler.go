// Code generated with the command 'go generate', DO NOT EDIT.

package dao

import (
	"xorm.io/xorm"
)

type transactionHandler struct {
	session       *xorm.Session
	inTransaction bool
}

func (receiver *transactionHandler) StartTransaction() error {
	if receiver.IsInTransaction() {
		panic("mysql不支持事务嵌套")
	}
	receiver.inTransaction = true
	return (*receiver.session).Begin()
}

func (receiver *transactionHandler) Commit() error {
	if !receiver.IsInTransaction() {
		return nil
	}
	receiver.inTransaction = false
	return (*receiver.session).Commit()
}

func (receiver *transactionHandler) Rollback() error {
	receiver.inTransaction = false
	return (*receiver.session).Rollback()
}

func (receiver *transactionHandler) IsInTransaction() bool {
	return receiver.inTransaction
}
