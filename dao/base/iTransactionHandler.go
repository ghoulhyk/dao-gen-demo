// Code generated with the command 'go generate', DO NOT EDIT.

package base

type ITransactionHandler interface {
	StartTransaction() error
	Commit() error
	Rollback() error
	IsInTransaction() bool
}
