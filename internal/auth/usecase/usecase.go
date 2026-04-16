package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/bagusyanuar/go-simrs/internal/shared/config"
	"github.com/bagusyanuar/go-simrs/internal/user/domain"
	"github.com/bagusyanuar/go-simrs/pkg/jwt"
	"github.com/bagusyanuar/go-simrs/pkg/password"
)

type authUC struct {
	userRepo domain.UserRepository
	conf     *config.Config
}

func NewAuthUsecase(repo domain.UserRepository, conf *config.Config) domain.AuthUsecase {
	return &authUC{
		userRepo: repo,
		conf:     conf,
	}
}

func (u *authUC) Login(ctx context.Context, email, pwd string) (string, string, error) {
	// 1. Find User
	user, err := u.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", "", fmt.Errorf("auth_uc.Login (user not found): %w", err)
	}

	// 2. Verify Password
	if err := password.VerifyPassword(pwd, user.Password); err != nil {
		return "", "", fmt.Errorf("auth_uc.Login (invalid password): %w", err)
	}

	// 3. Generate Access Token
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
		return "", "", fmt.Errorf("auth_uc.Login (failed generate AT): %w", err)
	}

	// 4. Generate Refresh Token
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
		return "", "", fmt.Errorf("auth_uc.Login (failed generate RT): %w", err)
	}

	return accessToken, refreshToken, nil
}

func (u *authUC) Refresh(ctx context.Context, rt string) (string, error) {
	// 1. Parse RT
	claims, err := jwt.ParseToken(rt, u.conf.JWTRefreshSecret)
	if err != nil {
		return "", fmt.Errorf("auth_uc.Refresh (invalid RT): %w", err)
	}

	// 2. Issuing new AT
	atExpiration := time.Duration(u.conf.JWTExpiration) * time.Minute
	accessToken, err := jwt.GenerateToken(
		claims.Subject,
		claims.Email,
		[]string{},
		u.conf.JWTSecret,
		u.conf.JWTIssuer,
		atExpiration,
	)
	if err != nil {
		return "", fmt.Errorf("auth_uc.Refresh (failed generate AT): %w", err)
	}

	return accessToken, nil
}

func (u *authUC) GetRefreshExpiration() int {
	return u.conf.JWTRefreshExpiration
}
