// Code generated with the command 'go generate', DO NOT EDIT.

package base

type IJoinSelectorParams interface {
	Joins() map[TableInfo]IJoinCond
	TableInfo() TableInfo
	ProvideResultData() interface{}
	GetWhereCondList() *[]IWhereCond
	GetPageCond() IPage
	GetOrderCondList() *[]IOrder
	GetFieldList() *[]IFieldList
}
