package menumodels

import (
	itemmodels "backend/models/item"
	"time"

	"github.com/google/uuid"
)

type Tabler interface {
	TableName() string
}

type Menu struct {
	ID uuid.UUID `gorm:"primaryKey;default:gen_random_uuid();type:uuid;not null" json:"id"`

	ItemID uuid.UUID       `json:"itemId"`
	Item   itemmodels.Item `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	CategoryID uuid.UUID           `json:"categoryID"`
	Category   itemmodels.Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	InitialQty   uint `gorm:"not null" json:"initialQty"`
	AvailableQty uint `gorm:"not null" json:"availableQty"`

	StartTime time.Time `gorm:"not null" json:"startTime"`
	EndTime   time.Time `gorm:"not null" json:"endTime"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`
	DeletedAt time.Time `json:"-"`
}

func (Menu) TableName() string {
	return "menu"
}
