// Code generated with the command 'go generate', DO NOT EDIT.

package do

import (
	"encoding/json"

	"time"
)

type User struct {
	Id         *uint32    `json:"id" xorm:"'id' pk autoincr"`
	Name       *string    `json:"name" xorm:"'name'"` // 姓名
	Gender     *uint8     `json:"gender" xorm:"'gender'"`
	CreateTime *time.Time `json:"createTime" xorm:"'create_time'"`
}

// region getterSetter

// region id

func (receiver *User) GetId() uint32 {
	if receiver.Id == nil {
		var _default uint32
		return _default
	}
	return *receiver.Id
}
func (receiver *User) GetIdPtr() *uint32 {
	return receiver.Id
}
func (receiver *User) IsIdNil() bool {
	return receiver.Id == nil
}
func (receiver *User) SetId(id uint32) {
	receiver.Id = &id
}
func (receiver *User) SetIdPtr(id *uint32) {
	receiver.Id = id
}

// endregion

// region name

// 姓名
func (receiver *User) GetName() string {
	if receiver.Name == nil {
		var _default string
		return _default
	}
	return *receiver.Name
}

// 姓名
func (receiver *User) GetNamePtr() *string {
	return receiver.Name
}

// 姓名
func (receiver *User) IsNameNil() bool {
	return receiver.Name == nil
}

// 姓名
func (receiver *User) SetName(name string) {
	receiver.Name = &name
}

// 姓名
func (receiver *User) SetNamePtr(name *string) {
	receiver.Name = name
}

// endregion

// region gender

func (receiver *User) GetGender() uint8 {
	if receiver.Gender == nil {
		var _default uint8
		return _default
	}
	return *receiver.Gender
}
func (receiver *User) GetGenderPtr() *uint8 {
	return receiver.Gender
}
func (receiver *User) IsGenderNil() bool {
	return receiver.Gender == nil
}
func (receiver *User) SetGender(gender uint8) {
	receiver.Gender = &gender
}
func (receiver *User) SetGenderPtr(gender *uint8) {
	receiver.Gender = gender
}

// endregion

// region createTime

func (receiver *User) GetCreateTime() time.Time {
	if receiver.CreateTime == nil {
		var _default time.Time
		return _default
	}
	return *receiver.CreateTime
}
func (receiver *User) GetCreateTimePtr() *time.Time {
	return receiver.CreateTime
}
func (receiver *User) IsCreateTimeNil() bool {
	return receiver.CreateTime == nil
}
func (receiver *User) SetCreateTime(createTime time.Time) {
	receiver.CreateTime = &createTime
}
func (receiver *User) SetCreateTimePtr(createTime *time.Time) {
	receiver.CreateTime = createTime
}

// endregion

// endregion

func (receiver *User) String() string {
	data, _ := json.Marshal(receiver)
	return string(data)
}
