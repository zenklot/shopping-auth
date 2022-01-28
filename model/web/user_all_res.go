package web

type UserAllResponse struct {
	Email    string `valdiate:"required" json:"email"`
	Username string `valdiate:"required" json:"username"`
}
