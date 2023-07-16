// Code generated with the command 'go generate', DO NOT EDIT.

package dao

import (
	"demo/dao/base"
)

type deleteHandler struct {
	c *Client
}

func (d deleteHandler) Delete(i base.IDeleterParams) (deleteRows uint64, err error) {
	whereCondList := i.GetWhereCondList()
	tableInfo := i.TableInfo()
	session := d.c.session.Table(tableInfo.FullName())
	if tableInfo.Alias != "" {
		session.Alias(tableInfo.Alias)
	}

	for _, item := range *whereCondList {
		if item.GetOp() == base.Op_OR {
			session = session.Or(item.GetWhereBuilder())
		} else {
			session = session.And(item.GetWhereBuilder())
		}
	}

	var cnt int64
	placeholderObj := placeholder{} // 啥用也没有，纯为满足xorm
	cnt, err = session.Delete(&placeholderObj)
	if err != nil {
		return
	}

	deleteRows = uint64(cnt)

	return
}

// 啥用也没有，纯为满足xorm
type placeholder struct {
}
