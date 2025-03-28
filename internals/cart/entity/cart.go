package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	ID        string      `json:"id" gorm:"unique;not null;index;primary_key"`
	UserID    string      `json:"user_id" gorm:"unique;not null;index"`
	Lines     []*CartLine `json:"lines"`
	User      *User
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (cart *Cart) BeforeCreate(tx *gorm.DB) error {
	cart.ID = uuid.New().String()

	return nil
}

func (cart *Cart) TableName() string {
	return "carts"
}
