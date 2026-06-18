package usecases

import (
	"attendance-api/internal/adapters/repositories"
	"attendance-api/internal/entities"
	"context"
	"log"
)

type (
	ParentService interface {
		CreateParent(ctx context.Context, parentReq *entities.ParentRequest) (*entities.ParentResponse, error)
		FindParentByID(ctx context.Context, id string) (*entities.ParentResponse, error)
	}

	parentService struct {
		parentRepository repositories.ParentRepository
		userService      UserService
	}
)

func NewParentService(parentRepository repositories.ParentRepository, userService UserService) ParentService {

	return &parentService{
		parentRepository: parentRepository,
		userService:      userService,
	}
}

func (p *parentService) CreateParent(ctx context.Context, parentReq *entities.ParentRequest) (*entities.ParentResponse, error) {
	log.Println("user id : ", parentReq.UserId)
	res, err := p.userService.FindById(ctx, parentReq.UserId)

	if err != nil {
		return nil, err
	}

	parent := &entities.Parent{
		UserId:     res.ID,
		Occupation: parentReq.Occupation,
		Address:    parentReq.Address,
	}

	if err := p.parentRepository.Create(ctx, parent); err != nil {
		return nil, err
	}

	return &entities.ParentResponse{
		Id: parent.ID,
		UserResponse: entities.UserResponse{
			Id:          res.ID,
			FullName:    res.FullName,
			Email:       res.Email,
			PhoneNumber: res.PhoneNumber,
			Password:    res.Password,
			Role:        res.Role,
			LastLogin:   res.LastLogin,
			IsActive:    res.IsActive,
			CreatedAt:   res.CreatedAt,
			UpdatedAt:   res.UpdatedAt,
		},
		Occupation: parentReq.Occupation,
		Address:    parentReq.Address,
	}, nil

}

func (p *parentService) FindParentByID(ctx context.Context, id string) (*entities.ParentResponse, error) {

	res, err := p.parentRepository.FindById(ctx, id)

	if err != nil {
		return nil, err
	}

	return &entities.ParentResponse{
		Id: res.ID,
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
		Occupation: res.Occupation,
		Address:    res.Occupation,
	}, nil
}
