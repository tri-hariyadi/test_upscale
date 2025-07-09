package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"net/http"
	"time"
	"todo-app/common/config"
	"todo-app/model"
)

type CustomClaims struct {
	ID     string `json:"id"`
	UserID int64  `json:"user-id"`
	Email  string `json:"email"`
	Name   string `json:"username"`
	jwt.RegisteredClaims
}

type TokenService interface {
	GenerateAccessToken(user model.User) (string, error)
}

type TokenAuthService struct{}

func NewTokenAuthService() *TokenAuthService {
	return &TokenAuthService{}
}

func (t *TokenAuthService) GenerateAccessToken(user model.User) (string, error) {
	claims := CustomClaims{
		ID:     uuid.NewString(),
		UserID: user.ID,
		Email:  user.Email,
		Name:   user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JWTSecret.AccessSecret))
}

func SetTokenToCookie(w http.ResponseWriter, accessToken string) {
	var secure bool = false
	var sameSite http.SameSite = http.SameSiteLaxMode
	var domain string = "http://localhost:5173"

	if config.AppConfig.App.Env == "production" {
		secure = true
		sameSite = http.SameSiteNoneMode
		domain = "https://test-upscale.vercel.app"
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		Domain:   domain,
		HttpOnly: true,
		Secure:   secure,
		SameSite: sameSite,
		MaxAge:   7 * 60 * 60,
	})
}
