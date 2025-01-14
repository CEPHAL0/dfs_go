package itemmodels

import (
	"time"

	uuid "github.com/google/uuid"
)

type Item struct {
	ID            uuid.UUID `gorm:"primaryKey;not null;default:gen_random_uuid()" json:"id"`
	Name          string    `gorm:"not null" json:"name"`
	Image         string    `gorm:"not null" json:"image"`
	InitialQty    uint      `gorm:"not null;" json:"initialQuantity"`
	Price         float64   `gorm:"not null;type:numeric(10,2)"`
	DiscountPrice float64   `gorm:"not null;type:numeric(10,2)"`
	CreatedAt     time.Time `gorm:"autoCreateTime;not null" json:"-"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime;not null" json:"-"`
	DeletedAt     time.Time `json:"-"`
}
