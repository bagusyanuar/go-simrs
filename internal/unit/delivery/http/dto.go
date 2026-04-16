package http

import "github.com/google/uuid"

type CreateUnitRequest struct {
	InstallationID uuid.UUID `json:"installation_id" validate:"required"`
	Code           string    `json:"code" validate:"required,min=2,max=50"`
	Name           string    `json:"name" validate:"required,min=3,max=255"`
	IsActive       bool      `json:"is_active"`
}

type UpdateUnitRequest struct {
	InstallationID uuid.UUID `json:"installation_id" validate:"required"`
	Code           string    `json:"code" validate:"required,min=2,max=50"`
	Name           string    `json:"name" validate:"required,min=3,max=255"`
	IsActive       bool      `json:"is_active"`
}

type UnitResponse struct {
	ID             uuid.UUID `json:"id"`
	InstallationID uuid.UUID `json:"installation_id"`
	Installation   any       `json:"installation,omitempty"` // populated if preloaded
	Code           string    `json:"code"`
	Name           string    `json:"name"`
	IsActive       bool      `json:"is_active"`
	CreatedAt      string    `json:"created_at"`
}
