package repository

import (
	"context"
	"fmt"

	"github.com/bagusyanuar/go-simrs/internal/installation/domain"
	"github.com/bagusyanuar/go-simrs/pkg/request"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type installationRepo struct {
	db *gorm.DB
}

func NewInstallationRepository(db *gorm.DB) domain.InstallationRepository {
	return &installationRepo{db: db}
}

func (r *installationRepo) Create(ctx context.Context, installation *domain.Installation) error {
	if err := r.db.WithContext(ctx).Create(installation).Error; err != nil {
		return fmt.Errorf("installationRepo.Create: %w", err)
	}
	return nil
}

func (r *installationRepo) FindAll(ctx context.Context, params request.PaginationParam) ([]domain.Installation, int64, error) {
	var installations []domain.Installation
	var total int64

	db := r.db.WithContext(ctx).Model(&domain.Installation{})

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("installationRepo.FindAll.Count: %w", err)
	}

	if err := db.Order(params.GetSort()).
		Limit(params.GetLimit()).
		Offset(params.GetOffset()).
		Find(&installations).Error; err != nil {
		return nil, 0, fmt.Errorf("installationRepo.FindAll.Find: %w", err)
	}

	return installations, total, nil
}

func (r *installationRepo) FindByID(ctx context.Context, id uuid.UUID) (*domain.Installation, error) {
	var installation domain.Installation
	if err := r.db.WithContext(ctx).First(&installation, id).Error; err != nil {
		return nil, fmt.Errorf("installationRepo.FindByID: %w", err)
	}
	return &installation, nil
}

func (r *installationRepo) FindByCode(ctx context.Context, code string) (*domain.Installation, error) {
	var installation domain.Installation
	if err := r.db.WithContext(ctx).Where("code = ?", code).First(&installation).Error; err != nil {
		return nil, fmt.Errorf("installationRepo.FindByCode: %w", err)
	}
	return &installation, nil
}

func (r *installationRepo) Update(ctx context.Context, installation *domain.Installation) error {
	if err := r.db.WithContext(ctx).Save(installation).Error; err != nil {
		return fmt.Errorf("installationRepo.Update: %w", err)
	}
	return nil
}

func (r *installationRepo) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Delete(&domain.Installation{}, id).Error; err != nil {
		return fmt.Errorf("installationRepo.Delete: %w", err)
	}
	return nil
}
