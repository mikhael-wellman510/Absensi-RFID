package usecases

import (
	"attendance-api/internal/adapters/repositories"
	"attendance-api/internal/entities"
	"context"
	"errors"
)

type (
	StudentParentService interface {
		CreateStudentParent(ctx context.Context, studentParentReq *entities.StudentParentRequest) (*entities.StudentParentResponse, error)
		FindStudentParentById(ctx context.Context, id string) (*entities.StudentParentResponse, error)
	}

	studentParentService struct {
		studentParentRepository repositories.StudentParentRepository
		studentService          StudentService
		parentService           ParentService
	}
)

func NewStudentParentService(studentParentRepository repositories.StudentParentRepository, studentService StudentService, parentService ParentService) StudentParentService {

	return &studentParentService{
		studentParentRepository: studentParentRepository,
		studentService:          studentService,
		parentService:           parentService,
	}
}

func (s *studentParentService) CreateStudentParent(ctx context.Context, studentParentReq *entities.StudentParentRequest) (*entities.StudentParentResponse, error) {
	// find by id student
	student, err := s.studentService.FindById(ctx, studentParentReq.StudentId)
	if err != nil {
		return nil, err
	}
	// find by id parent
	parent, err := s.parentService.FindById(ctx, studentParentReq.ParentId)

	if err != nil {
		return nil, err
	}

	if !studentParentReq.RelationType.IsValid() {
		return nil, errors.New("invalid relation type")
	}

	studentParent := entities.StudentParent{
		StudentId:    student.ID,
		ParentId:     parent.ID,
		RelationType: studentParentReq.RelationType,
	}

	if err := s.studentParentRepository.Create(ctx, &studentParent); err != nil {
		return nil, err
	}

	return &entities.StudentParentResponse{
		Id: studentParent.ID,
		StudentResponse: entities.StudentResponse{
			Id:             student.ID,
			Nis:            student.Nis,
			Nisn:           student.Nisn,
			FullName:       student.FullName,
			Gender:         student.Gender,
			BirthDate:      student.BirthDate,
			Address:        student.Address,
			EnrollmentDate: student.EnrollmentDate,
			Status:         student.Status,
			SchoolResponse: entities.SchoolResponse{
				Id:          student.School.ID,
				Npsn:        student.School.Npsn,
				SchoolName:  student.School.SchoolName,
				Address:     student.School.Address,
				SchoolLevel: student.School.SchoolLevel,
			},
		},

		ParentResponse: entities.ParentResponse{
			Id: parent.ID,
			UserResponse: entities.UserResponse{
				Id:          parent.User.ID,
				FullName:    parent.User.FullName,
				Email:       parent.User.Email,
				PhoneNumber: parent.User.PhoneNumber,
				Role:        parent.User.Role,
				LastLogin:   parent.User.LastLogin,
				IsActive:    parent.User.IsActive,
				CreatedAt:   parent.User.CreatedAt,
				UpdatedAt:   parent.User.UpdatedAt,
			},
			Occupation: parent.Occupation,
			Address:    parent.Address,
		},
		RelationType: studentParent.RelationType,
		CreatedAt:    studentParent.CreatedAt,
		UpdatedAt:    studentParent.UpdatedAt,
	}, nil

}

func (s *studentParentService) FindStudentParentById(ctx context.Context, id string) (*entities.StudentParentResponse, error) {

	studentParent, err := s.studentParentRepository.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	return &entities.StudentParentResponse{
		Id: studentParent.ID,
		StudentResponse: entities.StudentResponse{
			Id:             studentParent.Student.ID,
			Nis:            studentParent.Student.Nis,
			Nisn:           studentParent.Student.Nisn,
			FullName:       studentParent.Student.FullName,
			Gender:         studentParent.Student.Gender,
			BirthDate:      studentParent.Student.BirthDate,
			Address:        studentParent.Student.Address,
			EnrollmentDate: studentParent.Student.EnrollmentDate,
			Status:         studentParent.Student.Status,
			SchoolResponse: entities.SchoolResponse{
				Id:         studentParent.Student.School.ID,
				SchoolName: studentParent.Student.School.SchoolName,
			},
			CreatedAt: studentParent.Student.CreatedAt,
			UpdatedAt: studentParent.Student.UpdatedAt,
		},
		ParentResponse: entities.ParentResponse{
			Id: studentParent.ParentId,
			UserResponse: entities.UserResponse{
				Id:          studentParent.Parent.User.ID,
				FullName:    studentParent.Parent.User.FullName,
				Email:       studentParent.Parent.User.Email,
				PhoneNumber: studentParent.Parent.User.PhoneNumber,
				Role:        studentParent.Parent.User.Role,
				LastLogin:   studentParent.Parent.User.LastLogin,
				IsActive:    studentParent.Parent.User.IsActive,
				CreatedAt:   studentParent.Parent.User.CreatedAt,
				UpdatedAt:   studentParent.Parent.User.UpdatedAt,
			},
			Occupation: studentParent.Parent.Occupation,
			Address:    studentParent.Parent.Address,
		},
		RelationType: studentParent.RelationType,
	}, nil
}
