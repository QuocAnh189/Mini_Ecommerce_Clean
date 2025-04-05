package entity

import (
	productEntity "ecommerce_clean/internals/product/entity"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartLine struct {
	ID        string `json:"id" gorm:"unique;not null;index;primary_key"`
	CartID    string `json:"cart_id"`
	ProductID string `json:"product_id"`
	Product   *productEntity.Product
	Quantity  uint            `json:"quantity"`
	Price     float64         `json:"price"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (cartLine *CartLine) BeforeCreate(tx *gorm.DB) error {
	cartLine.ID = uuid.New().String()

	return nil
}

func (cartLine *CartLine) TableName() string {
	return "cart_lines"
}
