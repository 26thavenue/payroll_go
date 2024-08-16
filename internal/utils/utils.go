package utils

import (
	_ "errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your_secret_key")

type Claims struct {
    UserID uint   `json:"user_id"`
    Role   string `json:"role"`
    jwt.StandardClaims
}

func ComparePasswords (hashedPassword, plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))

	if err != nil{
		return err
	}

	return  nil
}

func HashPassword(plainPassword string) (string,error){

	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(plainPassword),bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	password := string(hashedPassword)

	return password, nil
}

func GenerateAccessToken(userID uint, role string) (string, error) {
    claims := &Claims{
        UserID: userID,
        Role:   role,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(5 * time.Hour).Unix(), 
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

func GenerateRefreshToken(userID uint) (string, error) {
    claims := &Claims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(3 * 24 * time.Hour).Unix(), 
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}