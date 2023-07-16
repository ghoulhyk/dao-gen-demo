// Code generated with the command 'go generate', DO NOT EDIT.

package dao

import (
	"demo/dao/base"
	"reflect"
	"strings"
)

type selectHandler struct {
	c *Client
}

func (s selectHandler) Select(i base.ISelectorParams) (list interface{}, total uint64, err error) {
	result := reflect.New(reflect.SliceOf(reflect.TypeOf(i.ProvideResultData()).Elem())).Interface()
	pageCond := i.GetPageCond()
	orderCondList := i.GetOrderCondList()
	whereCondList := i.GetWhereCondList()
	fieldList := i.GetFieldList()
	tableInfo := i.TableInfo()
	session := s.c.session.Table(tableInfo.FullName())
	if tableInfo.Alias != "" {
		session.Alias(tableInfo.Alias)
	}

	shouldUsePage := pageCond != nil && pageCond.HasLimit()
	if shouldUsePage {
		session.Limit(pageCond.PageSize(), pageCond.Offset())
	}

	for _, item := range *orderCondList {
		session.OrderBy(item.OrderSql())
	}

	for _, item := range *whereCondList {
		if item.GetOp() == base.Op_OR {
			session = session.Or(item.GetWhereBuilder())
		} else {
			session = session.And(item.GetWhereBuilder())
		}
	}

	var fieldStrList []string
	for _, item := range *fieldList {
		fieldStrList = append(fieldStrList, item.GetFieldList())
	}
	session.Select(strings.Join(fieldStrList, ", "))

	if shouldUsePage {
		var cnt int64
		cnt, err = session.FindAndCount(result)
		if err != nil {
			return
		}

		total = uint64(cnt)
	} else {
		err = session.Find(result)
		if err != nil {
			return
		}

		total = uint64(reflect.ValueOf(result).Elem().Len())
	}
	return reflect.ValueOf(result).Elem().Interface(), total, nil
}

func (s selectHandler) Single(i base.ISelectorParams) (exist bool, model interface{}, err error) {
	model = i.ProvideResultData()
	pageCond := i.GetPageCond()
	orderCondList := i.GetOrderCondList()
	whereCondList := i.GetWhereCondList()
	fieldList := i.GetFieldList()
	tableInfo := i.TableInfo()
	session := s.c.session.Table(tableInfo.FullName())
	if tableInfo.Alias != "" {
		session.Alias(tableInfo.Alias)
	}

	shouldUsePage := pageCond != nil && pageCond.HasLimit()
	if shouldUsePage {
		session.Limit(pageCond.PageSize(), pageCond.Offset())
	}

	for _, item := range *orderCondList {
		session.OrderBy(item.OrderSql())
	}

	for _, item := range *whereCondList {
		if item.GetOp() == base.Op_OR {
			session = session.Or(item.GetWhereBuilder())
		} else {
			session = session.And(item.GetWhereBuilder())
		}
	}

	var fieldStrList []string
	for _, item := range *fieldList {
		fieldStrList = append(fieldStrList, item.GetFieldList())
	}
	session.Select(strings.Join(fieldStrList, ", "))

	exist, err = session.Get(model)
	return
}

func (s selectHandler) Count(i base.ISelectorParams) (total uint64, err error) {
	whereCondList := i.GetWhereCondList()
	tableInfo := i.TableInfo()
	session := s.c.session.Table(tableInfo.FullName())
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

	cnt, err := session.Count()
	if err != nil {
		return
	}
	total = uint64(cnt)
	return
}
