package itemmodels

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID `gorm:"primaryKey;default:gen_random_uuid();type:uuid;not null" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	IsSpecial bool      `gorm:"default:false;not null" json:"is_special"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null" json:"-"`
	DeletedAt time.Time `json:"-"`
}
