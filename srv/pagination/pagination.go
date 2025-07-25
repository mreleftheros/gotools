package pagination

import (
	"math"
	"net/url"
	"strings"

	"github.com/mreleftheros/gotools/srv/request"
	"github.com/mreleftheros/gotools/srv/validator"
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

func NewPagination(v *validator.Validator, qs url.Values, sortSafeList ...string) *Pagination {
	p := &Pagination{
		Page:         request.ParseQueryInt(qs, "page", DEFAULT_PAGE),
		PageSize:     request.ParseQueryInt(qs, "page_size", DEFAULT_PAGE_SIZE),
		Sort:         request.ParseQueryString(qs, "sort", DEFAULT_SORT),
		SortSafeList: sortSafeList,
	}

	v.Between(p.Page, 1, 10_000_000, "page query")
	v.Between(p.PageSize, 1, 100, "page size query")
	v.Check(validator.In(p.Sort, p.SortSafeList...), "sort query", "sort query is not permitted")

	return p
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

func (p *Pagination) SortDirection() string {
	if strings.HasSuffix(p.Sort, "_asc") {
		return "ASC"
	}
	return "DESC"
}

func (p *Pagination) Limit() int {
	return p.PageSize
}

func (p *Pagination) Offset() int {
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
