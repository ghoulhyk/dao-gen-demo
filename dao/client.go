// Code generated with the command 'go generate', DO NOT EDIT.

package dao

import (
	"demo/dao/base"
	"demo/db"
	"fmt"
	"xorm.io/xorm"
)

type Client struct {
	session            *xorm.Session
	transactionHandler base.ITransactionHandler
	baseSelectHandler  base.IBaseSelectHandler
	insertHandler      base.IInsertHandler
	deleteHandler      base.IDeleteHandler
	updateHandler      base.IUpdateHandler
	inserterGroup      *inserterGroup
	bulkInserterGroup  *bulkInserterGroup
	selectorGroup      *selectorGroup
	// joinSelectorGroup  *joinSelectorGroup
	deleterGroup *deleterGroup
	updaterGroup *updaterGroup
}

// region base.ITransactionHandler

func (c *Client) Insert(model base.IInserter) (insertCnt int64, pk interface{}, err error) {
	return c.getInsertHandler().Insert(model)
}

func (c *Client) BulkInsert(model base.IBulkInserter) (insertCnt int64, pkList []interface{}, err error) {
	return c.getInsertHandler().BulkInsert(model)
}

// endregion

// region base.IBaseSelectHandler

func (c *Client) Select(selector base.IBaseSelectorParams) (list interface{}, total uint64, err error) {
	return c.getBaseSelectHandler().Select(selector)
}

func (c *Client) Single(selector base.IBaseSelectorParams) (exist bool, model interface{}, err error) {
	return c.getBaseSelectHandler().Single(selector)
}

func (c *Client) Count(selector base.IBaseSelectorParams) (total uint64, err error) {
	return c.getBaseSelectHandler().Count(selector)
}

// endregion

// region base.IDeleteHandler

func (c *Client) Delete(deleter base.IDeleterParams) (deleteRows uint64, err error) {
	return c.getDeleteHandler().Delete(deleter)
}

// endregion

// region base.IUpdateHandler

func (c *Client) Update(updater base.IUpdaterParams) (updateRows uint64, err error) {
	return c.getUpdateHandler().Update(updater)
}

// endregion

// region base.ITransactionHandler

func (c *Client) StartTransaction() error {
	return c.getTransactionHandler().StartTransaction()
}

func (c *Client) Commit() error {
	return c.getTransactionHandler().Commit()
}

// 发生错误直接panic
func (c *Client) MustStartTransaction() {
	err := c.getTransactionHandler().StartTransaction()
	if err != nil {
		panic(fmt.Sprintf("StartTransaction fail [%v]", err.Error()))
	}
}

// Deprecated: 使用 MustStartTransaction
func (c *Client) StartTransactionOrPanic() {
	c.MustStartTransaction()
}

// 发生错误直接panic
func (c *Client) MustCommit() {
	err := c.getTransactionHandler().Commit()
	if err != nil {
		if c.IsInTransaction() {
			c.Rollback()
		}
		panic(fmt.Sprintf("CommitTransaction fail [%v]", err.Error()))
	}
}

// Deprecated: 使用 MustCommit
func (c *Client) CommitOrPanic() {
	c.MustCommit()
}

func (c *Client) Rollback() error {
	return c.getTransactionHandler().Rollback()
}

func (c *Client) IsInTransaction() bool {
	return c.getTransactionHandler().IsInTransaction()
}

// endregion

func (c *Client) Inserter() *inserterGroup {
	if c.inserterGroup == nil {
		c.inserterGroup = &inserterGroup{client: c}
	}
	return c.inserterGroup
}

func (c *Client) BulkInserter() *bulkInserterGroup {
	if c.bulkInserterGroup == nil {
		c.bulkInserterGroup = &bulkInserterGroup{client: c}
	}
	return c.bulkInserterGroup
}

func (c *Client) Selector() *selectorGroup {
	if c.selectorGroup == nil {
		c.selectorGroup = &selectorGroup{client: c}
	}
	return c.selectorGroup
}

// func (c *Client) JoinSelector() *joinSelectorGroup {
// 	if c.joinSelectorGroup == nil {
// 		c.joinSelectorGroup = &joinSelectorGroup{client: c}
// 	}
// 	return c.joinSelectorGroup
// }

func (c *Client) Deleter() *deleterGroup {
	if c.deleterGroup == nil {
		c.deleterGroup = &deleterGroup{client: c}
	}
	return c.deleterGroup
}

func (c *Client) Updater() *updaterGroup {
	if c.updaterGroup == nil {
		c.updaterGroup = &updaterGroup{client: c}
	}
	return c.updaterGroup
}

// region 内部对象生成

func (c *Client) getTransactionHandler() base.ITransactionHandler {
	if c.transactionHandler == nil {
		c.transactionHandler = &transactionHandler{session: c.session}
	}
	return c.transactionHandler
}

func (c *Client) getInsertHandler() base.IInsertHandler {
	if c.insertHandler == nil {
		c.insertHandler = &insertHandler{c: c}
	}
	return c.insertHandler
}

func (c *Client) getBaseSelectHandler() base.IBaseSelectHandler {
	if c.baseSelectHandler == nil {
		c.baseSelectHandler = &baseSelectHandler{
			c:                 c,
			selectHandler:     selectHandler{c: c},
			joinSelectHandler: joinSelectHandler{c: c},
		}
	}
	return c.baseSelectHandler
}

func (c *Client) getDeleteHandler() base.IDeleteHandler {
	if c.deleteHandler == nil {
		c.deleteHandler = &deleteHandler{c: c}
	}
	return c.deleteHandler
}

func (c *Client) getUpdateHandler() base.IUpdateHandler {
	if c.updateHandler == nil {
		c.updateHandler = &updateHandler{c: c}
	}
	return c.updateHandler
}

// endregion

func (c *Client) Release() error {
	c.session.Close()
	return nil
}

func (c *Client) TransactionFunc(fn func(client *Client) error) (err error) {
	c.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			c.TryRollbackAndRelease()
			if err1, ok := r.(error); ok {
				err = err1
			} else {
				err = fmt.Errorf("%v", err1)
			}
		}
		c.Release()
	}()
	err = fn(c)
	if err != nil {
		c.Rollback()
		return err
	}
	err = c.Commit()
	return err
}

func (c *Client) TryRollbackAndRelease() error {
	var err error
	if c.IsInTransaction() {
		err = c.getTransactionHandler().Rollback()
	}
	c.Release()
	return err
}

func (c *Client) GetSession() *xorm.Session {
	return c.session
}

func NewClient() *Client {
	return NewClientWithSession(db.NewSession())
}

// region TransactionFunc

func TransactionFunc(fn func(client *Client) error) error {
	return NewClient().TransactionFunc(fn)
}

func TransactionFunc1[A any](fn func(client *Client) (A, error)) (_ A, err error) {
	client := NewClient()
	client.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			client.TryRollbackAndRelease()
			if err1, ok := r.(error); ok {
				err = err1
			} else {
				err = fmt.Errorf("%v", err1)
			}
		}
		client.Release()
	}()
	a, err := fn(client)
	if err != nil {
		client.Rollback()
		return a, err
	}
	err = client.Commit()
	return a, err
}
func TransactionFunc2[A, B any](fn func(client *Client) (A, B, error)) (_ A, _ B, err error) {
	client := NewClient()
	client.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			client.TryRollbackAndRelease()
			if err1, ok := r.(error); ok {
				err = err1
			} else {
				err = fmt.Errorf("%v", err1)
			}
		}
		client.Release()
	}()
	a, b, err := fn(client)
	if err != nil {
		client.Rollback()
		return a, b, err
	}
	err = client.Commit()
	return a, b, err
}
func TransactionFunc3[A, B, C any](fn func(client *Client) (A, B, C, error)) (_ A, _ B, _ C, err error) {
	client := NewClient()
	client.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			client.TryRollbackAndRelease()
			if err1, ok := r.(error); ok {
				err = err1
			} else {
				err = fmt.Errorf("%v", err1)
			}
		}
		client.Release()
	}()
	a, b, c, err := fn(client)
	if err != nil {
		client.Rollback()
		return a, b, c, err
	}
	err = client.Commit()
	return a, b, c, err
}
func TransactionFunc4[A, B, C, D any](fn func(client *Client) (A, B, C, D, error)) (_ A, _ B, _ C, _ D, err error) {
	client := NewClient()
	client.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			client.TryRollbackAndRelease()
			if err1, ok := r.(error); ok {
				err = err1
			} else {
				err = fmt.Errorf("%v", err1)
			}
		}
		client.Release()
	}()
	a, b, c, d, err := fn(client)
	if err != nil {
		client.Rollback()
		return a, b, c, d, err
	}
	err = client.Commit()
	return a, b, c, d, err
}
func TransactionFunc5[A, B, C, D, E any](fn func(client *Client) (A, B, C, D, E, error)) (_ A, _ B, _ C, _ D, _ E, err error) {
	client := NewClient()
	client.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			client.TryRollbackAndRelease()
			if err1, ok := r.(error); ok {
				err = err1
			} else {
				err = fmt.Errorf("%v", err1)
			}
		}
		client.Release()
	}()
	a, b, c, d, e, err := fn(client)
	if err != nil {
		client.Rollback()
		return a, b, c, d, e, err
	}
	err = client.Commit()
	return a, b, c, d, e, err
}
func TransactionFunc6[A, B, C, D, E, F any](fn func(client *Client) (A, B, C, D, E, F, error)) (_ A, _ B, _ C, _ D, _ E, _ F, err error) {
	client := NewClient()
	client.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			client.TryRollbackAndRelease()
			if err1, ok := r.(error); ok {
				err = err1
			} else {
				err = fmt.Errorf("%v", err1)
			}
		}
		client.Release()
	}()
	a, b, c, d, e, f, err := fn(client)
	if err != nil {
		client.Rollback()
		return a, b, c, d, e, f, err
	}
	err = client.Commit()
	return a, b, c, d, e, f, err
}
func TransactionFunc7[A, B, C, D, E, F, G any](fn func(client *Client) (A, B, C, D, E, F, G, error)) (_ A, _ B, _ C, _ D, _ E, _ F, _ G, err error) {
	client := NewClient()
	client.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			client.TryRollbackAndRelease()
			if err1, ok := r.(error); ok {
				err = err1
			} else {
				err = fmt.Errorf("%v", err1)
			}
		}
		client.Release()
	}()
	a, b, c, d, e, f, g, err := fn(client)
	if err != nil {
		client.Rollback()
		return a, b, c, d, e, f, g, err
	}
	err = client.Commit()
	return a, b, c, d, e, f, g, err
}
func TransactionFunc8[A, B, C, D, E, F, G, H any](fn func(client *Client) (A, B, C, D, E, F, G, H, error)) (_ A, _ B, _ C, _ D, _ E, _ F, _ G, _ H, err error) {
	client := NewClient()
	client.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			client.TryRollbackAndRelease()
			if err1, ok := r.(error); ok {
				err = err1
			} else {
				err = fmt.Errorf("%v", err1)
			}
		}
		client.Release()
	}()
	a, b, c, d, e, f, g, h, err := fn(client)
	if err != nil {
		client.Rollback()
		return a, b, c, d, e, f, g, h, err
	}
	err = client.Commit()
	return a, b, c, d, e, f, g, h, err
}
func TransactionFunc9[A, B, C, D, E, F, G, H, I any](fn func(client *Client) (A, B, C, D, E, F, G, H, I, error)) (_ A, _ B, _ C, _ D, _ E, _ F, _ G, _ H, _ I, err error) {
	client := NewClient()
	client.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			client.TryRollbackAndRelease()
			if err1, ok := r.(error); ok {
				err = err1
			} else {
				err = fmt.Errorf("%v", err1)
			}
		}
		client.Release()
	}()
	a, b, c, d, e, f, g, h, i, err := fn(client)
	if err != nil {
		client.Rollback()
		return a, b, c, d, e, f, g, h, i, err
	}
	err = client.Commit()
	return a, b, c, d, e, f, g, h, i, err
}

// endregion

func NewClientWithSession(session *xorm.Session) *Client {
	return &Client{
		session: session,
	}
}
