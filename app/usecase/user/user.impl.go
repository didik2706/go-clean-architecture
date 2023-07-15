package usecase

import (
	"context"
	"database/sql"
	"latihan-restful-api-2/app/entity"
	repository "latihan-restful-api-2/app/repository/user"
	"latihan-restful-api-2/exception"
	"latihan-restful-api-2/pkg"

	"github.com/go-playground/validator/v10"
)

type UserUsecaseImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserUsecase(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserUsecase {
	return &UserUsecaseImpl{
		UserRepository: userRepository,
		DB: DB,
		Validate: validate,
	}
}

func (usecase *UserUsecaseImpl) Create(ctx context.Context, request UserUsecaseRequestCreate) UserUsecaseResponse {
	err := usecase.Validate.Struct(request)
	pkg.PanicIfError(err)

	tx, err := usecase.DB.Begin()
	pkg.PanicIfError(err)
	defer pkg.CommitOrRollback(tx)

	user := entity.User{
		Name: request.Name,
		Username: request.Username,
		Password: request.Password,
	}
	data := usecase.UserRepository.Create(ctx, tx, user)
	user.Id = data.Id

	return ToUserResponse(user)
}

func (usecase *UserUsecaseImpl) Update(ctx context.Context, request UserUsecaseRequestUpdate) UserUsecaseResponse {
	err := usecase.Validate.Struct(request)
	pkg.PanicIfError(err)

	tx, err := usecase.DB.Begin()
	pkg.PanicIfError(err)
	defer pkg.CommitOrRollback(tx)

	// check existing data
	user, err := usecase.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user.Name 		= request.Name
	user.Username = request.Username
	user.Password = request.Password

	usecase.UserRepository.Update(ctx, tx, user)

	return ToUserResponse(user)
}

func (usecase *UserUsecaseImpl) Delete(ctx context.Context, id int) {
	tx, err := usecase.DB.Begin()
	pkg.PanicIfError(err)
	defer pkg.CommitOrRollback(tx)

	user, err := usecase.UserRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	usecase.UserRepository.Delete(ctx, tx, user.Id)
}

func (usecase *UserUsecaseImpl) FindAll(ctx context.Context) []UserUsecaseResponse {
	tx, err := usecase.DB.Begin()
	pkg.PanicIfError(err)
	defer pkg.CommitOrRollback(tx)

	users := usecase.UserRepository.FindAll(ctx, tx)
	return ToUsersResponse(users)
}

func (usecase *UserUsecaseImpl) FindById(ctx context.Context, id int) UserUsecaseResponse {
	tx, err := usecase.DB.Begin()
	pkg.PanicIfError(err)
	defer pkg.CommitOrRollback(tx)

	user, err := usecase.UserRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return ToUserResponse(user)
}
