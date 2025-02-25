package dto

type RegisterRequest struct {
  Name                  string `json:"name"`
  Email                 string `json:"email"`
  Password              string `json:"password"`
  PasswordConfirmation  string `json:"passsword_confirm"`
  Gender                string `json:"gender"`
}
