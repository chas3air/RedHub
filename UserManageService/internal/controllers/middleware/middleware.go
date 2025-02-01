package middleware

import (
	"fmt"
	"net/http"
	constants "userManageService/config"
	"userManageService/internal/lib/logger"
	"userManageService/internal/lib/logger/sl"
)

const adminRoleName = "admin"

func AdminMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	logger := logger.SetupLogger(constants.EnvLocal)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userRole := r.Header.Get("X-User-Role")
		if userRole != adminRoleName {
			logger.Error("Access Forbidden", sl.Err(fmt.Errorf("%d", http.StatusForbidden)))
			http.Error(w, "Fordidden", http.StatusForbidden)
			return
		}

		next(w, r)
	})
}
