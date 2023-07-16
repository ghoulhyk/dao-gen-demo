// Code generated with the command 'go generate', DO NOT EDIT.

package dao

import (
	"demo/dao/base"
	"demo/dao/deleter"
)

type deleterGroup struct {
	client base.IClient
}

// User
func (receiver deleterGroup) User() *deleter.UserWrapper {
	return (&deleter.UserWrapper{}).Init(receiver.client)
}
