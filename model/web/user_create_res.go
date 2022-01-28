package web

type UserCreateResponse struct {
	Email    string `valdiate:"required" json:"email"`
	Token    string `valdiate:"required" json:"token"`
	Username string `valdiate:"required" json:"username"`
}
