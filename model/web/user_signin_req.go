package web

type UserSignRequest struct {
	Email    string `valdiate:"required" json:"email"`
	Password string `valdiate:"required" json:"password"`
}
