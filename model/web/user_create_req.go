package web

type UserCreateDetail struct {
	Username          string `valdiate:"required" json:"username"`
	Email             string `valdiate:"required" json:"email"`
	EncryptedPassword string `valdiate:"required" json:"encrypted_password"`
	Phone             string `valdiate:"required" json:"phone"`
	Address           string `valdiate:"required" json:"address"`
	City              string `valdiate:"required" json:"city"`
	Country           string `valdiate:"required" json:"country"`
	Name              string `valdiate:"required" json:"name"`
	Postcode          string `valdiate:"required" json:"postcode"`
}

type UserCreateRequest struct {
	User UserCreateDetail `valdiate:"required" json:"user"`
}
