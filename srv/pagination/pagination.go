package pagination

import (
	"math"
	"net/url"
	"strings"
)

const (
	DEFAULT_PAGE      = 1
	DEFAULT_PAGE_SIZE = 5
	DEFAULT_SORT      = "id"
)

type Pagination struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafeList []string
}

func NewPagination(qs url.Values, v *Validator, sortSafeList ...string) *Pagination {
	p := &Pagination{}

	p.Page = ParseQueryInt(qs, "page", DEFAULT_PAGE, v)
	p.PageSize = ParseQueryInt(qs, "page_size", DEFAULT_PAGE_SIZE, v)
	p.Sort = ParseQueryString(qs, "sort", DEFAULT_SORT)
	p.SortSafeList = sortSafeList

	p.Validate(v)

	return p
}

func (p *Pagination) Validate(v *Validator) {
	v.Between("page", p.Page, 1, 10_000_000)
	v.Between("pageSize", p.PageSize, 1, 100)
	v.Check(validator.In(p.Sort, p.SortSafeList...), "sort", "sort parameter is not permitted")
}

func (p *Pagination) GetSortColumn() string {
	for _, v := range p.SortSafeList {
		if v == p.Sort {
			if strings.Contains(p.Sort, "_asc") {
				return strings.TrimSuffix(p.Sort, "_asc")
			} else {
				return strings.TrimSuffix(p.Sort, "_desc")
			}
		}
	}
	panic("unsafe sort query parameter " + p.Sort)
}

func (p *Pagination) sortDirection() string {
	if strings.HasSuffix(p.Sort, "_asc") {
		return "ASC"
	}
	return "DESC"
}

func (p *Pagination) limit() int {
	return p.PageSize
}

func (p *Pagination) offset() int {
	return (p.Page - 1) * p.PageSize
}

type PaginationMetadata struct {
	CurrentPage  int `json:"current_page,omitempty"`
	PageSize     int `json:"page_size,omitempty"`
	FirstPage    int `json:"first_page,omitempty"`
	LastPage     int `json:"last_page,omitempty"`
	TotalRecords int `json:"total_records,omitempty"`
}

func NewPaginationMetadata(totalRecords, page, pageSize int) PaginationMetadata {
	if totalRecords == 0 {
		return PaginationMetadata{}
	}

	return PaginationMetadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     int(math.Ceil(float64(totalRecords) / float64(pageSize))),
		TotalRecords: totalRecords,
	}
}
