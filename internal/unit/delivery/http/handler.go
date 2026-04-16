package http

import (
	"math"

	"github.com/bagusyanuar/go-simrs/internal/unit/domain"
	"github.com/bagusyanuar/go-simrs/pkg/request"
	"github.com/bagusyanuar/go-simrs/pkg/response"
	"github.com/bagusyanuar/go-simrs/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UnitHandler struct {
	uc domain.UnitUsecase
}

func NewUnitHandler(uc domain.UnitUsecase) *UnitHandler {
	return &UnitHandler{uc: uc}
}

func (h *UnitHandler) Register(router fiber.Router) {
	units := router.Group("/units")
	units.Post("/", h.Create)
	units.Get("/", h.GetAll)
	units.Get("/:id", h.GetByID)
	units.Put("/:id", h.Update)
	units.Delete("/:id", h.Delete)
}

func (h *UnitHandler) Create(c *fiber.Ctx) error {
	var req CreateUnitRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid request body"))
	}

	if errs := validator.ValidateStruct(req); len(errs) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.ValidationError(errs))
	}

	unit := &domain.Unit{
		InstallationID: req.InstallationID,
		Code:           req.Code,
		Name:           req.Name,
		IsActive:       req.IsActive,
	}

	if err := h.uc.Create(c.Context(), unit); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(response.Success(unit, "unit created"))
}

func (h *UnitHandler) GetAll(c *fiber.Ctx) error {
	var params request.PaginationParam
	if err := c.QueryParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid query parameters"))
	}

	var installationID *uuid.UUID
	if instIDStr := c.Query("installation_id"); instIDStr != "" {
		if id, err := uuid.Parse(instIDStr); err == nil {
			installationID = &id
		}
	}

	units, total, err := h.uc.GetAll(c.Context(), params, installationID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err.Error()))
	}

	limit := params.GetLimit()
	totalPage := int(math.Ceil(float64(total) / float64(limit)))

	paginationResponse := response.Pagination{
		CurrentPage: params.GetPage(),
		Limit:       limit,
		TotalData:   total,
		TotalPage:   totalPage,
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessWithPagination(units, paginationResponse, "units retrieved"))
}

func (h *UnitHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid id format"))
	}

	unit, err := h.uc.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.Error("unit not found"))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(unit, "unit retrieved"))
}

func (h *UnitHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid id format"))
	}

	var req UpdateUnitRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid request body"))
	}

	if errs := validator.ValidateStruct(req); len(errs) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.ValidationError(errs))
	}

	unit := &domain.Unit{
		InstallationID: req.InstallationID,
		Code:           req.Code,
		Name:           req.Name,
		IsActive:       req.IsActive,
	}

	if err := h.uc.Update(c.Context(), id, unit); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success[any](nil, "unit updated"))
}

func (h *UnitHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid id format"))
	}

	if err := h.uc.Delete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success[any](nil, "unit deleted"))
}
