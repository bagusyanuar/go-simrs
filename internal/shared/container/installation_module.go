package container

import (
	installationHandler "github.com/bagusyanuar/go-simrs/internal/installation/delivery/http"
	installationRepository "github.com/bagusyanuar/go-simrs/internal/installation/repository"
	installationUsecase "github.com/bagusyanuar/go-simrs/internal/installation/usecase"
	"gorm.io/gorm"
)

func (c *Container) wireInstallationModule(db *gorm.DB) {
	// 1. Repository
	installationRepo := installationRepository.NewInstallationRepository(db)

	// 2. Usecase
	installationUC := installationUsecase.NewInstallationUsecase(installationRepo)

	// 3. Handler
	c.InstallationHandler = installationHandler.NewInstallationHandler(installationUC)
}
