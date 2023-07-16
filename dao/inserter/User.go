// Code generated with the command 'go generate', DO NOT EDIT.

package inserter

import (
	"demo/dao/base"
	"demo/dao/daoerrors"
	"demo/dao/model/do"

	"demo/dao/databaseDef"
	"demo/dao/util"

	"time"
)

// region inserterDataModel

type userDataModel struct {
	do.User
	_insertColumns map[string]byte
}

func (receiver userDataModel) RealInsertData() any {
	return &receiver.User
}

func (receiver userDataModel) GetInsertCols() []string {
	var result []string
	for col := range receiver._insertColumns {
		result = append(result, col)
	}
	return result
}

// endregion

// region inserter

// region inserterParams

type userInserterParams struct {
	insertData userDataModel
}

func (receiver *userInserterParams) TableInfo() base.TableInfo {
	return base.TableInfo{
		DatabaseInfo: util.ToPtr(databaseDef.Db1()),
		TableName:    "user",
	}
}

func (receiver *userInserterParams) InsertData() base.IInserterDataModel {
	return &receiver.insertData
}

func (receiver *userInserterParams) HasAutoincrPk() bool {
	return true
}

// endregion

type UserWrapper struct {
	clientHolder base.IClientHolder
	params       *userInserterParams
}

// 内部方法，不应调用
func (receiver *UserWrapper) Init(client base.IClient) *UserWrapper {
	clientHolder := &base.ClientHolder{}
	clientHolder.Init(client)
	receiver.clientHolder = clientHolder
	receiver.params = &userInserterParams{}
	receiver.params.insertData._insertColumns = map[string]byte{}
	return receiver
}

// region column setter

func (receiver *UserWrapper) SetId(id uint32) *UserWrapper {
	receiver.params.insertData.Id = &id
	receiver.params.insertData._insertColumns["id"] = 1
	return receiver
}

// 姓名
func (receiver *UserWrapper) SetName(name string) *UserWrapper {
	receiver.params.insertData.Name = &name
	receiver.params.insertData._insertColumns["name"] = 1
	return receiver
}
func (receiver *UserWrapper) SetGender(gender uint8) *UserWrapper {
	receiver.params.insertData.Gender = &gender
	receiver.params.insertData._insertColumns["gender"] = 1
	return receiver
}
func (receiver *UserWrapper) SetCreateTime(createTime time.Time) *UserWrapper {
	receiver.params.insertData.CreateTime = &createTime
	receiver.params.insertData._insertColumns["create_time"] = 1
	return receiver
}

// endregion

func (receiver *UserWrapper) FillNotNilData(data do.User) *UserWrapper {
	if data.Id != nil {
		receiver.SetId(*data.Id)
	}
	if data.Name != nil {
		receiver.SetName(*data.Name)
	}
	if data.Gender != nil {
		receiver.SetGender(*data.Gender)
	}
	if data.CreateTime != nil {
		receiver.SetCreateTime(*data.CreateTime)
	}

	return receiver
}

func (receiver *UserWrapper) FillNotNilDataFunc(fun func(data *do.User)) *UserWrapper {
	data := do.User{}
	fun(&data)
	receiver.FillNotNilData(data)
	return receiver
}

func (receiver *UserWrapper) Insert() (*do.User, error) {
	_, pk, err := receiver.clientHolder.GetClient().Insert(receiver.params)
	if err != nil {
		return nil, err
	}
	modelObj := receiver.params.insertData.RealInsertData().(*do.User)
	modelObj.SetId(uint32(pk.(int64)))
	return modelObj, nil
}

// 发生错误直接panic
func (receiver *UserWrapper) MustInsert() *do.User {
	insertResult, err := receiver.Insert()
	if err != nil {
		if receiver.clientHolder.GetClient().IsInTransaction() {
			receiver.clientHolder.GetClient().Rollback()
		}
		panic(daoerrors.InsertErr("user", err))
	}

	return insertResult
}

// Deprecated: 使用 MustInsert
func (receiver *UserWrapper) InsertOrPanic() *do.User {
	return receiver.MustInsert()
}

// endregion

// region bulkInserter

// region bulkInserterParams

type userBulkInserterParams struct {
	insertData []userDataModel
}

func (receiver *userBulkInserterParams) TableInfo() base.TableInfo {
	return base.TableInfo{
		DatabaseInfo: util.ToPtr(databaseDef.Db1()),
		TableName:    "user",
	}
}

func (receiver *userBulkInserterParams) InsertData() []base.IInserterDataModel {
	result := []base.IInserterDataModel{}
	for _, data := range receiver.insertData {
		result = append(result, data)
	}
	return result
}

func (receiver *userBulkInserterParams) HasAutoincrPk() bool {
	return true
}

// endregion

type UserBulkWrapper struct {
	clientHolder base.IClientHolder
	params       *userBulkInserterParams
}

// 内部方法，不应调用
func (receiver *UserBulkWrapper) Init(client base.IClient) *UserBulkWrapper {
	clientHolder := &base.ClientHolder{}
	clientHolder.Init(client)
	receiver.params = &userBulkInserterParams{}
	receiver.clientHolder = clientHolder
	return receiver
}

func (receiver *UserBulkWrapper) AddData(datas ...do.User) *UserBulkWrapper {
	for _, data := range datas {
		inserterModel := userDataModel{}
		inserterModel._insertColumns = map[string]byte{}

		if data.Id != nil {
			inserterModel.Id = data.Id
			inserterModel._insertColumns["id"] = 1
		}

		if data.Name != nil {
			inserterModel.Name = data.Name
			inserterModel._insertColumns["name"] = 1
		}

		if data.Gender != nil {
			inserterModel.Gender = data.Gender
			inserterModel._insertColumns["gender"] = 1
		}

		if data.CreateTime != nil {
			inserterModel.CreateTime = data.CreateTime
			inserterModel._insertColumns["create_time"] = 1
		}

		receiver.params.insertData = append(receiver.params.insertData, inserterModel)
	}
	return receiver
}

func (receiver *UserBulkWrapper) AddDataFunc(fun func(data *do.User)) *UserBulkWrapper {
	data := do.User{}
	fun(&data)
	receiver.AddData(data)
	return receiver
}

func (receiver *UserBulkWrapper) Insert() ([]*do.User, error) {
	_, pkList, err := receiver.clientHolder.GetClient().BulkInsert(receiver.params)
	if err != nil {
		return nil, err
	}

	result := make([]*do.User, len(receiver.params.insertData))
	for i, data := range receiver.params.insertData {
		result[i] = data.RealInsertData().(*do.User)
		result[i].SetId(uint32(pkList[i].(int64)))
	}

	return result, nil
}

// 发生错误直接panic
func (receiver *UserBulkWrapper) MustInsert() []*do.User {
	insertResults, err := receiver.Insert()
	if err != nil {
		if receiver.clientHolder.GetClient().IsInTransaction() {
			receiver.clientHolder.GetClient().Rollback()
		}
		panic(daoerrors.InsertErr("user", err))
	}

	return insertResults
}

// Deprecated: 使用 MustInsert
func (receiver *UserBulkWrapper) InsertOrPanic() []*do.User {
	return receiver.MustInsert()
}

// endregion
