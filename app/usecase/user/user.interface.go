package usecase

import "context"

type UserUsecase interface {
	Create(ctx context.Context, request UserUsecaseRequestCreate) UserUsecaseResponse
	Update(ctx context.Context, request UserUsecaseRequestUpdate) UserUsecaseResponse
	Delete(ctx context.Context, id int)
	FindAll(ctx context.Context) []UserUsecaseResponse
	FindById(ctx context.Context, id int) UserUsecaseResponse
}