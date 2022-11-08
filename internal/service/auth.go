package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/UserNaMEeman/shops/app"
	"github.com/UserNaMEeman/shops/internal/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "qw3vvfgy6ffrs"
	signingKey = "dsda$$@ggdgs$#@#$f"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserGUID string `json:"user_guid"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user app.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(user app.User) (string, error) {
	// fmt.Println("do")
	user.Password = generatePasswordHash(user.Password)
	_, err := s.repo.GetUser(user)
	if err != nil {
		return "", err
	}
	// fmt.Println("do")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		"1",
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	jwtString := strings.Split(accessToken, "Bearer ")[1]
	token, err := jwt.ParseWithClaims(jwtString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	// fmt.Println("Raw: ", token.Raw)
	// fmt.Println("Method: ", token.Method)
	// fmt.Println("Header: ", token.Header)
	// fmt.Println("Claims: ", token.Claims)
	// fmt.Println("Signature: ", token.Signature)
	// fmt.Println("Valid: ", token.Valid)
	if err != nil {
		// fmt.Println("ERRRR: ", err)
		return "", err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserGUID, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func generateCookieHash(value string) string {
	hash := sha1.New()
	hash.Write([]byte(value))
	return fmt.Sprintf("%x", hash.Sum([]byte("")))
}
