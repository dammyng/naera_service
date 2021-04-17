package rest

var (
	InternalServerError         = "Internal server error. Please try later."
	InvalidRequest              = "Sorry this is an invalid request"
	ProcessingRequestError      = "Oops! an error occurred processing your request. Retry."
	UserCreationError           = "Account creation unsuccessful. Please try later."
	UserCreationSuccessful      = "Congratulations! You’re good to go for the Flairs experience"
	InvalidCredentialError      = "Oops! You entered wrong email or password. Please check and retry."
	UnverifiedEmailError        = "Sorry! Looks like you haven’t verified your email. Please check your mailbox to continue…"
	PasswordResetReqSuccessful  = "Good to go! Password  reset successful."
	InvalidPassword             = "Sorry! Password cannot be less than 8 characters."
	PassswordResetSuccessful    = "Great! Password  reset successful."
	UpdateRecordError           = "Oops! An error occurred updating your record. Please retry."
	WalletNotFound              = "Sorry! Wallet not found. Please retry."
	RecordUpdateSuccessful      = "Great! Your record update is successful."
	EmailVerificationSuccessful = "Great! Email verification successful"
	PhoneVerificationSuccessful = "Great! Phone verification successful"
	UserNotFound                = "Oops! User doesn’t exist. Please retry."
	DuplicateUserAccount        = "Oops! A Naera Pay user with this email already exist."
)
