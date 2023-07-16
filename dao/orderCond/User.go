// Code generated with the command 'go generate', DO NOT EDIT.

package orderCond

import (
	"demo/dao/base"
	"demo/dao/baseCond"
)

type User struct{ conds []base.IOrder }

// region basicOrder

func (receiver *User) Id() *baseCond.OrderBuilder {
	cond := baseCond.NewOrderBuilder("id")
	receiver.conds = append(receiver.conds, cond)
	return cond
}

// 姓名
func (receiver *User) Name() *baseCond.OrderBuilder {
	cond := baseCond.NewOrderBuilder("name")
	receiver.conds = append(receiver.conds, cond)
	return cond
}

func (receiver *User) Gender() *baseCond.OrderBuilder {
	cond := baseCond.NewOrderBuilder("gender")
	receiver.conds = append(receiver.conds, cond)
	return cond
}

func (receiver *User) CreateTime() *baseCond.OrderBuilder {
	cond := baseCond.NewOrderBuilder("create_time")
	receiver.conds = append(receiver.conds, cond)
	return cond
}

// endregion

func (receiver *User) GetOrderCondList() []base.IOrder {
	return receiver.conds
}

func (receiver *User) RAW(sql string) base.IOrder {
	cond := baseCond.OrderRaw(sql)
	receiver.conds = append(receiver.conds, cond)
	return cond
}
