package usecase

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/bagusyanuar/go-simrs/internal/shared/config"
	"github.com/bagusyanuar/go-simrs/internal/sso/domain"
	userDomain "github.com/bagusyanuar/go-simrs/internal/user/domain"
	"github.com/bagusyanuar/go-simrs/pkg/jwt"
	"github.com/bagusyanuar/go-simrs/pkg/password"
	"github.com/google/uuid"
)

type ssoUC struct {
	ssoRepo  domain.SSORepository
	userRepo userDomain.UserRepository
	conf     *config.Config
}

func NewSSOUsecase(ssoRepo domain.SSORepository, userRepo userDomain.UserRepository, conf *config.Config) domain.SSOUsecase {
	return &ssoUC{
		ssoRepo:  ssoRepo,
		userRepo: userRepo,
		conf:     conf,
	}
}

func (u *ssoUC) Authorize(ctx context.Context, req domain.AuthorizeRequest) (string, string, error) {
	// 1. Verify User Credentials
	user, err := u.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return "", "", errors.New("invalid email or password")
	}

	if err := password.VerifyPassword(req.Password, user.Password); err != nil {
		return "", "", errors.New("invalid email or password")
	}

	// 2. Validate Client
	_, err = u.ssoRepo.GetClientByID(ctx, req.ClientID)
	if err != nil {
		return "", "", errors.New("invalid client_id")
	}

	// 3. Create SSO Session (Long-lived for True SSO)
	sessionID := uuid.New()
	ssoSession := &domain.UserSession{
		ID:        sessionID,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(24 * time.Hour), // SSO Session lasts 24 hours
	}
	if err := u.ssoRepo.SaveSession(ctx, ssoSession); err != nil {
		return "", "", err
	}

	// 4. Generate Random Auth Code (Short-lived)
	code := uuid.New().String()

	// 5. Save Auth Code Session
	codeSession := &domain.AuthCodeSession{
		Code:          code,
		ClientID:      req.ClientID,
		UserID:        user.ID,
		CodeChallenge: req.CodeChallenge,
		RedirectURI:   req.RedirectURI,
		ExpiresAt:     time.Now().Add(5 * time.Minute),
	}

	if err := u.ssoRepo.SaveAuthCode(ctx, codeSession); err != nil {
		return "", "", err
	}

	return code, sessionID.String(), nil
}

func (u *ssoUC) AuthorizeSilent(ctx context.Context, sessionID string, req domain.AuthorizeSilentRequest) (string, error) {
	// 1. Verify Session
	ssoSession, err := u.ssoRepo.GetSession(ctx, sessionID)
	if err != nil {
		return "", errors.New("session not found")
	}

	if time.Now().After(ssoSession.ExpiresAt) {
		u.ssoRepo.DeleteSession(ctx, sessionID)
		return "", errors.New("session expired")
	}

	// 2. Validate Client
	_, err = u.ssoRepo.GetClientByID(ctx, req.ClientID)
	if err != nil {
		return "", errors.New("invalid client_id")
	}

	// 3. Generate Auth Code
	code := uuid.New().String()

	// 4. Save Auth Code Session
	codeSession := &domain.AuthCodeSession{
		Code:          code,
		ClientID:      req.ClientID,
		UserID:        ssoSession.UserID,
		CodeChallenge: req.CodeChallenge,
		RedirectURI:   req.RedirectURI,
		ExpiresAt:     time.Now().Add(5 * time.Minute),
	}

	if err := u.ssoRepo.SaveAuthCode(ctx, codeSession); err != nil {
		return "", err
	}

	return code, nil
}

func (u *ssoUC) ExchangeToken(ctx context.Context, req domain.TokenRequest) (*domain.TokenResponse, error) {
	// 1. Get & Validate Code
	session, err := u.ssoRepo.GetAuthCode(ctx, req.Code)
	if err != nil {
		return nil, errors.New("invalid or expired code")
	}

	if time.Now().After(session.ExpiresAt) {
		u.ssoRepo.DeleteAuthCode(ctx, req.Code)
		return nil, errors.New("code expired")
	}

	if session.ClientID != req.ClientID {
		return nil, errors.New("client_id mismatch")
	}

	// 2. Verify PKCE (S256 only for security)
	h := sha256.New()
	h.Write([]byte(req.CodeVerifier))
	computedChallenge := base64.RawURLEncoding.EncodeToString(h.Sum(nil))

	if computedChallenge != session.CodeChallenge {
		return nil, errors.New("invalid code_verifier")
	}

	// 3. Get User
	user, err := u.userRepo.FindByID(ctx, session.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// 4. Generate Tokens
	atExpiration := time.Duration(u.conf.JWTExpiration) * time.Minute
	accessToken, err := jwt.GenerateToken(
		user.ID.String(),
		user.Email,
		[]string{},
		u.conf.JWTSecret,
		u.conf.JWTIssuer,
		atExpiration,
	)
	if err != nil {
		return nil, fmt.Errorf("failed generate AT: %w", err)
	}

	rtExpiration := time.Duration(u.conf.JWTRefreshExpiration) * time.Hour
	refreshToken, err := jwt.GenerateToken(
		user.ID.String(),
		user.Email,
		[]string{},
		u.conf.JWTRefreshSecret,
		u.conf.JWTIssuer,
		rtExpiration,
	)
	if err != nil {
		return nil, fmt.Errorf("failed generate RT: %w", err)
	}

	// 5. Cleanup
	u.ssoRepo.DeleteAuthCode(ctx, req.Code)

	return &domain.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int(atExpiration.Seconds()),
	}, nil
}
