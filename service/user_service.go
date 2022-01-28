package service

import (
	"context"
	"database/sql"
	"shopping/model/domain"
	"shopping/model/web"
	"shopping/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type TokenClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserCreateResponse
	Sign(ctx context.Context, request web.UserSignRequest) web.UserCreateResponse
	FindAll(ctx context.Context) []web.UserAllResponse
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserCreateResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		panic(err)
	}
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}

	defer func() {
		err := recover()
		if err != nil {
			errRolback := tx.Rollback()
			if errRolback != nil {
				panic(errRolback)
			}
			panic(err)
		} else {
			errCommit := tx.Commit()
			if errCommit != nil {
				panic(errCommit)
			}
		}
	}()

	user := domain.User{
		Username: request.User.Username,
		Password: request.User.EncryptedPassword,
		Email:    request.User.Email,
		Phone:    request.User.Phone,
		Country:  request.User.Country,
		City:     request.User.City,
		Postcode: request.User.Postcode,
		Name:     request.User.Name,
		Address:  request.User.Address,
	}

	tokenClaims := TokenClaims{
		Email: user.Email,
	}
	tokenClaims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 30))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := token.SignedString([]byte("rahasia"))
	if err != nil {
		panic(err)
	}

	user = service.UserRepository.Save(ctx, tx, user)
	return web.UserCreateResponse{
		Email:    user.Email,
		Token:    tokenString,
		Username: user.Username,
	}
}

func (service *UserServiceImpl) Sign(ctx context.Context, request web.UserSignRequest) web.UserCreateResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		panic(err)
	}
	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}

	defer func() {
		err := recover()
		if err != nil {
			errRolback := tx.Rollback()
			if errRolback != nil {
				panic(errRolback)
			}
			panic(err)
		} else {
			errCommit := tx.Commit()
			if errCommit != nil {
				panic(errCommit)
			}
		}
	}()

	user := domain.User{
		Email:    request.Email,
		Password: request.Password,
	}

	tokenClaims := TokenClaims{
		Email: user.Email,
	}
	tokenClaims.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 30))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := token.SignedString([]byte("rahasia"))
	if err != nil {
		panic(err)
	}

	user = service.UserRepository.Sign(ctx, tx, user)
	return web.UserCreateResponse{
		Email:    user.Email,
		Token:    tokenString,
		Username: user.Username,
	}
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []web.UserAllResponse {
	panic("not implemented") // TODO: Implement
}
