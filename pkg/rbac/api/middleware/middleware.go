package middleware

import (
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/omc-college/management-system/pkg/rbac/authcache"
	"github.com/omc-college/management-system/pkg/rbac/opa"
)

type AuthorizationMiddleware struct {
	AuthCache *authcache.Cache
}

func NewAuthorizationMiddleware(authCache *authcache.Cache) *AuthorizationMiddleware {
	return &AuthorizationMiddleware{
		AuthCache: authCache,
	}
}

func (middleware *AuthorizationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestRegoInput := opa.RegoInput{
			Path:   r.URL.Path,
			Method: r.Method,
			Token:  r.Header.Get("Authorization"),
			Cache:  *middleware.AuthCache,
		}

		err := opa.GetDecision(r.Context(), requestRegoInput)
		if err != nil {
			if errors.Is(err, opa.ErrNotAuthorized) {
				http.Error(w, err.Error(), http.StatusForbidden)
				return
			}

			http.Error(w, "cannot get authorization decision", http.StatusInternalServerError)
			logrus.Errorf("cannot get authorization decision: %s", err.Error())
			return
		}

		next.ServeHTTP(w, r)
	})
}
