// Code generated with the command 'go generate', DO NOT EDIT.

package base

type TableInfo struct {
	DatabaseInfo *DatabaseInfo
	TableName    string
	Alias        string
}

func (receiver TableInfo) FullName() string {
	var fillName string
	if receiver.DatabaseInfo != nil {
		fillName += string(*receiver.DatabaseInfo)
		fillName += "."
	}
	fillName += receiver.TableName
	return fillName
}
