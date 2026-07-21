package usecases

import (
	"attendance-api/internal/adapters/repositories"
	"attendance-api/internal/entities"
	"attendance-api/internal/utils"
	"context"
)

type (
	ClassRoomService interface {
		CreateClassRoom(ctx context.Context, classRoomReq *entities.ClassRoomRequest) (*entities.ClassRoomResponse, error)
		FindClassRoomById(ctx context.Context, id string) (*entities.ClassRoomResponse, error)
		FindById(ctx context.Context, id string) (*entities.ClassRoom, error)
	}

	classRoomService struct {
		classRoomRepository repositories.ClassRoomRepository
		schoolService       SchoolService
		academicYearService AcademicYearService
		gradeService        GradeService
		teacherService      TeacherService
	}
)

func NewClassRoomService(classRoomRepository repositories.ClassRoomRepository, schoolService SchoolService, academicYearService AcademicYearService, gradeService GradeService, teacherService TeacherService) ClassRoomService {

	return &classRoomService{
		classRoomRepository: classRoomRepository,
		schoolService:       schoolService,
		academicYearService: academicYearService,
		gradeService:        gradeService,
		teacherService:      teacherService,
	}
}

func (c *classRoomService) CreateClassRoom(ctx context.Context, classRoomReq *entities.ClassRoomRequest) (*entities.ClassRoomResponse, error) {

	// todo -> school , academicYear, grade , teacher

	school, err := c.schoolService.FindById(ctx, classRoomReq.SchoolId)

	if err != nil {
		return nil, err
	}

	academicYear, err := c.academicYearService.FindById(ctx, classRoomReq.AcademicYearsId)

	if err != nil {
		return nil, err
	}

	grade, err := c.gradeService.FindById(ctx, classRoomReq.GradeId)

	if err != nil {
		return nil, err
	}

	teacher, err := c.teacherService.FindById(ctx, classRoomReq.TeacherId)

	if err != nil {
		return nil, err
	}

	classRoom := entities.ClassRoom{
		SchoolId:        school.ID,
		AcademicYearsId: academicYear.ID,
		GradeId:         grade.ID,
		TeacherId:       teacher.ID,
		ClassName:       classRoomReq.ClassName,
	}

	if err := c.classRoomRepository.Create(ctx, &classRoom); err != nil {
		return nil, err
	}

	return &entities.ClassRoomResponse{
		Id: classRoom.ID,
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
		AcademicYearsResponse: entities.AcademicYearResponse{
			Id:        academicYear.ID,
			YearName:  academicYear.YearName,
			StartDate: academicYear.StartDate,
			EndDate:   academicYear.EndDate,
			IsActive:  academicYear.IsActive,
			CreatedAt: academicYear.CreatedAt,
			UpdatedAt: academicYear.UpdatedAt,
		},
		GradeResponse: entities.GradeResponse{
			Id:          grade.ID,
			SchoolLevel: grade.SchoolLevel,
			GradeNumber: grade.GradeNumber,
			CreatedAt:   grade.CreatedAt,
			UpdatedAt:   grade.UpdatedAt,
		},
		TeacherResponse: entities.TeacherResponse{
			Id:        teacher.ID,
			Nip:       teacher.Nip,
			Gender:    teacher.Gender,
			BirthDate: utils.FormatDate(teacher.BirthDate),
			Address:   teacher.Address,
			IsActive:  teacher.IsActive,
		},
		ClassName: classRoomReq.ClassName,
	}, nil

}

func (c *classRoomService) FindClassRoomById(ctx context.Context, id string) (*entities.ClassRoomResponse, error) {

	res, err := c.classRoomRepository.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	return &entities.ClassRoomResponse{
		Id: res.ID,
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
		},
		AcademicYearsResponse: entities.AcademicYearResponse{
			Id:        res.AcademicYears.ID,
			YearName:  res.AcademicYears.YearName,
			StartDate: res.AcademicYears.StartDate,
			EndDate:   res.AcademicYears.EndDate,
			IsActive:  res.AcademicYears.IsActive,
		},
		GradeResponse: entities.GradeResponse{
			Id:          res.Grade.ID,
			SchoolLevel: res.Grade.SchoolLevel,
			GradeNumber: res.Grade.GradeNumber,
		},
		TeacherResponse: entities.TeacherResponse{
			Id:        res.Teacher.ID,
			Nip:       res.Teacher.Nip,
			Gender:    res.Teacher.Gender,
			BirthDate: utils.FormatDate(res.Teacher.BirthDate),
			Address:   res.Teacher.Address,
			IsActive:  res.Teacher.IsActive,
		},
		ClassName: res.ClassName,
	}, nil
}

func (c *classRoomService) FindById(ctx context.Context, id string) (*entities.ClassRoom, error) {

	return c.classRoomRepository.FindById(ctx, id)
}
