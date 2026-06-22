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
	TeacherService interface {
		CreateTeacher(ctx context.Context, teacherReq *entities.TeacherRequest) (*entities.TeacherResponse, error)
		FindTeacherById(ctx context.Context, id string) (*entities.TeacherResponse, error)
	}

	teacherService struct {
		teacherRepository repositories.TeacherRepository
		userService       UserService
		schoolService     SchoolService
	}
)

func NewTeacherService(teacherRepository repositories.TeacherRepository, userService UserService, schoolService SchoolService) TeacherService {

	return &teacherService{
		teacherRepository: teacherRepository,
		userService:       userService,
		schoolService:     schoolService,
	}
}

func (t *teacherService) CreateTeacher(ctx context.Context, teacherReq *entities.TeacherRequest) (*entities.TeacherResponse, error) {

	user, errUser := t.userService.FindById(ctx, teacherReq.UserId)

	if errUser != nil {
		log.Println("user service not found")
		return nil, errors.New("user is not found")
	}
	school, errSchool := t.schoolService.FindById(ctx, teacherReq.SchoolId)

	if errSchool != nil {
		log.Println("school service not found")
		return nil, errSchool
	}

	if !teacherReq.Gender.IsValid() {
		return nil, errors.New("invalid gender")
	}

	birthDate, err := utils.ParseDate(teacherReq.BirthDate)

	if err != nil {
		return nil, err
	}

	teacher := &entities.Teacher{
		UserId:    teacherReq.UserId,
		SchoolId:  teacherReq.SchoolId,
		Nip:       teacherReq.Nip,
		Gender:    teacherReq.Gender,
		BirthDate: birthDate,
		Address:   teacherReq.Address,
		IsActive:  true,
	}

	if err := t.teacherRepository.Create(ctx, teacher); err != nil {
		return nil, err
	}

	return &entities.TeacherResponse{
		Id:        teacher.ID,
		Nip:       teacher.Nip,
		Gender:    teacher.Gender,
		BirthDate: utils.FormatDate(birthDate),
		Address:   teacher.Address,
		IsActive:  teacher.IsActive,
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
		UserResponse: entities.UserResponse{
			Id:          user.ID,
			FullName:    user.FullName,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Password:    user.Password,
			Role:        user.Role,
			LastLogin:   user.LastLogin,
			IsActive:    user.IsActive,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
		},
		CreatedAt: teacher.CreatedAt,
		UpdatedAt: teacher.UpdatedAt,
	}, nil

}

func (t *teacherService) FindTeacherById(ctx context.Context, id string) (*entities.TeacherResponse, error) {

	res, err := t.teacherRepository.FindById(ctx, id)

	if err != nil {
		log.Println("teacher is not found")
		return nil, errors.New("teacher is not found")
	}

	return &entities.TeacherResponse{
		Id:        res.ID,
		Nip:       res.Nip,
		Gender:    res.Gender,
		BirthDate: utils.FormatDate(res.BirthDate),
		Address:   res.Address,
		IsActive:  res.IsActive,
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
			IsActive:    res.IsActive,
			CreatedAt:   res.CreatedAt,
			UpdatedAt:   res.UpdatedAt,
		},
		UserResponse: entities.UserResponse{
			Id:          res.User.ID,
			FullName:    res.User.FullName,
			Email:       res.User.Email,
			PhoneNumber: res.User.PhoneNumber,
			Password:    res.User.Password,
			Role:        res.User.Role,
			LastLogin:   res.User.LastLogin,
			IsActive:    res.User.IsActive,
			CreatedAt:   res.User.CreatedAt,
			UpdatedAt:   res.User.UpdatedAt,
		},
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil

}
