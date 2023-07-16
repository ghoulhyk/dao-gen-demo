// Code generated with the command 'go generate', DO NOT EDIT.

package base

type ISelectorParams interface {
	TableInfo() TableInfo
	ProvideResultData() interface{}
	GetWhereCondList() *[]IWhereCond
	GetPageCond() IPage
	GetOrderCondList() *[]IOrder
	GetFieldList() *[]IFieldList
}
