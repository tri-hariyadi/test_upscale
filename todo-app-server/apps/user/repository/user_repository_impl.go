package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"todo-app/common/database"
	"todo-app/helpers/exception"
	"todo-app/model"
	"todo-app/model/web"
)

type UserRepositoryImpl struct {
	DB database.DBQuerier
}

func NewUserRepositoryImpl(DB *pgxpool.Pool) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: DB}
}

func (u *UserRepositoryImpl) Create(ctx context.Context, data web.CreateRequest) (int64, error) {
	SQL := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	var lastInsertedID int64
	err := u.DB.QueryRow(ctx, SQL, data.Name, data.Email, data.Password).Scan(&lastInsertedID)
	return lastInsertedID, err
}

func (u *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}
	SQL := "SELECT id, email, name, password FROM users WHERE email = $1"
	rows, err := u.DB.Query(ctx, SQL, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Email, &user.Name, &user.Password)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, exception.NewNotFoundError("User not found")
}

