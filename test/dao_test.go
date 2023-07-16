package test

import (
	"demo/dao"
	"demo/dao/model/do"
	"demo/dao/util"
	"demo/dao/whereCond"
	"errors"
	"fmt"
	"testing"
)

func TestDao(t *testing.T) {
	initConf()
	connectDb()

	_, _, err := dao.TransactionFunc2(func(client *dao.Client) (int, int, error) {
		var isSuccess bool
		insertResult := client.Inserter().User().
			FillNotNilData(do.User{
				Name:   util.ToPtr("user_001"),
				Gender: util.ToPtr[uint8](1),
			}).
			MustInsert()
		isSuccess = insertResult.GetId() > 0
		if !isSuccess {
			return 0, 0, errors.New("插入失败")
		}
		insertResult1 := client.Inserter().User().
			FillNotNilData(do.User{
				Name: util.ToPtr("user_002"),
			}).
			SetGender(2).
			MustInsert()

		singleSelectResult := client.Selector().User().
			Where(func(cond *whereCond.User) {
				cond.And().Id().Equ(insertResult.GetId())
			}).
			MustSingle()
		isSuccess = singleSelectResult != nil
		if !isSuccess {
			return 0, 0, errors.New("single 查询失败 1")
		}
		singleSelectResult = client.Selector().User().MustSingleById(insertResult.GetId())
		isSuccess = singleSelectResult != nil
		if !isSuccess {
			return 0, 0, errors.New("single 查询失败 2")
		}
		fmt.Printf("%#v", singleSelectResult)

		listSelectResult := client.Selector().User().
			Where(func(cond *whereCond.User) {
				cond.And().Id().Equ(insertResult.GetId())
			}).
			OrWhere(func(cond *whereCond.User) {
				cond.And().Id().Equ(insertResult1.GetId())
			}).
			MustList()
		isSuccess = len(listSelectResult) == 2
		if !isSuccess {
			return 0, 0, errors.New("list 查询失败 1")
		}
		listSelectResult = client.Selector().User().
			Where(func(cond *whereCond.User) {
				cond.And().Id().Equ(insertResult.GetId())
				cond.Or().Id().Equ(insertResult1.GetId())
			}).
			MustList()
		isSuccess = len(listSelectResult) == 2
		if !isSuccess {
			return 0, 0, errors.New("list 查询失败 2")
		}
		listSelectResult = client.Selector().User().MustListByIds(insertResult.GetId(), insertResult1.GetId())
		isSuccess = len(listSelectResult) == 2
		if !isSuccess {
			return 0, 0, errors.New("list 查询失败 3")
		}
		fmt.Printf("%#v", listSelectResult)

		updateResult := client.Updater().User().
			Where(func(cond *whereCond.User) {
				cond.And().Id().Equ(insertResult.GetId())
			}).
			SetGender(3).
			MustUpdate()
		isSuccess = updateResult == 1
		if !isSuccess {
			return 0, 0, errors.New("更新失败")
		}

		//deleteRowResult := client.Deleter().User().
		//	Where(func(cond *whereCond.User) {
		//		cond.And().Id().Equ(insertResult.GetId())
		//	}).
		//	MustDelete()
		//isSuccess = deleteRowResult == 1
		//if !isSuccess {
		//	return  0,0,errors.New("删除失败 1")
		//}
		deleteResult := client.Deleter().User().MustDeleteById(insertResult1.GetId())
		isSuccess = deleteResult == true
		if !isSuccess {
			return 0, 0, errors.New("删除失败 2")
		}

		return 0, 0, nil
	})
	if err != nil {
		println(err.Error())
		return
	}
}
