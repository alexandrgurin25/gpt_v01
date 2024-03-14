package middlewares

import (
	"app/internal/common"
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type tokenData struct {
	jwt.RegisteredClaims        // техническое поле для парсинга
	UserId               string `json:"sub"`
	CreatedAt            int64  `json:"iat"`
}

func AuthMiddleware(next http.Handler) http.Handler {
	secretKeyString, exists := os.LookupEnv("AUTH_SECRET_KEY")

	secretKey := []byte(secretKeyString)

	if !exists {
		log.Fatal("AUTH_SECRET_KEY not founded")
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessTokenHeader := r.Header.Get("Authorization") // получение данных из заголовка

		if len(accessTokenHeader) == 0 || accessTokenHeader[:7] != "Bearer " { // проверка, что токен начинается с корректного обозначения типа
			log.Printf("Could not get token %s", accessTokenHeader)
			common.HandleHttpError(w, common.ForbiddenError)
			return
		}

		accessTokenString := accessTokenHeader[7:] // извлечение самой строки токена
		token, err := jwt.ParseWithClaims(accessTokenString, &tokenData{}, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		ctx := r.Context()

		if data, ok := token.Claims.(*tokenData); ok && token.Valid {
			expirationTime := time.Now().Add(common.ExpirationTime).Unix()
			if data.CreatedAt > expirationTime {
				log.Print("accessToken timed out")
				common.HandleHttpError(w, common.ForbiddenError)
				return
			}
			ctx = context.WithValue(ctx, "userId", data.UserId)
		} else {
			log.Printf("%v", err)
			common.HandleHttpError(w, common.ForbiddenError)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
