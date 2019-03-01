package utils

import (
	"io/ioutil"

	stdjwt "github.com/dgrijalva/jwt-go"
)

func GenerateJwtToken(claims stdjwt.StandardClaims, privateKeyPath string) (string, error) {
	key, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return "", err
	}
	parsedKey, err := stdjwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		return "", err
	}
	token, err := stdjwt.NewWithClaims(stdjwt.SigningMethodRS256, claims).SignedString(parsedKey)
	if err != nil {
		return "", err
	}
	return token, nil
}
