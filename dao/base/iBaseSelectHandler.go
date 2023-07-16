// Code generated with the command 'go generate', DO NOT EDIT.

package base

type IBaseSelectHandler interface {
	Select(IBaseSelectorParams) (list interface{}, total uint64, err error)
	Single(IBaseSelectorParams) (exist bool, model interface{}, err error)
	Count(IBaseSelectorParams) (total uint64, err error)
}
