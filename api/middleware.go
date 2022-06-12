package api

import (
	"context"
	"fmt"
	"net/http"
)

func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		cookie, err := request.Cookie("User-Id")
		if err != nil {
			fmt.Println(err)
			writer.WriteHeader(http.StatusNonAuthoritativeInfo)

		} else {
			ctx := context.WithValue(request.Context(), "UserId", cookie.Value)
			next.ServeHTTP(writer, request.WithContext(ctx))
		}
	})
}
