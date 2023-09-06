package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secret []byte
}

func (j *JWT) New(claims jwt.Claims) (signedJWT string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedJWT, _ = token.SignedString(j.secret)
	return signedJWT
}

func (j *JWT) KeyFunc(t *jwt.Token) (interface{}, error) { return j.secret, nil }

func New(secret []byte) *JWT {
	return &JWT{secret: secret}
}
