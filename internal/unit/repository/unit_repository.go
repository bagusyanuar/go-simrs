package repository

import (
	"context"
	"fmt"

	"github.com/bagusyanuar/go-simrs/internal/unit/domain"
	"github.com/bagusyanuar/go-simrs/pkg/request"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type unitRepo struct {
	db *gorm.DB
}

func NewUnitRepository(db *gorm.DB) domain.UnitRepository {
	return &unitRepo{db: db}
}

func (r *unitRepo) Create(ctx context.Context, unit *domain.Unit) error {
	if err := r.db.WithContext(ctx).Create(unit).Error; err != nil {
		return fmt.Errorf("unitRepo.Create: %w", err)
	}
	return nil
}

func (r *unitRepo) FindAll(ctx context.Context, params request.PaginationParam, installationID *uuid.UUID) ([]domain.Unit, int64, error) {
	var units []domain.Unit
	var total int64

	db := r.db.WithContext(ctx).Model(&domain.Unit{})

	if installationID != nil {
		db = db.Where("installation_id = ?", installationID)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("unitRepo.FindAll.Count: %w", err)
	}

	if err := db.Order(params.GetSort()).
		Limit(params.GetLimit()).
		Offset(params.GetOffset()).
		Preload("Installation").
		Find(&units).Error; err != nil {
		return nil, 0, fmt.Errorf("unitRepo.FindAll.Find: %w", err)
	}

	return units, total, nil
}

func (r *unitRepo) FindByID(ctx context.Context, id uuid.UUID) (*domain.Unit, error) {
	var unit domain.Unit
	if err := r.db.WithContext(ctx).Preload("Installation").First(&unit, id).Error; err != nil {
		return nil, fmt.Errorf("unitRepo.FindByID: %w", err)
	}
	return &unit, nil
}

func (r *unitRepo) FindByIDs(ctx context.Context, ids []uuid.UUID) ([]domain.Unit, error) {
	var units []domain.Unit
	if err := r.db.WithContext(ctx).Where("id IN ?", ids).Find(&units).Error; err != nil {
		return nil, fmt.Errorf("unitRepo.FindByIDs: %w", err)
	}
	return units, nil
}

func (r *unitRepo) FindByCode(ctx context.Context, code string) (*domain.Unit, error) {
	var unit domain.Unit
	if err := r.db.WithContext(ctx).Where("code = ?", code).First(&unit).Error; err != nil {
		return nil, fmt.Errorf("unitRepo.FindByCode: %w", err)
	}
	return &unit, nil
}

func (r *unitRepo) Update(ctx context.Context, unit *domain.Unit) error {
	if err := r.db.WithContext(ctx).Save(unit).Error; err != nil {
		return fmt.Errorf("unitRepo.Update: %w", err)
	}
	return nil
}

func (r *unitRepo) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Delete(&domain.Unit{}, id).Error; err != nil {
		return fmt.Errorf("unitRepo.Delete: %w", err)
	}
	return nil
}
