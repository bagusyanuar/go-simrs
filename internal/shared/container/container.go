package container

import (
	authHandler "github.com/bagusyanuar/go-simrs/internal/auth/delivery/http"
	"github.com/bagusyanuar/go-simrs/internal/shared/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Container struct {
	AuthHandler *authHandler.AuthHandler
}

func NewContainer(db *gorm.DB, conf *config.Config) *Container {
	c := &Container{}

	c.wireAuthModule(db, conf)

	return c
}

func (c *Container) RegisterRoutes(router fiber.Router, authMiddleware fiber.Handler) {
	c.AuthHandler.Register(router)
}
