package rest

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RegistrationPayload struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

type ForgetPasswordPayload struct {
	Email string `json:"email"`
}

type ResetPasswordPayload struct {
	Email    string `json:"email"`
	Token    string `json:"token"`
	Password string `json:"password"`
}

type ChangePasswordPayload struct {
	OldPassword string `json:"old"`
	NewPassword string `json:"new"`
}
