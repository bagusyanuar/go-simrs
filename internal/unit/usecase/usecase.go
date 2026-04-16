package usecase

import (
	"context"
	"fmt"

	"github.com/bagusyanuar/go-simrs/internal/unit/domain"
	"github.com/bagusyanuar/go-simrs/pkg/request"
	"github.com/google/uuid"
)

type unitUC struct {
	repo domain.UnitRepository
}

func NewUnitUsecase(repo domain.UnitRepository) domain.UnitUsecase {
	return &unitUC{repo: repo}
}

func (u *unitUC) Create(ctx context.Context, unit *domain.Unit) error {
	// Check if code already exists
	existing, _ := u.repo.FindByCode(ctx, unit.Code)
	if existing != nil {
		return fmt.Errorf("unit code %s already exists", unit.Code)
	}

	return u.repo.Create(ctx, unit)
}

func (u *unitUC) GetAll(ctx context.Context, params request.PaginationParam, installationID *uuid.UUID) ([]domain.Unit, int64, error) {
	return u.repo.FindAll(ctx, params, installationID)
}

func (u *unitUC) GetByID(ctx context.Context, id uuid.UUID) (*domain.Unit, error) {
	return u.repo.FindByID(ctx, id)
}

func (u *unitUC) Update(ctx context.Context, id uuid.UUID, unit *domain.Unit) error {
	existing, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("unit not found: %w", err)
	}

	// Update fields
	existing.InstallationID = unit.InstallationID
	existing.Code = unit.Code
	existing.Name = unit.Name
	existing.IsActive = unit.IsActive

	return u.repo.Update(ctx, existing)
}

func (u *unitUC) Delete(ctx context.Context, id uuid.UUID) error {
	return u.repo.Delete(ctx, id)
}
