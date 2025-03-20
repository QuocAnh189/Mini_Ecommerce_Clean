package paging

import "math"

const (
	DefaultPageSize int64 = 20
)

type Pagination struct {
	Page        int64 `json:"page"`
	Size        int64 `json:"size"`
	TakeAll     bool  `json:"take_all"`
	Skip        int64 `json:"skip"`
	TotalCount  int64 `json:"total_count"`
	TotalPages  int64 `json:"total_pages"`
	HasPrevious bool  `json:"has_previous"`
	HasNext     bool  `json:"has_next"`
}

func NewPagination(page int64, size int64, total int64) *Pagination {
	var pageInfo Pagination
	limit := DefaultPageSize
	if size > 0 && size <= 1000 {
		pageInfo.Size = size
	} else {
		pageInfo.Size = limit
	}

	totalPage := int64(math.Ceil(float64(total) / float64(pageInfo.Size)))
	pageInfo.TotalCount = total
	pageInfo.TotalPages = totalPage
	if page < 1 || totalPage == 0 {
		page = 1
	}

	pageInfo.Page = page
	pageInfo.Skip = (page - 1) * pageInfo.Size

	if page == 1 {
		pageInfo.HasPrevious = false
		pageInfo.HasNext = true
	}

	if page > 1 && page < totalPage {
		pageInfo.HasPrevious = true
		pageInfo.HasNext = true
	}

	if page == totalPage {
		pageInfo.HasPrevious = true
		pageInfo.HasNext = false
	}

	if totalPage == 1 {
		pageInfo.HasPrevious = false
		pageInfo.HasNext = false
	}

	return &pageInfo
}
