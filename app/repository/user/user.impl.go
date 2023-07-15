package repository

import (
	"context"
	"database/sql"
	"errors"
	"latihan-restful-api-2/app/entity"
	"latihan-restful-api-2/pkg"
)

type UserRepositoryImpl struct {}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user entity.User) entity.User {
	query 			:= "INSERT INTO users(name, username, password) VALUES (?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, user.Name, user.Username, user.Password)
	pkg.PanicIfError(err)

	// get user id
	id, err := result.LastInsertId()
	pkg.PanicIfError(err)
	user.Id = int(id)

	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user entity.User) entity.User {
	query  := "UPDATE users SET name = ?, username = ?, password = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, user.Name, user.Username, user.Password, user.Id)
	pkg.PanicIfError(err)
	
	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) {
	query  := "DELETE FROM users WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, id)
	pkg.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.User {
	query 	  := "SELECT * FROM users"
	rows, err := tx.QueryContext(ctx, query)
	pkg.PanicIfError(err)
	defer rows.Close()

	users := []entity.User{}

	for rows.Next() {
		user := entity.User{}
		rows.Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		users = append(users, user)
	}

	return users
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (entity.User, error) {
	query 		:= "SELECT * FROM users WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, id)
	pkg.PanicIfError(err)
	defer rows.Close()

	user := entity.User{}

	if rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		return user, nil
	} else {
		return user, errors.New("user id not found")
	}
}

