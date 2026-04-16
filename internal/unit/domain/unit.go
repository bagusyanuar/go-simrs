package domain

import (
	"context"
	"time"

	installationDomain "github.com/bagusyanuar/go-simrs/internal/installation/domain"
	"github.com/bagusyanuar/go-simrs/pkg/request"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Unit struct {
	ID             uuid.UUID                      `gorm:"type:uuid;primaryKey" json:"id"`
	InstallationID uuid.UUID                      `gorm:"type:uuid;not null" json:"installation_id"`
	Installation   *installationDomain.Installation `gorm:"foreignKey:InstallationID" json:"installation,omitempty"`
	Code           string                         `gorm:"type:varchar(50);unique;not null" json:"code"`
	Name           string                         `gorm:"type:varchar(255);not null" json:"name"`
	IsActive       bool                           `gorm:"default:true" json:"is_active"`
	CreatedAt      time.Time                      `json:"created_at"`
	UpdatedAt      time.Time                      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt                 `gorm:"index" json:"deleted_at,omitempty"`
}

func (u *Unit) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}

type UnitRepository interface {
	Create(ctx context.Context, unit *Unit) error
	FindAll(ctx context.Context, params request.PaginationParam, installationID *uuid.UUID) ([]Unit, int64, error)
	FindByID(ctx context.Context, id uuid.UUID) (*Unit, error)
	FindByCode(ctx context.Context, code string) (*Unit, error)
	Update(ctx context.Context, unit *Unit) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type UnitUsecase interface {
	Create(ctx context.Context, unit *Unit) error
	GetAll(ctx context.Context, params request.PaginationParam, installationID *uuid.UUID) ([]Unit, int64, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Unit, error)
	Update(ctx context.Context, id uuid.UUID, unit *Unit) error
	Delete(ctx context.Context, id uuid.UUID) error
}
