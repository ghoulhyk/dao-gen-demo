// Code generated with the command 'go generate', DO NOT EDIT.

package dao

import (
	"demo/dao/base"
	"demo/dao/selector"
)

type selectorGroup struct {
	client base.IClient
}

// User
func (receiver selectorGroup) User() *selector.UserWrapper {
	return (&selector.UserWrapper{}).Init(receiver.client)
}
