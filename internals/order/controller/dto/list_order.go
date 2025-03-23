package dto

import (
	"ecommerce_clean/pkgs/paging"
)

type ListOrdersRequest struct {
	UserID    string `json:"-"`
	Code      string `json:"code,omitempty" form:"code"`
	Status    string `json:"status,omitempty" form:"status"`
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"limit"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
}

type ListOrdersResponse struct {
	Orders     []*Order           `json:"orders,omitempty"`
	Pagination *paging.Pagination `json:"pagination,omitempty"`
}
