// Code generated with the command 'go generate', DO NOT EDIT.

package base

type IUpdateValItem interface {
	ColumnName() string
	GetVal() interface{}
	GetRawSql() string
	IsRawSql() bool
}
