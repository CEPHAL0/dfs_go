package userwallettransactionmodels

import (
	"backend/enums"
	userWalletmodels "backend/models/userWallet"
	"time"

	"github.com/google/uuid"
)

type UserWalletTransaction struct {
	ID           uuid.UUID                   `gorm:"primaryKey;not null;default:gen_random_uuid();type:uuid" json:"id"`
	UserWalletID uuid.UUID                   `gorm:"not null" json:"userWalletId"`
	UserWallet   userWalletmodels.UserWallet `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	Amount      float64               `gorm:"not null;type:numeric(10,2)"`
	Date        time.Time             `gorm:"not null" json:"date"`
	Type        enums.TransactionMode `gorm:"not null" json:"type"`
	PaymentMode string                `gorm:"not null" json:"paymentMode"`
	Remarks     string                `gorm:"type:text" json:"remarks"`

	CreatedAt time.Time `gorm:"autoCreateTime;not null" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null" json:"-"`
	DeletedAt time.Time `json:"-"`
}
