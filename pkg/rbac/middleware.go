package rbac

import (
	"context"
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Middleware struct {
	cache  *Cache
	decide func(context.Context, Input) error
}

func NewRBACMiddleware(cache *Cache, decide func(context.Context, Input) error) *Middleware {
	return &Middleware{
		cache:  cache,
		decide: decide,
	}
}

func (middleware *Middleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestRegoInput := Input{
			Path:   r.URL.Path,
			Method: r.Method,
			Token:  r.Header.Get("Authorization"),
			Cache:  middleware.cache,
		}

		err := middleware.decide(r.Context(), requestRegoInput)
		if err != nil {
			if errors.Is(err, ErrNotAuthorized) {
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
