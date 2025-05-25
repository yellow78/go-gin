package repository

import "go-gin/internal/domain/game/model/entity" // Adjust import path as needed

type DigimonSpeciesRepository interface {
	FindByID(id string) (*entity.DigimonSpecies, error)
	FindByName(name string) (*entity.DigimonSpecies, error)
	ListAll(offset int, limit int) ([]*entity.DigimonSpecies, error) // Added pagination params
	// Add Create, Update, Delete methods if species can be managed via API later
}
