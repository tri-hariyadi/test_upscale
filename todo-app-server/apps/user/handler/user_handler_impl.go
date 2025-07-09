package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"todo-app/apps/user/service"
	"todo-app/common/config"
	"todo-app/common/router"
	"todo-app/helpers/auth"
	"todo-app/helpers/resp_handler"
	"todo-app/middleware"
	"todo-app/model/web"
)

type UserHandlerImpl struct {
	UserService service.UserService
}

func NewUserHandlerImpl(userService service.UserService) router.UserRegisterHandlerFunc {
	return func(router *mux.Router) {
		userHandler := &UserHandlerImpl{UserService: userService}
		authRouter := router.NewRoute().Subrouter()
		authRouter.Use(middleware.TokenVerifyMiddleware())

		router.HandleFunc("/auth/login", userHandler.Login).Methods("POST")
		router.HandleFunc("/auth/register", userHandler.Register).Methods("POST")
		router.HandleFunc("/auth/logout", userHandler.Logout).Methods("GET")
		authRouter.HandleFunc("/auth/verify-token", userHandler.VerifyToken).Methods("GET")
	}
}

func (handler UserHandlerImpl) Login(w http.ResponseWriter, r *http.Request) {
	loginRequest := web.LoginRequest{}
	resp_handler.HandleRequestBody(r, &loginRequest)

	result := handler.UserService.Login(r.Context(), loginRequest)
	auth.SetTokenToCookie(w, result.AccessToken);
	resp_handler.HandleSuccess(w, http.StatusOK, nil)
}

func (handler UserHandlerImpl) Register(w http.ResponseWriter, r *http.Request) {
	registerRequest := web.CreateRequest{}
	resp_handler.HandleRequestBody(r, &registerRequest)
	fmt.Println(registerRequest.Password)

	result := handler.UserService.Register(r.Context(), registerRequest)
	auth.SetTokenToCookie(w, result.AccessToken);
	resp_handler.HandleSuccess(w, http.StatusOK, nil)
}

func (handler UserHandlerImpl) Logout(w http.ResponseWriter, r *http.Request) {
	var secure bool = false
	var sameSite http.SameSite = http.SameSiteLaxMode
	var domain string = "localhost:5173"

	if config.AppConfig.App.Env == "production" {
		secure = true
		sameSite = http.SameSiteNoneMode
		domain = "test-upscale.vercel.app"
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		Domain:   domain,
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   secure,
		SameSite: sameSite,
	})
}

func (handler UserHandlerImpl) VerifyToken(w http.ResponseWriter, r *http.Request) {
	result := handler.UserService.VerifyToken(r.Context())
	//time.Sleep(5 * time.Second)
	resp_handler.HandleSuccess(w, http.StatusOK, result)
}

