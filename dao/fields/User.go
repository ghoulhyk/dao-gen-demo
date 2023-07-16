// Code generated with the command 'go generate', DO NOT EDIT.

package fields

import (
	"demo/dao/base"
	"demo/dao/baseCond"
	"fmt"
	"strings"
)

type User struct {
	_tableInfo *base.TableInfo
	fields     []*baseCond.FieldCond
}

func NewUser(tableInfo *base.TableInfo) *User {
	return &User{
		_tableInfo: tableInfo,
	}
}

func (receiver User) GetFieldList() string {
	var columns []string
	for _, field := range receiver.fields {
		columns = append(columns, field.Sql())
	}
	return strings.Join(columns, ", ")
}

func (receiver *User) tableName() string {
	tableName := ""
	if receiver._tableInfo != nil && receiver._tableInfo.Alias != "" {
		tableName = receiver._tableInfo.Alias
		tableName = fmt.Sprintf("`%s`.", tableName)
	}
	return tableName
}

func (receiver *User) Id() *baseCond.FieldCond {
	tableName := receiver.tableName()
	fieldExpr := baseCond.NewFieldCond(tableName + "id")
	receiver.fields = append(receiver.fields, fieldExpr)
	return fieldExpr
}

func (receiver *User) Name() *baseCond.FieldCond {
	tableName := receiver.tableName()
	fieldExpr := baseCond.NewFieldCond(tableName + "name")
	receiver.fields = append(receiver.fields, fieldExpr)
	return fieldExpr
}

func (receiver *User) Gender() *baseCond.FieldCond {
	tableName := receiver.tableName()
	fieldExpr := baseCond.NewFieldCond(tableName + "gender")
	receiver.fields = append(receiver.fields, fieldExpr)
	return fieldExpr
}

func (receiver *User) CreateTime() *baseCond.FieldCond {
	tableName := receiver.tableName()
	fieldExpr := baseCond.NewFieldCond(tableName + "create_time")
	receiver.fields = append(receiver.fields, fieldExpr)
	return fieldExpr
}
