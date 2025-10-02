package middlewares

import (
	"apiGo/utils"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return  func(w http.ResponseWriter, r *http.Request) {

		authReader := r.Header.Get("Authorization")

		if authReader == ""{
			http.Error(w, "token não fornecido", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authReader, "Bearer ")

		claims, err := utils.ValidateToken(tokenStr)

		if err != nil{
			http.Error(w, "token inválido", http.StatusForbidden)
			return
		}

		role, ok := claims["role"].(string)

		if !ok{
			http.Error(w, "role inválida", http.StatusUnauthorized)
			return
		}
		
		if role != "admin"{
			http.Error(w, "acesso inválido", http.StatusUnauthorized)
			return
		}

		

		next(w, r)
	}
}