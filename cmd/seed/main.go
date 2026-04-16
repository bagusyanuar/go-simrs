package main

import (
	"log"

	"github.com/bagusyanuar/go-simrs/internal/shared/config"
	"github.com/bagusyanuar/go-simrs/internal/user/domain"
	"github.com/bagusyanuar/go-simrs/pkg/password"
	"github.com/google/uuid"
)

func main() {
	// 1. Load Config
	conf := config.LoadConfig()

	// 2. Initialize DB
	db := config.InitDB(conf)

	// 3. Prepare Dummy User
	hashedPassword, err := password.HashPassword("password123")
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	adminUser := domain.User{
		ID:       uuid.New(),
		Email:    "admin@simrs.com",
		Username: "admin",
		Password: hashedPassword,
	}

	// 4. Seed to DB
	log.Println("Seeding admin user...")
	if err := db.FirstOrCreate(&adminUser, domain.User{Username: "admin"}).Error; err != nil {
		log.Fatalf("Failed to seed admin user: %v", err)
	}

	log.Println("Seeding completed successfully!")
}
