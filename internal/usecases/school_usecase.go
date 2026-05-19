package usecases

import (
	"attendance-api/internal/adapters/repositories"
	"attendance-api/internal/entities"
)

type (
	SchoolService interface {
		CreateSchool(schoolReq *entities.SchoolRequest) (*entities.SchoolResponse, error)
	}

	schoolService struct {
		SchoolRepository repositories.SchoolRepository
	}
)

func NewSchoolService(schoolRepository repositories.SchoolRepository) SchoolService {

	return &schoolService{
		SchoolRepository: schoolRepository,
	}
}

func (ss *schoolService) CreateSchool(schoolReq *entities.SchoolRequest) (*entities.SchoolResponse, error) {

	school := &entities.School{
		SchoolName:  schoolReq.SchoolName,
		Address:     schoolReq.Address,
		SchoolLevel: schoolReq.SchoolLevel,
		PhoneNumber: schoolReq.PhoneNumber,
		Email:       schoolReq.Email,
		City:        schoolReq.City,
		Province:    schoolReq.Province,
		IsActive:    true,
	}

	if err := ss.SchoolRepository.Create(school); err != nil {

		return nil, err
	}

	return &entities.SchoolResponse{
		Id:          school.ID,
		SchoolName:  school.SchoolName,
		Address:     school.Address,
		SchoolLevel: school.SchoolLevel,
		PhoneNumber: school.PhoneNumber,
		Email:       school.Email,
		City:        school.City,
		Province:    school.Province,
		IsActive:    school.IsActive,
		CreatedAt:   school.CreatedAt,
		UpdatedAt:   school.UpdatedAt,
	}, nil
}
