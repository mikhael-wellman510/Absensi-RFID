package usecases

import (
	"attendance-api/internal/adapters/repositories"
	"attendance-api/internal/entities"
	"attendance-api/internal/utils"
	"context"
	"log"
)

type (
	AcademicYearService interface {
		CreateAcademicYear(ctx context.Context, academicYearReq *entities.AcademicYearRequest) (*entities.AcademicYearResponse, error)
		FindAcademicYearById(ctx context.Context, id string) (*entities.AcademicYearResponse, error)
		FindById(ctx context.Context, id string) (*entities.AcademicYear, error)
	}

	academicYearService struct {
		academicYearRepository repositories.AcademicYearRepository
		schoolService          SchoolService
	}
)

func NewAcademicYearService(academicYearRepository repositories.AcademicYearRepository, schoolService SchoolService) AcademicYearService {

	return &academicYearService{
		academicYearRepository: academicYearRepository,
		schoolService:          schoolService,
	}
}

func (a *academicYearService) CreateAcademicYear(ctx context.Context, academicYearReq *entities.AcademicYearRequest) (*entities.AcademicYearResponse, error) {
	school, err := a.schoolService.FindById(ctx, academicYearReq.SchoolId)

	if err != nil {
		return nil, err
	}

	startDate, err := utils.ParseDate(academicYearReq.StartDate)

	if err != nil {
		return nil, err
	}

	endDate, err := utils.ParseDate(academicYearReq.EndDate)

	if err != nil {
		return nil, err
	}

	academicYear := &entities.AcademicYear{
		SchoolId:  school.ID,
		YearName:  academicYearReq.YearName,
		StartDate: startDate,
		EndDate:   endDate,
		IsActive:  true,
	}

	if err := a.academicYearRepository.Create(ctx, academicYear); err != nil {
		return nil, err
	}

	return &entities.AcademicYearResponse{
		Id:        academicYear.ID,
		YearName:  academicYear.YearName,
		StartDate: academicYear.StartDate,
		EndDate:   academicYear.EndDate,
		IsActive:  academicYear.IsActive,
		CreatedAt: academicYear.CreatedAt,
		UpdatedAt: academicYear.UpdatedAt,
	}, nil
}

func (a *academicYearService) FindAcademicYearById(ctx context.Context, id string) (*entities.AcademicYearResponse, error) {

	res, err := a.academicYearRepository.FindById(ctx, id)

	if err != nil {
		log.Println("AcademicYear not found")
		return nil, err
	}

	return &entities.AcademicYearResponse{
		Id:        res.ID,
		YearName:  res.YearName,
		StartDate: res.StartDate,
		EndDate:   res.EndDate,
		IsActive:  res.IsActive,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (a *academicYearService) FindById(ctx context.Context, id string) (*entities.AcademicYear, error) {

	return a.academicYearRepository.FindById(ctx, id)
}
