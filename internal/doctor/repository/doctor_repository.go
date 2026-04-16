package repository

import (
	"context"
	"fmt"

	"github.com/bagusyanuar/go-simrs/internal/doctor/domain"
	"github.com/bagusyanuar/go-simrs/pkg/request"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type doctorRepo struct {
	db *gorm.DB
}

func NewDoctorRepository(db *gorm.DB) domain.DoctorRepository {
	return &doctorRepo{db: db}
}

func (r *doctorRepo) Create(ctx context.Context, doctor *domain.Doctor) error {
	if err := r.db.WithContext(ctx).Create(doctor).Error; err != nil {
		return fmt.Errorf("doctorRepo.Create: %w", err)
	}
	return nil
}

func (r *doctorRepo) FindAll(ctx context.Context, params request.PaginationParam, specialtyID, unitID *uuid.UUID) ([]domain.Doctor, int64, error) {
	var doctors []domain.Doctor
	var total int64

	db := r.db.WithContext(ctx).Model(&domain.Doctor{})

	if specialtyID != nil {
		db = db.Where("specialty_id = ?", specialtyID)
	}

	if unitID != nil {
		db = db.Joins("JOIN doctor_units ON doctor_units.doctor_id = doctors.id").
			Where("doctor_units.unit_id = ?", unitID)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("doctorRepo.FindAll.Count: %w", err)
	}

	if err := db.Order(params.GetSort()).
		Limit(params.GetLimit()).
		Offset(params.GetOffset()).
		Preload("Specialty").
		Preload("Units").
		Find(&doctors).Error; err != nil {
		return nil, 0, fmt.Errorf("doctorRepo.FindAll.Find: %w", err)
	}

	return doctors, total, nil
}

func (r *doctorRepo) FindByID(ctx context.Context, id uuid.UUID) (*domain.Doctor, error) {
	var doctor domain.Doctor
	if err := r.db.WithContext(ctx).
		Preload("Specialty").
		Preload("Units").
		First(&doctor, id).Error; err != nil {
		return nil, fmt.Errorf("doctorRepo.FindByID: %w", err)
	}
	return &doctor, nil
}

func (r *doctorRepo) FindByNIK(ctx context.Context, nik string) (*domain.Doctor, error) {
	var doctor domain.Doctor
	if err := r.db.WithContext(ctx).Where("nik = ?", nik).First(&doctor).Error; err != nil {
		return nil, fmt.Errorf("doctorRepo.FindByNIK: %w", err)
	}
	return &doctor, nil
}

func (r *doctorRepo) Update(ctx context.Context, doctor *domain.Doctor) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Update main record
		if err := tx.Save(doctor).Error; err != nil {
			return err
		}

		// Update Many-to-Many associations
		if err := tx.Model(doctor).Association("Units").Replace(doctor.Units); err != nil {
			return err
		}

		return nil
	})
}

func (r *doctorRepo) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Delete(&domain.Doctor{}, id).Error; err != nil {
		return fmt.Errorf("doctorRepo.Delete: %w", err)
	}
	return nil
}
