package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name     string    `gorm:"type:varchar(100);not null"`
	Email    string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string    `gorm:"not null"`
	Role     string    `gorm:"type:varchar(20);default:'user';"`
	Photo    string    `gorm:"default:'default.png';"`

	Verified  bool      `gorm:"default:false;"`
	Provider  string    `gorm:"default:'local';"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (user *User) BeforeCreate(db *gorm.DB) error {
	user.ID = uuid.NewV4()

	return nil
}

type RegisterUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Role      string `json:"role,omitempty"`
	Provider  string `json:"provider,omitempty"`
	Photo     string `json:"photo,omitempty"`
	Verified  bool   `json:"verified,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func FilteredResponse(user *User) UserResponse {
	return UserResponse{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		Provider:  user.Provider,
		Photo:     user.Photo,
		Verified:  user.Verified,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
