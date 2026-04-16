package http

import (
	"time"

	"github.com/google/uuid"
)

type CreateDoctorRequest struct {
	SpecialtyID   uuid.UUID   `json:"specialty_id" validate:"required"`
	UnitIDs       []uuid.UUID `json:"unit_ids" validate:"required,min=1"`
	NIK           string      `json:"nik" validate:"required,min=16,max=50"`
	SIP           string      `json:"sip" validate:"required,min=5,max=100"`
	SIPExpiryDate string      `json:"sip_expiry_date" validate:"required" example:"2025-12-31"`
	Name          string      `json:"name" validate:"required,min=3,max=255"`
	Phone         string      `json:"phone" validate:"required"`
	Email         string      `json:"email" validate:"required,email"`
	IsActive      bool        `json:"is_active"`
}

type UpdateDoctorRequest struct {
	SpecialtyID   uuid.UUID   `json:"specialty_id" validate:"required"`
	UnitIDs       []uuid.UUID `json:"unit_ids" validate:"required,min=1"`
	NIK           string      `json:"nik" validate:"required,min=16,max=50"`
	SIP           string      `json:"sip" validate:"required,min=5,max=100"`
	SIPExpiryDate string      `json:"sip_expiry_date" validate:"required" example:"2025-12-31"`
	Name          string      `json:"name" validate:"required,min=3,max=255"`
	Phone         string      `json:"phone" validate:"required"`
	Email         string      `json:"email" validate:"required,email"`
	IsActive      bool        `json:"is_active"`
}

type DoctorResponse struct {
	ID            uuid.UUID `json:"id"`
	SpecialtyID   uuid.UUID `json:"specialty_id"`
	Specialty     any       `json:"specialty,omitempty"`
	Units         any       `json:"units,omitempty"`
	NIK           string    `json:"nik"`
	SIP           string    `json:"sip"`
	SIPExpiryDate time.Time `json:"sip_expiry_date"`
	Name          string    `json:"name"`
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
	IsActive      bool      `json:"is_active"`
}
