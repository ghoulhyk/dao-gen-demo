// Code generated with the command 'go generate', DO NOT EDIT.

package selector

import (
	"demo/dao/base"
	"demo/dao/baseCond"
	"demo/dao/baseModel"
	"demo/dao/daoerrors"
	"demo/dao/fields"
	"demo/dao/model/do"
	"demo/dao/orderCond"
	"demo/dao/whereCond"

	"demo/dao/databaseDef"
	"demo/dao/util"
)

// region selectorParams

type userSelectorParams struct {
	whereCondList []base.IWhereCond
	page          base.IPage
	orderList     []base.IOrder
	fieldList     []base.IFieldList
}

func (receiver *userSelectorParams) TableInfo() base.TableInfo {
	return base.TableInfo{
		DatabaseInfo: util.ToPtr(databaseDef.Db1()),
		TableName:    "user",
	}
}

func (receiver *userSelectorParams) ProvideResultData() interface{} {
	return &do.User{}
}

func (receiver *userSelectorParams) GetWhereCondList() *[]base.IWhereCond {
	return &receiver.whereCondList
}

func (receiver *userSelectorParams) GetPageCond() base.IPage {
	return receiver.page
}

func (receiver *userSelectorParams) GetOrderCondList() *[]base.IOrder {
	return &receiver.orderList
}

func (receiver *userSelectorParams) GetFieldList() *[]base.IFieldList {
	return &receiver.fieldList
}

// endregion

type UserWrapper struct {
	_count
	_exist
	clientHolder base.IClientHolder
	params       *userSelectorParams
}

func (receiver *UserWrapper) Init(client base.IClient) *UserWrapper {
	clientHolder := &base.ClientHolder{}
	clientHolder.Init(client)
	receiver.clientHolder = clientHolder
	receiver.params = &userSelectorParams{}
	receiver._count = _count{
		clientHolder: receiver.clientHolder,
		params:       receiver.params,
	}
	receiver._exist = _exist{
		_count: receiver._count,
	}
	return receiver
}

func (receiver *UserWrapper) Where(fun func(cond *whereCond.User)) *UserWrapper {
	return receiver.AndWhere(fun)
}

func (receiver *UserWrapper) AndWhere(fun func(cond *whereCond.User)) *UserWrapper {
	cond := whereCond.New_User(base.Op_AND, util.ToPtr(receiver.params.TableInfo()))
	fun(cond)
	receiver.params.whereCondList = append(receiver.params.whereCondList, cond)
	return receiver
}

func (receiver *UserWrapper) OrWhere(fun func(cond *whereCond.User)) *UserWrapper {
	cond := whereCond.New_User(base.Op_OR, util.ToPtr(receiver.params.TableInfo()))
	fun(cond)
	receiver.params.whereCondList = append(receiver.params.whereCondList, cond)
	return receiver
}

func (receiver *UserWrapper) ClearWhereCond() *UserWrapper {
	receiver.params.whereCondList = []base.IWhereCond{}
	return receiver
}

func (receiver *UserWrapper) WithField(fun func(fields *fields.User)) *UserWrapper {
	cond := fields.NewUser(util.ToPtr(receiver.params.TableInfo()))
	fun(cond)
	receiver.params.fieldList = append(receiver.params.fieldList, cond)
	return receiver
}

func (receiver *UserWrapper) WithoutField(fun func(fields *fields.User)) *UserWrapper {
	cond := fields.NewUser(util.ToPtr(receiver.params.TableInfo()))
	fun(cond)
	receiver.params.fieldList = append(receiver.params.fieldList, cond)
	return receiver
}

func (receiver *UserWrapper) Page(page base.IPage) *UserWrapper {
	receiver.params.page = page
	return receiver
}

func (receiver *UserWrapper) OrderBy(fun func(orders *orderCond.User)) *UserWrapper {
	cond := &orderCond.User{}
	fun(cond)
	receiver.params.orderList = append(receiver.params.orderList, cond.GetOrderCondList()...)
	return receiver
}

func (receiver *UserWrapper) List() ([]do.User, error) {
	list, _, err := receiver.clientHolder.GetClient().Select(receiver.params)
	if err != nil {
		return nil, err
	}
	selectResult := list.([]do.User)
	if selectResult == nil {
		selectResult = []do.User{}
	}
	return selectResult, nil
}

func (receiver *UserWrapper) Single() (exist bool, model *do.User, err error) {
	exist, single, err := receiver.clientHolder.GetClient().Single(receiver.params)
	if err != nil {
		return
	}
	if !exist {
		return
	}
	model = single.(*do.User)
	return
}

func (receiver *UserWrapper) Pagination() (*baseModel.Pagination[do.User], error) {
	list, total, err := receiver.clientHolder.GetClient().Select(receiver.params)
	if err != nil {
		return nil, err
	}
	selectResult := list.([]do.User)
	if selectResult == nil {
		selectResult = []do.User{}
	}
	return baseModel.NewPagination(selectResult, total, receiver.params.page), nil
}

// 发生错误直接panic
// 长度为0 不会 panic
func (receiver *UserWrapper) MustList() []do.User {
	result, err := receiver.List()
	if err != nil {
		if receiver.clientHolder.GetClient().IsInTransaction() {
			receiver.clientHolder.GetClient().Rollback()
		}
		panic(daoerrors.SelectErr("user", err))
	}
	return result
}

// 发生错误直接panic
// 未找到 不会 panic
func (receiver *UserWrapper) MustSingle() *do.User {
	_, result, err := receiver.Single()
	if err != nil {
		if receiver.clientHolder.GetClient().IsInTransaction() {
			receiver.clientHolder.GetClient().Rollback()
		}
		panic(daoerrors.SelectErr("user", err))
	}
	return result
}

// 发生错误直接panic
// 长度为0 不会 panic
func (receiver *UserWrapper) MustPagination() *baseModel.Pagination[do.User] {
	result, err := receiver.Pagination()
	if err != nil {
		if receiver.clientHolder.GetClient().IsInTransaction() {
			receiver.clientHolder.GetClient().Rollback()
		}
		panic(daoerrors.SelectErr("user", err))
	}
	return result
}

// Deprecated: 使用 MustList
func (receiver *UserWrapper) ListOrPanic() []do.User {
	return receiver.MustList()
}

// Deprecated: 使用 MustSingle
func (receiver *UserWrapper) SingleOrPanic() *do.User {
	return receiver.MustSingle()
}

// Deprecated: 使用 MustPagination
func (receiver *UserWrapper) PaginationOrPanic() *baseModel.Pagination[do.User] {
	return receiver.MustPagination()
}

func (receiver *UserWrapper) Chunk(cb func(list []do.User) (stop bool), pageSize int) {
	page := 0
	for true {
		pageCdt := baseCond.PageCdtLimited(page, pageSize)
		list, _ := receiver.Page(pageCdt).List()
		if len(list) == 0 {
			break
		}
		stop := cb(list)
		if stop {
			break
		}
		page++
	}
}

func (receiver *UserWrapper) SingleById(id uint32) (exist bool, model *do.User, err error) {
	localVar := UserWrapper{}

	localVar.Init(receiver.clientHolder.GetClient())
	localVar.params.page = receiver.params.page
	localVar.params.orderList = receiver.params.orderList
	localVar.params.fieldList = receiver.params.fieldList

	return localVar.AndWhere(func(cond *whereCond.User) {
		cond.And().Id().Equ(id)
	}).Single()
}

// 发生错误直接panic
// 未找到 不会 panic
func (receiver *UserWrapper) MustSingleById(id uint32) *do.User {
	_, result, err := receiver.SingleById(id)
	if err != nil {
		panic(daoerrors.SelectErr("user", err))
	}
	return result
}

// Deprecated: 使用 MustSingleById
func (receiver *UserWrapper) SingleByIdOrPanic(id uint32) *do.User {
	return receiver.MustSingleById(id)
}

func (receiver *UserWrapper) ListByIds(ids ...uint32) (model []do.User, err error) {
	localVar := UserWrapper{}

	localVar.Init(receiver.clientHolder.GetClient())
	localVar.params.page = receiver.params.page
	localVar.params.orderList = receiver.params.orderList
	localVar.params.fieldList = receiver.params.fieldList

	return localVar.AndWhere(func(cond *whereCond.User) {
		cond.And().Id().In(ids...)
	}).List()
}

// 发生错误直接panic
// 长度为0 不会 panic
func (receiver *UserWrapper) MustListByIds(ids ...uint32) []do.User {
	result, err := receiver.ListByIds(ids...)
	if err != nil {
		panic(daoerrors.SelectErr("user", err))
	}
	return result
}

// Deprecated: 使用 MustListByIds
func (receiver *UserWrapper) ListByIdsOrPanic(ids ...uint32) []do.User {
	return receiver.MustListByIds(ids...)
}
