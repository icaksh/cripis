package utils

import (
	"github.com/icaksh/cripis/app/models"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateNewAuthToken(par *models.JwtAuthModel) (*models.JwtTokenDetails, error) {
	td := &models.JwtTokenDetails{}
	var err error
	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_ACCESS_TIME_KEY_EXPIRE_MINUTES_COUNT"))
	accessTime := time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()
	atm := &JWTAccess{
		AccessUuid: par.AccessId,
		User:       par.UserId,
		Role:       par.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessTime,
		},
	}
	accessToken, err := GenerateNewAccessToken(atm)

	rtm := &JWTRefresh{
		RefreshUuid: par.RefreshId,
		User:        par.UserId,
		Email:       par.Email,
		FirstName:   par.FirstName,
		LastName:    par.LastName,
		Role:        par.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: par.Duration,
		},
	}
	refreshToken, err := GenerateNewRefreshToken(rtm)

	if err != nil {
		return nil, err
	}

	td = &models.JwtTokenDetails{
		AccessUuid:   par.AccessId.String(),
		AccessToken:  accessToken,
		RefreshUuid:  par.RefreshId.String(),
		RefreshToken: refreshToken,
		AtExpires:    accessTime,
		RtExpires:    par.Duration,
	}
	return td, err
}

func GenerateNewAccessToken(atm *JWTAccess) (string, error) {
	secret := os.Getenv("JWT_SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atm)

	accessToken, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}
	return accessToken, nil
}

func GenerateNewRefreshToken(rtm *JWTRefresh) (string, error) {
	secret := os.Getenv("JWT_REFRESH_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, rtm)

	refreshToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "nil", err
	}
	return refreshToken, nil
}
