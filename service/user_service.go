package service

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/constant"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/dto"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/models"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/repository"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserService interface {
		RegisterUser(ctx context.Context, req dto.RegisterUserRequest) (dto.RegisterUserResponse, error)
		Verify(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error)
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

func (s *userService) Verify(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	var user models.User

	// Step 1: Check if email exists in the database
	user, exists, err := s.userRepo.CheckEmail(ctx, nil, req.Email)
	if err != nil {
		return dto.UserLoginResponse{}, dto.ErrInternalServer
	}
	if !exists {
		return dto.UserLoginResponse{}, dto.ErrInvalidCredentials
	}

	// Step 2: Ensure the account is verified
	if !user.IsVerified {
		return dto.UserLoginResponse{}, dto.ErrAccountNotVerified
	}

	// Step 3: Validate the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return dto.UserLoginResponse{}, dto.ErrInvalidCredentials
	}

	// Step 4: Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user.ID.String(),
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	// Step 5: Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return dto.UserLoginResponse{}, dto.ErrFailedCreateToken
	}

	// Step 6: Return the response
	return dto.UserLoginResponse{
		Token: tokenString,
		Role:  user.Role,
	}, nil
}
