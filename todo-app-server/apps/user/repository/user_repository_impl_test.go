package repository

import (
	"context"
	"fmt"
	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"todo-app/model"
	"todo-app/model/web"
)

var dbMock pgxmock.PgxPoolIface
var userRepositoryImpl UserRepositoryImpl;

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")

	mock, err := pgxmock.NewPool()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer mock.Close()
	dbMock = mock
	userRepositoryImpl = UserRepositoryImpl{
		DB: dbMock,
	}
	m.Run()

	fmt.Println("AFTER UNIT TEST")
}

func TestUserRepositoryImpl_Create(t *testing.T) {
	t.Run("Positive Case", func(t *testing.T) {
		expectedID := int64(1)
		createReq := web.CreateRequest{
			Email:    "trihariyadi24@gmail.com",
			Password: "hashed_password",
			Name:     "trihariyadi24",
		}

		dbMock.ExpectQuery(`INSERT INTO users \(name, email, password\) VALUES \(\$1, \$2, \$3\) RETURNING id`).
			WithArgs(createReq.Name, createReq.Email, createReq.Password).
			WillReturnRows(pgxmock.NewRows([]string{"id"}).AddRow(expectedID))

		id, err := userRepositoryImpl.Create(context.Background(), createReq)

		assert.NoError(t, err)
		assert.Equal(t, expectedID, id)
		assert.NoError(t, dbMock.ExpectationsWereMet())
	})

	t.Run("Negative Case - Error from DB", func(t *testing.T) {
		createReq := web.CreateRequest{
			Name:     "Error Case",
			Email:    "fail@example.com",
			Password: "wrongpass",
		}

		dbMock.ExpectQuery(`INSERT INTO users \(name, email, password\) VALUES \(\$1, \$2, \$3\) RETURNING id`).
			WithArgs(createReq.Name, createReq.Email, createReq.Password).
			WillReturnError(fmt.Errorf("insert failed"))

		id, err := userRepositoryImpl.Create(context.Background(), createReq)

		assert.Error(t, err)
		assert.Equal(t, int64(0), id)
		assert.NoError(t, dbMock.ExpectationsWereMet())
	})
}

func TestUserRepositoryImpl_FindByEmail(t *testing.T) {
	t.Run("Positive Case", func(t *testing.T) {
		expectedPassword := "hashed_password"
		loginReq := web.LoginRequest{
			Email:    "trihariyadi24@gmail.com",
		}

		rows := pgxmock.NewRows([]string{"id", "email", "name", "password"}).AddRow(5, loginReq.Email, "Tri Hariyadi", expectedPassword)
		dbMock.ExpectQuery("SELECT id, email, name, password FROM users WHERE email = \\$1").
			WithArgs(loginReq.Email).
			WillReturnRows(rows)

		user, err := userRepositoryImpl.FindByEmail(context.Background(), loginReq.Email)

		assert.NoError(t, err)
		assert.Equal(t, expectedPassword, user.Password)
		assert.NoError(t, dbMock.ExpectationsWereMet())
	})

	t.Run("Negative Case - Not Found", func(t *testing.T) {
		mock, err := pgxmock.NewPool()
		assert.NoError(t, err)
		defer mock.Close()

		repo := UserRepositoryImpl{DB: mock}

		loginReq := web.LoginRequest{Email: "notfound@example.com"}

		rows := pgxmock.NewRows([]string{"password"})

		mock.ExpectQuery("SELECT id, email, name, password FROM users WHERE email = \\$1").
			WithArgs(loginReq.Email).
			WillReturnRows(rows)

		user, err := repo.FindByEmail(context.Background(), loginReq.Email)

		assert.Error(t, err)
		assert.Equal(t, "User not found", err.Error())
		var expected *model.User = nil
		assert.Equal(t, expected, user)
	})
}
