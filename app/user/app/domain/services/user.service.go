package services

import (
	"context"
	"errors"
	"user/app/domain/dto"
	"user/app/domain/entities"
	"user/app/infrastructure/persistent/postgresql/mapper"
	"user/app/infrastructure/persistent/postgresql/repository"
	"user/package/encrypt_password"
	"user/package/jwt_utils"
	"user/package/logger"
	"user/package/xtypes"

	"go.uber.org/zap"
)

type UserService interface {
	Login(ctx context.Context, request *dto.LoginRequest) (*dto.LoginResponse, error)
	CreateUser(ctx context.Context, request *dto.CreateUserRequest) (*dto.CreateUserResponse, error)
}

type userService struct {
	userRepository repository.UserRepository
	jwt_utils      jwt_utils.JWTUtils
}

func NewUserService(userRepository repository.UserRepository, jwt_utils jwt_utils.JWTUtils) UserService {
	return &userService{
		userRepository: userRepository,
		jwt_utils:      jwt_utils,
	}
}

func (s *userService) Login(ctx context.Context, request *dto.LoginRequest) (*dto.LoginResponse, error) {
	log := logger.FromContext(ctx)
	if err := request.Validate(); err != nil {
		log.Error("invalid login request", zap.Error(err))
		return nil, err
	}

	user, err := s.userRepository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		log.Error("failed to get user by email", zap.Error(err))
		return nil, err
	}

	if user == nil {
		log.Error("user not found")
		return nil, errors.New("user not found")
	}

	check, err := encrypt_password.VerifyPassword(user.Password, request.Password)
	if err != nil {
		log.Error("failed to get verify password", zap.Error(err))
		return nil, err
	}
	if !check {
		log.Error("user not found")
		return nil, errors.New("wrong password")
	}

	at, rt, aexp, rexp, err := s.jwt_utils.GenerateToken(mapper.UserToEntity(user))
	if err != nil {
		log.Error("failed to get generate token", zap.Error(err))
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:           at,
		RefreshToken:          rt,
		AccessTokenExpiresIn:  aexp,
		RefreshTokenExpiresIn: rexp,
	}, nil
}

func (s *userService) CreateUser(ctx context.Context, request *dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	log := logger.FromContext(ctx)
	if err := request.Validate(); err != nil {
		log.Error("invalid login request", zap.Error(err))
		return nil, err
	}
	hashedPassword, err := encrypt_password.HashPassword(request.Password)
	if err != nil {
		log.Error("failed to hash password", zap.Error(err))
		return nil, err
	}
	userEntities := entities.DefaultUser()

	userEntities.SetEmail(request.Email)
	userEntities.SetPasswordHash(hashedPassword)
	userEntities.SetFirstName(request.FirstName)
	userEntities.SetLastName(request.LastName)
	userEntities.SetPhoneNumber(request.PhoneNumber)
	userEntities.SetBirthDate(&request.BirthDate)
	userEntities.SetTier(xtypes.TierBronze)

	_, err = s.userRepository.Create(ctx, mapper.UserToModel(userEntities))
	if err != nil {
		log.Error("invalid create user request", zap.Error(err))
		return nil, err
	}

	return &dto.CreateUserResponse{
		Message: "success",
	}, nil
}
