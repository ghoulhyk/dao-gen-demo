// Code generated with the command 'go generate', DO NOT EDIT.

package baseCond

import (
	"fmt"
	"regexp"
	"strings"
)

type OrderBuilder struct {
	column    string
	fields    []any
	tableName string
	orderType string
}

func NewOrderBuilder(column string) *OrderBuilder {
	o := OrderBuilder{}
	o.setColumn(column)
	return &o
}

func (o *OrderBuilder) setColumn(column string) {
	tableNameRegexResult := regexp.MustCompile("^((\\w+)\\.)?([\\w\\W]+)$").FindStringSubmatch(column)
	tableName := ""
	if tableNameRegexResult[1] != "" {
		tableName = strings.TrimRight(tableNameRegexResult[1], ".")
		column = tableNameRegexResult[3]
	}
	o.column = column
	o.tableName = tableName
}

func (o *OrderBuilder) TableName(tableName string) *OrderBuilder {
	o.tableName = tableName
	return o
}

func (o *OrderBuilder) Desc() *OrderBuilder {
	o.orderType = "DESC"
	return o
}

func (o *OrderBuilder) Asc() *OrderBuilder {
	o.orderType = "ASC"
	return o
}

func (o *OrderBuilder) FieldDesc(fields ...any) *OrderBuilder {
	o.Field(fields)
	o.Desc()
	return o
}

func (o *OrderBuilder) Field(fields ...any) *OrderBuilder {
	o.fields = fields
	return o
}

func (o *OrderBuilder) OrderSql() string {
	shouldQuoteList := []string{"order"}
	column := o.column
	tableName := o.tableName
	for _, shouldQuote := range shouldQuoteList {
		if strings.ToLower(shouldQuote) == column {
			column = "`" + column + "`"
			break
		}
	}
	if tableName != "" {
		tableName = "`" + tableName + "`" + "."
	}
	if len(o.fields) == 0 {
		return tableName + column + " " + o.orderType
	}
	fields := make([]string, len(o.fields)+1)
	fields[0] = column
	for i, f := range o.fields {
		fields[i+1] = fmt.Sprintf("\"%v\"", f)
	}
	return fmt.Sprintf("FIELD(%v) %v", strings.Join(fields, ","), o.orderType)
}

func OrderDesc(column string) *OrderBuilder {
	return &OrderBuilder{
		column:    column,
		orderType: "DESC",
	}
}

func OrderAsc(column string) *OrderBuilder {
	return &OrderBuilder{
		column:    column,
		orderType: "ASC",
	}
}

func OrderFieldDesc(column string, fields ...any) *OrderBuilder {
	v := make([]string, len(fields)+1)
	v[0] = column
	for i, f := range fields {
		v[i+1] = fmt.Sprintf("\"%v\"", f)
	}
	return &OrderBuilder{
		column:    fmt.Sprintf("FIELD(%v)", strings.Join(v, ",")),
		orderType: "DESC",
	}
}

func OrderFieldAsc(column string, fields ...any) *OrderBuilder {
	v := make([]string, len(fields)+1)
	v[0] = column
	for i, f := range fields {
		v[i+1] = fmt.Sprintf("\"%v\"", f)
	}
	return &OrderBuilder{
		column:    fmt.Sprintf("FIELD(%v)", strings.Join(v, ",")),
		orderType: "ASC",
	}
}

type orderRaw struct {
	sql string
}

func (o orderRaw) OrderSql() string {
	return o.sql
}

func OrderRaw(sql string) *orderRaw {
	return &orderRaw{
		sql: sql,
	}
}
