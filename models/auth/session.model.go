package authModels

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	SessionID string    `gorm:"not null" json:"sessionID"`
	UserID    uuid.UUID `gorm:"not null;type:uuid" json:"userID"`
	Expires   time.Time `gorm:"not null" json:"expires"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null" json:"-"`
}
