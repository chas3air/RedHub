package middleware

import (
	"net/http"
)

const adminRoleName = "admin"

func AdminMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: тут написать парсинг jwt токена
	})
}
