// Code generated with the command 'go generate', DO NOT EDIT.

package base

type IUpdateHandler interface {
	Update(IUpdaterParams) (updateRows uint64, err error)
}
