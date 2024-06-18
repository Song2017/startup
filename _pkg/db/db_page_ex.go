package pkg

import "fmt"

type Page struct {
	PageNum  int64
	PageSize int64
}

func (p Page) GetSqlText() string {
	return fmt.Sprintf(" OFFSET %d LIMIT %d ", (p.PageNum-1)*p.PageSize, p.PageSize)
}
