package _db1

import (
	"time"
)

// User
// @database db_1_test
// @tableName user
// @comment 用户信息表
type User struct {
	Id         *uint32    `json:"id" colAttr:"'id' pk autoincr"`
	Name       *string    `json:"name" colAttr:"'name'"` // 姓名
	Gender     *uint8     `json:"gender" colAttr:"'gender'"`
	CreateTime *time.Time `json:"createTime" colAttr:"'create_time'"`
}
