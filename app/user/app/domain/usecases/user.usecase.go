package usecases

import (
	"context"
	"errors"
	"user/app/application/commands/command"
	"user/app/domain/entities"
	"user/app/domain/value_object"
	"user/app/infrastructure/persistent/postgresql/mapper"
	"user/app/infrastructure/persistent/postgresql/repository"
	"user/package/encrypt_password"
	"user/package/jwt_utils"
	"user/package/logger"
	"user/package/xtypes"

	"go.uber.org/zap"
)

type UserUsecase interface {
	Login(ctx context.Context, command *command.LoginCommand) (*command.LoginCommandResult, error)
	CreateUser(ctx context.Context, command *command.CreateUserCommand) (*command.CreateUserCommandResult, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
	jwt_utils      jwt_utils.JWTUtils
}

func NewUserUsecase(userRepository repository.UserRepository, jwt_utils jwt_utils.JWTUtils) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		jwt_utils:      jwt_utils,
	}
}

func (s *userUsecase) Login(ctx context.Context, cmd *command.LoginCommand) (*command.LoginCommandResult, error) {
	log := logger.FromContext(ctx)

	user, err := s.userRepository.GetUserByEmail(ctx, cmd.Email)
	if err != nil {
		log.Error("failed to get user by email", zap.Error(err))
		return nil, err
	}

	if user == nil {
		log.Error("user not found")
		return nil, errors.New("user not found")
	}

	check, err := encrypt_password.VerifyPassword(user.Password, cmd.Password)
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

	return &command.LoginCommandResult{
		AccessToken:           at,
		RefreshToken:          rt,
		AccessTokenExpiresIn:  aexp,
		RefreshTokenExpiresIn: rexp,
	}, nil
}

func (s *userUsecase) CreateUser(ctx context.Context, cmd *command.CreateUserCommand) (*command.CreateUserCommandResult, error) {
	log := logger.FromContext(ctx)
	userEntities := entities.DefaultUser()

	userEntities.SetEmail(cmd.Email)
	userEntities.SetPassword(value_object.Password(cmd.Password))
	userEntities.SetFirstName(cmd.FirstName)
	userEntities.SetLastName(cmd.LastName)
	userEntities.SetPhoneNumber(cmd.PhoneNumber)
	userEntities.SetBirthDate(&cmd.BirthDate)
	userEntities.SetTier(xtypes.TierBronze)

	if !userEntities.Password().IsValid() {
		log.Error("invalid password")
		return nil, errors.New("invalid password")
	}

	hashedPassword, err := encrypt_password.HashPassword(cmd.Password)
	if err != nil {
		log.Error("failed to hash password", zap.Error(err))
		return nil, err
	}

	userEntities.SetPasswordHash(hashedPassword)

	_, err = s.userRepository.Create(ctx, mapper.UserToModel(userEntities))
	if err != nil {
		log.Error("invalid create user command", zap.Error(err))
		return nil, err
	}

	return &command.CreateUserCommandResult{
		Message: "success",
	}, nil
}
