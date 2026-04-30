package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type OAuthClient struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	ClientID     string    `gorm:"type:varchar(100);unique;not null"`
	ClientSecret string    `gorm:"type:varchar(255)"`
	Name         string    `gorm:"type:varchar(100)"`
	RedirectURIs string    `gorm:"type:text"` // Comma separated
	CreatedAt    time.Time
}

func (OAuthClient) TableName() string {
	return "oauth_clients"
}

type AuthCodeSession struct {
	Code          string    `gorm:"primaryKey"`
	ClientID      string    `gorm:"type:varchar(100)"`
	UserID        uuid.UUID `gorm:"type:uuid"`
	CodeChallenge string    `gorm:"type:text"`
	RedirectURI   string    `gorm:"type:text"`
	ExpiresAt     time.Time
}

func (AuthCodeSession) TableName() string {
	return "auth_code_sessions"
}

type UserSession struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	ExpiresAt time.Time
	CreatedAt time.Time
}

func (UserSession) TableName() string {
	return "sso_sessions"
}

type SSOUsecase interface {
	Authorize(ctx context.Context, req AuthorizeRequest) (string, string, error) // Returns code, sessionID
	AuthorizeSilent(ctx context.Context, sessionID string, req AuthorizeSilentRequest) (string, error)
	ExchangeToken(ctx context.Context, req TokenRequest) (*TokenResponse, error)
	RefreshToken(ctx context.Context, req TokenRequest) (*TokenResponse, error)
}

type AuthorizeRequest struct {
	Email         string
	Password      string
	ClientID      string
	CodeChallenge string
	RedirectURI   string
	State         string
}

type AuthorizeSilentRequest struct {
	ClientID      string
	CodeChallenge string
	RedirectURI   string
	State         string
	ResponseType  string
}

type TokenRequest struct {
	GrantType    string
	ClientID     string
	Code         string
	CodeVerifier string
	RedirectURI  string
	RefreshToken string
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"` // Omit if we only use cookie
	ExpiresIn    int    `json:"expires_in"`
}

type SSORepository interface {
	GetClientByID(ctx context.Context, clientID string) (*OAuthClient, error)
	SaveAuthCode(ctx context.Context, session *AuthCodeSession) error
	GetAuthCode(ctx context.Context, code string) (*AuthCodeSession, error)
	DeleteAuthCode(ctx context.Context, code string) error

	SaveSession(ctx context.Context, session *UserSession) error
	GetSession(ctx context.Context, sessionID string) (*UserSession, error)
	DeleteSession(ctx context.Context, sessionID string) error
}
