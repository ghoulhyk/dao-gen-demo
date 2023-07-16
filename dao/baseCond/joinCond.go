// Code generated with the command 'go generate', DO NOT EDIT.

package baseCond

import "demo/dao/base"

type JoinCond struct {
	joinIndex uint8
	op        base.JoinOp
	cond      string
}

func (receiver JoinCond) GetOp() base.JoinOp {
	return receiver.op
}

func (receiver JoinCond) GetCond() string {
	return receiver.cond
}

func (receiver JoinCond) GetJoinIndex() uint8 {
	return receiver.joinIndex
}

func LeftJoin(joinIndex uint8, cond string) base.IJoinCond {
	return &JoinCond{
		joinIndex: joinIndex,
		op:        base.JoinOpLeft,
		cond:      cond,
	}
}

func RightJoin(joinIndex uint8, cond string) base.IJoinCond {
	return &JoinCond{
		joinIndex: joinIndex,
		op:        base.JoinOpRight,
		cond:      cond,
	}
}

func InnerJoin(joinIndex uint8, cond string) base.IJoinCond {
	return &JoinCond{
		joinIndex: joinIndex,
		op:        base.JoinOpInner,
		cond:      cond,
	}
}
