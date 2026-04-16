package container

import (
	authHandler "github.com/bagusyanuar/go-simrs/internal/auth/delivery/http"
	installationHandler "github.com/bagusyanuar/go-simrs/internal/installation/delivery/http"
	"github.com/bagusyanuar/go-simrs/internal/shared/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Container struct {
	AuthHandler         *authHandler.AuthHandler
	InstallationHandler *installationHandler.InstallationHandler
}

func NewContainer(db *gorm.DB, conf *config.Config) *Container {
	c := &Container{}

	c.wireAuthModule(db, conf)
	c.wireInstallationModule(db)

	return c
}

func (c *Container) RegisterRoutes(router fiber.Router, authMiddleware fiber.Handler) {
	c.AuthHandler.Register(router)

	// Protected Routes
	protected := router.Group("/", authMiddleware)
	c.InstallationHandler.Register(protected)
}
