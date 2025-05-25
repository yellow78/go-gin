package usecase

import (
	"errors" // For custom errors
	dto "go-gin/internal/application/dto/game"
	digimonEntity "go-gin/internal/domain/game/model/entity"
	digimonRepo "go-gin/internal/domain/game/repository"
	"log" // For logging
	"time" // For AcquiredAt, and formatting response

	"github.com/google/uuid" // For generating Digimon instance ID
	// accountRepo "go-gin/internal/domain/account/repository" // If needed
)

// Errors for Digimon Usecases
var (
	ErrSpeciesNotFound        = errors.New("digimon species not found")
	ErrMaxDigimonSlotsReached = errors.New("maximum digimon slots reached for player") // Example custom error
	ErrDigimonNotFound        = errors.New("digimon not found")                        // Specific error for this usecase
	ErrDigimonNotOwned        = errors.New("player does not own this digimon")         // Specific error
)

// AcquireDigimonUsecase handles acquiring a new Digimon.
type AcquireDigimonUsecase struct {
	digimonRepo        digimonRepo.DigimonRepository
	digimonSpeciesRepo digimonRepo.DigimonSpeciesRepository
	// userRepo           accountRepo.UserRepository // To validate player exists or has slots
}

func NewAcquireDigimonUsecase(
	dr digimonRepo.DigimonRepository,
	dsr digimonRepo.DigimonSpeciesRepository,
	// ur accountRepo.UserRepository,
) *AcquireDigimonUsecase {
	return &AcquireDigimonUsecase{
		digimonRepo:        dr,
		digimonSpeciesRepo: dsr,
		// userRepo:           ur,
	}
}

func (uc *AcquireDigimonUsecase) Execute(playerID string, req *dto.AcquireDigimonRequest) (*dto.DigimonResponse, error) {
	// 1. Fetch the DigimonSpecies
	species, err := uc.digimonSpeciesRepo.FindByID(req.SpeciesID) // Or FindByName if request uses name
	if err != nil {
		log.Printf("Error fetching species %s for player %s: %v", req.SpeciesID, playerID, err)
		return nil, ErrSpeciesNotFound
	}

	// 2. (Optional) Validate if player can acquire this Digimon
	//    e.g., check player level, items, or if they already have this species, max slots.
	//    For now, we'll skip complex validation like max slots.
	//    Example:
	//    playerDigimonCount, err := uc.digimonRepo.CountByPlayerID(playerID)
	//    if err != nil { return nil, err }
	//    playerMaxSlots := 10 // This could come from User entity or config
	//    if playerDigimonCount >= int64(playerMaxSlots) { return nil, ErrMaxDigimonSlotsReached }

	// 3. Create the Digimon instance
	newDigimonID := uuid.NewString()
	acquiredAt := time.Now() // Explicitly set AcquiredAt and UpdatedAt for clarity
	updatedAt := acquiredAt

	newDigimon := &digimonEntity.Digimon{
		ID:                newDigimonID,
		PlayerID:          playerID,
		SpeciesID:         species.ID,
		Nickname:          req.Nickname, // Use provided nickname, or default if empty
		CurrentLevel:      1,            // Starting level
		CurrentAttack:     species.BaseAttack, // Start with base stats
		CurrentDefense:    species.BaseDefense,
		CurrentSpeed:      species.BaseSpeed,
		ExperiencePoints:  0,
		AcquiredAt:        acquiredAt, // Set explicitly
		UpdatedAt:         updatedAt,  // Set explicitly
	}
	if newDigimon.Nickname == "" {
		newDigimon.Nickname = species.Name // Default nickname to species name
	}

	err = uc.digimonRepo.Create(newDigimon)
	if err != nil {
		log.Printf("Error creating digimon for player %s: %v", playerID, err)
		return nil, errors.New("failed to acquire digimon")
	}

	// 4. Prepare and return the response DTO
	return &dto.DigimonResponse{
		ID:                newDigimon.ID,
		PlayerID:          newDigimon.PlayerID,
		Nickname:          newDigimon.Nickname,
		CurrentLevel:      newDigimon.CurrentLevel,
		CurrentAttack:     newDigimon.CurrentAttack,
		CurrentDefense:    newDigimon.CurrentDefense,
		CurrentSpeed:      newDigimon.CurrentSpeed,
		ExperiencePoints:  newDigimon.ExperiencePoints,
		AcquiredAt:        newDigimon.AcquiredAt.Format(time.RFC3339),
		Species: dto.DigimonSpeciesResponse{ // Populate embedded species details
			ID:          species.ID,
			Name:        species.Name,
			Attribute:   species.Attribute,
			Stage:       species.Stage,
			BaseAttack:  species.BaseAttack,
			BaseDefense: species.BaseDefense,
			BaseSpeed:   species.BaseSpeed,
			SpriteURL:   species.SpriteURL,
			Description: species.Description,
		},
	}, nil
}

// ListPlayerDigimonUsecase handles listing a player's Digimon.
type ListPlayerDigimonUsecase struct {
	digimonRepo        digimonRepo.DigimonRepository
	digimonSpeciesRepo digimonRepo.DigimonSpeciesRepository // To enrich DigimonResponse with species details
}

func NewListPlayerDigimonUsecase(
	dr digimonRepo.DigimonRepository,
	dsr digimonRepo.DigimonSpeciesRepository,
) *ListPlayerDigimonUsecase {
	return &ListPlayerDigimonUsecase{
		digimonRepo:        dr,
		digimonSpeciesRepo: dsr,
	}
}

func (uc *ListPlayerDigimonUsecase) Execute(playerID string, offset int, limit int) ([]*dto.DigimonResponse, int64, error) {
	digimonList, err := uc.digimonRepo.ListByPlayerID(playerID, offset, limit)
	if err != nil {
		log.Printf("Error listing player %s digimon: %v", playerID, err)
		return nil, 0, errors.New("failed to list player digimon")
	}

	var totalCount int64 // Assuming repository or another method provides total count for pagination
	// For simplicity, if digimonRepo doesn't have CountByPlayerID, we can estimate or omit totalCount for now.
	// If CountByPlayerID is implemented in repo:
	// totalCount, err = uc.digimonRepo.CountByPlayerID(playerID)
	// if err != nil {
	//     log.Printf("Error counting player %s digimon: %v", playerID, err)
	//     return nil, 0, errors.New("failed to count player digimon")
	// }

	responseDTOs := make([]*dto.DigimonResponse, 0, len(digimonList))
	for _, digimon := range digimonList {
		species, err := uc.digimonSpeciesRepo.FindByID(digimon.SpeciesID)
		if err != nil {
			log.Printf("Error fetching species %s for digimon %s (player %s): %v", digimon.SpeciesID, digimon.ID, playerID, err)
			// Decide how to handle: skip this digimon, return partial list, or error out?
			// For now, let's skip if species not found, but this indicates data integrity issue.
			continue
		}

		responseDTOs = append(responseDTOs, &dto.DigimonResponse{
			ID:                digimon.ID,
			PlayerID:          digimon.PlayerID,
			Nickname:          digimon.Nickname,
			CurrentLevel:      digimon.CurrentLevel,
			CurrentAttack:     digimon.CurrentAttack,
			CurrentDefense:    digimon.CurrentDefense,
			CurrentSpeed:      digimon.CurrentSpeed,
			ExperiencePoints:  digimon.ExperiencePoints,
			AcquiredAt:        digimon.AcquiredAt.Format(time.RFC3339),
			Species: dto.DigimonSpeciesResponse{
				ID:          species.ID,
				Name:        species.Name,
				Attribute:   species.Attribute,
				Stage:       species.Stage,
				BaseAttack:  species.BaseAttack,
				BaseDefense: species.BaseDefense,
				BaseSpeed:   species.BaseSpeed,
				SpriteURL:   species.SpriteURL,
				Description: species.Description,
			},
		})
	}
	// If totalCount is not available from a direct repo call, it might be len(responseDTOs) if not paginating,
	// or requires a separate COUNT query. For now, returning 0 if not implemented.
	// totalCount = int64(len(responseDTOs)) // This is incorrect if paginating beyond what's fetched.

	return responseDTOs, totalCount, nil
}

// GetDigimonDetailsUsecase handles fetching details of a specific Digimon.
type GetDigimonDetailsUsecase struct {
	digimonRepo        digimonRepo.DigimonRepository
	digimonSpeciesRepo digimonRepo.DigimonSpeciesRepository
}

func NewGetDigimonDetailsUsecase(
	dr digimonRepo.DigimonRepository,
	dsr digimonRepo.DigimonSpeciesRepository,
) *GetDigimonDetailsUsecase {
	return &GetDigimonDetailsUsecase{
		digimonRepo:        dr,
		digimonSpeciesRepo: dsr,
	}
}

func (uc *GetDigimonDetailsUsecase) Execute(digimonID string, playerID string) (*dto.DigimonResponse, error) {
	digimon, err := uc.digimonRepo.FindByID(digimonID)
	if err != nil {
		log.Printf("Error fetching digimon %s (requested by player %s): %v", digimonID, playerID, err)
		return nil, ErrDigimonNotFound
	}

	// Important: Check if the digimon belongs to the requesting player
	if digimon.PlayerID != playerID {
		log.Printf("Player %s attempted to access digimon %s owned by %s", playerID, digimonID, digimon.PlayerID)
		return nil, ErrDigimonNotOwned
	}

	species, err := uc.digimonSpeciesRepo.FindByID(digimon.SpeciesID)
	if err != nil {
		log.Printf("Error fetching species %s for digimon %s (player %s): %v", digimon.SpeciesID, digimon.ID, playerID, err)
		return nil, ErrSpeciesNotFound // Or a more generic "failed to get digimon details"
	}

	return &dto.DigimonResponse{
		ID:                digimon.ID,
		PlayerID:          digimon.PlayerID,
		Nickname:          digimon.Nickname,
		CurrentLevel:      digimon.CurrentLevel,
		CurrentAttack:     digimon.CurrentAttack,
		CurrentDefense:    digimon.CurrentDefense,
		CurrentSpeed:      digimon.CurrentSpeed,
		ExperiencePoints:  digimon.ExperiencePoints,
		AcquiredAt:        digimon.AcquiredAt.Format(time.RFC3339),
		Species: dto.DigimonSpeciesResponse{
			ID:          species.ID,
			Name:        species.Name,
			Attribute:   species.Attribute,
			Stage:       species.Stage,
			BaseAttack:  species.BaseAttack,
			BaseDefense: species.BaseDefense,
			BaseSpeed:   species.BaseSpeed,
			SpriteURL:   species.SpriteURL,
			Description: species.Description,
		},
	}, nil
}

// UpdateDigimonNicknameUsecase handles updating a Digimon's nickname.
type UpdateDigimonNicknameUsecase struct {
	digimonRepo        digimonRepo.DigimonRepository
	digimonSpeciesRepo digimonRepo.DigimonSpeciesRepository // Added dependency
}

func NewUpdateDigimonNicknameUsecase(
	dr digimonRepo.DigimonRepository,
	dsr digimonRepo.DigimonSpeciesRepository, // Added parameter
) *UpdateDigimonNicknameUsecase {
	return &UpdateDigimonNicknameUsecase{
		digimonRepo:        dr,
		digimonSpeciesRepo: dsr, // Assign dependency
	}
}

func (uc *UpdateDigimonNicknameUsecase) Execute(playerID string, digimonID string, newNickname string) (*dto.DigimonResponse, error) {
	digimon, err := uc.digimonRepo.FindByID(digimonID)
	if err != nil {
		log.Printf("Error fetching digimon %s for nickname update (player %s): %v", digimonID, playerID, err)
		return nil, ErrDigimonNotFound
	}

	if digimon.PlayerID != playerID {
		log.Printf("Player %s attempted to update nickname for digimon %s owned by %s", playerID, digimonID, digimon.PlayerID)
		return nil, ErrDigimonNotOwned
	}

	digimon.Nickname = newNickname
	// digimon.UpdatedAt = time.Now() // GORM's autoUpdateTime should handle this if correctly tagged in entity

	err = uc.digimonRepo.Update(digimon)
	if err != nil {
		log.Printf("Error updating digimon %s nickname for player %s: %v", digimonID, playerID, err)
		return nil, errors.New("failed to update nickname")
	}

	species, err := uc.digimonSpeciesRepo.FindByID(digimon.SpeciesID)
	if err != nil {
		log.Printf("Error fetching species %s for digimon %s after nickname update (player %s): %v", digimon.SpeciesID, digimon.ID, playerID, err)
		return nil, errors.New("failed to fetch species details for response after update")
	}

	return &dto.DigimonResponse{
		ID:                digimon.ID,
		PlayerID:          digimon.PlayerID,
		Nickname:          digimon.Nickname,
		CurrentLevel:      digimon.CurrentLevel,
		CurrentAttack:     digimon.CurrentAttack,
		CurrentDefense:    digimon.CurrentDefense,
		CurrentSpeed:      digimon.CurrentSpeed,
		ExperiencePoints:  digimon.ExperiencePoints,
		AcquiredAt:        digimon.AcquiredAt.Format(time.RFC3339),
		Species: dto.DigimonSpeciesResponse{
			ID:          species.ID,
			Name:        species.Name,
			Attribute:   species.Attribute,
			Stage:       species.Stage,
			BaseAttack:  species.BaseAttack,
			BaseDefense: species.BaseDefense,
			BaseSpeed:   species.BaseSpeed,
			SpriteURL:   species.SpriteURL,
			Description: species.Description,
		},
	}, nil
}
