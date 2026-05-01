package http

import (
	"time"

	"github.com/bagusyanuar/go-simrs/internal/shared/config"
	"github.com/bagusyanuar/go-simrs/internal/sso/domain"
	"github.com/bagusyanuar/go-simrs/pkg/response"
	"github.com/bagusyanuar/go-simrs/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type SSOHandler struct {
	ssoUC domain.SSOUsecase
	conf  *config.Config
}

func NewSSOHandler(uc domain.SSOUsecase, conf *config.Config) *SSOHandler {
	return &SSOHandler{ssoUC: uc, conf: conf}
}

func (h *SSOHandler) Register(router fiber.Router) {
	sso := router.Group("/sso")
	sso.Post("/authorize", h.Authorize)
	sso.Get("/authorize", h.AuthorizeSilent) // For True SSO (Silent Login)
	sso.Post("/token", h.ExchangeToken)
}

func (h *SSOHandler) Authorize(c *fiber.Ctx) error {
	var req AuthorizeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid request body"))
	}

	if errs := validator.ValidateStruct(req); len(errs) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.ValidationError(errs))
	}

	code, sessionID, err := h.ssoUC.Authorize(c.Context(), domain.AuthorizeRequest{
		Email:         req.Email,
		Password:      req.Password,
		ClientID:      req.ClientID,
		CodeChallenge: req.CodeChallenge,
		RedirectURI:   req.RedirectURI,
		State:         req.State,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err.Error()))
	}

	// Set SSO Session Cookie
	secure := h.conf.AppEnv == "production"
	sameSite := "Lax"
	if secure {
		sameSite = "None"
	}

	domain := h.conf.AppDomain
	// If domain doesn't start with dot and it's not localhost, add it for cross-subdomain support
	if domain != "" && domain != "localhost" && domain[0] != '.' {
		domain = "." + domain
	}

	c.Cookie(&fiber.Cookie{
		Name:     "sso_session",
		Value:    sessionID,
		HTTPOnly: true,
		Secure:   secure,
		SameSite: sameSite,
		Path:     "/",
		Domain:   domain,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	return c.Status(fiber.StatusOK).JSON(response.Success(fiber.Map{
		"code":  code,
		"state": req.State,
	}, "authorize success"))
}

func (h *SSOHandler) AuthorizeSilent(c *fiber.Ctx) error {
	var req AuthorizeSilentRequest
	if err := c.QueryParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid query parameters"))
	}

	if errs := validator.ValidateStruct(req); len(errs) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.ValidationError(errs))
	}

	if req.ResponseType != "code" {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("unsupported response_type"))
	}

	sessionID := c.Cookies("sso_session")
	if sessionID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "login_required",
			"error":   "login_required",
		})
	}

	code, err := h.ssoUC.AuthorizeSilent(c.Context(), sessionID, domain.AuthorizeSilentRequest{
		ClientID:      req.ClientID,
		CodeChallenge: req.CodeChallenge,
		RedirectURI:   req.RedirectURI,
		State:         req.State,
		ResponseType:  req.ResponseType,
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "login_required",
			"error":   "login_required",
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.Success(fiber.Map{
		"code":  code,
		"state": req.State,
	}, "silent authorize success"))
}

func (h *SSOHandler) ExchangeToken(c *fiber.Ctx) error {
	var req TokenRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("invalid request body"))
	}

	if errs := validator.ValidateStruct(req); len(errs) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.ValidationError(errs))
	}

	var res *domain.TokenResponse
	var err error

	switch req.GrantType {
	case "authorization_code":
		res, err = h.ssoUC.ExchangeToken(c.Context(), domain.TokenRequest{
			GrantType:    req.GrantType,
			ClientID:     req.ClientID,
			Code:         req.Code,
			CodeVerifier: req.CodeVerifier,
			RedirectURI:  req.RedirectURI,
		})
	case "refresh_token":
		refreshToken := c.Cookies("refresh_token")
		if refreshToken == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(response.Error("refresh token not found"))
		}

		res, err = h.ssoUC.RefreshToken(c.Context(), domain.TokenRequest{
			GrantType:    req.GrantType,
			ClientID:     req.ClientID,
			RefreshToken: refreshToken,
		})
	default:
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("unsupported grant_type"))
	}

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Error(err.Error()))
	}

	// Set Refresh Token Cookie
	secure := h.conf.AppEnv == "production"
	sameSite := "Lax"
	if secure {
		sameSite = "None"
	}

	domain := h.conf.AppDomain
	if domain != "" && domain != "localhost" && domain[0] != '.' {
		domain = "." + domain
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    res.RefreshToken,
		HTTPOnly: true,
		Secure:   secure,
		SameSite: sameSite,
		Path:     "/",
		Domain:   domain,
		Expires:  time.Now().Add(time.Duration(h.conf.JWTRefreshExpiration) * time.Hour),
	})

	// Hide refresh token from JSON response if you want it only in cookie
	res.RefreshToken = ""

	return c.Status(fiber.StatusOK).JSON(response.Success(res, "token success"))
}
