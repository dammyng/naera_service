package events

type UserCreatedEvent struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type WelcomeUserEvent struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type ResendEmailEvent struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
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