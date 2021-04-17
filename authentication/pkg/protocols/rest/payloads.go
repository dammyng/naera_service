package rest

import (
	"github.com/asaskevich/govalidator"
)

func init() {
  govalidator.SetFieldsRequiredByDefault(true)
}


type LoginPayload struct {
	Email    string `json:"email" valid:"email"`
	Password string `json:"password" valid:"type(string)~Your password is required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RegistrationPayload struct {
	Email     string `json:"email" valid:"email"`
	Password  string `json:"password" valid:"length(7|255)~Your password should not be less than seven characters"`
	FirstName string `json:"first_name" valid:"type(string)~Your First name is required"`
	LastName  string `json:"last_name" valid:"type(string)~Your Last name is required"`
	Phone     string `json:"phone" valid:"type(string)~Your phone number is required"`
}

type ForgetPasswordPayload struct {
	Email string `json:"email" valid:"email"`
}

type ResetPasswordPayload struct {
	Email    string `json:"email" valid:"email"`
	Token    string `json:"token"`
	Password string `json:"password"`
}

type ChangePasswordPayload struct {
	OldPassword string `json:"old"`
	NewPassword string `json:"new"`
}


