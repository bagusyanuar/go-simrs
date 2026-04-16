package container

import (
	authHandler "github.com/bagusyanuar/go-simrs/internal/auth/delivery/http"
	authUsecase "github.com/bagusyanuar/go-simrs/internal/auth/usecase"
	"github.com/bagusyanuar/go-simrs/internal/shared/config"
	userRepository "github.com/bagusyanuar/go-simrs/internal/user/repository"
	"gorm.io/gorm"
)

func (c *Container) wireAuthModule(db *gorm.DB, conf *config.Config) {
	// 1. Repository
	userRepo := userRepository.NewUserRepository(db)

	// 2. Usecase
	authUC := authUsecase.NewAuthUsecase(userRepo, conf)

	// 3. Handler
	c.AuthHandler = authHandler.NewAuthHandler(authUC, conf)
}
