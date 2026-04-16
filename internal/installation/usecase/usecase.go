package usecase

import (
	"context"
	"fmt"

	"github.com/bagusyanuar/go-simrs/internal/installation/domain"
	"github.com/google/uuid"
)

type installationUC struct {
	repo domain.InstallationRepository
}

func NewInstallationUsecase(repo domain.InstallationRepository) domain.InstallationUsecase {
	return &installationUC{repo: repo}
}

func (u *installationUC) Create(ctx context.Context, installation *domain.Installation) error {
	// Check if code already exists
	existing, _ := u.repo.FindByCode(ctx, installation.Code)
	if existing != nil {
		return fmt.Errorf("installation code %s already exists", installation.Code)
	}

	return u.repo.Create(ctx, installation)
}

func (u *installationUC) GetAll(ctx context.Context, filter domain.InstallationFilter) ([]domain.Installation, int64, error) {
	return u.repo.FindAll(ctx, filter)
}

func (u *installationUC) GetByID(ctx context.Context, id uuid.UUID) (*domain.Installation, error) {
	return u.repo.FindByID(ctx, id)
}

func (u *installationUC) Update(ctx context.Context, id uuid.UUID, installation *domain.Installation) error {
	existing, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("installation not found: %w", err)
	}

	// Update fields
	existing.Code = installation.Code
	existing.Name = installation.Name
	existing.IsMedical = installation.IsMedical
	existing.IsActive = installation.IsActive

	return u.repo.Update(ctx, existing)
}

func (u *installationUC) Delete(ctx context.Context, id uuid.UUID) error {
	return u.repo.Delete(ctx, id)
}
