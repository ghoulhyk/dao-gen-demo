// Code generated with the command 'go generate', DO NOT EDIT.

package dao

import (
	"demo/dao/base"
	"demo/dao/inserter"
)

type bulkInserterGroup struct {
	client base.IClient
}

// User
func (receiver bulkInserterGroup) User() *inserter.UserBulkWrapper {
	return (&inserter.UserBulkWrapper{}).Init(receiver.client)
}
