package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"xm-task/packages/storage"
)

type CompHandler struct {
	l *log.Logger
	s *storage.CompStorage
}

func NewCompanies(l *log.Logger, s *storage.CompStorage) *CompHandler {
	return &CompHandler{
		l: l,
		s: s,
	}
}

func (ch *CompHandler) ShowCompanies(w http.ResponseWriter, r *http.Request) {
	d := ch.s.GetCompanies()

	err := d.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (ch *CompHandler) ShowOneCompany(w http.ResponseWriter, r *http.Request) {}

func (ch *CompHandler) AddCompany(w http.ResponseWriter, r *http.Request) {
	cmp := r.Context().Value(KeyCompany{}).(storage.Company)

	ch.s.AddCompany(cmp)
}

func (ch *CompHandler) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	cmp := r.Context().Value(KeyCompany{}).(storage.Company)

	err = ch.s.UpdateCompany(id, &cmp)
	if err == storage.ErrCompanyNotFound {
		http.Error(w, "Company not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Company not found", http.StatusInternalServerError)
		return
	}

	return
}
