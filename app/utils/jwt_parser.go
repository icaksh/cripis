package utils

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type JWTAccess struct {
	AccessUuid uuid.UUID `json:"accessUuid"`
	User       uuid.UUID `json:"userId"`
	Role       int16     `json:"role"`
	jwt.StandardClaims
}

type JWTRefresh struct {
	RefreshUuid uuid.UUID `json:"refreshUuid"`
	User        uuid.UUID `json:"userId"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	Role        int16     `json:"role"`
	jwt.StandardClaims
}

func ExtractRefreshToken(t string) (*JWTRefresh, error) {
	token, err := jwt.ParseWithClaims(t, &JWTRefresh{}, jwtKeyFunc)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JWTRefresh)
	if ok && token.Valid {
		return &JWTRefresh{
			RefreshUuid: claims.RefreshUuid,
			User:        claims.User,
			Email:       claims.Email,
			FirstName:   claims.FirstName,
			LastName:    claims.LastName,
			Role:        claims.Role,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: claims.ExpiresAt,
			},
		}, nil
	}
	return nil, err
}

func ExtractRefreshTokenMetadata(c *fiber.Ctx) (*JWTRefresh, error) {
	token, err := verifyRefreshToken(c)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(*JWTRefresh)
	if ok && token.Valid {
		// Expires time.
		return &JWTRefresh{
			RefreshUuid: claims.RefreshUuid,
			User:        claims.User,
			Email:       claims.Email,
			FirstName:   claims.FirstName,
			LastName:    claims.LastName,
			Role:        claims.Role,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: claims.ExpiresAt,
			},
		}, nil
	}

	return nil, err
}

func ExtractTokenMetadata(c *fiber.Ctx) (*JWTAccess, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTAccess)
	if ok && token.Valid {
		return &JWTAccess{
			AccessUuid: claims.AccessUuid,
			User:       claims.User,
			Role:       claims.Role,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: claims.ExpiresAt,
			},
		}, nil
	}

	return nil, err
}

func TokenValid(c *fiber.Ctx) error {
	token, err := verifyToken(c)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	// Normally Authorization HTTP header.
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.ParseWithClaims(tokenString, &JWTAccess{}, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func verifyRefreshToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.ParseWithClaims(tokenString, &JWTRefresh{}, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}
func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}
func jwtRefreshKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_REFRESH_KEY")), nil
}
