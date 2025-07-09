package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo-app/apps/user/service"
	"todo-app/common/logger"
	"todo-app/middleware"
	"todo-app/model/web"
)

var userServiceMock *service.UserServiceMock
var r *mux.Router

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")

	logger.InitLogger()
	userServiceMock = new(service.UserServiceMock)
	r = mux.NewRouter()
	r.Use(middleware.RecoverMiddleware)
	NewUserHandlerImpl(userServiceMock)(r)
	m.Run()

	fmt.Println("AFTER UNIT TEST")
}

func TestUserHandlerImpl_Login(t *testing.T) {
	t.Run("Test Success", func(t *testing.T) {
		body := `{"email":"test@example.com","password":"secret123"}`
		req := httptest.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		expectedReq := web.LoginRequest{
			Email:    "test@example.com",
			Password: "secret123",
		}

		expectedResp := web.LoginResponse{
			AccessToken:  "mock-access-token",
		}

		userServiceMock.On("Login", mock.Anything, expectedReq).Return(expectedResp)
		r.ServeHTTP(w, req)

		fmt.Println("w.Body", w.Body)
		responseBody := web.ResponseWrapper{Data: web.LoginResponse{}}
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, true, responseBody.Status)
		assert.Equal(t, "success", responseBody.Message)

		cookie := w.Result().Cookies()
		require.NotEmpty(t, cookie)
		var found bool
		for _, c := range cookie {
			if c.Name == "access_token" {
				found = true
				assert.Equal(t, "mock-access-token", c.Value)
			}
		}
		assert.True(t, found, "Expected access_token cookie to be set")
	})
}
