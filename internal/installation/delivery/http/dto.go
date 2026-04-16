package http

import "github.com/google/uuid"

type CreateInstallationRequest struct {
	Code      string `json:"code" validate:"required,min=2,max=50"`
	Name      string `json:"name" validate:"required,min=3,max=255"`
	IsMedical bool   `json:"is_medical"`
	IsActive  bool   `json:"is_active"`
}

type UpdateInstallationRequest struct {
	Code      string `json:"code" validate:"required,min=2,max=50"`
	Name      string `json:"name" validate:"required,min=3,max=255"`
	IsMedical bool   `json:"is_medical"`
	IsActive  bool   `json:"is_active"`
}

type InstallationResponse struct {
	ID        uuid.UUID `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	IsMedical bool      `json:"is_medical"`
	IsActive  bool      `json:"is_active"`
	CreatedAt string    `json:"created_at"`
}
