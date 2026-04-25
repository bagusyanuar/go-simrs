package main

import (
	"log"

	"github.com/bagusyanuar/go-simrs/internal/shared/config"
	ssoDomain "github.com/bagusyanuar/go-simrs/internal/sso/domain"
	"github.com/google/uuid"
)

func main() {
	// 1. Load Config
	conf := config.LoadConfig()

	// 2. Initialize DB
	db := config.InitDB(conf)

	// 3. Prepare OAuth Client
	log.Println("Seeding SSO Whitelist Client...")

	clients := []ssoDomain.OAuthClient{
		{
			ID:           uuid.New(),
			ClientID:     "sso-app",
			Name:         "SIMRS SSO",
			RedirectURIs: "http://neurovi-simulation.test:5173/callback,http://neurovi-simulation.test:5174/callback,http://neurovi-simulation.test:5175/callback",
		},
		{
			ID:           uuid.New(),
			ClientID:     "master-data-app",
			Name:         "SIMRS Master Data App",
			RedirectURIs: "http://neurovi-simulation.test:5173/callback,http://neurovi-simulation.test:5174/callback,http://neurovi-simulation.test:5175/callback",
		},
		{
			ID:           uuid.New(),
			ClientID:     "simrs-mobile-app",
			Name:         "SIMRS Mobile",
			RedirectURIs: "simrs://callback",
		},
	}

	for _, client := range clients {
		if err := db.FirstOrCreate(&client, ssoDomain.OAuthClient{ClientID: client.ClientID}).Error; err != nil {
			log.Fatalf("Failed to seed client %s: %v", client.ClientID, err)
		}
		log.Printf("Client %s seeded successfully!\n", client.ClientID)
	}

	log.Println("SSO Seeding completed successfully!")
}
