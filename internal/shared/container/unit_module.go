package container

import (
	unitHandler "github.com/bagusyanuar/go-simrs/internal/unit/delivery/http"
	unitRepository "github.com/bagusyanuar/go-simrs/internal/unit/repository"
	unitUsecase "github.com/bagusyanuar/go-simrs/internal/unit/usecase"
	"gorm.io/gorm"
)

func (c *Container) wireUnitModule(db *gorm.DB) {
	// 1. Repository
	unitRepo := unitRepository.NewUnitRepository(db)

	// 2. Usecase
	unitUC := unitUsecase.NewUnitUsecase(unitRepo)

	// 3. Handler
	c.UnitHandler = unitHandler.NewUnitHandler(unitUC)
}
