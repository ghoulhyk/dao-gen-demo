// Code generated with the command 'go generate', DO NOT EDIT.

package base

type IUpdaterParams interface {
	TableInfo() TableInfo
	GetWhereCondList() *[]IWhereCond
	GetUpdateValItemList() *[]IUpdateValItem
	GetLimit() int
}
