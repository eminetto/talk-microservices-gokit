package security

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	JWTSecret   = "d7830ad5791dsdsds"
	JwtExpHour  =1
	JwtExpMin = 0
	JwtExpSec = 30
)

//NewToken create a new token
func NewToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   email,
		"nbf":       time.Now().Unix(),
		"iat":       time.Now().Unix(),
		"exp":       time.Now().Local().Add(time.Hour*time.Duration(JwtExpHour) + time.Minute*time.Duration(JwtExpMin) + time.Second*time.Duration(JwtExpSec)).Unix(),
	})
	// Sign and get the complete encoded token as a string using the secret
	sToken, err := token.SignedString([]byte(JWTSecret))
	if err != nil {
		return "", err
	}

	return sToken, nil
}

//ParseToken parse a token
func ParseToken(tokenString string) (*jwt.Token, error) {
	var token *jwt.Token
	var err error
	token, err = parseHS256(tokenString, token)
	if err != nil && err.Error() != "Token is expired" {
		token, err = parseHS256(tokenString, token)
	}

	return token, err
}

func parseHS256(tokenString string, token *jwt.Token) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWTSecret), nil
	})
	return token, err
}

//GetClaims get claims information
func GetClaims(token *jwt.Token) (jwt.MapClaims, error) {
	if !token.Valid {
		return nil, fmt.Errorf("unauthorized")
	}
	err := token.Claims.(jwt.MapClaims).Valid()
	if err != nil {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}