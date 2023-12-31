// Code generated with the command 'go generate', DO NOT EDIT.

package baseCond

import (
	"fmt"
)

type FieldCond struct {
	column string
	alias  string

	exprList []*Expr[any]
}

// region Constructor

func NewFieldCond(column string) *FieldCond {
	return &FieldCond{
		column: column,
	}
}

// endregion

func (receiver *FieldCond) As(alias string) *FieldCond {
	receiver.alias = alias
	return receiver
}

func (receiver *FieldCond) Sum() *FieldCond {
	receiver.exprList = append(receiver.exprList, Sum[any]())
	return receiver
}

func (receiver *FieldCond) Avg() *FieldCond {
	receiver.exprList = append(receiver.exprList, Avg[any]())
	return receiver
}

func (receiver *FieldCond) Distinct() *FieldCond {
	receiver.exprList = append(receiver.exprList, Distinct[any]())
	return receiver
}

func (receiver *FieldCond) Count() *FieldCond {
	receiver.exprList = append(receiver.exprList, Count[any]())
	return receiver
}

func (receiver *FieldCond) Sql() string {
	column := receiver.column
	for _, expr := range receiver.exprList {
		switch expr.expr {
		case ExprSum:
			column = fmt.Sprintf("SUM(%s)", column)
			break
		case ExprAvg:
			column = fmt.Sprintf("AVG(%s)", column)
			break
		case ExprDistinct:
			column = fmt.Sprintf("DISTINCT(%s)", column)
			break
		case ExprCount:
			column = fmt.Sprintf("COUNT(%s)", column)
			break
		default:
			panic(fmt.Sprintf("FieldCond[%v]无法在 FieldCond 中使用", expr.expr))
		}
	}
	if receiver.alias != "" {
		column = column + " AS " + receiver.alias
	}
	return column
}
