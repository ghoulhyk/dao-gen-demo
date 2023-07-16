// Code generated with the command 'go generate', DO NOT EDIT.

package base

type IOrder interface {
	OrderSql() string
}

type IOrdersWrapper interface {
	OrderList() []IOrder
}
