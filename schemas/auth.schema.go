package schemas


type LoginSchema struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterSchema struct {
	Username        string
	Email           string
	Password        string
	ConfirmPassword string
}
