package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secret []byte
}

func (j *JWT) New(claims jwt.Claims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedJWT, err := token.SignedString(j.secret)
	if err != nil {
		panic(err)
	}
	return signedJWT
}

func (j *JWT) KeyFunc(t *jwt.Token) (interface{}, error) { return j.secret, nil }

func New(secret []byte) *JWT {
	return &JWT{secret: secret}
}
