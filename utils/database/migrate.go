package database

import (
	authmodels "backend/models/auth"
	itemmodels "backend/models/item"
	menumodels "backend/models/menu"
	ordermodels "backend/models/order"
	userWalletmodels "backend/models/userWallet"
	userwallettransactionmodels "backend/models/userWalletTransaction"
	walletmodels "backend/models/wallet"

	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(&authmodels.User{})
	DB.AutoMigrate(&itemmodels.Item{})
	DB.AutoMigrate(&itemmodels.Category{})
	DB.AutoMigrate(&walletmodels.WalletType{})
	DB.AutoMigrate(&menumodels.Menu{})
	DB.AutoMigrate(&userWalletmodels.UserWallet{})
	DB.AutoMigrate(&userwallettransactionmodels.UserWalletTransaction{})
	DB.AutoMigrate(&ordermodels.Order{})
}
