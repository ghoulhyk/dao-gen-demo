// Code generated with the command 'go generate', DO NOT EDIT.

package base

type IDeleteHandler interface {
	Delete(IDeleterParams) (deleteRows uint64, err error)
}
