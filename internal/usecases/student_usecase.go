package usecases

import (
	"attendance-api/internal/adapters/repositories"
	"attendance-api/internal/entities"
	"attendance-api/internal/utils"
	"context"
	"errors"
	"log"
)

type (
	StudentService interface {
		CreateStudent(ctx context.Context, studentReq *entities.StudentRequest) (*entities.StudentResponse, error)
		FindStudentById(ctx context.Context, id string) (*entities.StudentResponse, error)
	}

	studentService struct {
		studentRepository repositories.StudentRepository
		schoolService     SchoolService
	}
)

func NewStudentService(studentRepository repositories.StudentRepository, schoolService SchoolService) StudentService {

	return &studentService{
		studentRepository: studentRepository,
		schoolService:     schoolService,
	}
}

func (s *studentService) CreateStudent(ctx context.Context, studentReq *entities.StudentRequest) (*entities.StudentResponse, error) {

	school, err := s.schoolService.FindById(ctx, studentReq.SchoolId)

	if err != nil {
		log.Println("School not found")
		return nil, err
	}

	if !studentReq.Gender.IsValid() {

		return nil, errors.New("invalid gender")
	}

	if !studentReq.Status.IsValid() {

		return nil, errors.New("invalid status")
	}

	birthDate, err := utils.ParseDate(studentReq.BirthDate)

	if err != nil {
		return nil, err
	}

	enrollmentDate, err := utils.ParseDate(studentReq.EnrollmentDate)

	if err != nil {
		return nil, err
	}
	student := entities.Student{
		SchoolId:       school.ID,
		Nis:            studentReq.Nis,
		Nisn:           studentReq.Nisn,
		FullName:       studentReq.FullName,
		Gender:         studentReq.Gender,
		BirthDate:      birthDate,
		Address:        studentReq.Address,
		EnrollmentDate: enrollmentDate,
		Status:         studentReq.Status,
	}

	if err := s.studentRepository.Create(ctx, &student); err != nil {
		return nil, err
	}

	return &entities.StudentResponse{
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
		},
		CreatedAt: student.CreatedAt,
		UpdatedAt: student.UpdatedAt,
	}, nil

}

func (s *studentService) FindStudentById(ctx context.Context, id string) (*entities.StudentResponse, error) {

	res, err := s.studentRepository.FindById(ctx, id)

	if err != nil {
		log.Println("Student not found")
		return nil, err
	}

	return &entities.StudentResponse{
		Id:             res.ID,
		Nis:            res.Nis,
		Nisn:           res.Nisn,
		FullName:       res.FullName,
		Gender:         res.Gender,
		BirthDate:      res.BirthDate,
		Address:        res.Address,
		EnrollmentDate: res.EnrollmentDate,
		Status:         res.Status,
		SchoolResponse: entities.SchoolResponse{
			Id:          res.School.ID,
			Npsn:        res.School.Npsn,
			SchoolName:  res.School.SchoolName,
			Address:     res.School.Address,
			SchoolLevel: res.School.SchoolLevel,
			Email:       res.School.Email,
			City:        res.School.City,
			Province:    res.School.Province,
			PhoneNumber: res.School.PhoneNumber,
			IsActive:    res.School.IsActive,
			CreatedAt:   res.School.CreatedAt,
			UpdatedAt:   res.School.UpdatedAt,
		},
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil

}
