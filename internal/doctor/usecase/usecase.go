package usecase

import (
	"context"
	"fmt"

	"github.com/bagusyanuar/go-simrs/internal/doctor/domain"
	unitDomain "github.com/bagusyanuar/go-simrs/internal/unit/domain"
	"github.com/bagusyanuar/go-simrs/pkg/request"
	"github.com/google/uuid"
)

type doctorUC struct {
	repo     domain.DoctorRepository
	unitRepo unitDomain.UnitRepository
}

func NewDoctorUsecase(repo domain.DoctorRepository, unitRepo unitDomain.UnitRepository) domain.DoctorUsecase {
	return &doctorUC{
		repo:     repo,
		unitRepo: unitRepo,
	}
}

func (u *doctorUC) Create(ctx context.Context, doctor *domain.Doctor, unitIDs []uuid.UUID) error {
	// Check if NIK already exists
	existing, _ := u.repo.FindByNIK(ctx, doctor.NIK)
	if existing != nil {
		return fmt.Errorf("doctor with NIK %s already exists", doctor.NIK)
	}

	// Map Unit IDs to Unit entities (Batch)
	units, err := u.unitRepo.FindByIDs(ctx, unitIDs)
	if err != nil {
		return fmt.Errorf("failed to fetch units: %w", err)
	}

	if len(units) != len(unitIDs) {
		return fmt.Errorf("some units are invalid or not found")
	}

	doctor.Units = units

	return u.repo.Create(ctx, doctor)
}

func (u *doctorUC) GetAll(ctx context.Context, params request.PaginationParam, specialtyID, unitID *uuid.UUID) ([]domain.Doctor, int64, error) {
	return u.repo.FindAll(ctx, params, specialtyID, unitID)
}

func (u *doctorUC) GetByID(ctx context.Context, id uuid.UUID) (*domain.Doctor, error) {
	return u.repo.FindByID(ctx, id)
}

func (u *doctorUC) Update(ctx context.Context, id uuid.UUID, doctor *domain.Doctor, unitIDs []uuid.UUID) error {
	existing, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("doctor not found: %w", err)
	}

	// Update basic fields
	existing.SpecialtyID = doctor.SpecialtyID
	existing.NIK = doctor.NIK
	existing.SIP = doctor.SIP
	existing.SIPExpiryDate = doctor.SIPExpiryDate
	existing.Name = doctor.Name
	existing.Phone = doctor.Phone
	existing.Email = doctor.Email
	existing.IsActive = doctor.IsActive

	// Map Unit IDs to Unit entities (Batch)
	units, err := u.unitRepo.FindByIDs(ctx, unitIDs)
	if err != nil {
		return fmt.Errorf("failed to fetch units: %w", err)
	}

	if len(units) != len(unitIDs) {
		return fmt.Errorf("some units are invalid or not found")
	}
	existing.Units = units

	return u.repo.Update(ctx, existing)
}

func (u *doctorUC) Delete(ctx context.Context, id uuid.UUID) error {
	return u.repo.Delete(ctx, id)
}
