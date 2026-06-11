package repositories

import (
	"attendance-api/internal/entities"
	"context"

	"gorm.io/gorm"
)

type (
	StudyProgramRepository interface {
		Create(ctx context.Context, studyProgramRepository *entities.StudyProgram) error
		FindById(ctx context.Context, id string) (*entities.StudyProgram, error)
	}

	studyProgramRepository struct {
		db *gorm.DB
	}
)

func NewStudyProgramRepository(db *gorm.DB) StudyProgramRepository {

	return &studyProgramRepository{
		db: db,
	}
}

func (s *studyProgramRepository) Create(ctx context.Context, studyProgramRepository *entities.StudyProgram) error {

	return s.db.WithContext(ctx).Create(studyProgramRepository).Error
}

func (s *studyProgramRepository) FindById(ctx context.Context, id string) (*entities.StudyProgram, error) {
	studyProgram := &entities.StudyProgram{}

	err := s.db.WithContext(ctx).Where("id = ?", id).First(studyProgram).Error

	return studyProgram, err
}
