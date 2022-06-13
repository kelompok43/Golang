package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kelompok43/Golang/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JWTCustomClaim struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (cJWT ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JWTCustomClaim{},
		SigningKey: []byte(auth.SECRET_KEY),
	}
}

func (cJWT ConfigJWT) GenerateToken(userID int) (token string, err error) {
	claims := JWTCustomClaim{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(cJWT.ExpiresDuration))).Unix(),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = t.SignedString([]byte(auth.SECRET_KEY))

	if err != nil {
		return token, err
	}

	return token, nil
}

func GetUser(ctx echo.Context) *JWTCustomClaim {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTCustomClaim)
	return claims
}
