package authModels

import (
	"backend/enums"
	"time"

	uuid "github.com/google/uuid"
)

type User struct {
	ID                uuid.UUID  `gorm:"primaryKey;default:gen_random_uuid();type:uuid;not null" json:"id"`
	Name              string     `gorm:"not null" json:"name"`
	Email             string     `gorm:"not null;unique" json:"email"`
	Password          string     `gorm:"not null" json:"-"`
	Role              enums.Role `gorm:"not null;default:customer" json:"role"`
	IsPasswordExpired bool       `gorm:"default:false;not null" json:"is_password_expired"`
	CreatedAt         time.Time  `gorm:"autoCreateTime;not null" json:"-"`
	UpdatedAt         time.Time  `gorm:"autoUpdateTime;not null" json:"-"`
	DeletedAt         time.Time  `gorm:"default:null" json:"-"`
}
