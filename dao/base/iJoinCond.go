// Code generated with the command 'go generate', DO NOT EDIT.

package base

type IJoinCond interface {
	GetOp() JoinOp

	GetCond() string

	GetJoinIndex() uint8
}

type JoinOp string

const (
	JoinOpLeft  JoinOp = "LEFT"
	JoinOpRight JoinOp = "RIGHT"
	JoinOpInner JoinOp = "INNER"
)
