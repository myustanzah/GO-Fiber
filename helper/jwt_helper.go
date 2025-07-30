package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("local123") // Replace with your actual secret key

func GenerateJwtToken(email string) (string, error) {
	claims := &jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token valid for 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 	return nil, jwt.ErrSignatureInvalid
		// }
		return jwtSecret, nil
	})
	// if err != nil {
	// 	return nil, err
	// }
	// return token, nil
}
