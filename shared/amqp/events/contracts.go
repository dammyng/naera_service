package events

type UserCreatedEvent struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
	Host  string `json:"host"`
}

type WelcomeUserEvent struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Host     string `json:"host"`
	Username string `json:"username"`
}

type ResendEmailEvent struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
	Host  string `json:"host"`
}

func (e *UserCreatedEvent) EventName() string {
	return "user.created"
}

func (e *ResendEmailEvent) EventName() string {
	return "user.resend_email"
}

func (e *WelcomeUserEvent) EventName() string {
	return "user.welcome"
}