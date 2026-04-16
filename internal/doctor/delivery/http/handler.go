package http

import (
	"math"
	"time"

	"github.com/bagusyanuar/go-simrs/internal/doctor/domain"
	"github.com/bagusyanuar/go-simrs/pkg/request"
	"github.com/bagusyanuar/go-simrs/pkg/response"
	"github.com/bagusyanuar/go-simrs/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type DoctorHandler struct {
	uc domain.DoctorUsecase
}

func NewDoctorHandler(uc domain.DoctorUsecase) *DoctorHandler {
	return &DoctorHandler{uc: uc}
}

func (h *DoctorHandler) Register(router fiber.Router) {
	doctors := router.Group("/doctors")
	doctors.Post("/", h.Create)
	doctors.Get("/", h.GetAll)
	doctors.Get("/:id", h.GetByID)
	doctors.Put("/:id", h.Update)
	doctors.Delete("/:id", h.Delete)
}

func (h *DoctorHandler) Create(c *fiber.Ctx) error {
	var req CreateDoctorRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid request body"))
	}

	if errs := validator.ValidateStruct(req); len(errs) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.ValidationError(errs))
	}

	sipExpiry, err := time.Parse("2006-01-02", req.SIPExpiryDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid sip_expiry_date format, use YYYY-MM-DD"))
	}

	doctor := &domain.Doctor{
		SpecialtyID:   req.SpecialtyID,
		NIK:           req.NIK,
		SIP:           req.SIP,
		SIPExpiryDate: sipExpiry,
		Name:          req.Name,
		Phone:         req.Phone,
		Email:         req.Email,
		IsActive:      req.IsActive,
	}

	if err := h.uc.Create(c.Context(), doctor, req.UnitIDs); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(response.Success(doctor, "doctor created"))
}

func (h *DoctorHandler) GetAll(c *fiber.Ctx) error {
	var params request.PaginationParam
	if err := c.QueryParser(&params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid query parameters"))
	}

	var specialtyID *uuid.UUID
	if sIDStr := c.Query("specialty_id"); sIDStr != "" {
		if id, err := uuid.Parse(sIDStr); err == nil {
			specialtyID = &id
		}
	}

	var unitID *uuid.UUID
	if uIDStr := c.Query("unit_id"); uIDStr != "" {
		if id, err := uuid.Parse(uIDStr); err == nil {
			unitID = &id
		}
	}

	doctors, total, err := h.uc.GetAll(c.Context(), params, specialtyID, unitID)
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

	return c.Status(fiber.StatusOK).JSON(response.SuccessWithPagination(doctors, paginationResponse, "doctors retrieved"))
}

func (h *DoctorHandler) GetByID(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid id format"))
	}

	doctor, err := h.uc.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(response.Error("doctor not found"))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(doctor, "doctor retrieved"))
}

func (h *DoctorHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid id format"))
	}

	var req UpdateDoctorRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid request body"))
	}

	if errs := validator.ValidateStruct(req); len(errs) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.ValidationError(errs))
	}

	sipExpiry, err := time.Parse("2006-01-02", req.SIPExpiryDate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid sip_expiry_date format, use YYYY-MM-DD"))
	}

	doctor := &domain.Doctor{
		SpecialtyID:   req.SpecialtyID,
		NIK:           req.NIK,
		SIP:           req.SIP,
		SIPExpiryDate: sipExpiry,
		Name:          req.Name,
		Phone:         req.Phone,
		Email:         req.Email,
		IsActive:      req.IsActive,
	}

	if err := h.uc.Update(c.Context(), id, doctor, req.UnitIDs); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success[any](nil, "doctor updated"))
}

func (h *DoctorHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid id format"))
	}

	if err := h.uc.Delete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(response.Success[any](nil, "doctor deleted"))
}
