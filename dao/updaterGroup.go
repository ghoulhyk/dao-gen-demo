// Code generated with the command 'go generate', DO NOT EDIT.

package dao

import (
	"demo/dao/base"
	"demo/dao/updater"
)

type updaterGroup struct {
	client base.IClient
}

// User
func (receiver updaterGroup) User() *updater.UserWrapper {
	return (&updater.UserWrapper{}).Init(receiver.client)
}
