package middlewares

import (
	"context"
	"fmt"
	"gql-tools/graph/utils"
	"net/http"
)

var userCtxKey = &contextKey{name: "user"}

type contextKey struct {
	name string
}

type UserAuth struct {
	Name  string
	Token string
}

func JWTMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			token := header

			name, err := utils.ParseToken(token)
			print("are you dying here")
			fmt.Println(name, err)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}
			userAuth := UserAuth{
				Name:  name,
				Token: token,
			}
			ctx := context.WithValue(r.Context(), userCtxKey, &userAuth)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func GetAuthFromContext(ctx context.Context) *UserAuth {
	raw, _ := ctx.Value(userCtxKey).(*UserAuth)
	return raw
}
