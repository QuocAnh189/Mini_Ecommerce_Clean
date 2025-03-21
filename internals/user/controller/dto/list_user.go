package dto

import (
	"ecommerce_clean/internals/user/entity"
	"ecommerce_clean/pkgs/paging"
)

type ListUserRequest struct {
	Search    string `json:"search,omitempty" form:"search"`
	Page      int64  `json:"-" form:"page"`
	Limit     int64  `json:"-" form:"size"`
	OrderBy   string `json:"-" form:"order_by"`
	OrderDesc bool   `json:"-" form:"order_desc"`
	TakeAll   bool   `json:"-" form:"take_all"`
}
type ListUserResponse struct {
	Users      []*entity.User     `json:"items"`
	Pagination *paging.Pagination `json:"metadata"`
}
