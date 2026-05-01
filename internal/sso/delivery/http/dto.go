package http

type AuthorizeRequest struct {
	Email         string `json:"email" validate:"required,email"`
	Password      string `json:"password" validate:"required"`
	ClientID      string `json:"client_id" validate:"required"`
	CodeChallenge string `json:"code_challenge" validate:"required"`
	RedirectURI   string `json:"redirect_uri" validate:"required"`
	State         string `json:"state"`
}

type AuthorizeSilentRequest struct {
	ClientID      string `query:"client_id" validate:"required"`
	CodeChallenge string `query:"code_challenge" validate:"required"`
	RedirectURI   string `query:"redirect_uri" validate:"required"`
	State         string `query:"state"`
	ResponseType  string `query:"response_type" validate:"required"`
	Prompt        string `query:"prompt"`
}

type TokenRequest struct {
	GrantType    string `json:"grant_type" validate:"required"`
	ClientID     string `json:"client_id" validate:"required"`
	Code         string `json:"code" validate:"required_if=GrantType authorization_code"`
	CodeVerifier string `json:"code_verifier" validate:"required_if=GrantType authorization_code"`
	RedirectURI  string `json:"redirect_uri" validate:"required_if=GrantType authorization_code"`
}
