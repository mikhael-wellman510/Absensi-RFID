package repositories

import (
	"attendance-api/internal/entities"
	"context"

	"gorm.io/gorm"
)

type (
	ParentRepository interface {
		Create(ctx context.Context, parent *entities.Parent) error
		FindById(ctx context.Context, id string) (*entities.Parent, error)
	}

	parentRepository struct {
		db *gorm.DB
	}
)

func NewParentRepository(db *gorm.DB) ParentRepository {
	return &parentRepository{
		db: db,
	}
}

func (p *parentRepository) Create(ctx context.Context, parent *entities.Parent) error {

	return p.db.WithContext(ctx).Create(parent).Error
}

func (p *parentRepository) FindById(ctx context.Context, id string) (*entities.Parent, error) {

	parent := &entities.Parent{}

	err := p.db.WithContext(ctx).Where("id = ?", id).First(parent).Error

	return parent, err
}
