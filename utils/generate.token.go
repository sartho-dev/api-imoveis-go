package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const  JWT_SECRET_KEY = "293829452039stghsetg"

func GenerateToken(userId int, email string, role string) (string, error){
	claims := jwt.MapClaims{
		"user_id": userId,
		"role" : role,
		"email": email, 
		"exp" : time.Now().Add(time.Hour * 8).Unix(),
		"iat":  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err  := token.SignedString([]byte(JWT_SECRET_KEY))

	if err != nil{
		return "", err
	}

	return tokenStr, nil 
}

func ValidateToken(tokenStr string) (jwt.MapClaims, error){

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, fmt.Errorf("token inválido")
		}
		return []byte(JWT_SECRET_KEY), nil
	})

	if err != nil{
		return nil, fmt.Errorf("token inválido")
	}

	if !token.Valid || token == nil{
		return nil, fmt.Errorf("token inválido")
	}


	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok{
		return nil, fmt.Errorf("token invalido")
	}	

	return claims, nil

	
}