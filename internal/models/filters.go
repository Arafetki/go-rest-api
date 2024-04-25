package models

type Filters struct {
	Page     int `validate:"gt=0,lte=10_000_000"`
	PageSize int `validate:"gt=0,lte=100"`
}

func (f Filters) Limit() int {
	return f.PageSize
}

func (f Filters) Offset() int {
	return (f.Page - 1) * f.PageSize
}
