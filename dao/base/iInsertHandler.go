// Code generated with the command 'go generate', DO NOT EDIT.

package base

type IInsertHandler interface {
	Insert(IInserter) (insertCnt int64, pk interface{}, err error)
	BulkInsert(IBulkInserter) (insertCnt int64, pkList []interface{}, err error)
}
