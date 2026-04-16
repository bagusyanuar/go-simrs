package container

import (
	authHandler "github.com/bagusyanuar/go-simrs/internal/auth/delivery/http"
	doctorHandler "github.com/bagusyanuar/go-simrs/internal/doctor/delivery/http"
	installationHandler "github.com/bagusyanuar/go-simrs/internal/installation/delivery/http"
	specialtyHandler "github.com/bagusyanuar/go-simrs/internal/specialty/delivery/http"
	unitHandler "github.com/bagusyanuar/go-simrs/internal/unit/delivery/http"
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
}

func NewContainer(db *gorm.DB, conf *config.Config) *Container {
	c := &Container{}

	c.wireAuthModule(db, conf)
	c.wireInstallationModule(db)
	c.wireUnitModule(db)
	c.wireSpecialtyModule(db)
	c.wireDoctorModule(db)

	return c
}

func (c *Container) RegisterRoutes(router fiber.Router, authMiddleware fiber.Handler) {
	c.AuthHandler.Register(router)

	// Protected Routes
	protected := router.Group("/", authMiddleware)
	c.InstallationHandler.Register(protected)
	c.UnitHandler.Register(protected)
	c.SpecialtyHandler.Register(protected)
	c.DoctorHandler.Register(protected)
}
