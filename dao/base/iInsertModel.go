// Code generated with the command 'go generate', DO NOT EDIT.

package base

type IInserter interface {
	TableInfo() TableInfo
	InsertData() IInserterDataModel
	HasAutoincrPk() bool
}

type IBulkInserter interface {
	TableInfo() TableInfo
	InsertData() []IInserterDataModel
	HasAutoincrPk() bool
}

type IInserterDataModel interface {
	GetInsertCols() []string
	RealInsertData() any
}
