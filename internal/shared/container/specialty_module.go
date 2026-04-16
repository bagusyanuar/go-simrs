package container

import (
	specialtyHandler "github.com/bagusyanuar/go-simrs/internal/specialty/delivery/http"
	specialtyRepository "github.com/bagusyanuar/go-simrs/internal/specialty/repository"
	specialtyUsecase "github.com/bagusyanuar/go-simrs/internal/specialty/usecase"
	"gorm.io/gorm"
)

func (c *Container) wireSpecialtyModule(db *gorm.DB) {
	// 1. Repository
	specialtyRepo := specialtyRepository.NewSpecialtyRepository(db)

	// 2. Usecase
	specialtyUC := specialtyUsecase.NewSpecialtyUsecase(specialtyRepo)

	// 3. Handler
	c.SpecialtyHandler = specialtyHandler.NewSpecialtyHandler(specialtyUC)
}
