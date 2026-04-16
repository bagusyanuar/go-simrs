package http

import (
	"math"

	"github.com/bagusyanuar/go-simrs/internal/installation/domain"
	"github.com/bagusyanuar/go-simrs/pkg/request"
	"github.com/bagusyanuar/go-simrs/pkg/response"
	"github.com/bagusyanuar/go-simrs/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type InstallationHandler struct {
	uc domain.InstallationUsecase
}

func NewInstallationHandler(uc domain.InstallationUsecase) *InstallationHandler {
	return &InstallationHandler{uc: uc}
}

func (h *InstallationHandler) Register(router fiber.Router) {
	installations := router.Group("/installations")
	installations.Post("/", h.Create)
	installations.Get("/", h.GetAll)
	installations.Get("/:id", h.GetByID)
	installations.Put("/:id", h.Update)
	installations.Delete("/:id", h.Delete)
}

func (h *InstallationHandler) Create(c *fiber.Ctx) error {
	var req CreateInstallationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid request body"))
	}

	if errs := validator.ValidateStruct(req); len(errs) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.ValidationError(errs))
	}

	installation := &domain.Installation{
		Code:      req.Code,
		Name:      req.Name,
		IsMedical: req.IsMedical,
		IsActive:  req.IsActive,
	}

	if err := h.uc.Create(c.Context(), installation); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(response.Success(installation, "installation created"))
}

func (h *InstallationHandler) GetAll(c *fiber.Ctx) error {
	var params request.PaginationParam
	if err := c.QueryParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid query parameters"))
	}

	installations, total, err := h.uc.GetAll(c.Context(), params)
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

	return c.Status(fiber.StatusOK).JSON(response.SuccessWithPagination(installations, paginationResponse, "installations retrieved"))
}

func (h *InstallationHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid id format"))
	}

	installation, err := h.uc.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.Error("installation not found"))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(installation, "installation retrieved"))
}

func (h *InstallationHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid id format"))
	}

	var req UpdateInstallationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid request body"))
	}

	if errs := validator.ValidateStruct(req); len(errs) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.ValidationError(errs))
	}

	installation := &domain.Installation{
		Code:      req.Code,
		Name:      req.Name,
		IsMedical: req.IsMedical,
		IsActive:  req.IsActive,
	}

	if err := h.uc.Update(c.Context(), id, installation); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success[any](nil, "installation updated"))
}

func (h *InstallationHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid id format"))
	}

	if err := h.uc.Delete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success[any](nil, "installation deleted"))
}
