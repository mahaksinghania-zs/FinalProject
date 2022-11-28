package company

import (
	"Project3/Placement/svc"
	"net/http"
)

type Handler struct {
	service svc.Company
}

func New(service svc.Company) Handler {
	return Handler{service: service}
}

func (h Handler) Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetById(w, r)
	case http.MethodPost:
		h.Create(w, r)
	case http.MethodPut:
		h.Put(w, r)
	case http.MethodDelete:
		h.Delete(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
func (h Handler) GetById(w http.ResponseWriter, r *http.Request) {

}
func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
}
func (h Handler) Put(w http.ResponseWriter, r *http.Request) {

}
func (h Handler) Delete(w http.ResponseWriter, r *http.Request) {

}
