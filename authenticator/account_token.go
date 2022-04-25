package authenticator

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	"gokost.com/m/delivery/appresponse"
	"time"
)

type Token interface {
	CreateToken(dataLogin appresponse.LoginResponse) (string, error)
	VerifAccessToken(tokenString string) (jwt.MapClaims, error)
	GetAppName() string
	CheckTokenAvailable(tokenString string) (bool, error)
	UpdateToken(tokenString string)
}

type TokenConfig struct {
	AplicationName      string
	JwtSignatureKey     string
	JwtSignatureMethod  *jwt.SigningMethodHMAC
	AccessTokenDuration time.Duration
}

type token struct {
	config TokenConfig
	rdb    *redis.Client
	ctx    context.Context
}

func (t *token) GetAppName() string {
	return t.config.AplicationName
}

func (t *token) UpdateToken(tokenString string) {
	t.rdb.Set(t.ctx, "token", tokenString, t.config.AccessTokenDuration)
}

func (t *token) CreateToken(dataLogin appresponse.LoginResponse) (string, error) {
	claims := MyClaims{ // Menyiapkan struct dengan isi yg dibutuhkan
		StandardClaims: jwt.StandardClaims{
			Issuer: t.config.AplicationName,
		},
		Username: dataLogin.Username,
		Name:     dataLogin.Name,
	}
	token := jwt.NewWithClaims(t.config.JwtSignatureMethod, claims) // membuat jwt dengan format method dan claim berupa struct
	tokenString, err := token.SignedString([]byte(t.config.JwtSignatureKey))
	t.UpdateToken(tokenString)
	return tokenString, err // mendapatkan token
}

func (t *token) CheckTokenAvailable(tokenString string) (bool, error) {
	tokenAuth, err := t.rdb.Get(t.ctx, "token").Result()
	if err != nil {
		return false, err
	}
	if tokenString != tokenAuth {
		return false, nil
	}
	return true, nil
}

func (t *token) VerifAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signid Method invalid")
		} else if method != t.config.JwtSignatureMethod {
			return nil, fmt.Errorf("signid Method invalid")
		}
		return []byte(t.config.JwtSignatureKey), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}

func NewToken(config TokenConfig, ctx context.Context, rdb *redis.Client) Token {
	return &token{
		config,
		rdb,
		ctx,
	}
}
