// Code generated with the command 'go generate', DO NOT EDIT.

package dao

import (
	"demo/dao/base"
	"demo/dao/inserter"
)

type inserterGroup struct {
	client base.IClient
}

// User
func (receiver inserterGroup) User() *inserter.UserWrapper {
	return (&inserter.UserWrapper{}).Init(receiver.client)
}
