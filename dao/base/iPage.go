// Code generated with the command 'go generate', DO NOT EDIT.

package base

type IPage interface {
	HasLimit() bool
	PageSize() int
	Offset() int
}
