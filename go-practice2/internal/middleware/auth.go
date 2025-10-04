package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Логируем метод и путь
		log.Printf("%s %s", r.Method, r.URL.Path)

		// Проверяем заголовок
		apiKey := r.Header.Get("X-API-Key")
		if apiKey != "secret123" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
			return
		}

		// Если ключ верный → пропускаем дальше
		next.ServeHTTP(w, r)
	})
}
