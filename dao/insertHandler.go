// Code generated with the command 'go generate', DO NOT EDIT.

package dao

import (
	"demo/dao/base"
	"demo/dao/util"
	"errors"
	"strconv"
	"sync"
)

type insertHandler struct {
	c  *Client
	mu sync.Mutex
}

func (i *insertHandler) Insert(model base.IInserter) (insertCnt int64, pk interface{}, err error) {
	i.mu.Lock()
	defer i.mu.Unlock()
	tableInfo := model.TableInfo()
	session := i.c.session.Table(tableInfo.FullName())
	if tableInfo.Alias != "" {
		session.Alias(tableInfo.Alias)
	}
	insertData := model.InsertData()
	insertColStrings := insertData.GetInsertCols()

	insertCnt, err = session.Cols(insertColStrings...).Insert(insertData.RealInsertData())
	if err != nil {
		return
	}
	if insertCnt == 0 {
		err = errors.New("插入失败,插入数量为0")
		return
	}
	if model.HasAutoincrPk() {
		var results []map[string][]byte
		results, err = i.c.session.Query("SELECT LAST_INSERT_ID() AS lastId")
		var lastId int64
		lastId, err = strconv.ParseInt(string(results[0]["lastId"]), 10, 0)
		if err != nil {
			return
		}
		pk = lastId
	}
	return
}

func (i *insertHandler) BulkInsert(model base.IBulkInserter) (insertCnt int64, pkList []interface{}, err error) {
	i.mu.Lock()
	defer i.mu.Unlock()
	tableInfo := model.TableInfo()
	session := i.c.session.Table(tableInfo.FullName())
	if tableInfo.Alias != "" {
		session.Alias(tableInfo.Alias)
	}
	insertData := model.InsertData()
	realInsertData := []any{}
	insertColStrings := []string{}

	for _, datum := range insertData {
		insertColStrings = append(insertColStrings, datum.GetInsertCols()...)
		realInsertData = append(realInsertData, datum.RealInsertData())
	}
	insertColStrings = util.Uniq(insertColStrings)

	if !i.c.IsInTransaction() {
		i.c.StartTransaction()
		defer func() {
			if e := recover(); e != nil {
				i.c.Rollback()
				panic(e)
			}
			if err == nil {
				i.c.Commit()
			} else {
				i.c.Rollback()
			}
		}()
	}

	insertCnt, err = session.Cols(insertColStrings...).Insert(realInsertData)
	if err != nil {
		return
	}
	if insertCnt == 0 {
		err = errors.New("插入失败,插入数量为0")
		return
	}
	if insertCnt != int64(len(insertData)) {
		err = errors.New("插入失败,实际插入数量与预期不符")
		return
	}
	if model.HasAutoincrPk() {
		var results []map[string][]byte
		results, err = i.c.session.Query("SELECT LAST_INSERT_ID() AS lastId")
		var lastId int64
		lastId, err = strconv.ParseInt(string(results[0]["lastId"]), 10, 0)
		if err != nil {
			return
		}
		for _ = range insertData {
			pkList = append(pkList, lastId)
			lastId++
		}
	}
	return
}
