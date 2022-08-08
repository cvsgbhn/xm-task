package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"xm-task/packages/domain"
	"xm-task/packages/dtl"
	"xm-task/packages/entities"
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
	filterMap := r.URL.Query()
	d, err := ch.d.ShowMany(r.Context(), dtl.FilterToStruct(filterMap))
	if err != nil {
		http.Error(w, "Unable to get companies", http.StatusInternalServerError)
		return
	}

	err = d.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

func (ch *CompHandler) ShowCompanyByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	code, ok := vars["code"]
	if !ok {
		http.Error(w, "Unable to get code", http.StatusBadRequest)
		return
	}

	d, err := ch.d.ShowByCode(r.Context(), code)
	if err != nil {
		http.Error(w, "Unable to get company", http.StatusInternalServerError)
		return
	}

	err = d.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

func (ch *CompHandler) AddCompany(w http.ResponseWriter, r *http.Request) {
	cmp := r.Context().Value(KeyCompany{})

	cmp, err := ch.d.Create(r.Context(), cmp.(entities.Company))
	if err != nil {
		http.Error(w, "Unable to create a new company", http.StatusInternalServerError)
		return
	}
}

func (ch *CompHandler) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	code, ok := vars["code"]
	if !ok {
		http.Error(w, "Unable to get code", http.StatusBadRequest)
		return
	}

	cmp := r.Context().Value(KeyCompany{}).(entities.Company)

	cmp, err := ch.d.Update(r.Context(), code, cmp)
	if err != nil {
		http.Error(w, "Unable to update company", http.StatusInternalServerError)
		return
	}

	return
}

func (ch *CompHandler) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	code, ok := vars["code"]
	if !ok {
		http.Error(w, "Unable to get code", http.StatusBadRequest)
		return
	}

	err := ch.d.DeleteByCode(r.Context(), code)
	if err != nil {
		http.Error(w, "Unable to delete company", http.StatusInternalServerError)
		return
	}

	return
}
