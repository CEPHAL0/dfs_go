package userwalletModels

import (
	authModels "backend/models/auth"
	walletModels "backend/models/wallet"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserWallet struct {
	gorm.Model
	ID           uuid.UUID               `gorm:"type:uuid;default:gen_random_uuid();primaryKey;not null" json:"id"`
	UserID       uuid.UUID               `gorm:"not null" json:"userId"`
	User         authModels.User         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	WalletTypeID uuid.UUID               `gorm:"not null" json:"walletTypeId"`
	WalletType   walletModels.WalletType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	DfsID        uint                    `gorm:"not null;unique" json:"dfsId"`
	CreatedAt    time.Time               `gorm:"autoCreateTime;not null" json:"-"`
	UpdatedAt    time.Time               `gorm:"autoUpdateTime;not null" json:"-"`
	DeletedAt    time.Time               `json:"-"`
}
