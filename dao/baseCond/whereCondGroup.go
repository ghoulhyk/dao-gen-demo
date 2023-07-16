// Code generated with the command 'go generate', DO NOT EDIT.

package baseCond

import (
	"demo/dao/base"
	"xorm.io/builder"
)

type WhereCondGroup struct {
	op    base.Op
	group []base.IWhereCond
}

func (receiver WhereCondGroup) GetOp() base.Op {
	return receiver.op
}

func (receiver WhereCondGroup) GetWhereBuilder() builder.Cond {
	cond := builder.NewCond()
	for _, whereCond := range receiver.group {
		condItem := whereCond.GetWhereBuilder()
		if whereCond.GetOp() == base.Op_AND {
			cond = cond.And(condItem)
		} else {
			cond = cond.Or(condItem)
		}
	}
	return cond
}
