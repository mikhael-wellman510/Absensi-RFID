package usecases

import (
	"attendance-api/internal/adapters/repositories"
	"attendance-api/internal/entities"
	"context"
	"errors"
)

type (
	GradeService interface {
		CreateGrade(ctx context.Context, gradeReq *entities.GradeRequest) (*entities.GradeResponse, error)
		FindGradeById(ctx context.Context, id string) (*entities.GradeResponse, error)
	}

	gradeService struct {
		gradeRepository repositories.GradeRepository
	}
)

func NewGradeService(gradeRepository repositories.GradeRepository) GradeService {
	return &gradeService{
		gradeRepository: gradeRepository,
	}
}

func (g *gradeService) CreateGrade(ctx context.Context, gradeReq *entities.GradeRequest) (*entities.GradeResponse, error) {

	if !gradeReq.SchoolLevel.IsValid() {
		return nil, errors.New("invalid school level")
	}

	grade := &entities.Grade{
		SchoolLevel: gradeReq.SchoolLevel,
		GradeNumber: gradeReq.GradeNumber,
	}

	if err := g.gradeRepository.Create(ctx, grade); err != nil {
		return nil, err
	}

	return &entities.GradeResponse{
		Id:          grade.ID,
		SchoolLevel: grade.SchoolLevel,
		GradeNumber: grade.GradeNumber,
		CreatedAt:   grade.CreatedAt,
		UpdatedAt:   grade.UpdatedAt,
	}, nil
}

func (g *gradeService) FindGradeById(ctx context.Context, id string) (*entities.GradeResponse, error) {

	res, err := g.gradeRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entities.GradeResponse{
		Id:          res.ID,
		SchoolLevel: res.SchoolLevel,
		GradeNumber: res.GradeNumber,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}, nil
}
