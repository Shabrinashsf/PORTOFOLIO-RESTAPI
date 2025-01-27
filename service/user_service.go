package service

import (
	"context"
	"sync"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/constant"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/dto"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/models"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/repository"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserService interface {
		RegisterUser(ctx context.Context, req dto.RegisterUserRequest) (dto.RegisterUserResponse, error)
	}

	userService struct {
		userRepo repository.UserRepository
	}
)

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

var (
	mu sync.Mutex
)

func (s *userService) RegisterUser(ctx context.Context, req dto.RegisterUserRequest) (dto.RegisterUserResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	_, flag, _ := s.userRepo.CheckEmail(ctx, nil, req.Email)
	if flag {
		return dto.RegisterUserResponse{}, dto.ErrEmailAlreadyExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return dto.RegisterUserResponse{}, dto.ErrHashPass
	}

	user := models.User{
		Name:       req.Name,
		Email:      req.Email,
		Password:   string(hash),
		NoTelp:     req.NoTelp,
		Role:       constant.ROLE_USER,
		IsVerified: false,
	}

	userReg, err := s.userRepo.RegisterUser(ctx, nil, user)
	if err != nil {
		return dto.RegisterUserResponse{}, dto.ErrCreateUser
	}

	return dto.RegisterUserResponse{
		Name:   userReg.Name,
		Email:  userReg.Email,
		NoTelp: userReg.NoTelp,
		Role:   userReg.Role,
	}, nil

}
