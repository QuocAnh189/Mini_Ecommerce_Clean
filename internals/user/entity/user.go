package entity

import (
	"ecommerce_clean/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        string         `json:"id" gorm:"unique;not null;index;primary_key"`
	Email     string         `json:"email" gorm:"unique;not null;index;primary_key"`
	Name      string         `json:"name" gorm:"unique;not null;index;primary_key"`
	AvatarUrl string         `json:"avatar_url"`
	Password  string         `json:"password" gorm:"not null;"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New().String()
	user.Password = utils.HashAndSalt([]byte(user.Password))

	return nil
}

func (user *User) TableName() string {
	return "users"
}
