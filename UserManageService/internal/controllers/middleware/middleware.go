package middleware

import (
	"net/http"
)

const adminRoleName = "admin"

func AdminMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: тут надо распарсить jwt токен, я пока не знаю как это сделать
	})
}
