// Code generated with the command 'go generate', DO NOT EDIT.

package baseModel

import (
	"demo/dao/base"
	"math"
)

type Pagination[T any] struct {
	Data    []T        `json:"list"`
	Total   uint64     `json:"total"`
	PageCdt base.IPage `json:"-"`
}

func (receiver Pagination[T]) PageCount() uint32 {
	if receiver.PageCdt == nil || !receiver.PageCdt.HasLimit() {
		return 1
	}
	return uint32(math.Ceil(float64(receiver.Total) / float64(receiver.PageCdt.PageSize())))
}

func (receiver Pagination[T]) Each(fn func(item T)) {
	for _, item := range receiver.Data {
		fn(item)
	}
}

func (receiver *Pagination[T]) AddDataItem(item T) {
	receiver.Data = append([]T(receiver.Data), item)
}

func (receiver *Pagination[T]) DataSize() uint64 {
	return uint64(len([]T(receiver.Data)))
}

func NewPagination[T any](data []T, total uint64, pageCdt base.IPage) *Pagination[T] {
	pagination := new(Pagination[T])
	pagination.Data = data
	pagination.Total = total
	pagination.PageCdt = pageCdt
	return pagination
}
