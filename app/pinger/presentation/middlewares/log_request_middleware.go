package middlewares

import (
	"log"
	"net/http"
)

type LogRequestMiddleware struct{}

func NewLogRequestMiddleware() *LogRequestMiddleware {
	return &LogRequestMiddleware{}
}

func (l *LogRequestMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(
			"Request: ",
			map[string]interface{}{
				"host":   r.URL.Host,
				"URI":    r.RequestURI,
				"method": r.Method,
			})

		next.ServeHTTP(w, r)
	})
}
