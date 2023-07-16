// Code generated with the command 'go generate', DO NOT EDIT.

package dao

import (
	"demo/dao/base"
)

type baseSelectHandler struct {
	c                 *Client
	selectHandler     selectHandler
	joinSelectHandler joinSelectHandler
}

func (s baseSelectHandler) Select(i base.IBaseSelectorParams) (list interface{}, total uint64, err error) {
	if joinSelectorParams, ok := i.(base.IJoinSelectorParams); ok {
		return s.joinSelectHandler.Select(joinSelectorParams)
	}
	if selectorParams, ok := i.(base.ISelectorParams); ok {
		return s.selectHandler.Select(selectorParams)
	}
	panic("IBaseSelectorParams 不能直接传入 baseSelectHandler")
}

func (s baseSelectHandler) Single(i base.IBaseSelectorParams) (exist bool, model interface{}, err error) {
	if joinSelectorParams, ok := i.(base.IJoinSelectorParams); ok {
		return s.joinSelectHandler.Single(joinSelectorParams)
	}
	if selectorParams, ok := i.(base.ISelectorParams); ok {
		return s.selectHandler.Single(selectorParams)
	}
	panic("IBaseSelectorParams 不能直接传入 baseSelectHandler")
}

func (s baseSelectHandler) Count(i base.IBaseSelectorParams) (total uint64, err error) {
	if joinSelectorParams, ok := i.(base.IJoinSelectorParams); ok {
		return s.joinSelectHandler.Count(joinSelectorParams)
	}
	if selectorParams, ok := i.(base.ISelectorParams); ok {
		return s.selectHandler.Count(selectorParams)
	}
	panic("IBaseSelectorParams 不能直接传入 baseSelectHandler")
}
