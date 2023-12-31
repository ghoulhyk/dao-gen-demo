// Code generated with the command 'go generate', DO NOT EDIT.

package daoerrors

import "fmt"

type DatabaseError struct {
	TableName string
	Operation string
	Err       error
}

func InsertErr(tableName string, err error) DatabaseError {
	return DatabaseError{
		TableName: tableName,
		Operation: "insert",
		Err:       err,
	}
}

func UpdateErr(tableName string, err error) DatabaseError {
	return DatabaseError{
		TableName: tableName,
		Operation: "update",
		Err:       err,
	}
}

func DeleteErr(tableName string, err error) DatabaseError {
	return DatabaseError{
		TableName: tableName,
		Operation: "delete",
		Err:       err,
	}
}

func SelectErr(tableName string, err error) DatabaseError {
	return DatabaseError{
		TableName: tableName,
		Operation: "select",
		Err:       err,
	}
}

func (receiver DatabaseError) Error() string {
	return fmt.Sprintf("%s %s 失败 (%s)", receiver.Operation, receiver.TableName, receiver.Err.Error())
}
