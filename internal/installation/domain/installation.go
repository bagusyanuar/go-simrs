package domain

import (
	"context"
	"time"

	"github.com/bagusyanuar/go-simrs/pkg/request"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Installation struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Code      string         `gorm:"type:varchar(50);unique;not null" json:"code"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	IsMedical bool           `gorm:"default:true" json:"is_medical"`
	IsActive  bool           `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (i *Installation) BeforeCreate(tx *gorm.DB) (err error) {
	if i.ID == uuid.Nil {
		i.ID = uuid.New()
	}
	return
}

type InstallationRepository interface {
	Create(ctx context.Context, installation *Installation) error
	FindAll(ctx context.Context, params request.PaginationParam) ([]Installation, int64, error)
	FindByID(ctx context.Context, id uuid.UUID) (*Installation, error)
	FindByCode(ctx context.Context, code string) (*Installation, error)
	Update(ctx context.Context, installation *Installation) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type InstallationUsecase interface {
	Create(ctx context.Context, installation *Installation) error
	GetAll(ctx context.Context, params request.PaginationParam) ([]Installation, int64, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Installation, error)
	Update(ctx context.Context, id uuid.UUID, installation *Installation) error
	Delete(ctx context.Context, id uuid.UUID) error
}
