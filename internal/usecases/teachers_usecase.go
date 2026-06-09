package usecases

import (
	"attendance-api/internal/adapters/repositories"
	"attendance-api/internal/entities"
	"context"
)

type (
	TeachersService interface {
		CreateTeachers(ctx context.Context, teachersReq *entities.TeachersRequest) (*entities.TeachersResponse, error)
	}

	teachersService struct {
		TeacherRepository repositories.TeachersRepository
		SchoolService     SchoolService
	}
)

func NewTeachersService(teachersRepository repositories.TeachersRepository, schoolService SchoolService) TeachersService {
	return &teachersService{
		TeacherRepository: teachersRepository,
		SchoolService:     schoolService,
	}
}

func (t teachersService) CreateTeachers(ctx context.Context, teachersReq *entities.TeachersRequest) (*entities.TeachersResponse, error) {

	school, err := t.SchoolService.FindSchoolByID(ctx, teachersReq.SchoolId)

	if err != nil {
		return nil, err
	}

	teachers := &entities.Teachers{
		Nip:         teachersReq.Nip,
		FullName:    teachersReq.FullName,
		Email:       teachersReq.Email,
		Address:     teachersReq.Address,
		PhoneNumber: teachersReq.PhoneNumber,
		UserName:    teachersReq.UserName,
		Password:    teachersReq.Password,
		IsActive:    true,
		SchoolId:    school.ID,
	}

	if err := t.TeacherRepository.Create(ctx, teachers); err != nil {
		return nil, err
	}

	return &entities.TeachersResponse{
		Id:          teachers.ID,
		Nip:         teachers.Nip,
		FullName:    teachers.FullName,
		Email:       teachers.Email,
		Address:     teachers.Address,
		PhoneNumber: teachers.PhoneNumber,
		UserName:    teachers.UserName,
		Password:    teachers.Password,
		IsActive:    teachers.IsActive,
		SchoolResponse: entities.SchoolResponse{
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
		},
	}, nil
}
