package orderModels

import (
	"backend/enums"
	menuModels "backend/models/menu"
	userwalletModels "backend/models/userWallet"
	userWalletTransactionModels "backend/models/userWalletTransaction"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID uuid.UUID `gorm:"primaryKey;default:gen_random_uuid();type:uuid;not null" json:"id"`

	MenuID uuid.UUID       `gorm:"not null" json:"menuId"`
	Menu   menuModels.Menu `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	UserWalletID uuid.UUID                   `gorm:"not null" json:"userWalletId"`
	UserWallet   userwalletModels.UserWallet `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;not null"`

	ItemName   string            `gorm:"not null" json:"itemName"`
	Quantity   uint              `gorm:"not null" json:"quantity"`
	TotalPrice float64           `gorm:"not null;type:numeric(10,2)" json:"totalPrice"`
	Status     enums.OrderStatus `gorm:"not null" json:"status"`

	UserWalletTransactionId uuid.UUID                                         `json:"userWalletTransactionId"`
	UserWalletTransaction   userWalletTransactionModels.UserWalletTransaction `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`

	CreatedAt time.Time `gorm:"autoCreateTime;not null" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null" json:"-"`
	DeletedAt time.Time `json:"-"`
}
