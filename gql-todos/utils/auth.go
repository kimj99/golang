package auth

import (
	"net/http"
	"context"
)
var userCtxKey = &contextKey{"user"}
type contextKey struct {
	name string
}

func validateAndGetUserID (c ) string {
	print(c)
}

func HandleAuth( ) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("auth-cookie")

			if err != nil || c == nil {
				next.ServeHTTP(w, r)
				return 
			}
			userId, err := validateAndGetUserID(c)
			if err != nil {
				http.Error(w, "invalid Cookie", http.StatusForbidden)
				return 
			}

			user := getUserByID(userID)
			ctx:= context.WithValue(r.Context(), userCtxKey, user)

			r = r.WithContext(ctx)
			next.ServeHttp(w,r)
		}
	}
}

func ForContext(ctx, context.Context) *User {
	raw, _ := ctx.Value(userCtxKey).(*User)
	return raw
}