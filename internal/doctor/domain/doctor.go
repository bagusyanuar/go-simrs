package domain

import (
	"context"
	"time"

	specialtyDomain "github.com/bagusyanuar/go-simrs/internal/specialty/domain"
	unitDomain "github.com/bagusyanuar/go-simrs/internal/unit/domain"
	"github.com/bagusyanuar/go-simrs/pkg/request"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Doctor struct {
	ID            uuid.UUID                 `gorm:"type:uuid;primaryKey" json:"id"`
	SpecialtyID   uuid.UUID                 `gorm:"type:uuid;not null" json:"specialty_id"`
	Specialty     *specialtyDomain.Specialty `gorm:"foreignKey:SpecialtyID" json:"specialty,omitempty"`
	Units         []unitDomain.Unit         `gorm:"many2many:doctor_units;" json:"units,omitempty"`
	NIK           string                    `gorm:"type:varchar(50);unique;not null" json:"nik"`
	SIP           string                    `gorm:"column:sip;type:varchar(100);unique;not null" json:"sip"`
	SIPExpiryDate time.Time                 `gorm:"column:sip_expiry_date" json:"sip_expiry_date"`
	Name          string                    `gorm:"type:varchar(255);not null" json:"name"`
	Phone         string                    `gorm:"type:varchar(20)" json:"phone"`
	Email         string                    `gorm:"type:varchar(100)" json:"email"`
	IsActive      bool                      `gorm:"default:true" json:"is_active"`
	CreatedAt     time.Time                 `json:"created_at"`
	UpdatedAt     time.Time                 `json:"updated_at"`
	DeletedAt     gorm.DeletedAt            `gorm:"index" json:"deleted_at,omitempty"`
}

func (d *Doctor) BeforeCreate(tx *gorm.DB) (err error) {
	if d.ID == uuid.Nil {
		d.ID = uuid.New()
	}
	return
}

type DoctorRepository interface {
	Create(ctx context.Context, doctor *Doctor) error
	FindAll(ctx context.Context, params request.PaginationParam, specialtyID, unitID *uuid.UUID) ([]Doctor, int64, error)
	FindByID(ctx context.Context, id uuid.UUID) (*Doctor, error)
	FindByNIK(ctx context.Context, nik string) (*Doctor, error)
	Update(ctx context.Context, doctor *Doctor) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type DoctorUsecase interface {
	Create(ctx context.Context, doctor *Doctor, unitIDs []uuid.UUID) error
	GetAll(ctx context.Context, params request.PaginationParam, specialtyID, unitID *uuid.UUID) ([]Doctor, int64, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Doctor, error)
	Update(ctx context.Context, id uuid.UUID, doctor *Doctor, unitIDs []uuid.UUID) error
	Delete(ctx context.Context, id uuid.UUID) error
}
