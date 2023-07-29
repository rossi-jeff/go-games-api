package payloads

type CredentialsPayload struct {
	UserName string
	Password string
}

type LoginResponse struct {
	UserName string
	Token    string
}
