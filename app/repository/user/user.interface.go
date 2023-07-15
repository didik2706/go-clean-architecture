package repository

import (
	"context"
	"database/sql"
	"latihan-restful-api-2/app/entity"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user entity.User) entity.User
	Update(ctx context.Context, tx *sql.Tx, user entity.User) entity.User
	Delete(ctx context.Context, tx *sql.Tx, id int)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.User
	FindById(ctx context.Context, tx *sql.Tx, id int) (entity.User, error)
}