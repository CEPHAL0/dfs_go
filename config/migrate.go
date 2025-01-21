package config

import (
	authModels "backend/models/auth"
	itemModels "backend/models/item"
	menuModels "backend/models/menu"
	orderModels "backend/models/order"
	userwalletModels "backend/models/userWallet"
	userWalletTransactionModels "backend/models/userWalletTransaction"
	walletModels "backend/models/wallet"

	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) {
	DB.AutoMigrate(&authModels.User{})
	DB.AutoMigrate(&itemModels.Item{})
	DB.AutoMigrate(&itemModels.Category{})
	DB.AutoMigrate(&walletModels.WalletType{})
	DB.AutoMigrate(&menuModels.Menu{})
	DB.AutoMigrate(&userwalletModels.UserWallet{})
	DB.AutoMigrate(&userWalletTransactionModels.UserWalletTransaction{})
	DB.AutoMigrate(&orderModels.Order{})
	DB.AutoMigrate(&authModels.Session{})
}
