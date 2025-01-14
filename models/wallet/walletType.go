package walletmodels

import (
	"backend/enums"
	"time"

	uuid "github.com/google/uuid"
)

type WalletType struct {
	ID              uuid.UUID             `gorm:"default:gen_random_uuid();not null;primaryKey;type:uuid" json:"id"`
	Organization    enums.Organization    `gorm:"not null" json:"organization"`
	MemberType      enums.MemberType      `gorm:"not null" json:"memberType"`
	TransactionMode enums.TransactionMode `gorm:"not null" json:"transactionMode"`
	IsDiscounted    bool                  `gorm:"not null" json:"is_discounted"`
	CreatedAt       time.Time             `gorm:"autoCreateTime;not null" json:"-"`
	UpdatedAt       time.Time             `gorm:"autoUpdateTime;not null" json:"-"`
	DeletedAt       time.Time             `json:"-"`
}
