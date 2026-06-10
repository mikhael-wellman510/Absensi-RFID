package usecases

import (
	"attendance-api/internal/adapters/repositories"
	"attendance-api/internal/entities"
	"context"
	"log"
)

type (
	EducationLevelsService interface {
		CreateEducationLevels(ctx context.Context, educationLevelReq *entities.EducationLevelsRequest) (*entities.EducationLevelsResponse, error)
		FindEducationLevelByID(ctx context.Context, id string) (*entities.EducationLevel, error)
	}

	educationLevelsService struct {
		educationLevelsRepository repositories.EducationLevelsRepository
	}
)

func NewEducationLevelsService(educationLevelsRepository repositories.EducationLevelsRepository) EducationLevelsService {

	return &educationLevelsService{
		educationLevelsRepository: educationLevelsRepository,
	}
}

func (e *educationLevelsService) CreateEducationLevels(ctx context.Context, educationLevelsReq *entities.EducationLevelsRequest) (*entities.EducationLevelsResponse, error) {
	educationLevels := &entities.EducationLevel{
		Name: educationLevelsReq.Name,
		Code: educationLevelsReq.Code,
	}

	if err := e.educationLevelsRepository.Create(ctx, educationLevels); err != nil {
		return nil, err
	}

	return &entities.EducationLevelsResponse{
		Id:        educationLevels.ID,
		Name:      educationLevels.Name,
		Code:      educationLevels.Code,
		CreatedAt: educationLevels.CreatedAt,
		UpdatedAt: educationLevels.UpdatedAt,
	}, nil

}

func (e *educationLevelsService) FindEducationLevelByID(ctx context.Context, id string) (*entities.EducationLevel, error) {

	log.Println("id : ", id)
	res, err := e.educationLevelsRepository.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	return res, nil
}
