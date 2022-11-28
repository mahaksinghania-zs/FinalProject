package student

import (
	"Project3/Placement/svc"
	"net/http"
)

type Handler struct {
	service svc.Student
}

func New(service svc.Student) Handler {
	return Handler{service: service}
}

func (h Handler) Handler(w http.ResponseWriter, r *http.Request) {
	//switch r.Method {
	//case http.MethodGet:
	//	h.get(w, r)
	//case http.MethodPost:
	//	h.create(w, r)
	//case http.MethodPut:
	//	h.put(w, r)
	//case http.MethodDelete:
	//	h.delete(w, r)
	//default:
	//	w.WriteHeader(http.StatusMethodNotAllowed)
	//}
}
func (h Handler) GetById(w http.ResponseWriter, r *http.Request) {

}
func (h Handler) Get(w http.ResponseWriter, r *http.Request) {

}
func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
}
func (h Handler) Put(w http.ResponseWriter, r *http.Request) {

}
func (h Handler) Delete(w http.ResponseWriter, r *http.Request) {

}
