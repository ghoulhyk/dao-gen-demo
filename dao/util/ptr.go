// Code generated with the command 'go generate', DO NOT EDIT.

package util

func ToPtr[T any](item T) *T {
	return &item
}

func FromPtr[T any](item *T) T {
	return *item
}
