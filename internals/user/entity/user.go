package entity

import (
	cartEntity "ecommerce_clean/internals/cart/entity"
	"ecommerce_clean/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string          `json:"id" gorm:"unique;not null;index;primary_key"`
	Email     string          `json:"email" gorm:"uniqueIndex:unique_user_email;not null"`
	Name      string          `json:"name" gorm:"uniqueIndex:unique_user_name;not null"`
	AvatarUrl string          `json:"avatar_url" gorm:"unique:unique_user_avatar;not null"`
	Password  string          `json:"password" gorm:"not null;"`
	Role      string          `json:"role" gorm:"default:'customer';not null"`
	CreatedAt time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New().String()
	user.Password = utils.HashAndSalt([]byte(user.Password))

	return nil
}

func (user *User) AfterCreate(tx *gorm.DB) error {
	cart := cartEntity.Cart{
		ID:     uuid.New().String(),
		UserID: user.ID,
	}

	if err := tx.Create(&cart).Error; err != nil {
		return err
	}

	return nil
}

func (user *User) TableName() string {
	return "users"
}
