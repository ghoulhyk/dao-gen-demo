// Code generated with the command 'go generate', DO NOT EDIT.

package base

import (
	"xorm.io/builder"
)

type IWhereCond interface {
	GetOp() Op
	GetWhereBuilder() builder.Cond
}

type Op string

const (
	Op_AND Op = "AND"
	Op_OR  Op = "OR"
)
