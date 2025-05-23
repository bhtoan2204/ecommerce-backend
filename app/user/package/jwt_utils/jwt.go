package jwt_utils

import (
	"errors"
	"time"
	"user/app/domain/entities"

	"github.com/golang-jwt/jwt/v5"
)

type JWTUtils struct {
	AccessSecret      string
	RefreshSecret     string
	AccessExpiration  int32
	RefreshExpiration int32
}

func NewJWTUtils(accessSecret, refreshSecret string, accessExpiration, refreshExpiration int32) *JWTUtils {
	return &JWTUtils{
		AccessSecret:      accessSecret,
		RefreshSecret:     refreshSecret,
		AccessExpiration:  accessExpiration,
		RefreshExpiration: refreshExpiration,
	}
}

func (j *JWTUtils) GenerateToken(user *entities.User) (string, string, int64, int64, error) {
	accessSecret := []byte(j.AccessSecret)
	accessExpiration := time.Now().Add(time.Duration(j.AccessExpiration) * time.Second).Unix()
	refreshSecret := []byte(j.RefreshSecret)
	refreshExpiration := time.Now().Add(time.Duration(j.RefreshExpiration) * time.Second).Unix()

	accessClaims := jwt.MapClaims{
		"id":    user.ID(),
		"email": user.Email(),
		"exp":   accessExpiration,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	signedAccessToken, err := accessToken.SignedString(accessSecret)

	if err != nil {
		return "", "", 0, 0, err
	}

	refreshClaims := jwt.MapClaims{
		"id":  user.ID(),
		"exp": refreshExpiration,
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefreshToken, err := refreshToken.SignedString(refreshSecret)
	if err != nil {
		return "", "", 0, 0, err
	}

	return signedAccessToken, signedRefreshToken, int64(accessExpiration), int64(refreshExpiration), nil
}

func (j *JWTUtils) VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	return j.parseToken(tokenString, []byte(j.AccessSecret))
}

func (j *JWTUtils) VerifyRefreshToken(tokenString string) (jwt.MapClaims, error) {
	return j.parseToken(tokenString, []byte(j.RefreshSecret))
}

func (j *JWTUtils) parseToken(tokenString string, secret []byte) (jwt.MapClaims, error) {
	tok, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if !tok.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}
	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return nil, jwt.ErrInvalidKey
	}
	return claims, nil
}

func (j *JWTUtils) ExtractAccessClaims(tokenString string) (jwt.MapClaims, error) {
	return j.VerifyAccessToken(tokenString)
}

func (j *JWTUtils) RefreshToken(user *entities.User, refreshTokenString string) (newAccessToken, newRefreshToken string, newAccessExp, newRefreshExp int64, err error) {
	claims, err := j.VerifyRefreshToken(refreshTokenString)
	if err != nil {
		return "", "", 0, 0, err
	}

	expFloat, ok := claims["exp"].(float64)
	if !ok {
		return "", "", 0, 0, errors.New("invalid exp claim in refresh token")
	}
	if time.Now().Unix() > int64(expFloat) {
		return "", "", 0, 0, jwt.ErrTokenExpired
	}

	return j.GenerateToken(user)
}
