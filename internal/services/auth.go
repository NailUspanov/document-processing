package services

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"sstu/internal/models"
	"sstu/internal/repos"
	"time"
)

type Authorization struct {
	repo repos.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId primitive.ObjectID `json:"user_id"`
}

var salt = os.Getenv("SALT")
var signingKey = os.Getenv("SIGNING_KEY")

const tokenTTL = 12 * time.Hour

func NewAuthService(repo repos.Authorization) *Authorization {
	return &Authorization{repo: repo}
}

func (s *Authorization) ParseToken(accessToken string) (primitive.ObjectID, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return primitive.ObjectID{}, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return primitive.ObjectID{}, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (s *Authorization) CreateUser(user models.UserRequest) (any, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *Authorization) GenerateJWT(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
