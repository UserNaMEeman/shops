package service

import (
	"crypto/sha1"
	"fmt"
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
	UserId int `json:"user_guid"`
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
	fmt.Println("do")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		1,
	})
	fmt.Println(token)
	t, err := token.SignedString([]byte(signingKey))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(t)
	return t, nil
	// return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
