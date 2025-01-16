package repositories

import (
	authModels "backend/models/auth"
)

type UserRepository interface {
	Insert(username string, password string, email string) authModels.User
}
