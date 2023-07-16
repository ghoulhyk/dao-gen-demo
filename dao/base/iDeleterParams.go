// Code generated with the command 'go generate', DO NOT EDIT.

package base

type IDeleterParams interface {
	TableInfo() TableInfo
	GetWhereCondList() *[]IWhereCond
}
