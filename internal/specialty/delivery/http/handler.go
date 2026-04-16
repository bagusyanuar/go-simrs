package http

import (
	"math"

	"github.com/bagusyanuar/go-simrs/internal/specialty/domain"
	"github.com/bagusyanuar/go-simrs/pkg/request"
	"github.com/bagusyanuar/go-simrs/pkg/response"
	"github.com/bagusyanuar/go-simrs/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type SpecialtyHandler struct {
	uc domain.SpecialtyUsecase
}

func NewSpecialtyHandler(uc domain.SpecialtyUsecase) *SpecialtyHandler {
	return &SpecialtyHandler{uc: uc}
}

func (h *SpecialtyHandler) Register(router fiber.Router) {
	specialties := router.Group("/specialties")
	specialties.Post("/", h.Create)
	specialties.Get("/", h.GetAll)
	specialties.Get("/:id", h.GetByID)
	specialties.Put("/:id", h.Update)
	specialties.Delete("/:id", h.Delete)
}

func (h *SpecialtyHandler) Create(c *fiber.Ctx) error {
	var req CreateSpecialtyRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid request body"))
	}

	if errs := validator.ValidateStruct(req); len(errs) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.ValidationError(errs))
	}

	specialty := &domain.Specialty{
		Code: req.Code,
		Name: req.Name,
	}

	if err := h.uc.Create(c.Context(), specialty); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(response.Success(specialty, "specialty created"))
}

func (h *SpecialtyHandler) GetAll(c *fiber.Ctx) error {
	var params request.PaginationParam
	if err := c.QueryParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid query parameters"))
	}

	specialties, total, err := h.uc.GetAll(c.Context(), params)
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

	return c.Status(fiber.StatusOK).JSON(response.SuccessWithPagination(specialties, paginationResponse, "specialties retrieved"))
}

func (h *SpecialtyHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid id format"))
	}

	specialty, err := h.uc.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.Error("specialty not found"))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(specialty, "specialty retrieved"))
}

func (h *SpecialtyHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid id format"))
	}

	var req UpdateSpecialtyRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid request body"))
	}

	if errs := validator.ValidateStruct(req); len(errs) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.ValidationError(errs))
	}

	specialty := &domain.Specialty{
		Code: req.Code,
		Name: req.Name,
	}

	if err := h.uc.Update(c.Context(), id, specialty); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success[any](nil, "specialty updated"))
}

func (h *SpecialtyHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid id format"))
	}

	if err := h.uc.Delete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success[any](nil, "specialty deleted"))
}
