package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"xm-task/packages/entities"
)

type KeyCompany struct{}

func (ch *CompHandler) MiddlewareCompanyValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cmp := entities.Company{}

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

func (ch *CompHandler) MiddlewareIPFilter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		IP := strings.Split(r.RemoteAddr, ":")[0]

		resp, err := http.Get(fmt.Sprintf("https://ipapi.co/%s/country/", IP))
		if err != nil {
			http.Error(w, "Unable to identify location", http.StatusForbidden)
			return
		}

		var cntr []byte

		_, err = resp.Body.Read(cntr)
		if err != nil {
			http.Error(w, "Unable to read response body", http.StatusBadRequest)
			return
		}

		if string(cntr) != "CY" {
			if err != nil {
				http.Error(w, "Access denied", http.StatusForbidden)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
