package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"todo-app/common/database"
	"todo-app/helpers/exception"
	"todo-app/model/web"
)

type TodoRepositoryImpl struct {
	DB database.DBQuerier
}

func NewTodoRepositoryImpl(DB *pgxpool.Pool) *TodoRepositoryImpl {
	return &TodoRepositoryImpl{DB: DB}
}

func (todo TodoRepositoryImpl) Create(ctx context.Context, req *web.CreateTodoRequest, userId int64) (*web.TodoResponse, error) {
	SQL := "INSERT INTO todos (title, description, tag, due_date, user_id) VALUES ($1, $2, $3, $4, $5) RETURNING id, title, description, tag, due_date, created_at, updated_at"
	result := &web.TodoResponse{}
	err := todo.DB.QueryRow(ctx, SQL, req.Title, req.Description, req.Tag, req.DueDate, userId).Scan(
		&result.ID,
		&result.Title,
		&result.Description,
		&result.Tag,
		&result.DueDate,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	return result, err
}

func (todo TodoRepositoryImpl) Get(ctx context.Context, id int64) (*web.TodoResponse, error) {
	t := &web.TodoResponse{}
	SQL := "SELECT id, title, description, tag, due_date, created_at, updated_at FROM todos WHERE  id = $1"
	rows, err := todo.DB.Query(ctx, SQL, id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&t.ID, &t.Title, &t.Description, &t.Tag, &t.DueDate, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		return t, nil
	}
	return nil, exception.NewNotFoundError("Todo not found")
}

func (todo TodoRepositoryImpl) GetAll(ctx context.Context, userId int64) ([]*web.TodoResponse, error) {
	SQL := "SELECT id, title, description, tag, due_date, created_at, updated_at FROM todos WHERE  user_id = $1"
	rows, err := todo.DB.Query(ctx, SQL, pgx.QueryExecModeSimpleProtocol, userId)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var todos []*web.TodoResponse
	for rows.Next() {
		t := &web.TodoResponse{}
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Tag, &t.DueDate, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

func (todo TodoRepositoryImpl) Update(ctx context.Context, req *web.CreateTodoRequest, id int64) (*web.CreateTodoRequest, error) {
	SQL := "UPDATE todos SET title = $1, description = $2, tag = $3, due_date = $4 WHERE id = $5"
	_, err := todo.DB.Exec(ctx, SQL, req.Title, req.Description, req.Tag, req.DueDate, id)
	return req, err
}

func (todo TodoRepositoryImpl) Delete(ctx context.Context, id int64) (error) {
	SQL := "DELETE FROM todos WHERE id = $1"
	_, err := todo.DB.Exec(ctx, SQL, id)
	return err
}

func (todo TodoRepositoryImpl) FilterByTag(ctx context.Context, search string, tag interface{}) ([]*web.TodoResponse, error) {
	//SQL := "SELECT id, title, description, tag, due_date, created_at, updated_at FROM todos WHERE  tag = $1"
	SQL := "SELECT id, title, description, tag, due_date, created_at, updated_at FROM todos WHERE (title ILIKE '%' || $1 || '%' OR description ILIKE '%' || $1 || '%' OR tag::text ILIKE '%' || $1 || '%') AND ($2::tag_enum IS NULL OR tag = $2::tag_enum) ORDER BY due_date ASC"
	rows, err := todo.DB.Query(ctx, SQL, search, tag)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var todos []*web.TodoResponse
	for rows.Next() {
		t := &web.TodoResponse{}
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Tag, &t.DueDate, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

