package schemas

type LoginSchema struct {
	UserName string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterSchema struct {
	Username        string `json:"username" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
}
