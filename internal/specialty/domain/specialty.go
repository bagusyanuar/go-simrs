package domain

import (
	"context"
	"time"

	"github.com/bagusyanuar/go-simrs/pkg/request"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Specialty struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Code      string         `gorm:"type:varchar(50);unique;not null" json:"code"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (s *Specialty) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return
}

type SpecialtyRepository interface {
	Create(ctx context.Context, specialty *Specialty) error
	FindAll(ctx context.Context, params request.PaginationParam) ([]Specialty, int64, error)
	FindByID(ctx context.Context, id uuid.UUID) (*Specialty, error)
	FindByCode(ctx context.Context, code string) (*Specialty, error)
	Update(ctx context.Context, specialty *Specialty) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type SpecialtyUsecase interface {
	Create(ctx context.Context, specialty *Specialty) error
	GetAll(ctx context.Context, params request.PaginationParam) ([]Specialty, int64, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Specialty, error)
	Update(ctx context.Context, id uuid.UUID, specialty *Specialty) error
	Delete(ctx context.Context, id uuid.UUID) error
}
