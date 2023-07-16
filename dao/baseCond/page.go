// Code generated with the command 'go generate', DO NOT EDIT.

package baseCond

type Page struct {
	noLimit bool

	page     int
	pageSize int
}

// PageCdtLimited
// page从 0 开始
func PageCdtLimited(page int, pageSize int) *Page {
	return &Page{
		noLimit:  false,
		page:     page,
		pageSize: pageSize,
	}
}

func PageCdtNoLimit() *Page {
	return &Page{
		noLimit: true,
	}
}

func (receiver Page) HasLimit() bool {
	return !receiver.noLimit
}

func (receiver Page) PageSize() int {
	return receiver.pageSize
}

func (receiver Page) Offset() int {
	return receiver.page * receiver.pageSize
}
