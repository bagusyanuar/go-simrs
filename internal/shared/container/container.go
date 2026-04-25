package container

import (
	authHandler "github.com/bagusyanuar/go-simrs/internal/auth/delivery/http"
	doctorHandler "github.com/bagusyanuar/go-simrs/internal/doctor/delivery/http"
	installationHandler "github.com/bagusyanuar/go-simrs/internal/installation/delivery/http"
	specialtyHandler "github.com/bagusyanuar/go-simrs/internal/specialty/delivery/http"
	unitHandler "github.com/bagusyanuar/go-simrs/internal/unit/delivery/http"
	ssoHandler "github.com/bagusyanuar/go-simrs/internal/sso/delivery/http"
	ssoUsecase "github.com/bagusyanuar/go-simrs/internal/sso/usecase"
	ssoRepository "github.com/bagusyanuar/go-simrs/internal/sso/repository"
	userRepository "github.com/bagusyanuar/go-simrs/internal/user/repository"
	"github.com/bagusyanuar/go-simrs/internal/shared/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Container struct {
	AuthHandler         *authHandler.AuthHandler
	InstallationHandler *installationHandler.InstallationHandler
	UnitHandler         *unitHandler.UnitHandler
	SpecialtyHandler    *specialtyHandler.SpecialtyHandler
	DoctorHandler       *doctorHandler.DoctorHandler
	SSOHandler          *ssoHandler.SSOHandler
}

func NewContainer(db *gorm.DB, conf *config.Config) *Container {
	c := &Container{}

	c.wireAuthModule(db, conf)
	c.wireInstallationModule(db)
	c.wireUnitModule(db)
	c.wireSpecialtyModule(db)
	c.wireDoctorModule(db)
	c.wireSSOModule(db, conf)

	return c
}

func (c *Container) RegisterRoutes(router fiber.Router, authMiddleware fiber.Handler) {
	c.AuthHandler.Register(router)
	c.SSOHandler.Register(router)

	// Protected Routes
	protected := router.Group("/", authMiddleware)
	c.InstallationHandler.Register(protected)
	c.UnitHandler.Register(protected)
	c.SpecialtyHandler.Register(protected)
	c.DoctorHandler.Register(protected)
}

func (c *Container) wireSSOModule(db *gorm.DB, conf *config.Config) {
	userRepo := userRepository.NewUserRepository(db)
	repo := ssoRepository.NewSSORepository(db)
	uc := ssoUsecase.NewSSOUsecase(repo, userRepo, conf)
	c.SSOHandler = ssoHandler.NewSSOHandler(uc, conf)
}
