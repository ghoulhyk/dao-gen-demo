// Code generated with the command 'go generate', DO NOT EDIT.

package dao

import (
	"demo/dao/base"
	"reflect"
	"sort"
	"strings"
	"xorm.io/xorm"
)

type joinSelectHandler struct {
	c *Client
}

func (s joinSelectHandler) Select(i base.IJoinSelectorParams) (list interface{}, total uint64, err error) {
	list = reflect.New(reflect.SliceOf(reflect.TypeOf(i.ProvideResultData()).Elem())).Interface()
	pageCond := i.GetPageCond()
	orderCondList := i.GetOrderCondList()
	whereCondList := i.GetWhereCondList()
	fieldList := i.GetFieldList()
	joins := i.Joins()
	tableInfo := i.TableInfo()
	session := s.c.session.Table(tableInfo.FullName())
	if tableInfo.Alias != "" {
		session.Alias(tableInfo.Alias)
	}

	join(session, joins)

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
		cnt, err = session.FindAndCount(list)
		if err != nil {
			return
		}

		total = uint64(cnt)
	} else {
		err = session.Find(list)
		if err != nil {
			return
		}

		total = uint64(reflect.ValueOf(list).Elem().Len())
	}
	return
}

func (s joinSelectHandler) Single(i base.IJoinSelectorParams) (exist bool, model interface{}, err error) {
	model = i.ProvideResultData()
	pageCond := i.GetPageCond()
	orderCondList := i.GetOrderCondList()
	whereCondList := i.GetWhereCondList()
	fieldList := i.GetFieldList()
	joins := i.Joins()
	tableInfo := i.TableInfo()
	session := s.c.session.Table(tableInfo.FullName())
	if tableInfo.Alias != "" {
		session.Alias(tableInfo.Alias)
	}

	join(session, joins)

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

func (s joinSelectHandler) Count(i base.IJoinSelectorParams) (total uint64, err error) {
	whereCondList := i.GetWhereCondList()
	joins := i.Joins()
	tableInfo := i.TableInfo()
	session := s.c.session.Table(tableInfo.FullName())
	if tableInfo.Alias != "" {
		session.Alias(tableInfo.Alias)
	}

	join(session, joins)

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

func join(session *xorm.Session, joins map[base.TableInfo]base.IJoinCond) {
	joinList := []struct {
		t base.TableInfo
		j base.IJoinCond
	}{}
	for tableInfo, joinCond := range joins {
		joinList = append(joinList, struct {
			t base.TableInfo
			j base.IJoinCond
		}{tableInfo, joinCond})
	}
	sort.Slice(joinList, func(i, j int) bool {
		return joinList[i].j.GetJoinIndex() < joinList[j].j.GetJoinIndex()
	})
	for _, join := range joinList {
		tableInfo := join.t
		joinCond := join.j
		tableName := tableInfo.FullName()
		if tableInfo.Alias != "" {
			tableName += " " + tableInfo.Alias
		}
		session.Join(string(joinCond.GetOp()), tableName, joinCond.GetCond())
	}
}
