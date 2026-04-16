package container

import (
	doctorHandler "github.com/bagusyanuar/go-simrs/internal/doctor/delivery/http"
	doctorRepository "github.com/bagusyanuar/go-simrs/internal/doctor/repository"
	doctorUsecase "github.com/bagusyanuar/go-simrs/internal/doctor/usecase"
	unitRepository "github.com/bagusyanuar/go-simrs/internal/unit/repository"
	"gorm.io/gorm"
)

func (c *Container) wireDoctorModule(db *gorm.DB) {
	// 1. Dependency Repositories
	unitRepo := unitRepository.NewUnitRepository(db)

	// 2. Repository
	doctorRepo := doctorRepository.NewDoctorRepository(db)

	// 3. Usecase
	doctorUC := doctorUsecase.NewDoctorUsecase(doctorRepo, unitRepo)

	// 4. Handler
	c.DoctorHandler = doctorHandler.NewDoctorHandler(doctorUC)
}
