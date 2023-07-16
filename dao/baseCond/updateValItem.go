// Code generated with the command 'go generate', DO NOT EDIT.

package baseCond

type UpdateValItem struct {
	columnName string
	val        interface{}
	rawSql     string
	isRawSql   bool
}

func NewUpdateValItem(column string) *UpdateValItem {
	return &UpdateValItem{
		columnName: column,
	}
}

func (u UpdateValItem) ColumnName() string {
	return u.columnName
}

func (u UpdateValItem) GetVal() interface{} {
	return u.val
}

func (u UpdateValItem) GetRawSql() string {
	return u.rawSql
}

func (u UpdateValItem) IsRawSql() bool {
	return u.isRawSql
}

func (u *UpdateValItem) SetVal(val interface{}) {
	u.SetPtr(&val)
}

func (u *UpdateValItem) SetPtr(val interface{}) {
	u.isRawSql = false
	u.val = val
}

func (u *UpdateValItem) SetRawSql(sql string) {
	u.isRawSql = true
	u.rawSql = sql
}
