package repository

import (
	"context"

	"github.com/bagusyanuar/go-simrs/internal/sso/domain"
	"gorm.io/gorm"
)

type ssoRepo struct {
	db *gorm.DB
}

func NewSSORepository(db *gorm.DB) domain.SSORepository {
	return &ssoRepo{db: db}
}

func (r *ssoRepo) GetClientByID(ctx context.Context, clientID string) (*domain.OAuthClient, error) {
	var client domain.OAuthClient
	err := r.db.WithContext(ctx).Where("client_id = ?", clientID).First(&client).Error
	return &client, err
}

func (r *ssoRepo) SaveAuthCode(ctx context.Context, session *domain.AuthCodeSession) error {
	return r.db.WithContext(ctx).Create(session).Error
}

func (r *ssoRepo) GetAuthCode(ctx context.Context, code string) (*domain.AuthCodeSession, error) {
	var session domain.AuthCodeSession
	err := r.db.WithContext(ctx).Where("code = ?", code).First(&session).Error
	return &session, err
}

func (r *ssoRepo) DeleteAuthCode(ctx context.Context, code string) error {
	return r.db.WithContext(ctx).Where("code = ?", code).Delete(&domain.AuthCodeSession{}).Error
}

func (r *ssoRepo) SaveSession(ctx context.Context, session *domain.UserSession) error {
	return r.db.WithContext(ctx).Create(session).Error
}

func (r *ssoRepo) GetSession(ctx context.Context, sessionID string) (*domain.UserSession, error) {
	var session domain.UserSession
	err := r.db.WithContext(ctx).Where("id = ?", sessionID).First(&session).Error
	return &session, err
}

func (r *ssoRepo) DeleteSession(ctx context.Context, sessionID string) error {
	return r.db.WithContext(ctx).Where("id = ?", sessionID).Delete(&domain.UserSession{}).Error
}
