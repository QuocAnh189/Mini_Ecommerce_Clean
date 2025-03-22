package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"ecommerce_clean/utils"
)

type Product struct {
	ID          string         `json:"id" gorm:"unique;not null;index;primary_key"`
	Code        string         `json:"code" gorm:"uniqueIndex:unique_product_code,not null"`
	Name        string         `json:"name" gorm:"uniqueIndex:unique_product_name,not null"`
	ImageUrl    string         `json:"image_url" gorm:"unique:unique_product_image,not null"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	Active      bool           `json:"active" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (m *Product) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.New().String()
	m.Code = utils.GenerateCode("P")
	m.Active = true
	return nil
}

func (m *Product) TableName() string {
	return "products"
}
