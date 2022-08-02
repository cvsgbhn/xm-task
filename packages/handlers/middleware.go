package handlers

import (
	"context"
	"net/http"
	"xm-task/packages/domain"
)

type KeyCompany struct{}

func (ch *CompHandler) MiddlewareCompanyValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cmp := &domain.Company{}

		err := cmp.FromJSON(r.Body)
		if err != nil {
			http.Error(w, "Unable to unmarshal JSON", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyCompany{}, cmp)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
