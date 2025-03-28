package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	userEntity "ecommerce_clean/internals/user/entity"
	"ecommerce_clean/utils"
)

type Order struct {
	ID         string `json:"id" gorm:"unique;not null;index;primary_key"`
	Code       string `json:"code"`
	UserID     string `json:"user_id"`
	User       *userEntity.User
	Lines      []*OrderLine      `json:"lines"`
	TotalPrice float64           `json:"total_price"`
	Status     utils.OrderStatus `json:"status"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
	DeletedAt  *gorm.DeletedAt   `json:"deleted_at" gorm:"index"`
}

func (order *Order) BeforeCreate(tx *gorm.DB) error {
	order.ID = uuid.New().String()
	order.Code = utils.GenerateCode("SO")

	if order.Status == "" {
		order.Status = utils.OrderStatusNew
	}

	return nil
}

func (order *Order) TableName() string {
	return "orders"
}
