package seeds

import (
	"newsportal-backend/internal/core/domain/model"
	"newsportal-backend/lib/conv"

	"github.com/rs/zerolog/log"

	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	users := []model.User{
		{Name: "John Doe", Email: "john@example.com", Password: "123456"},
		{Name: "Jane Smith", Email: "jane@example.com", Password: "123456"},
		{Name: "Alice Johnson", Email: "alice@example.com", Password: "123456"},
		{Name: "Bob Brown", Email: "bob@example.com", Password: "123456"},
		{Name: "Charlie White", Email: "charlie@example.com", Password: "123456"},
		{Name: "David Black", Email: "david@example.com", Password: "123456"},
		{Name: "Eve Green", Email: "eve@example.com", Password: "123456"},
		{Name: "Frank Blue", Email: "frank@example.com", Password: "123456"},
		{Name: "Grace Red", Email: "grace@example.com", Password: "123456"},
		{Name: "Hank Yellow", Email: "hank@example.com", Password: "123456"},
	}

	for i := range users {
		bytes, err := conv.HashPassword(users[i].Password)
		if err != nil {
			log.Fatal().Err(err).Msg("Error hashing password")
		}
		users[i].Password = string(bytes)
	}

	if err := db.Create(&users).Error; err != nil {
		log.Fatal().Err(err).Msg("Error seeding users")
	} else {
		log.Info().Msg("Successfully seeded")
	}
}
