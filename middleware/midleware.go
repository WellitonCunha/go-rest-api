package middleware

import "net/http"

// esse middleware serve como um intermediario antes de chamar o controller
// ele chama a função para transformar a requisição em json
func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}
