package repository

import (
	"context"
	"fmt"

	"github.com/bagusyanuar/go-simrs/internal/specialty/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type specialtyRepo struct {
	db *gorm.DB
}

func NewSpecialtyRepository(db *gorm.DB) domain.SpecialtyRepository {
	return &specialtyRepo{db: db}
}

func (r *specialtyRepo) Create(ctx context.Context, specialty *domain.Specialty) error {
	if err := r.db.WithContext(ctx).Create(specialty).Error; err != nil {
		return fmt.Errorf("specialtyRepo.Create: %w", err)
	}
	return nil
}

func (r *specialtyRepo) FindAll(ctx context.Context, filter domain.SpecialtyFilter) ([]domain.Specialty, int64, error) {
	var specialties []domain.Specialty
	var total int64

	db := r.db.WithContext(ctx).Model(&domain.Specialty{})

	if filter.Search != "" {
		searchTerm := "%" + filter.Search + "%"
		db = db.Where("name ILIKE ? OR code ILIKE ?", searchTerm, searchTerm)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("specialtyRepo.FindAll.Count: %w", err)
	}

	if err := db.Order(filter.GetSort()).
		Limit(filter.GetLimit()).
		Offset(filter.GetOffset()).
		Find(&specialties).Error; err != nil {
		return nil, 0, fmt.Errorf("specialtyRepo.FindAll.Find: %w", err)
	}

	return specialties, total, nil
}

func (r *specialtyRepo) FindByID(ctx context.Context, id uuid.UUID) (*domain.Specialty, error) {
	var specialty domain.Specialty
	if err := r.db.WithContext(ctx).First(&specialty, id).Error; err != nil {
		return nil, fmt.Errorf("specialtyRepo.FindByID: %w", err)
	}
	return &specialty, nil
}

func (r *specialtyRepo) FindByCode(ctx context.Context, code string) (*domain.Specialty, error) {
	var specialty domain.Specialty
	if err := r.db.WithContext(ctx).Where("code = ?", code).First(&specialty).Error; err != nil {
		return nil, fmt.Errorf("specialtyRepo.FindByCode: %w", err)
	}
	return &specialty, nil
}

func (r *specialtyRepo) Update(ctx context.Context, specialty *domain.Specialty) error {
	if err := r.db.WithContext(ctx).Save(specialty).Error; err != nil {
		return fmt.Errorf("specialtyRepo.Update: %w", err)
	}
	return nil
}

func (r *specialtyRepo) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Delete(&domain.Specialty{}, id).Error; err != nil {
		return fmt.Errorf("specialtyRepo.Delete: %w", err)
	}
	return nil
}
