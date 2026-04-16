package usecase

import (
	"context"
	"fmt"

	"github.com/bagusyanuar/go-simrs/internal/specialty/domain"
	"github.com/bagusyanuar/go-simrs/pkg/request"
	"github.com/google/uuid"
)

type specialtyUC struct {
	repo domain.SpecialtyRepository
}

func NewSpecialtyUsecase(repo domain.SpecialtyRepository) domain.SpecialtyUsecase {
	return &specialtyUC{repo: repo}
}

func (u *specialtyUC) Create(ctx context.Context, specialty *domain.Specialty) error {
	// Check if code already exists
	existing, _ := u.repo.FindByCode(ctx, specialty.Code)
	if existing != nil {
		return fmt.Errorf("specialty code %s already exists", specialty.Code)
	}

	return u.repo.Create(ctx, specialty)
}

func (u *specialtyUC) GetAll(ctx context.Context, params request.PaginationParam) ([]domain.Specialty, int64, error) {
	return u.repo.FindAll(ctx, params)
}

func (u *specialtyUC) GetByID(ctx context.Context, id uuid.UUID) (*domain.Specialty, error) {
	return u.repo.FindByID(ctx, id)
}

func (u *specialtyUC) Update(ctx context.Context, id uuid.UUID, specialty *domain.Specialty) error {
	existing, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("specialty not found: %w", err)
	}

	// Update fields
	existing.Code = specialty.Code
	existing.Name = specialty.Name

	return u.repo.Update(ctx, existing)
}

func (u *specialtyUC) Delete(ctx context.Context, id uuid.UUID) error {
	return u.repo.Delete(ctx, id)
}
