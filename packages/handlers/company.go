package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"xm-task/packages/domain"
	"xm-task/packages/storage"
)

type CompHandler struct {
	l *log.Logger
	d *domain.CompanyService
}

func NewCompanies(l *log.Logger, d *domain.CompanyService) *CompHandler {
	return &CompHandler{
		l: l,
		d: d,
	}
}

func (ch *CompHandler) ShowCompanies(w http.ResponseWriter, r *http.Request) {
	d := ch.d.ShowMany

	err := d.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (ch *CompHandler) ShowOneCompany(w http.ResponseWriter, r *http.Request) {}

func (ch *CompHandler) AddCompany(w http.ResponseWriter, r *http.Request) {
	cmp := r.Context().Value(KeyCompany{}).(domain.Company)

	cmp, err := ch.d.Create(cmp)
}

func (ch *CompHandler) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	cmp := r.Context().Value(KeyCompany{}).(domain.Company)

	cmp, err = ch.d.Update(code, cmp)
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
