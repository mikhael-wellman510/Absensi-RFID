package usecases

import (
	"attendance-api/internal/adapters/repositories"
	"attendance-api/internal/entities"
	"context"
)

type (
	SchoolService interface {
		CreateSchool(ctx context.Context, schoolReq *entities.SchoolRequest) (*entities.SchoolResponse, error)
		FindSchoolById(ctx context.Context, id string) (*entities.SchoolResponse, error)
		FindById(ctx context.Context, id string) (*entities.School, error)
	}

	schoolService struct {
		schoolRepository repositories.SchoolRepository
	}
)

func NewSchoolService(schoolRepository repositories.SchoolRepository) SchoolService {

	return &schoolService{
		schoolRepository: schoolRepository,
	}
}

func (s *schoolService) CreateSchool(ctx context.Context, schoolReq *entities.SchoolRequest) (*entities.SchoolResponse, error) {

	school := &entities.School{
		Npsn:        schoolReq.Npsn,
		SchoolName:  schoolReq.SchoolName,
		Address:     schoolReq.Address,
		SchoolLevel: schoolReq.SchoolLevel,
		Email:       schoolReq.Email,
		City:        schoolReq.City,
		Province:    schoolReq.Province,
		PhoneNumber: schoolReq.PhoneNumber,
		IsActive:    true,
	}

	if err := s.schoolRepository.Create(ctx, school); err != nil {
		return nil, err
	}

	return &entities.SchoolResponse{
		Id:          school.ID,
		Npsn:        school.Npsn,
		SchoolName:  school.SchoolName,
		Address:     school.Address,
		SchoolLevel: school.SchoolLevel,
		Email:       school.Email,
		City:        school.City,
		Province:    school.Province,
		PhoneNumber: school.PhoneNumber,
		IsActive:    school.IsActive,
		CreatedAt:   school.CreatedAt,
		UpdatedAt:   school.UpdatedAt,
	}, nil
}

func (s *schoolService) FindSchoolById(ctx context.Context, id string) (*entities.SchoolResponse, error) {

	res, err := s.schoolRepository.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	return &entities.SchoolResponse{
		Id:          res.ID,
		Npsn:        res.Npsn,
		SchoolName:  res.SchoolName,
		Address:     res.Address,
		SchoolLevel: res.SchoolLevel,
		Email:       res.Email,
		City:        res.City,
		Province:    res.Province,
		PhoneNumber: res.PhoneNumber,
		IsActive:    res.IsActive,
		CreatedAt:   res.CreatedAt,
		UpdatedAt:   res.UpdatedAt,
	}, nil
}

func (s *schoolService) FindById(ctx context.Context, id string) (*entities.School, error) {

	return s.schoolRepository.FindById(ctx, id)
}
