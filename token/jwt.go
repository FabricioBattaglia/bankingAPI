package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const SECRET_KEY = "CHAVE SECRETA"

const VALID_DAYS_AMOUT = 7

type MyCustomClaims struct {
	Account int `json:"account"`
	jwt.StandardClaims
}

func GetPrivateKey() []byte {
	return []byte(SECRET_KEY)
}

func GenerateJWT(account int) (string, error) {
	claims := MyCustomClaims{
		account,
		jwt.StandardClaims{
			//ExpiresAt: time.Now().Unix() - 15000,
			ExpiresAt: time.Now().AddDate(0, 0, VALID_DAYS_AMOUT).Unix(),
			Issuer:    "banking-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(GetPrivateKey())
	if err != nil {
		return "", err
	}

	return ss, nil
}

func ParseJWT(token string) (*MyCustomClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetPrivateKey()), nil
	})

	if claims, ok := parsedToken.Claims.(*MyCustomClaims); ok && parsedToken.Valid {
		return claims, nil
	} else {
		return &MyCustomClaims{}, err
	}
}
