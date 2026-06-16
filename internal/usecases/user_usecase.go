package usecases

import (
	"attendance-api/internal/adapters/repositories"
	"attendance-api/internal/entities"
	"context"
	"errors"
	"time"
)

type (
	UserService interface {
		CreateUser(ctx context.Context, userReq *entities.UserRequest) (*entities.UserResponse, error)
		FindUserById(ctx context.Context, id string) (*entities.UserResponse, error)
	}

	userService struct {
		userRepository repositories.UserRepository
	}
)

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) CreateUser(ctx context.Context, userReq *entities.UserRequest) (*entities.UserResponse, error) {

	if !userReq.Role.IsValid() {
		return nil, errors.New("invalid role")
	}

	user := &entities.User{
		FullName:    userReq.FullName,
		Email:       userReq.Email,
		PhoneNumber: userReq.PhoneNumber,
		Password:    userReq.Password,
		Role:        userReq.Role,
		LastLogin:   time.Now(),
		IsActive:    true,
	}

	if err := u.userRepository.Create(ctx, user); err != nil {
		return nil, err
	}

	return &entities.UserResponse{
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
	}, nil
}

func (u *userService) FindUserById(ctx context.Context, id string) (*entities.UserResponse, error) {

	res, err := u.userRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entities.UserResponse{
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
	}, nil
}
