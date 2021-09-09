package models

import "github.com/asaskevich/govalidator"

func init() {
  govalidator.SetFieldsRequiredByDefault(true)
}

type LoginRequestPayload struct {
	Email    string `json:"email" valid:"email"`
	Password string `json:"password" valid:"type(string)"`
}

type RegisterUserRequestPayload struct {
	Email     string `json:"email" valid:"email"`
	Password  string `json:"password" valid:"type(string)"`
	Firstname string `json:"firstname" valid:"type(string)"`
	Lastname  string `json:"lastname" valid:"type(string)"`
}
