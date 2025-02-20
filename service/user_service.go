package service

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/constant"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/dto"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/entity"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/middleware"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/repository"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/thanhpk/randstr"
)

type (
	UserService interface {
		RegisterUser(ctx context.Context, req dto.RegisterUserRequest) (dto.RegisterUserResponse, error)
		Verify(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error)
		GetAllUser(ctx context.Context) ([]dto.GetAllUser, error)
		VerifyEmail(ctx context.Context, code string) (dto.VerifyEmail, error)
		GetUserByID(ctx context.Context, id string) (dto.GetUserByID, error)
		UpdateUser(ctx *gin.Context, idParam string, req dto.UpdateUser) (dto.UpdateUser, error)
	}

	userService struct {
		userRepo repository.UserRepository
	}
)

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
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

	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		return dto.RegisterUserResponse{}, dto.ErrHashPass
	}

	// generate verification code
	code := randstr.String(20)
	verification_code := utils.Encode(code)

	now := time.Now()
	user := entity.User{
		Name:             req.Name,
		Email:            req.Email,
		Password:         hash,
		NoTelp:           req.NoTelp,
		Role:             constant.ROLE_USER,
		IsVerified:       false,
		VerificationCode: verification_code,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	userReg, err := s.userRepo.RegisterUser(ctx, nil, user)
	if err != nil {
		return dto.RegisterUserResponse{}, dto.ErrCreateUser
	}

	// send email
	ClientOrigin := os.Getenv("CLIENT_ORIGIN")
	name := user.Name

	emailData := utils.EmailData{
		URL:     ClientOrigin + "/verifyemail/" + code,
		Name:    name,
		Subject: "Your Account Verification Code To SHSF Server",
	}

	go utils.SendEmail(&userReg, &emailData)

	return dto.RegisterUserResponse{
		Name:   userReg.Name,
		Email:  userReg.Email,
		NoTelp: userReg.NoTelp,
		Role:   userReg.Role,
	}, nil
}

func (s *userService) VerifyEmail(ctx context.Context, code string) (dto.VerifyEmail, error) {
	verification_code := utils.Encode(code)

	user, err := s.userRepo.VerifyEmail(verification_code)
	if err != nil {
		return dto.VerifyEmail{}, dto.ErrInvalidVerificationCode
	}

	if user.IsVerified {
		return dto.VerifyEmail{}, dto.ErrUserAlreadyVerified
	}

	user.VerificationCode = ""
	user.IsVerified = true

	if err := s.userRepo.UpdateIsVerified(user); err != nil {
		return dto.VerifyEmail{}, dto.ErrUpdateIsVerified
	}

	return dto.VerifyEmail{
		Email:      user.Email,
		IsVerified: user.IsVerified,
	}, nil
}

func (s *userService) Verify(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	var user entity.User

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
	if err := utils.VerifyPassword(user.Password, req.Password); err != nil {
		return dto.UserLoginResponse{}, dto.ErrInvalidCredentials
	}

	// Step 4: Generate a JWT token
	privateKeyBytes, err := middleware.DecodePrivateKeyBase64()
	if err != nil {
		return dto.UserLoginResponse{}, dto.ErrFailedDecodePrivateKey
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return dto.UserLoginResponse{}, dto.ErrInvalidPrivateKeyFormat
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user": user.ID.String(),
		"role": user.Role,
		"exp":  time.Now().Add(15 * time.Minute).Unix(),
		"iat":  time.Now().Unix(),
	})

	// Step 5: Sign the token with the secret key
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return dto.UserLoginResponse{}, dto.ErrFailedCreateToken
	}

	// Step 6: Return the response
	return dto.UserLoginResponse{
		Token: tokenString,
		Role:  user.Role,
	}, nil
}

func (s *userService) GetAllUser(ctx context.Context) ([]dto.GetAllUser, error) {
	users, err := s.userRepo.GetAllUser(ctx)
	if err != nil {
		return []dto.GetAllUser{}, dto.ErrFailedGetUsers
	}

	var result []dto.GetAllUser
	for _, user := range users {
		result = append(result, dto.GetAllUser{
			Name:       user.Name,
			Email:      user.Email,
			NoTelp:     user.NoTelp,
			Role:       user.Role,
			IsVerified: user.IsVerified,
		})
	}

	return result, nil
}

func (s *userService) GetUserByID(ctx context.Context, id string) (dto.GetUserByID, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return dto.GetUserByID{}, dto.ErrInvalidUUID
	}

	user, err := s.userRepo.GetUserByID(parsedID)
	if err != nil {
		return dto.GetUserByID{}, dto.ErrFailedFindUser
	}

	return dto.GetUserByID{
		Name:       user.Name,
		Email:      user.Email,
		NoTelp:     user.NoTelp,
		Role:       user.Role,
		IsVerified: user.IsVerified,
	}, nil
}

func (s *userService) UpdateUser(ctx *gin.Context, idParam string, req dto.UpdateUser) (dto.UpdateUser, error) {
	parsedID, err := uuid.Parse(idParam)
	if err != nil {
		return dto.UpdateUser{}, dto.ErrInvalidUUID
	}

	authUser, _ := ctx.Get("user")
	authUserData := authUser.(entity.User)

	if idParam != authUserData.ID.String() {
		return dto.UpdateUser{}, dto.ErrUnauthorized
	}

	existingUser, err := s.userRepo.GetUserByID(parsedID)
	if err != nil {
		return dto.UpdateUser{}, dto.ErrFailedFindUser
	}

	if req.Name != "" {
		existingUser.Name = req.Name
	}
	if req.NoTelp != "" {
		existingUser.NoTelp = req.NoTelp
	}

	result, err := s.userRepo.UpdateUser(nil, existingUser)
	if err != nil {
		return dto.UpdateUser{}, dto.ErrFailedUpdateUser
	}

	return dto.UpdateUser{
		Name:   result.Name,
		NoTelp: result.NoTelp,
	}, nil
}
