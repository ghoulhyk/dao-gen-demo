// Code generated with the command 'go generate', DO NOT EDIT.

package dao

import (
	"demo/dao/base"
)

type updateHandler struct {
	c *Client
}

func (u updateHandler) Update(updater base.IUpdaterParams) (updateRows uint64, err error) {
	updateValItemList := updater.GetUpdateValItemList()
	whereCondList := updater.GetWhereCondList()
	limit := updater.GetLimit()

	tableInfo := updater.TableInfo()
	session := u.c.session.Table(tableInfo.FullName())
	if tableInfo.Alias != "" {
		session.Alias(tableInfo.Alias)
	}

	updateValMap := map[string]interface{}{}

	for _, item := range *updateValItemList {
		if item.IsRawSql() {
			session.SetExpr(item.ColumnName(), item.GetRawSql())
		} else {
			updateValMap[item.ColumnName()] = item.GetVal()
		}
	}

	for _, item := range *whereCondList {
		if item.GetOp() == base.Op_OR {
			session = session.Or(item.GetWhereBuilder())
		} else {
			session = session.And(item.GetWhereBuilder())
		}
	}

	if limit >= 0 {
		session.Limit(limit)
	}

	var cnt int64
	cnt, err = session.Update(updateValMap)
	if err != nil {
		return
	}

	updateRows = uint64(cnt)

	return
}
