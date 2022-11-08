package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
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
	UserID int `json:"user_guid"`
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
		1,
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) GenerateCookie(user app.User) *http.Cookie {
	cookieValue := user.Login + ":" + generateCookieHash(user.Login+strconv.Itoa(rand.Intn(100000000)))
	expire := time.Now().AddDate(0, 0, 1)
	return &http.Cookie{Name: "SessionID", Value: cookieValue, Expires: expire, HttpOnly: true}

	// timeNow := time.Now().Format("2006-01-02 15:04:05")
	// val := generateCookieHash(timeNow + user.Login)
	// cookie := http.Cookie{
	// 	Name:    "sessionID",
	// 	Value:   val,
	// 	Expires: time.Now().Add(12 * time.Hour),
	// 	MaxAge:  60 * 12 * 60,
	// }
	// return &cookie
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserID, nil
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
